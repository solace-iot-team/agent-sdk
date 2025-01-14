package oauth

import (
	"fmt"
	"time"

	"github.com/Axway/agent-sdk/pkg/cache"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

const (
	issuerKeyPrefix  = "issuer:"
	tokenEpKeyPrefix = "tokenEp:"
	authEpKeyPrefix  = "authEp:"
)

// ProviderRegistry - interface for provider registry
type ProviderRegistry interface {
	// RegisterProvider - registers the provider using the config
	RegisterProvider(idp corecfg.IDPConfig, tlsCfg corecfg.TLSConfig, proxyURL string, clientTimeout time.Duration) error
	// GetProviderByName - returns the provider from registry based on the name
	GetProviderByName(name string) (Provider, error)
	// GetProviderByIssuer - returns the provider from registry based on the IDP issuer
	GetProviderByIssuer(issuer string) (Provider, error)
	// GetProviderByTokenEndpoint - returns the provider from registry based on the IDP token endpoint
	GetProviderByTokenEndpoint(tokenEndpoint string) (Provider, error)
	// GetProviderByAuthorizationEndpoint - returns the provider from registry based on the IDP authorization endpoint
	GetProviderByAuthorizationEndpoint(authEndpoint string) (Provider, error)
}

type providerRegistry struct {
	logger      log.FieldLogger
	providerMap cache.Cache
}

// NewProviderRegistry - create a new provider registry
func NewProviderRegistry() ProviderRegistry {
	logger := log.NewFieldLogger().
		WithComponent("providerRegistry").
		WithPackage("sdk.agent.authz.oauth")
	return &providerRegistry{
		logger:      logger,
		providerMap: cache.New(),
	}
}

// RegisterProvider - registers the provider using the config
func (r *providerRegistry) RegisterProvider(idp corecfg.IDPConfig, tlsCfg corecfg.TLSConfig, proxyURL string, clientTimeout time.Duration) error {
	p, err := NewProvider(idp, tlsCfg, proxyURL, clientTimeout)
	if err != nil {
		return err
	}

	name := p.GetName()
	issuer := p.GetIssuer()
	tokenEndpoint := p.GetTokenEndpoint()
	authEndPoint := p.GetAuthorizationEndpoint()

	r.logger.
		WithField("name", name).
		WithField("issuer", issuer).
		WithField("token-endpoint", tokenEndpoint).
		WithField("authorization-endpoint", authEndPoint).
		Debug("registered IDP provider")

	r.providerMap.Set(name, p)
	r.providerMap.SetSecondaryKey(name, issuerKeyPrefix+issuer)
	r.providerMap.SetSecondaryKey(name, tokenEpKeyPrefix+tokenEndpoint)
	r.providerMap.SetSecondaryKey(name, authEpKeyPrefix+authEndPoint)

	return nil
}

// GetProviderByName - returns the provider from registry based on the name
func (r *providerRegistry) GetProviderByName(name string) (Provider, error) {
	p, err := r.providerMap.Get(name)
	if err != nil {
		return nil, err
	}

	prov, ok := p.(Provider)
	if !ok {
		return nil, fmt.Errorf("unexpected provider entry for %s", name)
	}
	return prov, nil
}

// GetProviderByTokenEndpoint - returns the provider from registry based on the IDP token endpoint
func (r *providerRegistry) GetProviderByIssuer(issuer string) (Provider, error) {
	return r.getProviderBySecondaryKey(issuerKeyPrefix + issuer)
}

// GetProviderByTokenEndpoint - returns the provider from registry based on the IDP token endpoint
func (r *providerRegistry) GetProviderByTokenEndpoint(tokenEndpoint string) (Provider, error) {
	return r.getProviderBySecondaryKey(tokenEpKeyPrefix + tokenEndpoint)
}

// GetProviderByAuthorizationEndpoint - returns the provider from registry based on the IDP authorization endpoint
func (r *providerRegistry) GetProviderByAuthorizationEndpoint(authEndpoint string) (Provider, error) {
	return r.getProviderBySecondaryKey(authEpKeyPrefix + authEndpoint)
}

func (r *providerRegistry) getProviderBySecondaryKey(key string) (Provider, error) {
	p, err := r.providerMap.GetBySecondaryKey(key)
	if err != nil {
		return nil, err
	}

	prov, ok := p.(Provider)
	if !ok {
		return nil, fmt.Errorf("unexpected provider entry for %s", key)
	}
	return prov, nil
}
