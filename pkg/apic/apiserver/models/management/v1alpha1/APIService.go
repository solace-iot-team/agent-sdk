/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_APIServiceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APIService",
		},
		APIVersion: "v1alpha1",
	}

	APIServiceScopes = []string{"Environment"}
)

const APIServiceResourceName = "apiservices"

func APIServiceGVK() apiv1.GroupVersionKind {
	return _APIServiceGVK
}

func init() {
	apiv1.RegisterGVK(_APIServiceGVK, APIServiceScopes[0], APIServiceResourceName)
}

// APIService Resource
type APIService struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	Spec ApiServiceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a APIService
func (res *APIService) FromInstance(ri *apiv1.ResourceInstance) error {
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

// APIServiceFromInstanceArray converts a []*ResourceInstance to a []*APIService
func APIServiceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APIService, error) {
	newArray := make([]*APIService, 0)
	for _, item := range fromArray {
		res := &APIService{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APIService, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APIService to a ResourceInstance
func (res *APIService) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = APIServiceGVK()
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
