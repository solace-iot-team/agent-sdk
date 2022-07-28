/*
 * This file is automatically generated
 */

package management

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

type AssetMappingMergeFunc func(*m.AssetMapping, *m.AssetMapping) (*m.AssetMapping, error)

// AssetMappingMerge builds a merge option for an update operation
func AssetMappingMerge(f AssetMappingMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.AssetMapping{}, &m.AssetMapping{}

		switch t := prev.(type) {
		case *m.AssetMapping:
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
		case *m.AssetMapping:
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

// AssetMappingClient - rest client for AssetMapping resources that have a defined resource scope
type AssetMappingClient struct {
	client v1.Scoped
}

// UnscopedAssetMappingClient - rest client for AssetMapping resources that do not have a defined scope
type UnscopedAssetMappingClient struct {
	client v1.Unscoped
}

// NewAssetMappingClient - creates a client that is not scoped to any resource
func NewAssetMappingClient(c v1.Base) (*UnscopedAssetMappingClient, error) {

	client, err := c.ForKind(m.AssetMappingGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAssetMappingClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedAssetMappingClient) WithScope(scope string) *AssetMappingClient {
	return &AssetMappingClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedAssetMappingClient) Get(name string) (*m.AssetMapping, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.AssetMapping{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedAssetMappingClient) Update(res *m.AssetMapping, opts ...v1.UpdateOption) (*m.AssetMapping, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.AssetMapping{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *AssetMappingClient) List(options ...v1.ListOptions) ([]*m.AssetMapping, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.AssetMapping, len(riList))

	for i := range riList {
		result[i] = &m.AssetMapping{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *AssetMappingClient) Get(name string) (*m.AssetMapping, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.AssetMapping{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *AssetMappingClient) Delete(res *m.AssetMapping) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *AssetMappingClient) Create(res *m.AssetMapping, opts ...v1.CreateOption) (*m.AssetMapping, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.AssetMapping{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *AssetMappingClient) Update(res *m.AssetMapping, opts ...v1.UpdateOption) (*m.AssetMapping, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.AssetMapping{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
