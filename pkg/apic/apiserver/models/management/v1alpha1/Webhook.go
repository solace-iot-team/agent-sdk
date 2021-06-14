/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_WebhookGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Webhook",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	WebhookScope = "Environment"

	WebhookResource = "webhooks"
)

func WebhookGVK() apiv1.GroupVersionKind {
	return _WebhookGVK
}

func init() {
	apiv1.RegisterGVK(_WebhookGVK, WebhookScope, WebhookResource)
}

// Webhook Resource
type Webhook struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec WebhookSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Webhook
func (res *Webhook) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &WebhookSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = Webhook{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a Webhook to a ResourceInstance
func (res *Webhook) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = WebhookGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
