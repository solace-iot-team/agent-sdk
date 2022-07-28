/*
 * This file is automatically generated
 */

package catalog

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1"
)

type ProductOverviewMergeFunc func(*m.ProductOverview, *m.ProductOverview) (*m.ProductOverview, error)

// ProductOverviewMerge builds a merge option for an update operation
func ProductOverviewMerge(f ProductOverviewMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.ProductOverview{}, &m.ProductOverview{}

		switch t := prev.(type) {
		case *m.ProductOverview:
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
		case *m.ProductOverview:
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

// ProductOverviewClient - rest client for ProductOverview resources that have a defined resource scope
type ProductOverviewClient struct {
	client v1.Scoped
}

// UnscopedProductOverviewClient - rest client for ProductOverview resources that do not have a defined scope
type UnscopedProductOverviewClient struct {
	client v1.Unscoped
}

// NewProductOverviewClient - creates a client that is not scoped to any resource
func NewProductOverviewClient(c v1.Base) (*UnscopedProductOverviewClient, error) {

	client, err := c.ForKind(m.ProductOverviewGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedProductOverviewClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedProductOverviewClient) WithScope(scope string) *ProductOverviewClient {
	return &ProductOverviewClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedProductOverviewClient) Get(name string) (*m.ProductOverview, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.ProductOverview{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedProductOverviewClient) Update(res *m.ProductOverview, opts ...v1.UpdateOption) (*m.ProductOverview, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.ProductOverview{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *ProductOverviewClient) List(options ...v1.ListOptions) ([]*m.ProductOverview, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.ProductOverview, len(riList))

	for i := range riList {
		result[i] = &m.ProductOverview{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *ProductOverviewClient) Get(name string) (*m.ProductOverview, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.ProductOverview{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *ProductOverviewClient) Delete(res *m.ProductOverview) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *ProductOverviewClient) Create(res *m.ProductOverview, opts ...v1.CreateOption) (*m.ProductOverview, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.ProductOverview{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *ProductOverviewClient) Update(res *m.ProductOverview, opts ...v1.UpdateOption) (*m.ProductOverview, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.ProductOverview{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
