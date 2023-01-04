/*
 * This file is automatically generated
 */

package management

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

var (
	EnvironmentCtx log.ContextField = "environment"

	_EnvironmentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Environment",
		},
		APIVersion: "v1alpha1",
	}

	EnvironmentScopes = []string{""}
)

const EnvironmentResourceName = "environments"

func EnvironmentGVK() apiv1.GroupVersionKind {
	return _EnvironmentGVK
}

func init() {
	apiv1.RegisterGVK(_EnvironmentGVK, EnvironmentScopes[0], EnvironmentResourceName)
	log.RegisterContextField(EnvironmentCtx)
}

// Environment Resource
type Environment struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner    `json:"owner"`
	Spec  EnvironmentSpec `json:"spec"`
}

// NewEnvironment creates an empty *Environment
func NewEnvironment(name string) *Environment {
	return &Environment{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _EnvironmentGVK,
		},
	}
}

// EnvironmentFromInstanceArray converts a []*ResourceInstance to a []*Environment
func EnvironmentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Environment, error) {
	newArray := make([]*Environment, 0)
	for _, item := range fromArray {
		res := &Environment{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Environment, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Environment to a ResourceInstance
func (res *Environment) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = EnvironmentGVK()
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

// FromInstance converts a ResourceInstance to a Environment
func (res *Environment) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *Environment) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["owner"] = res.Owner
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Environment) UnmarshalJSON(data []byte) error {
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

	return nil
}

// PluralName returns the plural name of the resource
func (res *Environment) PluralName() string {
	return EnvironmentResourceName
}
