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
	SubscriptionInvoiceCtx log.ContextField = "subscriptionInvoice"

	_SubscriptionInvoiceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "SubscriptionInvoice",
		},
		APIVersion: "v1alpha1",
	}

	SubscriptionInvoiceScopes = []string{"Subscription"}
)

const (
	SubscriptionInvoiceResourceName               = "subscriptioninvoices"
	SubscriptionInvoiceBillingSubResourceName     = "billing"
	SubscriptionInvoiceMarketplaceSubResourceName = "marketplace"
	SubscriptionInvoiceStateSubResourceName       = "state"
	SubscriptionInvoiceStatusSubResourceName      = "status"
)

func SubscriptionInvoiceGVK() apiv1.GroupVersionKind {
	return _SubscriptionInvoiceGVK
}

func init() {
	apiv1.RegisterGVK(_SubscriptionInvoiceGVK, SubscriptionInvoiceScopes[0], SubscriptionInvoiceResourceName)
	log.RegisterContextField(SubscriptionInvoiceCtx)
}

// SubscriptionInvoice Resource
type SubscriptionInvoice struct {
	apiv1.ResourceMeta
	Billing     SubscriptionInvoiceBilling     `json:"billing"`
	Marketplace SubscriptionInvoiceMarketplace `json:"marketplace"`
	Owner       *apiv1.Owner                   `json:"owner"`
	Spec        SubscriptionInvoiceSpec        `json:"spec"`
	State       SubscriptionInvoiceState       `json:"state"`
	// Status      SubscriptionInvoiceStatus      `json:"status"`
	Status *apiv1.ResourceStatus `json:"status"`
}

// NewSubscriptionInvoice creates an empty *SubscriptionInvoice
func NewSubscriptionInvoice(name, scopeName string) *SubscriptionInvoice {
	return &SubscriptionInvoice{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _SubscriptionInvoiceGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: SubscriptionInvoiceScopes[0],
				},
			},
		},
	}
}

// SubscriptionInvoiceFromInstanceArray converts a []*ResourceInstance to a []*SubscriptionInvoice
func SubscriptionInvoiceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*SubscriptionInvoice, error) {
	newArray := make([]*SubscriptionInvoice, 0)
	for _, item := range fromArray {
		res := &SubscriptionInvoice{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*SubscriptionInvoice, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a SubscriptionInvoice to a ResourceInstance
func (res *SubscriptionInvoice) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = SubscriptionInvoiceGVK()
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

// FromInstance converts a ResourceInstance to a SubscriptionInvoice
func (res *SubscriptionInvoice) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *SubscriptionInvoice) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["billing"] = res.Billing
	out["marketplace"] = res.Marketplace
	out["owner"] = res.Owner
	out["spec"] = res.Spec
	out["state"] = res.State
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *SubscriptionInvoice) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Billing
	if v, ok := aux.SubResources["billing"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "billing")
		err = json.Unmarshal(sr, &res.Billing)
		if err != nil {
			return err
		}
	}

	// marshalling subresource Marketplace
	if v, ok := aux.SubResources["marketplace"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "marketplace")
		err = json.Unmarshal(sr, &res.Marketplace)
		if err != nil {
			return err
		}
	}

	// marshalling subresource State
	if v, ok := aux.SubResources["state"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "state")
		err = json.Unmarshal(sr, &res.State)
		if err != nil {
			return err
		}
	}

	// marshalling subresource Status
	if v, ok := aux.SubResources["status"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "status")
		// err = json.Unmarshal(sr, &res.Status)
		res.Status = &apiv1.ResourceStatus{}
		err = json.Unmarshal(sr, res.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *SubscriptionInvoice) PluralName() string {
	return SubscriptionInvoiceResourceName
}
