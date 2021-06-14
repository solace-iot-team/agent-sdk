/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_VirtualAPIGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "VirtualAPI",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	VirtualAPIScope = ""

	VirtualAPIResource = "virtualapis"
)

func VirtualAPIGVK() apiv1.GroupVersionKind {
	return _VirtualAPIGVK
}

func init() {
	apiv1.RegisterGVK(_VirtualAPIGVK, VirtualAPIScope, VirtualAPIResource)
}

// VirtualAPI Resource
type VirtualAPI struct {
	apiv1.ResourceMeta

	Icon struct{} `json:"icon"`

	Owner struct{} `json:"owner"`

	Spec VirtualApiSpec `json:"spec"`

	State struct{} `json:"state"`
}

// FromInstance converts a ResourceInstance to a VirtualAPI
func (res *VirtualAPI) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &VirtualApiSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = VirtualAPI{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a VirtualAPI to a ResourceInstance
func (res *VirtualAPI) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = VirtualAPIGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
