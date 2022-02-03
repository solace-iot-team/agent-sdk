/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceDefinitionVersionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceDefinitionVersion",
		},
		APIVersion: "v1alpha1",
	}

	ResourceDefinitionVersionScopes = []string{"ResourceGroup"}
)

const ResourceDefinitionVersionResourceName = "resourceversions"

func ResourceDefinitionVersionGVK() apiv1.GroupVersionKind {
	return _ResourceDefinitionVersionGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDefinitionVersionGVK, ResourceDefinitionVersionScopes[0], ResourceDefinitionVersionResourceName)
}

// ResourceDefinitionVersion Resource
type ResourceDefinitionVersion struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner                  `json:"owner"`
	Spec  ResourceDefinitionVersionSpec `json:"spec"`
}

// ResourceDefinitionVersionFromInstanceArray converts a []*ResourceInstance to a []*ResourceDefinitionVersion
func ResourceDefinitionVersionFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ResourceDefinitionVersion, error) {
	newArray := make([]*ResourceDefinitionVersion, 0)
	for _, item := range fromArray {
		res := &ResourceDefinitionVersion{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ResourceDefinitionVersion, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ResourceDefinitionVersion to a ResourceInstance
func (res *ResourceDefinitionVersion) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ResourceDefinitionVersionGVK()
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

// FromInstance converts a ResourceInstance to a ResourceDefinitionVersion
func (res *ResourceDefinitionVersion) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ResourceDefinitionVersion) MarshalJSON() ([]byte, error) {
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
func (res *ResourceDefinitionVersion) UnmarshalJSON(data []byte) error {
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
