/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_EnvironmentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Environment",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	EnvironmentScope = ""

	EnvironmentResourceName = "environments"
)

func EnvironmentGVK() apiv1.GroupVersionKind {
	return _EnvironmentGVK
}

func init() {
	apiv1.RegisterGVK(_EnvironmentGVK, EnvironmentScope, EnvironmentResourceName)
}

// Environment Resource
type Environment struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec EnvironmentSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Environment
func (res *Environment) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	err := json.Unmarshal(ri.RawResource, res)
	return err
}

// EnvironmentFromInstanceArray converts a []*ResourceInstance to a []*Environment
func EnvironmentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Environment, error) {
	newArray := make([]*Environment, 0)
	for _, item := range fromArray {
		res := &Environment{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Environment, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Environment to a ResourceInstance
func (res *Environment) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = EnvironmentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
