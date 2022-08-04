package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	management "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	defs "github.com/Axway/agent-sdk/pkg/apic/definitions"
	prov "github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/authz/oauth"
	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
)

const (
	xAxwayEncrypted = "x-axway-encrypted"
	crFinalizer     = "agent.credential.provisioned"
)

type credProv interface {
	CredentialProvision(credentialRequest prov.CredentialRequest) (status prov.RequestStatus, credentails prov.Credential)
	CredentialDeprovision(credentialRequest prov.CredentialRequest) (status prov.RequestStatus)
}

type credentials struct {
	marketplaceHandler
	prov                credProv
	client              client
	encryptSchema       encryptSchemaFunc
	idpProviderRegistry oauth.ProviderRegistry
}

// encryptSchemaFunc func signature for encryptSchema
type encryptSchemaFunc func(schema, credData map[string]interface{}, key, alg, hash string) (map[string]interface{}, error)

// NewCredentialHandler creates a Handler for Credentials
func NewCredentialHandler(prov credProv, client client, providerRegistry oauth.ProviderRegistry) Handler {
	return &credentials{
		prov:                prov,
		client:              client,
		encryptSchema:       encryptSchema,
		idpProviderRegistry: providerRegistry,
	}
}

// Handle processes grpc events triggered for Credentials
func (h *credentials) Handle(ctx context.Context, meta *proto.EventMeta, resource *apiv1.ResourceInstance) error {
	action := GetActionFromContext(ctx)
	if resource.Kind != management.CredentialGVK().Kind || h.prov == nil || h.shouldIgnoreSubResourceUpdate(action, meta) {
		return nil
	}

	logger := getLoggerFromContext(ctx).WithComponent("credentialHandler")
	ctx = setLoggerInContext(ctx, logger)

	cr := &management.Credential{}
	err := cr.FromInstance(resource)
	if err != nil {
		logger.WithError(err).Error("could not handle credential request")
		return nil
	}

	if ok := isStatusFound(cr.Status); !ok {
		logger.Debugf("could not handle credential request as it did not have a status subresource")
		return nil
	}

	if ok := h.shouldProcessPending(cr.Status, cr.Metadata.State); ok {
		log.Trace("processing resource in pending status")
		ar := h.onPending(ctx, cr)
		err := h.client.CreateSubResource(cr.ResourceMeta, cr.SubResources)
		if err != nil {
			logger.WithError(err).Error("error creating subresources")
		}

		// update the status regardless of errors updating the other subresources
		statusErr := h.client.CreateSubResource(ar.ResourceMeta, map[string]interface{}{"status": ar.Status})
		if statusErr != nil {
			logger.WithError(statusErr).Error("error creating status subresources")
			return statusErr
		}

		return err
	}

	if ok := h.shouldProcessDeleting(cr.Status, cr.Metadata.State, cr.Finalizers); ok {
		logger.Trace("processing resource in deleting state")
		h.onDeleting(ctx, cr)
	}

	return nil
}

func (h *credentials) getReasonMetaAction(reasons []apiv1.ResourceStatusReason) string {
	if len(reasons) != 1 {
		return ""
	}
	if reasons[0].Meta == nil {
		return ""
	}
	if action, found := reasons[0].Meta["action"]; found {
		return fmt.Sprintf("%v", action)
	}
	return ""
}

// shouldProcessDeleting returns true when the resource is in a deleting state and has finalizers or when it is in Error and the only reason is CredentialExpired
func (h *credentials) shouldProcessDeleting(status *apiv1.ResourceStatus, state string, finalizers []apiv1.Finalizer) bool {
	switch {
	case len(finalizers) == 0:
		return false
	case h.marketplaceHandler.shouldProcessDeleting(status, state, finalizers):
		fallthrough
	case status.Level == prov.Error.String() && h.getReasonMetaAction(status.Reasons) == "CredentialExpired":
		return true
	default:
		return false
	}
}

func (h *credentials) onPending(ctx context.Context, cred *management.Credential) *management.Credential {
	logger := getLoggerFromContext(ctx)
	app, err := h.getManagedApp(ctx, cred)
	if err != nil {
		logger.WithError(err).Error("error getting managed app")
		h.onError(ctx, cred, err)
		return cred
	}

	// check the application status
	if app.Status.Level != prov.Success.String() {
		err = fmt.Errorf("cannot handle credential when application is not yet successful")
		h.onError(ctx, cred, err)
		return cred
	}

	crd, err := h.getCRD(ctx, cred)
	if err != nil {
		logger.WithError(err).Error("error getting credential request definition")
		h.onError(ctx, cred, err)
		return cred
	}

	provCreds, err := newProvCreds(cred, util.GetAgentDetails(app), nil, h.idpProviderRegistry)
	if err != nil {
		logger.WithError(err).Error("error preparing credential request")
		h.onError(ctx, cred, err)
		return cred
	}

	if provCreds.IsIDPCredential() {
		err := h.registerIDPClientCredential(provCreds)
		if err != nil {
			logger.WithError(err).Error("error provisioning credential request with IDP")
			h.onError(ctx, cred, err)
			return cred
		}
	}

	data := map[string]interface{}{}
	status, credentialData := h.prov.CredentialProvision(provCreds)

	if status.GetStatus() == prov.Success && credentialData != nil {
		credentialData = h.getProvisionedCredentialData(provCreds, credentialData)
		sec := app.Spec.Security
		d := credentialData.GetData()
		if crd.Spec.Provision == nil {
			data = d
		} else if d != nil {
			data, err = h.encryptSchema(
				crd.Spec.Provision.Schema,
				d,
				sec.EncryptionKey, sec.EncryptionAlgorithm, sec.EncryptionHash,
			)
		}

		if err != nil {
			status = prov.NewRequestStatusBuilder().
				SetMessage(fmt.Sprintf("error encrypting credential: %s", err.Error())).
				SetCurrentStatusReasons(cred.Status.Reasons).
				Failed()
		}
	}

	cred.Data = data
	cred.Status = prov.NewStatusReason(status)

	details := util.MergeMapStringString(util.GetAgentDetailStrings(cred), status.GetProperties())
	util.SetAgentDetails(cred, util.MapStringStringToMapStringInterface(details))

	ri, _ := cred.AsInstance()
	if cred.Status.Level == prov.Success.String() {
		// only add finalizer on success
		h.client.UpdateResourceFinalizer(ri, crFinalizer, "", true)
	}

	cred.SubResources = map[string]interface{}{
		defs.XAgentDetails: util.GetAgentDetails(cred),
		"data":             cred.Data,
	}

	return cred
}

func (h *credentials) onDeleting(ctx context.Context, cred *management.Credential) {
	logger := getLoggerFromContext(ctx)
	var provData map[string]interface{}
	if cred.Data != nil {
		if m, ok := cred.Data.(map[string]interface{}); ok {
			provData = m
		}
	}

	provCreds, err := newProvCreds(cred, map[string]interface{}{}, provData, h.idpProviderRegistry)
	if err != nil {
		logger.WithError(err).Error("error preparing credential request")
		h.onError(ctx, cred, err)
		return
	}

	status := h.prov.CredentialDeprovision(provCreds)
	if status.GetStatus() == prov.Success {
		if provCreds.IsIDPCredential() {
			err := h.unregisterIDPClientCredential(provCreds)
			if err != nil {
				logger.WithError(err).Error("error deprovisioning credential request from IDP")
				h.onError(ctx, cred, err)
				return
			}
		}

		ri, _ := cred.AsInstance()
		h.client.UpdateResourceFinalizer(ri, crFinalizer, "", false)

		// Delete the resource, since it was not in Deleting State
		if ri.Metadata.State != apiv1.ResourceDeleting {
			h.client.DeleteResourceInstance(ri)
		}
	} else {
		err := fmt.Errorf(status.GetMessage())
		logger.WithError(err).Error("request status was not Success, skipping")
		h.onError(ctx, cred, err)
		h.client.CreateSubResource(cred.ResourceMeta, cred.SubResources)
	}
}

// onError updates the AccessRequest with an error status
func (h *credentials) onError(_ context.Context, cred *management.Credential, err error) {
	ps := prov.NewRequestStatusBuilder()
	status := ps.SetMessage(err.Error()).SetCurrentStatusReasons(cred.Status.Reasons).Failed()
	cred.Status = prov.NewStatusReason(status)
	cred.SubResources = map[string]interface{}{
		"status": cred.Status,
	}
}

func (h *credentials) getManagedApp(ctx context.Context, cred *management.Credential) (*management.ManagedApplication, error) {
	app := management.NewManagedApplication(cred.Spec.ManagedApplication, cred.Metadata.Scope.Name)
	ri, err := h.client.GetResource(app.GetSelfLink())
	if err != nil {
		return nil, err
	}

	app = &management.ManagedApplication{}
	err = app.FromInstance(ri)
	return app, err
}

func (h *credentials) getCRD(ctx context.Context, cred *management.Credential) (*management.CredentialRequestDefinition, error) {
	crd := management.NewCredentialRequestDefinition(cred.Spec.CredentialRequestDefinition, cred.Metadata.Scope.Name)
	ri, err := h.client.GetResource(crd.GetSelfLink())
	if err != nil {
		return nil, err
	}

	crd = &management.CredentialRequestDefinition{}
	err = crd.FromInstance(ri)
	return crd, err
}

func (h *credentials) registerIDPClientCredential(cr *provCreds) error {
	p := cr.GetIDPProvider()
	idpCredData := cr.GetIDPCredentialData()

	formattedJWKS := strings.ReplaceAll(idpCredData.GetPublicKey(), "----- ", "-----\n")
	formattedJWKS = strings.ReplaceAll(formattedJWKS, " -----", "\n-----")

	// prepare external client metadata from CRD data
	clientMetadata, err := oauth.NewClientMetadataBuilder().
		SetClientName(cr.GetName()).
		SetScopes(idpCredData.GetScopes()).
		SetGrantTypes(idpCredData.GetGrantTypes()).
		SetTokenEndpointAuthMethod(idpCredData.GetTokenEndpointAuthMethod()).
		SetResponseType(idpCredData.GetResponseTypes()).
		SetRedirectURIs(idpCredData.GetRedirectURIs()).
		SetJWKS([]byte(formattedJWKS)).
		SetJWKSURI(idpCredData.GetJwksURI()).
		Build()
	if err != nil {
		return err
	}

	// provision external client
	resClientMetadata, err := p.RegisterClient(clientMetadata)
	if err != nil {
		return err
	}

	cr.idpCredData.clientID = resClientMetadata.GetClientID()
	cr.idpCredData.clientSecret = resClientMetadata.GetClientSecret()
	return nil
}

func (h *credentials) unregisterIDPClientCredential(cr *provCreds) error {
	p := cr.GetIDPProvider()
	err := p.UnregisterClient(cr.idpCredData.GetClientID())
	if err != nil {
		return err
	}

	cr.idpCredData.clientID = cr.idpCredData.GetClientID()
	return nil
}

func (h *credentials) getProvisionedCredentialData(provCreds *provCreds, credentialData prov.Credential) prov.Credential {
	if provCreds.IsIDPCredential() {
		return prov.NewCredentialBuilder().SetOAuthIDAndSecret(
			provCreds.GetIDPCredentialData().GetClientID(),
			provCreds.GetIDPCredentialData().GetClientSecret(),
		)
	}
	return credentialData
}

type provCreds struct {
	managedApp  string
	credType    string
	id          string
	name        string
	credData    map[string]interface{}
	credDetails map[string]interface{}
	appDetails  map[string]interface{}
	idpCredData *idpCredData
	idpProvider oauth.Provider
}

type idpCredData struct {
	clientID        string
	clientSecret    string
	scopes          []string
	grantTypes      []string
	tokenAuthMethod string
	responseTypes   []string
	redirectURLs    []string
	jwksURI         string
	publicKey       string
}

func newProvCreds(cr *management.Credential, appDetails map[string]interface{}, provData map[string]interface{}, idpProviderRegistry oauth.ProviderRegistry) (*provCreds, error) {
	credDetails := util.GetAgentDetails(cr)

	provCred := &provCreds{
		appDetails:  appDetails,
		credDetails: credDetails,
		credType:    cr.Spec.CredentialRequestDefinition,
		credData:    cr.Spec.Data,
		managedApp:  cr.Spec.ManagedApplication,
		id:          cr.Metadata.ID,
		name:        cr.Name,
	}

	// Setup external credential request data to be used for provisioning
	if idpTokenURL, ok := provCred.credData[prov.IDPTokenURL].(string); ok && idpProviderRegistry != nil {
		p, err := idpProviderRegistry.GetProviderByTokenEndpoint(idpTokenURL)
		if err != nil {
			return nil, fmt.Errorf("IDP provider not found for credential request")
		}
		provCred.idpProvider = p
		provCred.idpCredData = newIDPCredData(p, provCred.credData, provData)
	}

	return provCred, nil
}

// newIDPCredData - reads the idp client metadata from credential request
func newIDPCredData(p oauth.Provider, credData, provData map[string]interface{}) *idpCredData {
	cd := &idpCredData{}

	if provData != nil {
		cd.clientID = util.GetStringFromMapInterface(prov.OauthClientID, provData)
	}
	cd.scopes = util.GetStringArrayFromMapInterface(prov.OauthScopes, credData)
	cd.grantTypes = []string{util.GetStringFromMapInterface(prov.OauthGrantType, credData)}
	cd.redirectURLs = util.GetStringArrayFromMapInterface(prov.OauthRedirectURIs, credData)
	cd.tokenAuthMethod = util.GetStringFromMapInterface(prov.OauthTokenAuthMethod, credData)
	cd.publicKey = util.GetStringFromMapInterface(prov.OauthJwks, credData)
	cd.jwksURI = util.GetStringFromMapInterface(prov.OauthJwksURI, credData)

	return cd
}

// GetApplicationName gets the name of the managed application
func (c provCreds) GetApplicationName() string {
	return c.managedApp
}

// GetID gets the id of the credential resource
func (c provCreds) GetID() string {
	return c.id
}

// GetName gets the name of the credential resource
func (c provCreds) GetName() string {
	return c.name
}

// GetCredentialType gets the type of the credential
func (c provCreds) GetCredentialType() string {
	return c.credType
}

// GetCredentialData gets the data of the credential
func (c provCreds) GetCredentialData() map[string]interface{} {
	return c.credData
}

// IsIDPCredential returns boolean indicating if the credential request is for IDP provider
func (c provCreds) IsIDPCredential() bool {
	return c.idpProvider != nil
}

// GetIDPProvider returns the interface for IDP provider if the credential request is for IDP provider
func (c provCreds) GetIDPProvider() oauth.Provider {
	return c.idpProvider
}

// GetIDPCredentialData returns the credential data for IDP from the request
func (c provCreds) GetIDPCredentialData() prov.IDPCredentialData {
	return c.idpCredData
}

// GetCredentialDetailsValue returns a value found on the 'x-agent-details' sub resource of the Credentials.
func (c provCreds) GetCredentialDetailsValue(key string) string {
	if c.credDetails == nil {
		return ""
	}

	return util.ToString(c.credDetails[key])
}

// GetApplicationDetailsValue returns a value found on the 'x-agent-details' sub resource of the ManagedApplication.
func (c provCreds) GetApplicationDetailsValue(key string) string {
	if c.appDetails == nil {
		return ""
	}

	return util.ToString(c.appDetails[key])
}

// GetClientID - returns client ID
func (c *idpCredData) GetClientID() string {
	return c.clientID
}

// GetClientSecret - returns client secret
func (c *idpCredData) GetClientSecret() string {
	return c.clientSecret
}

// GetScopes - returns client scopes
func (c *idpCredData) GetScopes() []string {
	return c.scopes
}

// GetGrantTypes - returns grant types
func (c *idpCredData) GetGrantTypes() []string {
	return c.grantTypes
}

// GetTokenEndpointAuthMethod - returns token auth method
func (c *idpCredData) GetTokenEndpointAuthMethod() string {
	return c.tokenAuthMethod
}

// GetResponseTypes - returns token response type
func (c *idpCredData) GetResponseTypes() []string {
	return c.responseTypes
}

// GetRedirectURIs - Returns redirect urls
func (c *idpCredData) GetRedirectURIs() []string {
	return c.redirectURLs
}

// GetJwksURI - returns JWKS uri
func (c *idpCredData) GetJwksURI() string {
	return c.jwksURI
}

// GetPublicKey - returns the public key
func (c *idpCredData) GetPublicKey() string {
	return c.publicKey
}

// encryptSchema schema is the json schema. credData is the data that contains data to encrypt based on the key, alg and hash.
func encryptSchema(
	schema, credData map[string]interface{}, key, alg, hash string,
) (map[string]interface{}, error) {
	enc, err := util.NewEncryptor(key, alg, hash)
	if err != nil {
		return nil, err
	}

	schemaProps, ok := schema["properties"]
	if !ok {
		return nil, fmt.Errorf("properties field not found on schema")
	}

	props, ok := schemaProps.(map[string]interface{})
	if !ok {
		props = make(map[string]interface{})
	}

	data := encryptMap(enc, props, credData)
	return data, nil
}

// encryptMap loops through all data and checks the value against the provisioning schema to see if it should be encrypted.
func encryptMap(enc util.Encryptor, schema, data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		schemaValue := schema[key]
		v, ok := schemaValue.(map[string]interface{})
		if !ok {
			continue
		}

		if _, ok := v[xAxwayEncrypted]; ok {
			v, ok := value.(string)
			if !ok {
				continue
			}

			str, err := enc.Encrypt(v)
			if err != nil {

				log.Error(err)
				continue
			}

			data[key] = base64.StdEncoding.EncodeToString([]byte(str))
		}
	}

	return data
}
