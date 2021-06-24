/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_VirtualServiceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "VirtualService",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	VirtualServiceScope = "VirtualAPI"

	VirtualServiceResourceName = "virtualservices"
)

func VirtualServiceGVK() apiv1.GroupVersionKind {
	return _VirtualServiceGVK
}

func init() {
	apiv1.RegisterGVK(_VirtualServiceGVK, VirtualServiceScope, VirtualServiceResourceName)
}

// VirtualService Resource
type VirtualService struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec VirtualServiceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a VirtualService
func (res *VirtualService) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	err := json.Unmarshal(ri.RawResource, res)
	return err
}

// VirtualServiceFromInstanceArray converts a []*ResourceInstance to a []*VirtualService
func VirtualServiceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*VirtualService, error) {
	newArray := make([]*VirtualService, 0)
	for _, item := range fromArray {
		res := &VirtualService{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*VirtualService, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a VirtualService to a ResourceInstance
func (res *VirtualService) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = VirtualServiceGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
