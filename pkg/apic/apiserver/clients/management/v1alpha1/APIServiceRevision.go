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

type APIServiceRevisionMergeFunc func(*m.APIServiceRevision, *m.APIServiceRevision) (*m.APIServiceRevision, error)

// APIServiceRevisionMerge builds a merge option for an update operation
func APIServiceRevisionMerge(f APIServiceRevisionMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.APIServiceRevision{}, &m.APIServiceRevision{}

		switch t := prev.(type) {
		case *m.APIServiceRevision:
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
		case *m.APIServiceRevision:
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

// APIServiceRevisionClient - rest client for APIServiceRevision resources that have a defined resource scope
type APIServiceRevisionClient struct {
	client v1.Scoped
}

// UnscopedAPIServiceRevisionClient - rest client for APIServiceRevision resources that do not have a defined scope
type UnscopedAPIServiceRevisionClient struct {
	client v1.Unscoped
}

// NewAPIServiceRevisionClient - creates a client that is not scoped to any resource
func NewAPIServiceRevisionClient(c v1.Base) (*UnscopedAPIServiceRevisionClient, error) {

	client, err := c.ForKind(m.APIServiceRevisionGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAPIServiceRevisionClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedAPIServiceRevisionClient) WithScope(scope string) *APIServiceRevisionClient {
	return &APIServiceRevisionClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedAPIServiceRevisionClient) Get(name string) (*m.APIServiceRevision, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.APIServiceRevision{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedAPIServiceRevisionClient) Update(res *m.APIServiceRevision, opts ...v1.UpdateOption) (*m.APIServiceRevision, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.APIServiceRevision{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *APIServiceRevisionClient) List(options ...v1.ListOptions) ([]*m.APIServiceRevision, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.APIServiceRevision, len(riList))

	for i := range riList {
		result[i] = &m.APIServiceRevision{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *APIServiceRevisionClient) Get(name string) (*m.APIServiceRevision, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.APIServiceRevision{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *APIServiceRevisionClient) Delete(res *m.APIServiceRevision) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *APIServiceRevisionClient) Create(res *m.APIServiceRevision, opts ...v1.CreateOption) (*m.APIServiceRevision, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.APIServiceRevision{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *APIServiceRevisionClient) Update(res *m.APIServiceRevision, opts ...v1.UpdateOption) (*m.APIServiceRevision, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.APIServiceRevision{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
