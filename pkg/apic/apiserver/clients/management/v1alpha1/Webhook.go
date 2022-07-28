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

type WebhookMergeFunc func(*m.Webhook, *m.Webhook) (*m.Webhook, error)

// WebhookMerge builds a merge option for an update operation
func WebhookMerge(f WebhookMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Webhook{}, &m.Webhook{}

		switch t := prev.(type) {
		case *m.Webhook:
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
		case *m.Webhook:
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

// WebhookClient - rest client for Webhook resources that have a defined resource scope
type WebhookClient struct {
	client v1.Scoped
}

// UnscopedWebhookClient - rest client for Webhook resources that do not have a defined scope
type UnscopedWebhookClient struct {
	client v1.Unscoped
}

// NewWebhookClient - creates a client that is not scoped to any resource
func NewWebhookClient(c v1.Base) (*UnscopedWebhookClient, error) {

	client, err := c.ForKind(m.WebhookGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedWebhookClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedWebhookClient) WithScope(scope string) *WebhookClient {
	return &WebhookClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedWebhookClient) Get(name string) (*m.Webhook, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Webhook{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedWebhookClient) Update(res *m.Webhook, opts ...v1.UpdateOption) (*m.Webhook, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Webhook{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *WebhookClient) List(options ...v1.ListOptions) ([]*m.Webhook, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Webhook, len(riList))

	for i := range riList {
		result[i] = &m.Webhook{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *WebhookClient) Get(name string) (*m.Webhook, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Webhook{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *WebhookClient) Delete(res *m.Webhook) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *WebhookClient) Create(res *m.Webhook, opts ...v1.CreateOption) (*m.Webhook, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Webhook{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *WebhookClient) Update(res *m.Webhook, opts ...v1.UpdateOption) (*m.Webhook, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Webhook{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
