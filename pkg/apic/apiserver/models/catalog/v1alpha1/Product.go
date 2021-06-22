/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ProductGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Product",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ProductScope = ""

	ProductResourceName = "products"
)

func ProductGVK() apiv1.GroupVersionKind {
	return _ProductGVK
}

func init() {
	apiv1.RegisterGVK(_ProductGVK, ProductScope, ProductResourceName)
}

// Product Resource
type Product struct {
	apiv1.ResourceMeta

	AssetRelease struct{} `json:"assetrelease"`

	Icon struct{} `json:"icon"`

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec ProductSpec `json:"spec"`

	State struct{} `json:"state"`
}

// FromInstance converts a ResourceInstance to a Product
func (res *Product) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ProductSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	var assetrelease *struct{}
	err = json.Unmarshal(ri.SubResources["assetrelease"], assetrelease)
	if err != nil {
		return err
	}

	var icon *struct{}
	err = json.Unmarshal(ri.SubResources["icon"], icon)
	if err != nil {
		return err
	}

	var state *struct{}
	err = json.Unmarshal(ri.SubResources["state"], state)
	if err != nil {
		return err
	}

	*res = Product{ResourceMeta: ri.ResourceMeta, Spec: *spec, AssetRelease: *assetrelease, Icon: *icon, State: *state}

	return err
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
	meta.GroupVersionKind = ProductGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
