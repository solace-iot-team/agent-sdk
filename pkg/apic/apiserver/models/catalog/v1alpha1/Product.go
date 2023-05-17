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
	ProductCtx log.ContextField = "product"

	_ProductGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Product",
		},
		APIVersion: "v1alpha1",
	}

	ProductScopes = []string{""}
)

const ProductResourceName = "products"

func ProductGVK() apiv1.GroupVersionKind {
	return _ProductGVK
}

func init() {
	apiv1.RegisterGVK(_ProductGVK, ProductScopes[0], ProductResourceName)
	log.RegisterContextField(ProductCtx)
}

// Product Resource
type Product struct {
	apiv1.ResourceMeta
	Icon          interface{}          `json:"icon"`
	Latestrelease ProductLatestrelease `json:"latestrelease"`
	Owner         *apiv1.Owner         `json:"owner"`
	References    ProductReferences    `json:"references"`
	Spec          ProductSpec          `json:"spec"`
	State         ProductState         `json:"state"`
	// Status        ProductStatus        `json:"status"`
	Status *apiv1.ResourceStatus `json:"status"`
}

// NewProduct creates an empty *Product
func NewProduct(name string) *Product {
	return &Product{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ProductGVK,
		},
	}
}

// ProductFromInstanceArray converts a []*ResourceInstance to a []*Product
func ProductFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Product, error) {
	newArray := make([]*Product, 0)
	for _, item := range fromArray {
		res := &Product{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Product, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Product to a ResourceInstance
func (res *Product) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ProductGVK()
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

// FromInstance converts a ResourceInstance to a Product
func (res *Product) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *Product) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["icon"] = res.Icon
	out["latestrelease"] = res.Latestrelease
	out["owner"] = res.Owner
	out["references"] = res.References
	out["spec"] = res.Spec
	out["state"] = res.State
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Product) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Icon
	if v, ok := aux.SubResources["icon"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "icon")
		err = json.Unmarshal(sr, &res.Icon)
		if err != nil {
			return err
		}
	}

	// marshalling subresource Latestrelease
	if v, ok := aux.SubResources["latestrelease"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "latestrelease")
		err = json.Unmarshal(sr, &res.Latestrelease)
		if err != nil {
			return err
		}
	}

	// marshalling subresource References
	if v, ok := aux.SubResources["references"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "references")
		err = json.Unmarshal(sr, &res.References)
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
func (res *Product) PluralName() string {
	return ProductResourceName
}
