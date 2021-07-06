/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type StageMergeFunc func(*v1alpha1.Stage, *v1alpha1.Stage) (*v1alpha1.Stage, error)

// Merge builds a merge option for an update operation
func StageMerge(f StageMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.Stage{}, &v1alpha1.Stage{}

		switch t := prev.(type) {
		case *v1alpha1.Stage:
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
		case *v1alpha1.Stage:
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

// StageClient -
type StageClient struct {
	client v1.Scoped
}

// NewStageClient -
func NewStageClient(c v1.Base) (*StageClient, error) {

	client, err := c.ForKind(v1alpha1.StageGVK())
	if err != nil {
		return nil, err
	}

	return &StageClient{client}, nil

}

// List -
func (c *StageClient) List(options ...v1.ListOptions) ([]*v1alpha1.Stage, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.Stage, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.Stage{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *StageClient) Get(name string) (*v1alpha1.Stage, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.Stage{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *StageClient) Delete(res *v1alpha1.Stage) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *StageClient) Create(res *v1alpha1.Stage, opts ...v1.CreateOption) (*v1alpha1.Stage, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.Stage{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *StageClient) Update(res *v1alpha1.Stage, opts ...v1.UpdateOption) (*v1alpha1.Stage, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.Stage{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
