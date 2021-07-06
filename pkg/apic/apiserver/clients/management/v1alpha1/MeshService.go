/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

type MeshServiceMergeFunc func(*v1alpha1.MeshService, *v1alpha1.MeshService) (*v1alpha1.MeshService, error)

// Merge builds a merge option for an update operation
func MeshServiceMerge(f MeshServiceMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.MeshService{}, &v1alpha1.MeshService{}

		switch t := prev.(type) {
		case *v1alpha1.MeshService:
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
		case *v1alpha1.MeshService:
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

// MeshServiceClient -
type MeshServiceClient struct {
	client v1.Scoped
}

// UnscopedMeshServiceClient -
type UnscopedMeshServiceClient struct {
	client v1.Unscoped
}

// NewMeshServiceClient -
func NewMeshServiceClient(c v1.Base) (*UnscopedMeshServiceClient, error) {

	client, err := c.ForKind(v1alpha1.MeshServiceGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedMeshServiceClient{client}, nil

}

// WithScope -
func (c *UnscopedMeshServiceClient) WithScope(scope string) *MeshServiceClient {
	return &MeshServiceClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedMeshServiceClient) Get(name string) (*v1alpha1.MeshService, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.MeshService{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedMeshServiceClient) Update(res *v1alpha1.MeshService, opts ...v1.UpdateOption) (*v1alpha1.MeshService, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.MeshService{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *MeshServiceClient) List(options ...v1.ListOptions) ([]*v1alpha1.MeshService, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.MeshService, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.MeshService{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *MeshServiceClient) Get(name string) (*v1alpha1.MeshService, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.MeshService{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *MeshServiceClient) Delete(res *v1alpha1.MeshService) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *MeshServiceClient) Create(res *v1alpha1.MeshService, opts ...v1.CreateOption) (*v1alpha1.MeshService, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.MeshService{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *MeshServiceClient) Update(res *v1alpha1.MeshService, opts ...v1.UpdateOption) (*v1alpha1.MeshService, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.MeshService{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
