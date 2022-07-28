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

type DocumentMergeFunc func(*m.Document, *m.Document) (*m.Document, error)

// DocumentMerge builds a merge option for an update operation
func DocumentMerge(f DocumentMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Document{}, &m.Document{}

		switch t := prev.(type) {
		case *m.Document:
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
		case *m.Document:
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

// DocumentClient - rest client for Document resources that have a defined resource scope
type DocumentClient struct {
	client v1.Scoped
}

// UnscopedDocumentClient - rest client for Document resources that do not have a defined scope
type UnscopedDocumentClient struct {
	client v1.Unscoped
}

// NewDocumentClient - creates a client that is not scoped to any resource
func NewDocumentClient(c v1.Base) (*UnscopedDocumentClient, error) {

	client, err := c.ForKind(m.DocumentGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedDocumentClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedDocumentClient) WithScope(scope string) *DocumentClient {
	return &DocumentClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedDocumentClient) Get(name string) (*m.Document, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Document{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedDocumentClient) Update(res *m.Document, opts ...v1.UpdateOption) (*m.Document, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Document{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *DocumentClient) List(options ...v1.ListOptions) ([]*m.Document, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Document, len(riList))

	for i := range riList {
		result[i] = &m.Document{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *DocumentClient) Get(name string) (*m.Document, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Document{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *DocumentClient) Delete(res *m.Document) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *DocumentClient) Create(res *m.Document, opts ...v1.CreateOption) (*m.Document, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Document{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *DocumentClient) Update(res *m.Document, opts ...v1.UpdateOption) (*m.Document, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Document{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
