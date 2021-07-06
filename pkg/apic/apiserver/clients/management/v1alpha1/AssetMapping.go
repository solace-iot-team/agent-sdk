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

type AssetMappingMergeFunc func(*v1alpha1.AssetMapping, *v1alpha1.AssetMapping) (*v1alpha1.AssetMapping, error)

// Merge builds a merge option for an update operation
func AssetMappingMerge(f AssetMappingMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.AssetMapping{}, &v1alpha1.AssetMapping{}

		switch t := prev.(type) {
		case *v1alpha1.AssetMapping:
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
		case *v1alpha1.AssetMapping:
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

// AssetMappingClient -
type AssetMappingClient struct {
	client v1.Scoped
}

// UnscopedAssetMappingClient -
type UnscopedAssetMappingClient struct {
	client v1.Unscoped
}

// NewAssetMappingClient -
func NewAssetMappingClient(c v1.Base) (*UnscopedAssetMappingClient, error) {

	client, err := c.ForKind(v1alpha1.AssetMappingGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAssetMappingClient{client}, nil

}

// WithScope -
func (c *UnscopedAssetMappingClient) WithScope(scope string) *AssetMappingClient {
	return &AssetMappingClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedAssetMappingClient) Get(name string) (*v1alpha1.AssetMapping, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.AssetMapping{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedAssetMappingClient) Update(res *v1alpha1.AssetMapping, opts ...v1.UpdateOption) (*v1alpha1.AssetMapping, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.AssetMapping{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *AssetMappingClient) List(options ...v1.ListOptions) ([]*v1alpha1.AssetMapping, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.AssetMapping, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.AssetMapping{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *AssetMappingClient) Get(name string) (*v1alpha1.AssetMapping, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.AssetMapping{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *AssetMappingClient) Delete(res *v1alpha1.AssetMapping) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *AssetMappingClient) Create(res *v1alpha1.AssetMapping, opts ...v1.CreateOption) (*v1alpha1.AssetMapping, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.AssetMapping{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *AssetMappingClient) Update(res *v1alpha1.AssetMapping, opts ...v1.UpdateOption) (*v1alpha1.AssetMapping, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.AssetMapping{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
