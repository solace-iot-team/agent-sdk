package apic

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	coreapi "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/api"
	log "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/util/log"
	"github.com/tidwall/gjson"
)

// AddToAPIC -
func (c *ServiceClient) addCatalog(serviceBody ServiceBody) (string, error) {
	// add createdBy as a tag
	serviceBody.Tags["createdBy_"+serviceBody.CreatedBy] = ""

	serviceBody.ServiceExecution = addCatalog
	itemID, err := c.deployCatalog(serviceBody, http.MethodPost, c.cfg.GetCatalogItemsURL())
	if err != nil {
		return "", err
	}
	if serviceBody.Image != "" {
		serviceBody.ServiceExecution = addCatalogImage
		_, err = c.deployCatalog(serviceBody, http.MethodPost, c.cfg.GetCatalogItemImageURL(itemID))
		if err != nil {
			log.Warn("Unable to add image to the catalog item. " + err.Error())
		}
	}
	return itemID, nil
}

func (c *ServiceClient) deployCatalog(serviceBody ServiceBody, method, url string) (string, error) {
	if !isValidAuthPolicy(serviceBody.AuthPolicy) {
		return "", fmt.Errorf("Unsupported security policy '%v'. ", serviceBody.AuthPolicy)
	}

	buffer, err := c.createCatalogBody(serviceBody)
	if err != nil {
		log.Error("Error creating service item: ", err)
		return "", err
	}

	return c.catalogDeployAPI(method, url, buffer)
}

func (c *ServiceClient) createCatalogBody(serviceBody ServiceBody) ([]byte, error) {
	var spec []byte
	var err error
	switch serviceBody.ServiceExecution {
	case addCatalog:
		spec, err = c.marshalCatalogItemInit(serviceBody)
	case addCatalogImage:
		spec, err = c.marshalCatalogItemImage(serviceBody)
	case updateCatalog:
		spec, err = c.marshalCatalogItem(serviceBody)
	case updateCatalogRevision:
		spec, err = c.marshalCatalogItemRevision(serviceBody)
	default:
		return nil, errors.New("Invalid catalog operation")
	}
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func (c *ServiceClient) marshalCatalogItemInit(serviceBody ServiceBody) ([]byte, error) {
	enableSubscription := (serviceBody.AuthPolicy != Passthrough)

	// assume that we use the default schema unless it one is enabled and registered
	subSchema := c.DefaultSubscriptionSchema
	if enableSubscription {
		if c.RegisteredSubscriptionSchema != nil {
			subSchema = c.RegisteredSubscriptionSchema
		} else {
			enableSubscription = false
		}
	}

	catalogSubscriptionSchema, err := subSchema.rawJSON()
	if err != nil {
		return nil, err
	}

	definitionSubType, revisionPropertyKey := c.getDefinitionSubtypeAndRevisionKey(serviceBody)

	catalogProperties := []CatalogProperty{}
	if definitionSubType != Wsdl {
		catalogProperties = []CatalogProperty{
			{
				Key: "accessInfo",
				Value: CatalogPropertyValue{
					AuthPolicy: serviceBody.AuthPolicy,
					URL:        serviceBody.URL,
				},
			},
		}
	}

	newCatalogItem := CatalogItemInit{
		DefinitionType:     API,
		DefinitionSubType:  definitionSubType,
		DefinitionRevision: 1,
		Name:               serviceBody.NameToPush,
		OwningTeamID:       serviceBody.TeamID,
		Description:        serviceBody.Description,
		Properties:         catalogProperties,
		Tags:               c.mapToTagsArray(serviceBody.Tags),
		Visibility:         "RESTRICTED", // default value
		Subscription: CatalogSubscription{
			Enabled:         enableSubscription,
			AutoSubscribe:   true,
			AutoUnsubscribe: false,
			Properties: []CatalogRevisionProperty{{
				Key:   "profile",
				Value: catalogSubscriptionSchema,
			}},
		},
		Revision: CatalogItemInitRevision{
			Version: serviceBody.Version,
			State:   unpublishedState,
			Properties: []CatalogRevisionProperty{
				{
					Key:   "documentation",
					Value: json.RawMessage(string(serviceBody.Documentation)),
				},
				{
					Key:   revisionPropertyKey,
					Value: c.getRawMessageFromSwagger(serviceBody),
				},
			},
		},
	}

	return json.Marshal(newCatalogItem)
}

// marshal the CatalogItem -
func (c *ServiceClient) marshalCatalogItem(serviceBody ServiceBody) ([]byte, error) {

	definitionSubType, _ := c.getDefinitionSubtypeAndRevisionKey(serviceBody)

	newCatalogItem := CatalogItem{
		DefinitionType:    API,
		DefinitionSubType: definitionSubType,

		DefinitionRevision: 1,
		Name:               serviceBody.NameToPush,
		OwningTeamID:       serviceBody.TeamID,
		Description:        serviceBody.Description,
		Tags:               c.mapToTagsArray(serviceBody.Tags),
		Visibility:         "RESTRICTED",     // default value
		State:              unpublishedState, //default
		LatestVersionDetails: CatalogItemRevision{
			Version: serviceBody.Version,
			State:   publishedState,
		},
	}

	return json.Marshal(newCatalogItem)
}

// marshal the CatalogItem revision
func (c *ServiceClient) marshalCatalogItemRevision(serviceBody ServiceBody) ([]byte, error) {

	_, revisionPropertyKey := c.getDefinitionSubtypeAndRevisionKey(serviceBody)

	catalogItemRevision := CatalogItemInitRevision{
		Version: serviceBody.Version,
		State:   unpublishedState,
		Properties: []CatalogRevisionProperty{
			{
				Key:   "documentation",
				Value: json.RawMessage(string(serviceBody.Documentation)),
			},
			{
				Key:   revisionPropertyKey,
				Value: c.getRawMessageFromSwagger(serviceBody),
			},
		},
	}

	return json.Marshal(catalogItemRevision)
}

// marshals the catalog image body
func (c *ServiceClient) marshalCatalogItemImage(serviceBody ServiceBody) ([]byte, error) {
	catalogImage := CatalogItemImage{
		DataType:      serviceBody.ImageContentType,
		Base64Content: serviceBody.Image,
	}
	return json.Marshal(catalogImage)
}

func (c *ServiceClient) getDefinitionSubtypeAndRevisionKey(serviceBody ServiceBody) (definitionSubType, revisionPropertyKey string) {
	if serviceBody.ResourceType == Wsdl {
		definitionSubType = Wsdl
		revisionPropertyKey = Specification
	} else {
		oasVer := gjson.GetBytes(serviceBody.Swagger, "openapi")
		definitionSubType = SwaggerV2
		revisionPropertyKey = Swagger
		if oasVer.Exists() {
			// OAS v3
			definitionSubType = Oas3
			revisionPropertyKey = Specification
		}
	}
	return
}

func (c *ServiceClient) getRawMessageFromSwagger(serviceBody ServiceBody) (rawMsg json.RawMessage) {
	if serviceBody.ResourceType == Wsdl {
		str := base64.StdEncoding.EncodeToString(serviceBody.Swagger)
		rawMsg = json.RawMessage(strconv.Quote(str))
	} else {
		rawMsg = json.RawMessage(serviceBody.Swagger)
	}
	return
}

// UpdateCatalogItemRevisions -
func (c *ServiceClient) UpdateCatalogItemRevisions(ID string, serviceBody ServiceBody) (string, error) {
	serviceBody.ServiceExecution = updateCatalogRevision
	return c.deployCatalog(serviceBody, http.MethodPost, c.cfg.UpdateCatalogItemRevisions(ID))
}

// GetCatalogItemRevision -
func (c *ServiceClient) GetCatalogItemRevision(ID string) (string, error) {
	headers, err := c.createHeader()
	if err != nil {
		return "", err
	}

	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     c.cfg.GetCatalogItemByID(ID),
		Headers: headers,
	}

	response, err := c.apiClient.Send(request)
	if err != nil {
		return "", err
	}
	if !(response.Code == http.StatusOK) {
		logResponseErrors(response.Body)
		return "", errors.New(strconv.Itoa(response.Code))
	}

	revisions := gjson.Get(string(response.Body), "availableRevisions")
	availableRevisions := make([]int, 0)
	json.Unmarshal([]byte(revisions.Raw), &availableRevisions)
	revision := availableRevisions[len(availableRevisions)-1] // get the latest revsions
	return strconv.Itoa(revision), nil
}

func isValidAuthPolicy(auth string) bool {
	for _, item := range ValidPolicies {
		if item == auth {
			return true
		}
	}
	return false
}

// updateCatalog -
func (c *ServiceClient) updateCatalog(ID string, serviceBody ServiceBody) (string, error) {
	serviceBody.ServiceExecution = updateCatalog
	_, err := c.deployCatalog(serviceBody, http.MethodPut, c.cfg.GetCatalogItemsURL()+"/"+ID)
	if err != nil {
		return "", err
	}

	version, err := c.GetCatalogItemRevision(ID)
	i, err := strconv.Atoi(version)

	serviceBody.Version = strconv.Itoa(i + 1)
	_, err = c.UpdateCatalogItemRevisions(ID, serviceBody)
	if err != nil {
		return "", err
	}

	return ID, nil
}

// catalogDeployAPI -
func (c *ServiceClient) catalogDeployAPI(method, url string, buffer []byte) (string, error) {
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

	if !(response.Code == http.StatusOK || response.Code == http.StatusCreated) {
		logResponseErrors(response.Body)
		return "", errors.New(strconv.Itoa(response.Code))
	}

	itemID := gjson.Get(string(response.Body), "id").String()

	log.Debugf("HTTP response returning itemID: [%v]", itemID)
	return itemID, nil

}