/*
 * This file is automatically generated
 */

package catalog

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_CategoryGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Category",
		},
		APIVersion: "v1alpha1",
	}

	CategoryScopes = []string{""}
)

const CategoryResourceName = "categories"

func CategoryGVK() apiv1.GroupVersionKind {
	return _CategoryGVK
}

func init() {
	apiv1.RegisterGVK(_CategoryGVK, CategoryScopes[0], CategoryResourceName)
}

// Category Resource
type Category struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner `json:"owner"`
	Spec  CategorySpec `json:"spec"`
}

// NewCategory creates an empty *Category
func NewCategory(name string) *Category {
	return &Category{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _CategoryGVK,
		},
	}
}

// CategoryFromInstanceArray converts a []*ResourceInstance to a []*Category
func CategoryFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Category, error) {
	newArray := make([]*Category, 0)
	for _, item := range fromArray {
		res := &Category{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Category, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Category to a ResourceInstance
func (res *Category) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = CategoryGVK()
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

// FromInstance converts a ResourceInstance to a Category
func (res *Category) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *Category) MarshalJSON() ([]byte, error) {
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
func (res *Category) UnmarshalJSON(data []byte) error {
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
func (res *Category) PluralName() string {
	return CategoryResourceName
}
