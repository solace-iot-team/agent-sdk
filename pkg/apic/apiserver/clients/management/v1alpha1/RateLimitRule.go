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

type RateLimitRuleMergeFunc func(*m.RateLimitRule, *m.RateLimitRule) (*m.RateLimitRule, error)

// RateLimitRuleMerge builds a merge option for an update operation
func RateLimitRuleMerge(f RateLimitRuleMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.RateLimitRule{}, &m.RateLimitRule{}

		switch t := prev.(type) {
		case *m.RateLimitRule:
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
		case *m.RateLimitRule:
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

// RateLimitRuleClient - rest client for RateLimitRule resources that have a defined resource scope
type RateLimitRuleClient struct {
	client v1.Scoped
}

// UnscopedRateLimitRuleClient - rest client for RateLimitRule resources that do not have a defined scope
type UnscopedRateLimitRuleClient struct {
	client v1.Unscoped
}

// NewRateLimitRuleClient - creates a client that is not scoped to any resource
func NewRateLimitRuleClient(c v1.Base) (*UnscopedRateLimitRuleClient, error) {

	client, err := c.ForKind(m.RateLimitRuleGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedRateLimitRuleClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedRateLimitRuleClient) WithScope(scope string) *RateLimitRuleClient {
	return &RateLimitRuleClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedRateLimitRuleClient) Get(name string) (*m.RateLimitRule, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.RateLimitRule{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedRateLimitRuleClient) Update(res *m.RateLimitRule, opts ...v1.UpdateOption) (*m.RateLimitRule, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.RateLimitRule{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *RateLimitRuleClient) List(options ...v1.ListOptions) ([]*m.RateLimitRule, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.RateLimitRule, len(riList))

	for i := range riList {
		result[i] = &m.RateLimitRule{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *RateLimitRuleClient) Get(name string) (*m.RateLimitRule, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.RateLimitRule{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *RateLimitRuleClient) Delete(res *m.RateLimitRule) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *RateLimitRuleClient) Create(res *m.RateLimitRule, opts ...v1.CreateOption) (*m.RateLimitRule, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.RateLimitRule{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *RateLimitRuleClient) Update(res *m.RateLimitRule, opts ...v1.UpdateOption) (*m.RateLimitRule, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.RateLimitRule{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}