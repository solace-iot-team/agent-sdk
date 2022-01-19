/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AddOnGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AddOn",
		},
		APIVersion: "v1alpha1",
	}

	AddOnScopes = []string{"ProductPlan"}
)

const AddOnResourceName = "addons"

func AddOnGVK() apiv1.GroupVersionKind {
	return _AddOnGVK
}

func init() {
	apiv1.RegisterGVK(_AddOnGVK, AddOnScopes[0], AddOnResourceName)
}

// AddOn Resource
type AddOn struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	References AddOnReferences `json:"references"`

	Spec AddOnSpec `json:"spec"`

	Status AddOnStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a AddOn
func (res *AddOn) FromInstance(ri *apiv1.ResourceInstance) error {
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

// AddOnFromInstanceArray converts a []*ResourceInstance to a []*AddOn
func AddOnFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AddOn, error) {
	newArray := make([]*AddOn, 0)
	for _, item := range fromArray {
		res := &AddOn{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AddOn, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AddOn to a ResourceInstance
func (res *AddOn) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AddOnGVK()
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