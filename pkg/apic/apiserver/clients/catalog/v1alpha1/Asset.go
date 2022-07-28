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

type AssetMergeFunc func(*m.Asset, *m.Asset) (*m.Asset, error)

// AssetMerge builds a merge option for an update operation
func AssetMerge(f AssetMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Asset{}, &m.Asset{}

		switch t := prev.(type) {
		case *m.Asset:
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
		case *m.Asset:
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

// AssetClient - rest client for Asset resources that have a defined resource scope
type AssetClient struct {
	client v1.Scoped
}

// NewAssetClient - creates a client scoped to a particular resource
func NewAssetClient(c v1.Base) (*AssetClient, error) {

	client, err := c.ForKind(m.AssetGVK())
	if err != nil {
		return nil, err
	}

	return &AssetClient{client}, nil

}

// List - gets a list of resources
func (c *AssetClient) List(options ...v1.ListOptions) ([]*m.Asset, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Asset, len(riList))

	for i := range riList {
		result[i] = &m.Asset{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *AssetClient) Get(name string) (*m.Asset, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Asset{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *AssetClient) Delete(res *m.Asset) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *AssetClient) Create(res *m.Asset, opts ...v1.CreateOption) (*m.Asset, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Asset{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *AssetClient) Update(res *m.Asset, opts ...v1.UpdateOption) (*m.Asset, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Asset{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
