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
	PublishedProductCtx log.ContextField = "publishedProduct"

	_PublishedProductGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "PublishedProduct",
		},
		APIVersion: "v1alpha1",
	}

	PublishedProductScopes = []string{"Marketplace"}
)

const (
	PublishedProductResourceName              = "publishedproducts"
	PublishedProductReferencesSubResourceName = "references"
)

func PublishedProductGVK() apiv1.GroupVersionKind {
	return _PublishedProductGVK
}

func init() {
	apiv1.RegisterGVK(_PublishedProductGVK, PublishedProductScopes[0], PublishedProductResourceName)
	log.RegisterContextField(PublishedProductCtx)
}

// PublishedProduct Resource
type PublishedProduct struct {
	apiv1.ResourceMeta
	Owner      *apiv1.Owner               `json:"owner"`
	References PublishedProductReferences `json:"references"`
	Spec       PublishedProductSpec       `json:"spec"`
}

// NewPublishedProduct creates an empty *PublishedProduct
func NewPublishedProduct(name, scopeName string) *PublishedProduct {
	return &PublishedProduct{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _PublishedProductGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: PublishedProductScopes[0],
				},
			},
		},
	}
}

// PublishedProductFromInstanceArray converts a []*ResourceInstance to a []*PublishedProduct
func PublishedProductFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*PublishedProduct, error) {
	newArray := make([]*PublishedProduct, 0)
	for _, item := range fromArray {
		res := &PublishedProduct{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*PublishedProduct, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a PublishedProduct to a ResourceInstance
func (res *PublishedProduct) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = PublishedProductGVK()
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

// FromInstance converts a ResourceInstance to a PublishedProduct
func (res *PublishedProduct) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *PublishedProduct) MarshalJSON() ([]byte, error) {
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
	out["references"] = res.References
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *PublishedProduct) UnmarshalJSON(data []byte) error {
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

	return nil
}

// PluralName returns the plural name of the resource
func (res *PublishedProduct) PluralName() string {
	return PublishedProductResourceName
}
