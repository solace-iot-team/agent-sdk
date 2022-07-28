/*
 * This file is automatically generated
 */

package catalog

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type CredentialRequestDefinitionMergeFunc func(*m.CredentialRequestDefinition, *m.CredentialRequestDefinition) (*m.CredentialRequestDefinition, error)

// CredentialRequestDefinitionMerge builds a merge option for an update operation
func CredentialRequestDefinitionMerge(f CredentialRequestDefinitionMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.CredentialRequestDefinition{}, &m.CredentialRequestDefinition{}

		switch t := prev.(type) {
		case *m.CredentialRequestDefinition:
			p = t
		case *apiv1.ResourceInstance:
			err := p.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialise prev resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise prev resource, unxexpected resource type: %T", t)
		}

		switch t := new.(type) {
		case *m.CredentialRequestDefinition:
			n = t
		case *apiv1.ResourceInstance:
			err := n.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialize new resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise new resource, unxexpected resource type: %T", t)
		}

		return f(p, n)
	})
}

// CredentialRequestDefinitionClient - rest client for CredentialRequestDefinition resources that have a defined resource scope
type CredentialRequestDefinitionClient struct {
	client v1.Scoped
}

// UnscopedCredentialRequestDefinitionClient - rest client for CredentialRequestDefinition resources that do not have a defined scope
type UnscopedCredentialRequestDefinitionClient struct {
	client v1.Unscoped
}

// NewCredentialRequestDefinitionClient - creates a client that is not scoped to any resource
func NewCredentialRequestDefinitionClient(c v1.Base) (*UnscopedCredentialRequestDefinitionClient, error) {

	client, err := c.ForKind(m.CredentialRequestDefinitionGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedCredentialRequestDefinitionClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedCredentialRequestDefinitionClient) WithScope(scope string) *CredentialRequestDefinitionClient {
	return &CredentialRequestDefinitionClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedCredentialRequestDefinitionClient) Get(name string) (*m.CredentialRequestDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.CredentialRequestDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedCredentialRequestDefinitionClient) Update(res *m.CredentialRequestDefinition, opts ...v1.UpdateOption) (*m.CredentialRequestDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.CredentialRequestDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *CredentialRequestDefinitionClient) List(options ...v1.ListOptions) ([]*m.CredentialRequestDefinition, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.CredentialRequestDefinition, len(riList))

	for i := range riList {
		result[i] = &m.CredentialRequestDefinition{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *CredentialRequestDefinitionClient) Get(name string) (*m.CredentialRequestDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.CredentialRequestDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *CredentialRequestDefinitionClient) Delete(res *m.CredentialRequestDefinition) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *CredentialRequestDefinitionClient) Create(res *m.CredentialRequestDefinition, opts ...v1.CreateOption) (*m.CredentialRequestDefinition, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.CredentialRequestDefinition{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *CredentialRequestDefinitionClient) Update(res *m.CredentialRequestDefinition, opts ...v1.UpdateOption) (*m.CredentialRequestDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.CredentialRequestDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
