/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AccessRequestGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AccessRequest",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AccessRequestScope = "Environment"

	AccessRequestResourceName = "accessrequests"
)

func AccessRequestGVK() apiv1.GroupVersionKind {
	return _AccessRequestGVK
}

func init() {
	apiv1.RegisterGVK(_AccessRequestGVK, AccessRequestScope, AccessRequestResourceName)
}

// AccessRequest Resource
type AccessRequest struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec AccessRequestSpec `json:"spec"`

	State AccessRequestState `json:"state"`
}

// FromInstance converts a ResourceInstance to a AccessRequest
func (res *AccessRequest) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AccessRequestSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	var state *AccessRequestState
	err = json.Unmarshal(ri.SubResources["AccessRequestState"], state)
	if err != nil {
		return err
	}

	*res = AccessRequest{ResourceMeta: ri.ResourceMeta, Spec: *spec, State: *state}

	return err
}

// AccessRequestFromInstanceArray converts a []*ResourceInstance to a []*AccessRequest
func AccessRequestFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AccessRequest, error) {
	newArray := make([]*AccessRequest, 0)
	for _, item := range fromArray {
		res := &AccessRequest{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AccessRequest, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AccessRequest to a ResourceInstance
func (res *AccessRequest) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = AccessRequestGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
