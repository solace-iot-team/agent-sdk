/*
 * This file is automatically generated
 */

package catalog

import (
	"encoding/json"
	"fmt"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

var (
	ProductOverviewCtx log.ContextField = "productOverview"

	_ProductOverviewGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "ProductOverview",
		},
		APIVersion: "v1",
	}

	ProductOverviewScopes = []string{"Product", "ProductRelease"}
)

const ProductOverviewResourceName = "productoverviews"

func ProductOverviewGVK() apiv1.GroupVersionKind {
	return _ProductOverviewGVK
}

func init() {
	apiv1.RegisterGVK(_ProductOverviewGVK, ProductOverviewScopes[0], ProductOverviewResourceName)
	log.RegisterContextField(ProductOverviewCtx)
}

// ProductOverview Resource
type ProductOverview struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner        `json:"owner"`
	Spec  ProductOverviewSpec `json:"spec"`
}

// NewProductOverview creates an empty *ProductOverview
func NewProductOverview(name, scopeKind, scopeName string) (*ProductOverview, error) {
	validScope := false
	for _, s := range ProductOverviewScopes {
		if scopeKind == s {
			validScope = true
			break
		}
	}
	if !validScope {
		return nil, fmt.Errorf("scope '%s' not valid for ProductOverview kind", scopeKind)
	}

	return &ProductOverview{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ProductOverviewGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: scopeKind,
				},
			},
		},
	}, nil
}

// ProductOverviewFromInstanceArray converts a []*ResourceInstance to a []*ProductOverview
func ProductOverviewFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ProductOverview, error) {
	newArray := make([]*ProductOverview, 0)
	for _, item := range fromArray {
		res := &ProductOverview{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ProductOverview, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ProductOverview to a ResourceInstance
func (res *ProductOverview) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ProductOverviewGVK()
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

// FromInstance converts a ResourceInstance to a ProductOverview
func (res *ProductOverview) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ProductOverview) MarshalJSON() ([]byte, error) {
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
func (res *ProductOverview) UnmarshalJSON(data []byte) error {
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
func (res *ProductOverview) PluralName() string {
	return ProductOverviewResourceName
}
