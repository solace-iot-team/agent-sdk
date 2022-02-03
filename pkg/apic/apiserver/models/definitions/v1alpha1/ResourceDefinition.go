/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceDefinitionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceDefinition",
		},
		APIVersion: "v1alpha1",
	}

	ResourceDefinitionScopes = []string{"ResourceGroup"}
)

const ResourceDefinitionResourceName = "resources"

func ResourceDefinitionGVK() apiv1.GroupVersionKind {
	return _ResourceDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDefinitionGVK, ResourceDefinitionScopes[0], ResourceDefinitionResourceName)
}

// ResourceDefinition Resource
type ResourceDefinition struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner           `json:"owner"`
	Spec  ResourceDefinitionSpec `json:"spec"`
}

// ResourceDefinitionFromInstanceArray converts a []*ResourceInstance to a []*ResourceDefinition
func ResourceDefinitionFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ResourceDefinition, error) {
	newArray := make([]*ResourceDefinition, 0)
	for _, item := range fromArray {
		res := &ResourceDefinition{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ResourceDefinition, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ResourceDefinition to a ResourceInstance
func (res *ResourceDefinition) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ResourceDefinitionGVK()
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

// FromInstance converts a ResourceInstance to a ResourceDefinition
func (res *ResourceDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ResourceDefinition) MarshalJSON() ([]byte, error) {
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
func (res *ResourceDefinition) UnmarshalJSON(data []byte) error {
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
