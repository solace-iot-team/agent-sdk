/*
 * This file is automatically generated
 */

package v1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1"
)

type CategoryMergeFunc func(*m.Category, *m.Category) (*m.Category, error)

// CategoryMerge builds a merge option for an update operation
func CategoryMerge(f CategoryMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Category{}, &m.Category{}

		switch t := prev.(type) {
		case *m.Category:
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
		case *m.Category:
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

// CategoryClient - rest client for Category resources that have a defined resource scope
type CategoryClient struct {
	client v1.Scoped
}

// NewCategoryClient - creates a client scoped to a particular resource
func NewCategoryClient(c v1.Base) (*CategoryClient, error) {

	client, err := c.ForKind(m.CategoryGVK())
	if err != nil {
		return nil, err
	}

	return &CategoryClient{client}, nil

}

// List - gets a list of resources
func (c *CategoryClient) List(options ...v1.ListOptions) ([]*m.Category, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Category, len(riList))

	for i := range riList {
		result[i] = &m.Category{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *CategoryClient) Get(name string) (*m.Category, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Category{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *CategoryClient) Delete(res *m.Category) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *CategoryClient) Create(res *m.Category, opts ...v1.CreateOption) (*m.Category, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Category{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *CategoryClient) Update(res *m.Category, opts ...v1.UpdateOption) (*m.Category, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Category{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
