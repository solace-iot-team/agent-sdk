/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AmplifyRuntimeConfigGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AmplifyRuntimeConfig",
		},
		APIVersion: "v1alpha1",
	}

	AmplifyRuntimeConfigScopes = []string{"Environment"}
)

const AmplifyRuntimeConfigResourceName = "ampruntimeconfigs"

func AmplifyRuntimeConfigGVK() apiv1.GroupVersionKind {
	return _AmplifyRuntimeConfigGVK
}

func init() {
	apiv1.RegisterGVK(_AmplifyRuntimeConfigGVK, AmplifyRuntimeConfigScopes[0], AmplifyRuntimeConfigResourceName)
}

// AmplifyRuntimeConfig Resource
type AmplifyRuntimeConfig struct {
	apiv1.ResourceMeta
	Owner  *apiv1.Owner               `json:"owner"`
	Spec   AmplifyRuntimeConfigSpec   `json:"spec"`
	Status AmplifyRuntimeConfigStatus `json:"status"`
}

// NewAmplifyRuntimeConfig creates an empty *AmplifyRuntimeConfig
func NewAmplifyRuntimeConfig(name, scopeName string) *AmplifyRuntimeConfig {
	return &AmplifyRuntimeConfig{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _AmplifyRuntimeConfigGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: AmplifyRuntimeConfigScopes[0],
				},
			},
		},
	}
}

// AmplifyRuntimeConfigFromInstanceArray converts a []*ResourceInstance to a []*AmplifyRuntimeConfig
func AmplifyRuntimeConfigFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AmplifyRuntimeConfig, error) {
	newArray := make([]*AmplifyRuntimeConfig, 0)
	for _, item := range fromArray {
		res := &AmplifyRuntimeConfig{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AmplifyRuntimeConfig, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AmplifyRuntimeConfig to a ResourceInstance
func (res *AmplifyRuntimeConfig) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AmplifyRuntimeConfigGVK()
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

// FromInstance converts a ResourceInstance to a AmplifyRuntimeConfig
func (res *AmplifyRuntimeConfig) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *AmplifyRuntimeConfig) MarshalJSON() ([]byte, error) {
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
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *AmplifyRuntimeConfig) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Status
	if v, ok := aux.SubResources["status"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "status")
		err = json.Unmarshal(sr, &res.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *AmplifyRuntimeConfig) PluralName() string {
	return AmplifyRuntimeConfigResourceName
}
