/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

type VirtualServiceMergeFunc func(*m.VirtualService, *m.VirtualService) (*m.VirtualService, error)

// VirtualServiceMerge builds a merge option for an update operation
func VirtualServiceMerge(f VirtualServiceMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.VirtualService{}, &m.VirtualService{}

		switch t := prev.(type) {
		case *m.VirtualService:
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
		case *m.VirtualService:
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

// VirtualServiceClient - rest client for VirtualService resources that have a defined resource scope
type VirtualServiceClient struct {
	client v1.Scoped
}

// UnscopedVirtualServiceClient - rest client for VirtualService resources that do not have a defined scope
type UnscopedVirtualServiceClient struct {
	client v1.Unscoped
}

// NewVirtualServiceClient - creates a client that is not scoped to any resource
func NewVirtualServiceClient(c v1.Base) (*UnscopedVirtualServiceClient, error) {

	client, err := c.ForKind(m.VirtualServiceGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedVirtualServiceClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedVirtualServiceClient) WithScope(scope string) *VirtualServiceClient {
	return &VirtualServiceClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedVirtualServiceClient) Get(name string) (*m.VirtualService, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.VirtualService{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedVirtualServiceClient) Update(res *m.VirtualService, opts ...v1.UpdateOption) (*m.VirtualService, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.VirtualService{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *VirtualServiceClient) List(options ...v1.ListOptions) ([]*m.VirtualService, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.VirtualService, len(riList))

	for i := range riList {
		result[i] = &m.VirtualService{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *VirtualServiceClient) Get(name string) (*m.VirtualService, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.VirtualService{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *VirtualServiceClient) Delete(res *m.VirtualService) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *VirtualServiceClient) Create(res *m.VirtualService, opts ...v1.CreateOption) (*m.VirtualService, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.VirtualService{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *VirtualServiceClient) Update(res *m.VirtualService, opts ...v1.UpdateOption) (*m.VirtualService, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.VirtualService{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
