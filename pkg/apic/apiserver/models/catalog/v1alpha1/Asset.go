/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Asset",
		},
		APIVersion: "v1alpha1",
	}

	AssetScopes = []string{""}
)

const AssetResourceName = "assets"

func AssetGVK() apiv1.GroupVersionKind {
	return _AssetGVK
}

func init() {
	apiv1.RegisterGVK(_AssetGVK, AssetScopes[0], AssetResourceName)
}

// Asset Resource
type Asset struct {
	apiv1.ResourceMeta
	Icon       interface{}  `json:"icon"`
	Owner      *apiv1.Owner `json:"owner"`
	References interface{}  `json:"references"`
	Spec       AssetSpec    `json:"spec"`
	State      AssetState   `json:"state"`
}

// NewAsset creates an empty *Asset
func NewAsset(name string) *Asset {
	return &Asset{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _AssetGVK,
		},
	}
}

// AssetFromInstanceArray converts a []*ResourceInstance to a []*Asset
func AssetFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Asset, error) {
	newArray := make([]*Asset, 0)
	for _, item := range fromArray {
		res := &Asset{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Asset, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Asset to a ResourceInstance
func (res *Asset) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AssetGVK()
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

// FromInstance converts a ResourceInstance to a Asset
func (res *Asset) FromInstance(ri *apiv1.ResourceInstance) error {
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

// MarshalJSON custom marshaller to handle sub resources
func (res *Asset) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["icon"] = res.Icon
	out["owner"] = res.Owner
	out["references"] = res.References
	out["spec"] = res.Spec
	out["state"] = res.State

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Asset) UnmarshalJSON(data []byte) error {
	var err error

	aux := &apiv1.ResourceInstance{}
	err = json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	res.ResourceMeta = aux.ResourceMeta
	res.Owner = aux.Owner

	// ResourceInstance holds the spec as a map[string]interface{}.
	// Convert it to bytes, then convert to the spec type for the resource.
	sr, err := json.Marshal(aux.Spec)
	if err != nil {
		return err
	}

	err = json.Unmarshal(sr, &res.Spec)
	if err != nil {
		return err
	}

	// marshalling subresource Icon
	if v, ok := aux.SubResources["icon"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "icon")
		err = json.Unmarshal(sr, &res.Icon)
		if err != nil {
			return err
		}
	}

	// marshalling subresource References
	if v, ok := aux.SubResources["references"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "references")
		err = json.Unmarshal(sr, &res.References)
		if err != nil {
			return err
		}
	}

	// marshalling subresource State
	if v, ok := aux.SubResources["state"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "state")
		err = json.Unmarshal(sr, &res.State)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *Asset) PluralName() string {
	return AssetResourceName
}
