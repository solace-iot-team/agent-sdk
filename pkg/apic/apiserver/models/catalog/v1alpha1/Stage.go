/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_StageGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Stage",
		},
		APIVersion: "v1alpha1",
	}

	StageScopes = []string{""}
)

const StageResourceName = "stages"

func StageGVK() apiv1.GroupVersionKind {
	return _StageGVK
}

func init() {
	apiv1.RegisterGVK(_StageGVK, StageScopes[0], StageResourceName)
}

// Stage Resource
type Stage struct {
	apiv1.ResourceMeta

	Icon interface{} `json:"icon"`

	Owner *apiv1.Owner `json:"owner"`

	Spec StageSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Stage
func (res *Stage) FromInstance(ri *apiv1.ResourceInstance) error {
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

// StageFromInstanceArray converts a []*ResourceInstance to a []*Stage
func StageFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Stage, error) {
	newArray := make([]*Stage, 0)
	for _, item := range fromArray {
		res := &Stage{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Stage, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Stage to a ResourceInstance
func (res *Stage) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = StageGVK()
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
