package handler

import (
	"context"

	agentcache "github.com/Axway/agent-sdk/pkg/agent/cache"
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	environment "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	management "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
)

type environmentHandler struct {
	agentCacheManager agentcache.Manager
	credentialConfig  config.CredentialConfig
}

// NewEnvironmentHandler creates a Handler for Environments.
func NewEnvironmentHandler(agentCacheManager agentcache.Manager, credentialConfig config.CredentialConfig) Handler {
	return &environmentHandler{
		agentCacheManager: agentCacheManager,
		credentialConfig:  credentialConfig,
	}
}

func (c *environmentHandler) Handle(ctx context.Context, meta *proto.EventMeta, resource *v1.ResourceInstance) error {
	if resource.Kind != environment.EnvironmentGVK().Kind {
		return nil
	}

	// verify that action is subresource updated and meta subsresource is environment policy
	action := GetActionFromContext(ctx)
	if action != proto.Event_SUBRESOURCEUPDATED || meta.Subresource != management.EnvironmentPoliciesSubResourceName {
		return nil
	}

	log := getLoggerFromContext(ctx).WithComponent("environmentHandler")
	env := &management.Environment{}
	err := env.FromInstance(resource)
	if err != nil {
		log.WithError(err).Error("could not handle access request")
		return nil
	}

	// Set up credential config from environment resource policies
	c.credentialConfig.SetShouldDeprovisionExpired(env.Policies.Credentials.Expiry.Action == "deprovision")
	c.credentialConfig.SetExpirationDays(int(env.Policies.Credentials.Expiry.Period))

	return nil
}
