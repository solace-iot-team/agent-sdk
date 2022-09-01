package cache

import (
	"fmt"
	"os"
	"testing"

	defs "github.com/Axway/agent-sdk/pkg/apic/definitions"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/stretchr/testify/assert"
)

func createAPIService(apiID, apiName, primaryKey string) *v1.ResourceInstance {
	sub := map[string]interface{}{
		defs.AttrExternalAPIID:         apiID,
		defs.AttrExternalAPIName:       apiName,
		defs.AttrExternalAPIPrimaryKey: primaryKey,
	}

	if primaryKey != "" {
		sub[defs.AttrExternalAPIPrimaryKey] = primaryKey
	}

	return &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Name: fmt.Sprintf("name-%s", apiName),
			SubResources: map[string]interface{}{
				defs.XAgentDetails: sub,
			},
		},
	}
}

func createAPIServiceInstance(id, apiID, stage string) *v1.ResourceInstance {
	sub := map[string]interface{}{
		defs.AttrExternalAPIID:    apiID,
		defs.AttrExternalAPIStage: stage,
	}
	return &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Name: fmt.Sprintf("name-%s", id),
			Metadata: v1.Metadata{
				ID: id,
			},
			SubResources: map[string]interface{}{
				defs.XAgentDetails: sub,
			},
		},
	}
}

func createCategory(name, title string) *v1.ResourceInstance {
	return &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Name:  name,
			Title: title,
		},
	}
}

func createRI(id, name string) *v1.ResourceInstance {
	return &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Metadata: v1.Metadata{
				ID: id,
			},
			Name: name,
		},
	}
}

// add api service with externalAPIID, externalAPIName
// add api service with externalAPIPrimaryKey, externalAPIID, externalAPIName
// add existing api service with externalAPIID, externalAPIName
// get api service with APIID added by externalAPIID
// get api service with APIID added by externalAPIPrimaryKey
// get api service with Name added by externalAPIID
// get api service with Name added by externalAPIPrimaryKey
// get api service with invalid primary key
// get api service with primary key added by externalAPIPrimaryKey
// delete api service with APIID added by externalAPIID
// delete api service with APIID added by externalAPIPrimaryKey
func TestAPIServiceCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)

	apiCache := m.GetAPIServiceCache()
	assert.NotNil(t, apiCache)
	assert.Equal(t, apiCache.GetKeys(), m.GetAPIServiceKeys())

	api1 := createAPIService("id1", "api1", "")
	api2 := createAPIService("id2", "api2", "api2key")
	api1.Owner = &v1.Owner{Type: v1.TeamOwner, ID: "teamID1"}
	api2.Owner = &v1.Owner{Type: v1.TeamOwner, ID: "teamID2"}

	err := m.AddAPIService(api1)
	assert.Nil(t, err)

	err = m.AddAPIService(api2)
	assert.Nil(t, err)

	err = m.AddAPIService(api2)
	assert.Nil(t, err)

	teamIDs := m.GetTeamsIDsInAPIServices()
	assert.ElementsMatch(t, []string{"teamID1", "teamID2"}, teamIDs)

	cachedAPI := m.GetAPIServiceWithAPIID("id1")
	assert.Equal(t, api1, cachedAPI)

	cachedAPI = m.GetAPIServiceWithAPIID("id2")
	assert.Equal(t, api2, cachedAPI)

	cachedAPI = m.GetAPIServiceWithName("api1")
	assert.Equal(t, api1, cachedAPI)

	cachedAPI = m.GetAPIServiceWithName("api2")
	assert.Equal(t, api2, cachedAPI)

	cachedAPI = m.GetAPIServiceWithPrimaryKey("api1key")
	assert.Nil(t, cachedAPI)

	cachedAPI = m.GetAPIServiceWithPrimaryKey("api2key")
	assert.Equal(t, api2, cachedAPI)

	err = m.DeleteAPIService("api1")
	assert.Nil(t, err)
	cachedAPI = m.GetAPIServiceWithAPIID("api1")
	assert.Nil(t, cachedAPI)

	err = m.DeleteAPIService("api2")
	assert.Nil(t, err)
	cachedAPI = m.GetAPIServiceWithAPIID("api2")
	assert.Nil(t, cachedAPI)

	err = m.DeleteAPIService("api2")
	assert.NotNil(t, err)
}

// add api service instance
// get api service instance by ID
// get api service instance with invalid ID
// delete api service instance by ID
// delete all api service instance
func TestAPIServiceInstanceCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)
	assert.Equal(t, []string{}, m.GetAPIServiceInstanceKeys())

	api1 := createAPIService("apiID1", "api1", "api1key")
	instance1 := createAPIServiceInstance("id1", "apiID1", "stage1")
	instance2 := createAPIServiceInstance("id2", "apiID2", "stage2")

	m.AddAPIService(api1)
	m.AddAPIServiceInstance(instance1)
	m.AddAPIServiceInstance(instance2)
	assert.ElementsMatch(t, []string{"id1", "id2"}, m.GetAPIServiceInstanceKeys())

	allInstances := m.ListAPIServiceInstances()
	assert.ElementsMatch(t, []*v1.ResourceInstance{instance1, instance2}, allInstances)

	cachedInstance, err := m.GetAPIServiceInstanceByID("id1")
	assert.Nil(t, err)
	assert.Equal(t, instance1, cachedInstance)
	assert.Equal(t, 1, m.GetAPIServiceInstanceCount(api1.Name))

	cachedInstance, err = m.GetAPIServiceInstanceByName("name-id1")
	assert.Nil(t, err)
	assert.Equal(t, instance1, cachedInstance)
	assert.Equal(t, 1, m.GetAPIServiceInstanceCount(api1.Name))

	err = m.DeleteAPIServiceInstance("id1")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"id2"}, m.GetAPIServiceInstanceKeys())

	cachedInstance, err = m.GetAPIServiceInstanceByID("id1")
	assert.NotNil(t, err)
	assert.Nil(t, cachedInstance)
	assert.Equal(t, 0, m.GetAPIServiceInstanceCount(instance2.Name))

	cachedInstance, err = m.GetAPIServiceInstanceByName("name-id1")
	assert.NotNil(t, err)
	assert.Nil(t, cachedInstance)
	assert.Equal(t, 0, m.GetAPIServiceInstanceCount(instance2.Name))

	m.DeleteAllAPIServiceInstance()
	assert.ElementsMatch(t, []string{}, m.GetAPIServiceInstanceKeys())
}

// add category
// get category with name
// get category with title
// get category with invalid name
// delete category
func TestCategoryCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)

	categoryCache := m.GetCategoryCache()
	assert.NotNil(t, categoryCache)

	assert.Equal(t, []string{}, m.GetCategoryKeys())

	category1 := createCategory("c1", "category 1")
	category2 := createCategory("c2", "category 2")

	m.AddCategory(category1)
	assert.ElementsMatch(t, []string{"c1"}, m.GetCategoryKeys())
	m.AddCategory(category2)
	assert.ElementsMatch(t, []string{"c1", "c2"}, m.GetCategoryKeys())

	cachedCategory := m.GetCategory("c1")
	assert.Equal(t, category1, cachedCategory)

	cachedCategory = m.GetCategoryWithTitle("category 2")
	assert.Equal(t, category2, cachedCategory)

	err := m.DeleteCategory("c1")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"c2"}, m.GetCategoryKeys())

	cachedCategory = m.GetCategory("c1")
	assert.Nil(t, cachedCategory)

	err = m.DeleteCategory("c1")
	assert.NotNil(t, err)
}

// add sequence
// get sequence
// delete category
func TestSequenceCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)

	m.AddSequence("watch1", 1)
	assert.Equal(t, int64(1), m.GetSequence("watch1"))
	assert.Equal(t, int64(0), m.GetSequence("invalidwatch"))
	m.AddSequence("watch1", 2)
	assert.Equal(t, int64(2), m.GetSequence("watch1"))
}

// create manager
// add items to cache
// save cache
// create manager initialized with persisted cache
// validate all original cached items exists
func TestCachePersistenc(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{AgentName: "test", GRPCCfg: config.GRPCConfig{Enabled: true}}, true)
	assert.NotNil(t, m)

	api1 := createAPIService("id1", "apiID", "")
	err := m.AddAPIService(api1)
	assert.Nil(t, err)

	instance1 := createAPIServiceInstance("id1", "apiID", "stage")
	m.AddAPIServiceInstance(instance1)

	category1 := createCategory("c1", "category 1")
	m.AddCategory(category1)

	m.AddSequence("watch1", 1)

	defer func() {
		// Remove file if it exists
		_, err := os.Stat("./data")
		if !os.IsExist(err) {
			os.RemoveAll("./data")
		}
	}()

	m.SaveCache()

	m2 := NewAgentCacheManager(&config.CentralConfiguration{AgentName: "test", GRPCCfg: config.GRPCConfig{Enabled: true}}, true)

	persistedAPI := m2.GetAPIServiceWithAPIID("id1")
	assert.ElementsMatch(t, m.GetAPIServiceKeys(), m2.GetAPIServiceKeys())
	assertResourceInstance(t, api1, persistedAPI)
	assert.Equal(t, 1, m2.GetAPIServiceInstanceCount(api1.Name))

	persistedInstance, err := m2.GetAPIServiceInstanceByID("id1")
	assert.Nil(t, err)
	assert.ElementsMatch(t, m.GetAPIServiceInstanceKeys(), m2.GetAPIServiceInstanceKeys())
	assertResourceInstance(t, instance1, persistedInstance)

	persistedCategory := m2.GetCategory("c1")
	assert.ElementsMatch(t, m.GetCategoryKeys(), m2.GetCategoryKeys())
	assertResourceInstance(t, category1, persistedCategory)

	persistedSeq := m2.GetSequence("watch1")
	assert.Equal(t, int64(1), persistedSeq)
}

func assertResourceInstance(t *testing.T, expected *v1.ResourceInstance, actual *v1.ResourceInstance) {
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Group, actual.Group)
	assert.Equal(t, expected.Kind, actual.Kind)
	assert.Equal(t, expected.Metadata.ID, actual.Metadata.ID)
	assert.Equal(t, expected.Attributes, actual.Attributes)
	assert.Equal(t, expected.Spec, actual.Spec)
	assert.Equal(t, expected.SubResources, actual.SubResources)
}

func createRequestDefinition(name, id string) *v1.ResourceInstance {
	return &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Name: name,
			Metadata: v1.Metadata{
				ID: id,
			},
		},
	}
}

func TestAccessRequestDefinitionCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)

	ard1 := createRequestDefinition("name1", "id1")
	ard2 := createRequestDefinition("name2", "id2")

	m.AddAccessRequestDefinition(ard1)
	m.AddAccessRequestDefinition(ard2)

	cachedARD, err := m.GetAccessRequestDefinitionByName("name1")
	assert.Nil(t, err)
	assert.Equal(t, ard1, cachedARD)

	cachedARD, err = m.GetAccessRequestDefinitionByID("id1")
	assert.Nil(t, err)
	assert.Equal(t, ard1, cachedARD)

	err = m.DeleteAccessRequestDefinition("id1")
	assert.Nil(t, err)

	cachedARD, err = m.GetAccessRequestDefinitionByName("name1")
	assert.NotNil(t, err)
	assert.Nil(t, cachedARD)

	cachedARD, err = m.GetAccessRequestDefinitionByID("id1")
	assert.NotNil(t, err)
	assert.Nil(t, cachedARD)
}

func TestCredentialRequestDefinitionCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)

	crd1 := createRequestDefinition("name1", "id1")
	crd2 := createRequestDefinition("name2", "id2")

	m.AddCredentialRequestDefinition(crd1)
	m.AddCredentialRequestDefinition(crd2)

	crdKeys := m.GetCredentialRequestDefinitionKeys()
	assert.ElementsMatch(t, []string{"id1", "id2"}, crdKeys)

	cachedCRDs := m.ListCredentialRequestDefinitions()
	assert.ElementsMatch(t, []*v1.ResourceInstance{crd1, crd2}, cachedCRDs)

	cachedCRD, err := m.GetCredentialRequestDefinitionByName("name1")
	assert.Nil(t, err)
	assert.Equal(t, crd1, cachedCRD)

	cachedCRD, err = m.GetCredentialRequestDefinitionByID("id1")
	assert.Nil(t, err)
	assert.Equal(t, crd1, cachedCRD)

	err = m.DeleteCredentialRequestDefinition("id1")
	assert.Nil(t, err)

	cachedCRD, err = m.GetCredentialRequestDefinitionByName("name1")
	assert.NotNil(t, err)
	assert.Nil(t, cachedCRD)

	cachedCRD, err = m.GetCredentialRequestDefinitionByName("id1")
	assert.NotNil(t, err)
	assert.Nil(t, cachedCRD)
}

func TestFetchOnStartupCache(t *testing.T) {
	m := NewAgentCacheManager(&config.CentralConfiguration{}, false)
	assert.NotNil(t, m)

	res := m.GetAllFetchOnStartupResources()
	assert.Empty(t, res)

	ri1 := createRI("123", "foo")
	ri2 := createRI("456", "bar")
	ri3 := createRI("789", "baz")

	m.AddFetchOnStartupResources([]*v1.ResourceInstance{ri1})
	m.AddFetchOnStartupResources([]*v1.ResourceInstance{ri2, ri3})

	res = m.GetAllFetchOnStartupResources()
	assert.Len(t, res, 3)

	err := m.DeleteAllFetchOnStartupResources()
	assert.NoError(t, err)

	res = m.GetAllFetchOnStartupResources()
	assert.Empty(t, res)

}
