/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type AssetRequestDefinitionMergeFunc func(*m.AssetRequestDefinition, *m.AssetRequestDefinition) (*m.AssetRequestDefinition, error)

// AssetRequestDefinitionMerge builds a merge option for an update operation
func AssetRequestDefinitionMerge(f AssetRequestDefinitionMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.AssetRequestDefinition{}, &m.AssetRequestDefinition{}

		switch t := prev.(type) {
		case *m.AssetRequestDefinition:
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
		case *m.AssetRequestDefinition:
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

// AssetRequestDefinitionClient - rest client for AssetRequestDefinition resources that have a defined resource scope
type AssetRequestDefinitionClient struct {
	client v1.Scoped
}

// UnscopedAssetRequestDefinitionClient - rest client for AssetRequestDefinition resources that do not have a defined scope
type UnscopedAssetRequestDefinitionClient struct {
	client v1.Unscoped
}

// NewAssetRequestDefinitionClient - creates a client that is not scoped to any resource
func NewAssetRequestDefinitionClient(c v1.Base) (*UnscopedAssetRequestDefinitionClient, error) {

	client, err := c.ForKind(m.AssetRequestDefinitionGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAssetRequestDefinitionClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedAssetRequestDefinitionClient) WithScope(scope string) *AssetRequestDefinitionClient {
	return &AssetRequestDefinitionClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedAssetRequestDefinitionClient) Get(name string) (*m.AssetRequestDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.AssetRequestDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedAssetRequestDefinitionClient) Update(res *m.AssetRequestDefinition, opts ...v1.UpdateOption) (*m.AssetRequestDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.AssetRequestDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *AssetRequestDefinitionClient) List(options ...v1.ListOptions) ([]*m.AssetRequestDefinition, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.AssetRequestDefinition, len(riList))

	for i := range riList {
		result[i] = &m.AssetRequestDefinition{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *AssetRequestDefinitionClient) Get(name string) (*m.AssetRequestDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.AssetRequestDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *AssetRequestDefinitionClient) Delete(res *m.AssetRequestDefinition) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *AssetRequestDefinitionClient) Create(res *m.AssetRequestDefinition, opts ...v1.CreateOption) (*m.AssetRequestDefinition, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.AssetRequestDefinition{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *AssetRequestDefinitionClient) Update(res *m.AssetRequestDefinition, opts ...v1.UpdateOption) (*m.AssetRequestDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.AssetRequestDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
