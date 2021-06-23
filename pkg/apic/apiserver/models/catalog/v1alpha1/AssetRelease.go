/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetReleaseGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AssetRelease",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AssetReleaseScope = ""

	AssetReleaseResourceName = "assetreleases"
)

func AssetReleaseGVK() apiv1.GroupVersionKind {
	return _AssetReleaseGVK
}

func init() {
	apiv1.RegisterGVK(_AssetReleaseGVK, AssetReleaseScope, AssetReleaseResourceName)
}

// AssetRelease Resource
type AssetRelease struct {
	apiv1.ResourceMeta

	Icon interface{} `json:"icon"`

	Owner interface{} `json:"owner"`

	References AssetReleaseReferences `json:"references"`

	Spec AssetReleaseSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a AssetRelease
func (res *AssetRelease) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	err := json.Unmarshal(ri.RawResource, res)
	return err
}

// AssetReleaseFromInstanceArray converts a []*ResourceInstance to a []*AssetRelease
func AssetReleaseFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AssetRelease, error) {
	newArray := make([]*AssetRelease, 0)
	for _, item := range fromArray {
		res := &AssetRelease{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AssetRelease, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AssetRelease to a ResourceInstance
func (res *AssetRelease) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = AssetReleaseGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
