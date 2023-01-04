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
	AmplifyConfigCtx log.ContextField = "amplifyConfig"

	_AmplifyConfigGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AmplifyConfig",
		},
		APIVersion: "v1alpha1",
	}

	AmplifyConfigScopes = []string{"Environment"}
)

const AmplifyConfigResourceName = "ampconfigs"

func AmplifyConfigGVK() apiv1.GroupVersionKind {
	return _AmplifyConfigGVK
}

func init() {
	apiv1.RegisterGVK(_AmplifyConfigGVK, AmplifyConfigScopes[0], AmplifyConfigResourceName)
	log.RegisterContextField(AmplifyConfigCtx)
}

// AmplifyConfig Resource
type AmplifyConfig struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner      `json:"owner"`
	Spec  AmplifyConfigSpec `json:"spec"`
}

// NewAmplifyConfig creates an empty *AmplifyConfig
func NewAmplifyConfig(name, scopeName string) *AmplifyConfig {
	return &AmplifyConfig{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _AmplifyConfigGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: AmplifyConfigScopes[0],
				},
			},
		},
	}
}

// AmplifyConfigFromInstanceArray converts a []*ResourceInstance to a []*AmplifyConfig
func AmplifyConfigFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AmplifyConfig, error) {
	newArray := make([]*AmplifyConfig, 0)
	for _, item := range fromArray {
		res := &AmplifyConfig{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AmplifyConfig, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AmplifyConfig to a ResourceInstance
func (res *AmplifyConfig) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AmplifyConfigGVK()
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

// FromInstance converts a ResourceInstance to a AmplifyConfig
func (res *AmplifyConfig) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *AmplifyConfig) MarshalJSON() ([]byte, error) {
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
func (res *AmplifyConfig) UnmarshalJSON(data []byte) error {
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
func (res *AmplifyConfig) PluralName() string {
	return AmplifyConfigResourceName
}
