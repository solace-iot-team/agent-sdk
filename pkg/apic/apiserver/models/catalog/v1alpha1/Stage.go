/*
 * This file is automatically generated
 */

package catalog

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

var (
	StageCtx log.ContextField = "stage"

	_StageGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Stage",
		},
		APIVersion: "v1alpha1",
	}

	StageScopes = []string{""}
)

const (
	StageResourceName        = "stages"
	StageIconSubResourceName = "icon"
)

func StageGVK() apiv1.GroupVersionKind {
	return _StageGVK
}

func init() {
	apiv1.RegisterGVK(_StageGVK, StageScopes[0], StageResourceName)
	log.RegisterContextField(StageCtx)
}

// Stage Resource
type Stage struct {
	apiv1.ResourceMeta
	Icon  interface{}  `json:"icon"`
	Owner *apiv1.Owner `json:"owner"`
	Spec  StageSpec    `json:"spec"`
}

// NewStage creates an empty *Stage
func NewStage(name string) *Stage {
	return &Stage{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _StageGVK,
		},
	}
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

// MarshalJSON custom marshaller to handle sub resources
func (res *Stage) MarshalJSON() ([]byte, error) {
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
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Stage) UnmarshalJSON(data []byte) error {
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

	return nil
}

// PluralName returns the plural name of the resource
func (res *Stage) PluralName() string {
	return StageResourceName
}
