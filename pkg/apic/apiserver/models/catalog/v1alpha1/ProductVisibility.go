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
	ProductVisibilityCtx log.ContextField = "productVisibility"

	_ProductVisibilityGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "ProductVisibility",
		},
		APIVersion: "v1alpha1",
	}

	ProductVisibilityScopes = []string{"Marketplace"}
)

const (
	ProductVisibilityResourceName = "productvisibility"
)

func ProductVisibilityGVK() apiv1.GroupVersionKind {
	return _ProductVisibilityGVK
}

func init() {
	apiv1.RegisterGVK(_ProductVisibilityGVK, ProductVisibilityScopes[0], ProductVisibilityResourceName)
	log.RegisterContextField(ProductVisibilityCtx)
}

// ProductVisibility Resource
type ProductVisibility struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner          `json:"owner"`
	Spec  ProductVisibilitySpec `json:"spec"`
}

// NewProductVisibility creates an empty *ProductVisibility
func NewProductVisibility(name, scopeName string) *ProductVisibility {
	return &ProductVisibility{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ProductVisibilityGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: ProductVisibilityScopes[0],
				},
			},
		},
	}
}

// ProductVisibilityFromInstanceArray converts a []*ResourceInstance to a []*ProductVisibility
func ProductVisibilityFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ProductVisibility, error) {
	newArray := make([]*ProductVisibility, 0)
	for _, item := range fromArray {
		res := &ProductVisibility{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ProductVisibility, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ProductVisibility to a ResourceInstance
func (res *ProductVisibility) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ProductVisibilityGVK()
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

// FromInstance converts a ResourceInstance to a ProductVisibility
func (res *ProductVisibility) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ProductVisibility) MarshalJSON() ([]byte, error) {
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
func (res *ProductVisibility) UnmarshalJSON(data []byte) error {
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
func (res *ProductVisibility) PluralName() string {
	return ProductVisibilityResourceName
}
