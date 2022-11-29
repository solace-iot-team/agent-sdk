package agent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/Axway/agent-sdk/pkg/apic/definitions"
	"github.com/Axway/agent-sdk/pkg/apic/mock"
	"github.com/Axway/agent-sdk/pkg/util"

	"github.com/Axway/agent-sdk/pkg/util/log"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	management "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/stretchr/testify/assert"
)

func resetResources() {
	agent.agentResourceManager = nil
	if agent.cacheManager != nil {
		agent.cacheManager.ApplyResourceReadLock()
		defer agent.cacheManager.ReleaseResourceReadLock()
		agent.cacheManager = nil
	}
	agent.isInitialized = false
	agent.apicClient = nil
	agent.agentFeaturesCfg = nil
}

func createCentralCfg(url, env string) *config.CentralConfiguration {
	cfg := config.NewCentralConfig(config.DiscoveryAgent).(*config.CentralConfiguration)
	cfg.URL = url
	cfg.SingleURL = url
	cfg.TenantID = "123456"
	cfg.Environment = env
	cfg.APICDeployment = "apic"
	authCfg := cfg.Auth.(*config.AuthConfiguration)
	authCfg.URL = url + "/auth"
	authCfg.Realm = "Broker"
	authCfg.ClientID = "serviceaccount_1111"
	authCfg.PrivateKey = "../transaction/testdata/private_key.pem"
	authCfg.PublicKey = "../transaction/testdata/public_key"
	return cfg
}

func createOfflineCentralCfg(url, env string) *config.CentralConfiguration {
	cfg := config.NewCentralConfig(config.TraceabilityAgent).(*config.CentralConfiguration)
	cfg.EnvironmentID = "abc123"
	cfg.UsageReporting.(*config.UsageReportingConfiguration).Offline = true
	return cfg
}

func createDiscoveryAgentRes(id, name, dataplane, filter string) *v1.ResourceInstance {
	res := &management.DiscoveryAgent{
		ResourceMeta: v1.ResourceMeta{
			Name: name,
			Metadata: v1.Metadata{
				ID: id,
			},
		},
		Spec: management.DiscoveryAgentSpec{
			DataplaneType: dataplane,
			Config: management.DiscoveryAgentSpecConfig{
				Filter: filter,
			},
		},
	}
	instance, _ := res.AsInstance()
	return instance
}

func createTraceabilityAgentRes(id, name, dataplane string, processHeaders bool) *v1.ResourceInstance {
	res := &management.TraceabilityAgent{
		ResourceMeta: v1.ResourceMeta{
			Name: name,
			Metadata: v1.Metadata{
				ID: id,
			},
		},
		Spec: management.TraceabilityAgentSpec{
			DataplaneType: dataplane,
			Config: management.TraceabilityAgentSpecConfig{
				ProcessHeaders: processHeaders,
			},
		},
	}
	instance, _ := res.AsInstance()
	return instance
}

type TestConfig struct {
	resourceChanged bool
}

func (a *TestConfig) ApplyResources(agentResource *v1.ResourceInstance) error {
	a.resourceChanged = true
	return nil
}

func TestAgentInitialize(t *testing.T) {
	const (
		daName = "discovery"
		taName = "traceability"
	)

	teams := []definitions.PlatformTeam{
		{
			ID:      "123",
			Name:    "name",
			Default: true,
		},
	}
	environmentRes := &management.Environment{
		ResourceMeta: v1.ResourceMeta{
			Metadata: v1.Metadata{ID: "123"},
			Name:     "v7",
			Title:    "v7",
		},
	}
	discoveryAgentRes := createDiscoveryAgentRes("111", daName, "v7-dataplane", "")
	traceabilityAgentRes := createTraceabilityAgentRes("111", taName, "v7-dataplane", false)

	s := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.RequestURI, "/auth") {
			token := "{\"access_token\":\"somevalue\",\"expires_in\": 12235677}"
			resp.Write([]byte(token))
			return
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7/discoveryagents/"+daName) {
			buf, err := json.Marshal(discoveryAgentRes)
			log.Error(err)
			resp.Write(buf)
			return
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7/traceabilityagents/"+taName) {
			buf, err := json.Marshal(traceabilityAgentRes)
			log.Error(err)
			resp.Write(buf)
			return
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7") {
			buf, _ := json.Marshal(environmentRes)
			resp.Write(buf)
			return
		}

		if strings.Contains(req.RequestURI, "/api/v1/platformTeams") {
			buf, _ := json.Marshal(teams)
			resp.Write(buf)
			return
		}
	}))

	defer s.Close()

	cfg := createOfflineCentralCfg(s.URL, "v7")
	// Test with offline mode
	resetResources()
	err := Initialize(cfg)
	assert.Nil(t, err)
	da := GetAgentResource()
	assert.Nil(t, da)

	cfg = createCentralCfg(s.URL, "v7")
	// Test with no agent name - config to be validate successfully as no calls made to get agent and dataplane resource
	resetResources()
	err = Initialize(cfg)
	assert.Nil(t, err)
	da = GetAgentResource()
	assert.Nil(t, da)

	cfg.AgentType = config.DiscoveryAgent
	AgentResourceType = management.DiscoveryAgentResourceName
	cfg.AgentName = daName
	resetResources()
	err = Initialize(cfg)
	assert.Nil(t, err)

	da = GetAgentResource()
	assertResource(t, da, discoveryAgentRes)

	cfg.AgentType = config.TraceabilityAgent
	AgentResourceType = management.TraceabilityAgentResourceName
	cfg.AgentName = taName
	resetResources()
	err = Initialize(cfg)
	assert.Nil(t, err)

	da = GetAgentResource()
	assertResource(t, da, traceabilityAgentRes)

	agentCfg := &TestConfig{
		resourceChanged: false,
	}

	ApplyResourceToConfig(agentCfg)

	assert.True(t, agentCfg.resourceChanged)

	// Test for resource change
	traceabilityAgentRes = createTraceabilityAgentRes("111", taName, "v7-dataplane", true)
	resetResources()

	agentResChangeHandlerCall := 0
	OnAgentResourceChange(func() { agentResChangeHandlerCall++ })

	err = Initialize(cfg)
	assert.Nil(t, err)

	da = GetAgentResource()
	assertResource(t, da, traceabilityAgentRes)
	assert.Equal(t, 0, agentResChangeHandlerCall)
}

func TestInitEnvironment(t *testing.T) {
	teams := []definitions.PlatformTeam{
		{
			ID:      "123",
			Name:    "name",
			Default: true,
		},
	}
	environmentRes := management.NewEnvironment("v7")
	environmentRes.Title = "v7"
	environmentRes.Metadata.ID = "123"

	s := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.RequestURI, "/auth") {
			token := "{\"access_token\":\"somevalue\",\"expires_in\": 12235677}"
			resp.Write([]byte(token))
			return
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7") {
			if req.Method == "GET" {
				buf, _ := json.Marshal(environmentRes)
				resp.Write(buf)
			} else if req.Method == "PUT" {
				subRes := &management.Environment{}
				json.NewDecoder(req.Body).Decode(subRes)
				environmentRes.ResourceMeta.SubResources = subRes.ResourceMeta.SubResources
			}
			return
		}

		if strings.Contains(req.RequestURI, "/api/v1/platformTeams") {
			buf, _ := json.Marshal(teams)
			resp.Write(buf)
			return
		}
	}))

	defer s.Close()

	cfg := createCentralCfg(s.URL, "v7")
	cfg.AgentType = config.GenericService
	agent.cfg = cfg
	initializeTokenRequester(agent.cfg)
	apiClient := apic.New(agent.cfg, agent.tokenRequester, agent.cacheManager)
	// Test with no agent name - config to be validate successfully as no calls made to get agent and dataplane resource

	defer resetResources()
	err := initEnvResources(agent.cfg, apiClient)
	assert.Nil(t, err)
	xAgentDetail := util.GetAgentDetails(environmentRes)
	assert.Nil(t, xAgentDetail)

	cfg = createCentralCfg(s.URL, "v7")
	cfg.AgentType = config.DiscoveryAgent
	agent.cfg = cfg
	err = initEnvResources(agent.cfg, apiClient)
	assert.Nil(t, err)
	xAgentDetail = util.GetAgentDetails(environmentRes)
	assert.NotNil(t, xAgentDetail)
	assert.Equal(t, "true", xAgentDetail[config.DiscoveryAgent.ToString()+"-enabled"])

	cfg = createCentralCfg(s.URL, "v7")
	cfg.AgentType = config.TraceabilityAgent
	agent.cfg = cfg
	err = initEnvResources(agent.cfg, apiClient)
	assert.Nil(t, err)
	xAgentDetail = util.GetAgentDetails(environmentRes)
	assert.NotNil(t, xAgentDetail)
	assert.Equal(t, "true", xAgentDetail[config.DiscoveryAgent.ToString()+"-enabled"])
	assert.Equal(t, "true", xAgentDetail[config.TraceabilityAgent.ToString()+"-enabled"])
}

func TestAgentConfigOverride(t *testing.T) {
	const (
		daName = "discovery"
		taName = "traceability"
	)

	teams := []definitions.PlatformTeam{
		{
			ID:      "123",
			Name:    "name",
			Default: true,
		},
	}
	environmentRes := &management.Environment{
		ResourceMeta: v1.ResourceMeta{
			Metadata: v1.Metadata{ID: "123"},
			Name:     "v7",
			Title:    "v7",
		},
	}
	discoveryAgentRes := createDiscoveryAgentRes("111", daName, "v7-dataplane", "")
	traceabilityAgentRes := createTraceabilityAgentRes("111", taName, "v7-dataplane", false)

	s := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.RequestURI, "/auth") {
			token := "{\"access_token\":\"somevalue\",\"expires_in\": 12235677}"
			resp.Write([]byte(token))
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7/discoveryagents/"+daName) {
			buf, _ := json.Marshal(discoveryAgentRes)
			resp.Write(buf)
			return
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7/traceabilityagents/"+taName) {
			buf, _ := json.Marshal(traceabilityAgentRes)
			resp.Write(buf)
			return
		}

		if strings.Contains(req.RequestURI, "/apis/management/v1alpha1/environments/v7") {
			buf, _ := json.Marshal(environmentRes)
			resp.Write(buf)
			return
		}

		if strings.Contains(req.RequestURI, "/api/v1/platformTeams") {
			buf, _ := json.Marshal(teams)
			resp.Write(buf)
			return
		}
	}))

	defer s.Close()

	cfg := createCentralCfg(s.URL, "v7")

	AgentResourceType = management.DiscoveryAgentResourceName
	cfg.AgentName = "discovery"
	resetResources()
	err := Initialize(cfg)
	assert.Nil(t, err)

	da := GetAgentResource()
	assertResource(t, da, discoveryAgentRes)

}

func TestAgentAgentFeaturesOnByDefault(t *testing.T) {
	cfg := createCentralCfg("http://test", "v7")
	resetResources()
	err := Initialize(cfg)
	assert.NoError(t, err)

	// Assert the agent features are on by default
	assert.True(t, agent.agentFeaturesCfg.ConnectionToCentralEnabled())
	assert.True(t, agent.agentFeaturesCfg.ProcessSystemSignalsEnabled())
	assert.True(t, agent.agentFeaturesCfg.VersionCheckerEnabled())

	assert.NotNil(t, agent.apicClient)
}

func TestAgentAgentFeaturesDisabled(t *testing.T) {
	// Create invalid Central config
	cfg := config.NewCentralConfig(config.GenericService).(*config.CentralConfiguration)
	resetResources()
	agentFeatures := &config.AgentFeaturesConfiguration{
		ConnectToCentral:     false,
		ProcessSystemSignals: false,
		VersionChecker:       false,
	}
	err := InitializeWithAgentFeatures(cfg, agentFeatures)
	assert.NoError(t, err) // This asserts central config is not being validated as ConnectToCentral is false

	assert.False(t, agent.agentFeaturesCfg.ConnectionToCentralEnabled())
	assert.False(t, agent.agentFeaturesCfg.ProcessSystemSignalsEnabled())
	assert.False(t, agent.agentFeaturesCfg.VersionCheckerEnabled())

	// Assert no api client
	assert.Nil(t, agent.apicClient)
}

func Test_registerSubscriptionWebhook(t *testing.T) {
	err := registerSubscriptionWebhook(config.DiscoveryAgent, &mock.Client{})
	assert.Nil(t, err)

	err = registerSubscriptionWebhook(config.DiscoveryAgent, &mock.Client{
		RegisterSubscriptionWebhookMock: func() error {
			return fmt.Errorf("error")
		},
	})
	assert.NotNil(t, err)

	err = registerSubscriptionWebhook(config.TraceabilityAgent, &mock.Client{})
	assert.Nil(t, err)
}

func assertResource(t *testing.T, res, expectedRes *v1.ResourceInstance) {
	assert.Equal(t, expectedRes.Group, res.Group)
	assert.Equal(t, expectedRes.Kind, res.Kind)
	assert.Equal(t, expectedRes.Name, res.Name)
	assert.Equal(t, expectedRes.Metadata.ID, res.Metadata.ID)
	assert.Equal(t, expectedRes.Spec, res.Spec)
}
