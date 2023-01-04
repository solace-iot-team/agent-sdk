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

type ConsumerProductVisibilityMergeFunc func(*m.ConsumerProductVisibility, *m.ConsumerProductVisibility) (*m.ConsumerProductVisibility, error)

// ConsumerProductVisibilityMerge builds a merge option for an update operation
func ConsumerProductVisibilityMerge(f ConsumerProductVisibilityMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.ConsumerProductVisibility{}, &m.ConsumerProductVisibility{}

		switch t := prev.(type) {
		case *m.ConsumerProductVisibility:
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
		case *m.ConsumerProductVisibility:
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

// ConsumerProductVisibilityClient - rest client for ConsumerProductVisibility resources that have a defined resource scope
type ConsumerProductVisibilityClient struct {
	client v1.Scoped
}

// UnscopedConsumerProductVisibilityClient - rest client for ConsumerProductVisibility resources that do not have a defined scope
type UnscopedConsumerProductVisibilityClient struct {
	client v1.Unscoped
}

// NewConsumerProductVisibilityClient - creates a client that is not scoped to any resource
func NewConsumerProductVisibilityClient(c v1.Base) (*UnscopedConsumerProductVisibilityClient, error) {

	client, err := c.ForKind(m.ConsumerProductVisibilityGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedConsumerProductVisibilityClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedConsumerProductVisibilityClient) WithScope(scope string) *ConsumerProductVisibilityClient {
	return &ConsumerProductVisibilityClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedConsumerProductVisibilityClient) Get(name string) (*m.ConsumerProductVisibility, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.ConsumerProductVisibility{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedConsumerProductVisibilityClient) Update(res *m.ConsumerProductVisibility, opts ...v1.UpdateOption) (*m.ConsumerProductVisibility, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.ConsumerProductVisibility{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *ConsumerProductVisibilityClient) List(options ...v1.ListOptions) ([]*m.ConsumerProductVisibility, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.ConsumerProductVisibility, len(riList))

	for i := range riList {
		result[i] = &m.ConsumerProductVisibility{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *ConsumerProductVisibilityClient) Get(name string) (*m.ConsumerProductVisibility, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.ConsumerProductVisibility{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *ConsumerProductVisibilityClient) Delete(res *m.ConsumerProductVisibility) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *ConsumerProductVisibilityClient) Create(res *m.ConsumerProductVisibility, opts ...v1.CreateOption) (*m.ConsumerProductVisibility, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.ConsumerProductVisibility{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *ConsumerProductVisibilityClient) Update(res *m.ConsumerProductVisibility, opts ...v1.UpdateOption) (*m.ConsumerProductVisibility, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.ConsumerProductVisibility{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}