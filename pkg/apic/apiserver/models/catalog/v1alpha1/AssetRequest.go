/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/solace-iot-team/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetRequestGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AssetRequest",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AssetRequestScope = "Asset"

	AssetRequestResourceName = "assetrequests"
)

func AssetRequestGVK() apiv1.GroupVersionKind {
	return _AssetRequestGVK
}

func init() {
	apiv1.RegisterGVK(_AssetRequestGVK, AssetRequestScope, AssetRequestResourceName)
}

// AssetRequest Resource
type AssetRequest struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec AssetRequestSpec `json:"spec"`

	State AssetRequestState `json:"state"`
}

// FromInstance converts a ResourceInstance to a AssetRequest
func (res *AssetRequest) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AssetRequestSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = AssetRequest{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AssetRequestFromInstanceArray converts a []*ResourceInstance to a []*AssetRequest
func AssetRequestFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AssetRequest, error) {
	newArray := make([]*AssetRequest, 0)
	for _, item := range fromArray {
		res := &AssetRequest{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AssetRequest, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AssetRequest to a ResourceInstance
func (res *AssetRequest) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = AssetRequestGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
