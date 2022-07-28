/*
 * This file is automatically generated
 */

package definitions

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/definitions/v1alpha1"
)

type AccessControlListMergeFunc func(*m.AccessControlList, *m.AccessControlList) (*m.AccessControlList, error)

// AccessControlListMerge builds a merge option for an update operation
func AccessControlListMerge(f AccessControlListMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.AccessControlList{}, &m.AccessControlList{}

		switch t := prev.(type) {
		case *m.AccessControlList:
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
		case *m.AccessControlList:
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

// AccessControlListClient - rest client for AccessControlList resources that have a defined resource scope
type AccessControlListClient struct {
	client v1.Scoped
}

// UnscopedAccessControlListClient - rest client for AccessControlList resources that do not have a defined scope
type UnscopedAccessControlListClient struct {
	client v1.Unscoped
}

// NewAccessControlListClient - creates a client that is not scoped to any resource
func NewAccessControlListClient(c v1.Base) (*UnscopedAccessControlListClient, error) {

	client, err := c.ForKind(m.AccessControlListGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAccessControlListClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedAccessControlListClient) WithScope(scope string) *AccessControlListClient {
	return &AccessControlListClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedAccessControlListClient) Get(name string) (*m.AccessControlList, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.AccessControlList{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedAccessControlListClient) Update(res *m.AccessControlList, opts ...v1.UpdateOption) (*m.AccessControlList, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.AccessControlList{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *AccessControlListClient) List(options ...v1.ListOptions) ([]*m.AccessControlList, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.AccessControlList, len(riList))

	for i := range riList {
		result[i] = &m.AccessControlList{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *AccessControlListClient) Get(name string) (*m.AccessControlList, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.AccessControlList{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *AccessControlListClient) Delete(res *m.AccessControlList) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *AccessControlListClient) Create(res *m.AccessControlList, opts ...v1.CreateOption) (*m.AccessControlList, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.AccessControlList{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *AccessControlListClient) Update(res *m.AccessControlList, opts ...v1.UpdateOption) (*m.AccessControlList, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.AccessControlList{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
