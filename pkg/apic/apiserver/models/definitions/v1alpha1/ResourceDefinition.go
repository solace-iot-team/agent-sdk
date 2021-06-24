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
)

const (
	ResourceDefinitionScope = "ResourceGroup"

	ResourceDefinitionResourceName = "resources"
)

func ResourceDefinitionGVK() apiv1.GroupVersionKind {
	return _ResourceDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDefinitionGVK, ResourceDefinitionScope, ResourceDefinitionResourceName)
}

// ResourceDefinition Resource
type ResourceDefinition struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec ResourceDefinitionSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a ResourceDefinition
func (res *ResourceDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	err := json.Unmarshal(ri.GetRawResource(), res)
	return err
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
