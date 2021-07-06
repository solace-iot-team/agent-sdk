/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_IntegrationGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Integration",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	IntegrationScope = ""

	IntegrationResourceName = "integrations"
)

func IntegrationGVK() apiv1.GroupVersionKind {
	return _IntegrationGVK
}

func init() {
	apiv1.RegisterGVK(_IntegrationGVK, IntegrationScope, IntegrationResourceName)
}

// Integration Resource
type Integration struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec IntegrationSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Integration
func (res *Integration) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &IntegrationSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = Integration{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// IntegrationFromInstanceArray converts a []*ResourceInstance to a []*Integration
func IntegrationFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Integration, error) {
	newArray := make([]*Integration, 0)
	for _, item := range fromArray {
		res := &Integration{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Integration, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Integration to a ResourceInstance
func (res *Integration) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = IntegrationGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
