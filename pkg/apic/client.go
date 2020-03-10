package apic

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	coreapi "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/api"
	corecfg "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/config"
	"git.ecd.axway.int/apigov/service-mesh-agent/pkg/apicauth"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// constants for auth policy types
const (
	Apikey      = "verify-api-key"
	Passthrough = "pass-through"
	Oauth       = "verify-oauth-token"
)

// ValidPolicies - list of valid auth policies supported by Central.  Add to this list as more policies are supported.
var ValidPolicies = []string{Apikey, Passthrough, Oauth}

// SubscriptionProcessor - callback method type to process subscriptions
type SubscriptionProcessor func(subscription Subscription)

// SubscriptionValidator - callback method type to validate subscription for processing
type SubscriptionValidator func(subscription Subscription) bool

// Client - interface
type Client interface {
	CreateService(serviceBody ServiceBody) (string, error)
	UpdateService(ID string, serviceBody ServiceBody) (string, error)
	RegisterSubscriptionSchema(subscriptionSchema SubscriptionSchema) error
	GetSubscriptionManager() SubscriptionManager
	CheckHealth() error
}

type tokenGetter interface {
	GetToken() (string, error)
}

type platformTokenGetter struct {
	requester *apicauth.PlatformTokenGetter
}

func (p *platformTokenGetter) GetToken() (string, error) {
	return p.requester.GetToken()
}

// ServiceClient -
type ServiceClient struct {
	tokenRequester               tokenGetter
	cfg                          corecfg.CentralConfig
	apiClient                    coreapi.Client
	DefaultSubscriptionSchema    SubscriptionSchema
	RegisteredSubscriptionSchema SubscriptionSchema
	subscriptionMgr              SubscriptionManager
}

// New -
func New(cfg corecfg.CentralConfig) Client {
	tokenURL := cfg.GetAuthConfig().GetTokenURL()
	aud := cfg.GetAuthConfig().GetAudience()
	priKey := cfg.GetAuthConfig().GetPrivateKey()
	pubKey := cfg.GetAuthConfig().GetPublicKey()
	keyPwd := cfg.GetAuthConfig().GetKeyPassword()
	clientID := cfg.GetAuthConfig().GetClientID()
	authTimeout := cfg.GetAuthConfig().GetTimeout()
	platformTokenGetter := &platformTokenGetter{
		requester: apicauth.NewPlatformTokenGetter(priKey, pubKey, keyPwd, tokenURL, aud, clientID, authTimeout),
	}
	serviceClient := &ServiceClient{
		cfg:                       cfg,
		tokenRequester:            platformTokenGetter,
		apiClient:                 coreapi.NewClient(cfg.GetTLSConfig(), cfg.GetProxyURL()),
		DefaultSubscriptionSchema: NewSubscriptionSchema(),
	}
	serviceClient.subscriptionMgr = newSubscriptionManager(serviceClient)
	return serviceClient
}

// mapToTagsArray -
func (c *ServiceClient) mapToTagsArray(m map[string]interface{}) []string {
	strArr := []string{}

	for key, val := range m {
		var value string
		v, ok := val.(*string)
		if ok {
			value = *v
		} else {
			v, ok := val.(string)
			if ok {
				value = v
			}
		}
		if value == "" {
			strArr = append(strArr, key)
		} else {
			strArr = append(strArr, key+"_"+value)
		}
	}

	// Add any tags from config
	additionalTags := c.cfg.GetTagsToPublish()
	if additionalTags != "" {
		additionalTagsArray := strings.Split(additionalTags, ",")

		for _, tag := range additionalTagsArray {
			strArr = append(strArr, strings.TrimSpace(tag))
		}
	}

	return strArr
}

var log logrus.FieldLogger = logrus.WithField("package", "apic")

// SetLog sets the logger for the package.
func SetLog(newLog logrus.FieldLogger) {
	log = newLog
	return
}

func isUnitTesting() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}

// deployAPI -
func (c *ServiceClient) deployAPI(method, url string, buffer []byte) (string, error) {
	// Unit testing. For now just dummy up a return
	if isUnitTesting() {
		return "12345678", nil
	}

	headers, err := c.createHeader()
	if err != nil {
		return "", err
	}

	request := coreapi.Request{
		Method:      method,
		URL:         url,
		QueryParams: nil,
		Headers:     headers,
		Body:        buffer,
	}
	response, err := c.apiClient.Send(request)
	if err != nil {
		return "", err
	}
	//  Check to see if rollback was processed
	if method == http.MethodDelete && response.Code == http.StatusNoContent {
		log.Error("Rollback API service.  API has been removed.")
		logResponseErrors(response.Body)
		return "", errors.New(strconv.Itoa(response.Code))
	}

	if !(response.Code == http.StatusOK || response.Code == http.StatusCreated) {
		logResponseErrors(response.Body)
		return "", errors.New(strconv.Itoa(response.Code))
	}

	return c.handleResponse(response.Body)
}

type apiErrorResponse map[string][]apiError

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func logResponseErrors(body []byte) {
	detail := make(map[string]*json.RawMessage)
	json.Unmarshal(body, &detail)

	for k, v := range detail {
		buffer, _ := v.MarshalJSON()
		log.Debugf("HTTP response %v: %v", k, string(buffer))
	}
}

func (c *ServiceClient) handleResponse(body []byte) (string, error) {

	itemID := ""

	// Connected Mode
	if c.cfg.GetAgentMode() == corecfg.Connected {
		metadata := gjson.Get(string(body), "metadata").String()
		if metadata != "" {
			itemID = gjson.Get(string(metadata), "id").String()
		}
		// Disconnected Mode
	} else {
		itemID = gjson.Get(string(body), "id").String()
	}

	log.Debugf("HTTP response returning itemID: [%v]", itemID)
	return itemID, nil
}

func (c *ServiceClient) createHeader() (map[string]string, error) {
	token, err := c.tokenRequester.GetToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["X-Axway-Tenant-Id"] = c.cfg.GetTenantID()
	headers["Authorization"] = "Bearer " + token
	headers["Content-Type"] = "application/json"
	return headers, nil
}

// IsNewAPI -
func (c *ServiceClient) isNewAPI(serviceBody ServiceBody) bool {
	var token string
	apiName := strings.ToLower(serviceBody.APIName)
	request, err := http.NewRequest("GET", c.cfg.GetAPIServerServicesURL()+"/"+apiName, nil)

	if token, err = c.tokenRequester.GetToken(); err != nil {
		log.Error("Could not get token")
	}

	request.Header.Add("X-Axway-Tenant-Id", c.cfg.GetTenantID())
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode == http.StatusNotFound {
		log.Debugf("New api found to deploy: %s", apiName)
		return true
	}
	return false
}

// GetSubscriptionManager -
func (c *ServiceClient) GetSubscriptionManager() SubscriptionManager {
	return c.subscriptionMgr
}

// CheckHealth -
func (c *ServiceClient) CheckHealth() error {
	err := c.checkPlatformHealth()
	if err != nil {
		return err
	}

	return c.checkAPIServerHealth()
}

func (c *ServiceClient) checkPlatformHealth() error {
	_, err := c.tokenRequester.GetToken()
	if err != nil {
		return fmt.Errorf("error trying to get platform token: %s. Check AMPLIFY Central configuration for URL, AUTH_URL, AUTH_REALM, AUTH_CLIENTID, AUTH_PRIVATEKEY, and AUTH_PUBLICKEY", err.Error())
	}
	return nil
}

func (c *ServiceClient) checkAPIServerHealth() error {
	// do a request for catalog items
	headers, err := c.createHeader()
	if err != nil {
		return fmt.Errorf("error creating request header. %s", err.Error())
	}

	if c.cfg.GetAgentMode() == corecfg.Disconnected {
		sendErr := "error sending request to API Server: %s. Check AMPLIFY Central configuration for URL"
		statusErr := "error sending request to API Server - status code %d. Check AMPLIFY Central configuration"

		// do a request for catalog items
		return c.sendServerRequest(c.cfg.GetCatalogItemsURL(), headers, sendErr, statusErr)
	}

	if c.cfg.GetAgentMode() == corecfg.Connected {
		sendErr := "error sending request to API Server: %s. Check AMPLIFY Central configuration for URL and APISERVERENVIRONMENT"
		statusErr := "error sending request to API Server - status code %d. Check AMPLIFY Central configuration for APISERVERENVIRONMENT"

		// do a request for the environment
		return c.sendServerRequest(c.cfg.GetAPIServerURL()+c.cfg.GetEnvironmentName(), headers, sendErr, statusErr)
	}

	return nil
}

func (c *ServiceClient) sendServerRequest(url string, headers map[string]string, sendErr, statusErr string) error {
	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     url,
		Headers: headers,
	}
	response, err := c.apiClient.Send(request)
	if err != nil {
		return fmt.Errorf(sendErr, err.Error())
	}
	if response.Code != http.StatusOK {
		logResponseErrors(response.Body)
		return fmt.Errorf(statusErr, response.Code)
	}
	return nil
}
