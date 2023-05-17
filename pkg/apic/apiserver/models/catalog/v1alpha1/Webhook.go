/*
 * This file is automatically generated
 */

package catalog

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

var (
	WebhookCtx log.ContextField = "webhook"

	_WebhookGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Webhook",
		},
		APIVersion: "v1alpha1",
	}

	WebhookScopes = []string{"AuthorizationProfile"}
)

const (
	WebhookResourceName = "webhooks"
)

func WebhookGVK() apiv1.GroupVersionKind {
	return _WebhookGVK
}

func init() {
	apiv1.RegisterGVK(_WebhookGVK, WebhookScopes[0], WebhookResourceName)
	log.RegisterContextField(WebhookCtx)
}

// Webhook Resource
type Webhook struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner `json:"owner"`
	Spec  WebhookSpec  `json:"spec"`
}

// NewWebhook creates an empty *Webhook
func NewWebhook(name, scopeName string) *Webhook {
	return &Webhook{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _WebhookGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: WebhookScopes[0],
				},
			},
		},
	}
}

// WebhookFromInstanceArray converts a []*ResourceInstance to a []*Webhook
func WebhookFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Webhook, error) {
	newArray := make([]*Webhook, 0)
	for _, item := range fromArray {
		res := &Webhook{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Webhook, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Webhook to a ResourceInstance
func (res *Webhook) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = WebhookGVK()
	res.ResourceMeta = meta

	m, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	instance := apiv1.ResourceInstance{}
	err = json.Unmarshal(m, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}

// FromInstance converts a ResourceInstance to a Webhook
func (res *Webhook) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}
	var err error
	rawResource := ri.GetRawResource()
	if rawResource == nil {
		rawResource, err = json.Marshal(ri)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(rawResource, res)
	return err
}

// MarshalJSON custom marshaller to handle sub resources
func (res *Webhook) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["owner"] = res.Owner
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Webhook) UnmarshalJSON(data []byte) error {
	var err error

	aux := &apiv1.ResourceInstance{}
	err = json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	res.ResourceMeta = aux.ResourceMeta
	res.Owner = aux.Owner

	// ResourceInstance holds the spec as a map[string]interface{}.
	// Convert it to bytes, then convert to the spec type for the resource.
	sr, err := json.Marshal(aux.Spec)
	if err != nil {
		return err
	}

	err = json.Unmarshal(sr, &res.Spec)
	if err != nil {
		return err
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *Webhook) PluralName() string {
	return WebhookResourceName
}
