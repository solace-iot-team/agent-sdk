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

type MeshMergeFunc func(*m.Mesh, *m.Mesh) (*m.Mesh, error)

// MeshMerge builds a merge option for an update operation
func MeshMerge(f MeshMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Mesh{}, &m.Mesh{}

		switch t := prev.(type) {
		case *m.Mesh:
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
		case *m.Mesh:
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

// MeshClient - rest client for Mesh resources that have a defined resource scope
type MeshClient struct {
	client v1.Scoped
}

// NewMeshClient - creates a client scoped to a particular resource
func NewMeshClient(c v1.Base) (*MeshClient, error) {

	client, err := c.ForKind(m.MeshGVK())
	if err != nil {
		return nil, err
	}

	return &MeshClient{client}, nil

}

// List - gets a list of resources
func (c *MeshClient) List(options ...v1.ListOptions) ([]*m.Mesh, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Mesh, len(riList))

	for i := range riList {
		result[i] = &m.Mesh{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *MeshClient) Get(name string) (*m.Mesh, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Mesh{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *MeshClient) Delete(res *m.Mesh) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *MeshClient) Create(res *m.Mesh, opts ...v1.CreateOption) (*m.Mesh, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Mesh{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *MeshClient) Update(res *m.Mesh, opts ...v1.UpdateOption) (*m.Mesh, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Mesh{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
