/*
 * This file is automatically generated
 */

package definitions

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceGroupGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceGroup",
		},
		APIVersion: "v1alpha1",
	}

	ResourceGroupScopes = []string{""}
)

const ResourceGroupResourceName = "groups"

func ResourceGroupGVK() apiv1.GroupVersionKind {
	return _ResourceGroupGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceGroupGVK, ResourceGroupScopes[0], ResourceGroupResourceName)
}

// ResourceGroup Resource
type ResourceGroup struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner `json:"owner"`
	Spec  interface{}  `json:"spec"`
}

// NewResourceGroup creates an empty *ResourceGroup
func NewResourceGroup(name string) *ResourceGroup {
	return &ResourceGroup{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ResourceGroupGVK,
		},
	}
}

// ResourceGroupFromInstanceArray converts a []*ResourceInstance to a []*ResourceGroup
func ResourceGroupFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ResourceGroup, error) {
	newArray := make([]*ResourceGroup, 0)
	for _, item := range fromArray {
		res := &ResourceGroup{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ResourceGroup, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ResourceGroup to a ResourceInstance
func (res *ResourceGroup) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ResourceGroupGVK()
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

// FromInstance converts a ResourceInstance to a ResourceGroup
func (res *ResourceGroup) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ResourceGroup) MarshalJSON() ([]byte, error) {
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
func (res *ResourceGroup) UnmarshalJSON(data []byte) error {
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
func (res *ResourceGroup) PluralName() string {
	return ResourceGroupResourceName
}
