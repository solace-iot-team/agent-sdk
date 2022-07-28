/*
 * This file is automatically generated
 */

package management

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_APIServiceInstanceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APIServiceInstance",
		},
		APIVersion: "v1alpha1",
	}

	APIServiceInstanceScopes = []string{"Environment"}
)

const APIServiceInstanceResourceName = "apiserviceinstances"

func APIServiceInstanceGVK() apiv1.GroupVersionKind {
	return _APIServiceInstanceGVK
}

func init() {
	apiv1.RegisterGVK(_APIServiceInstanceGVK, APIServiceInstanceScopes[0], APIServiceInstanceResourceName)
}

// APIServiceInstance Resource
type APIServiceInstance struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner           `json:"owner"`
	Spec  ApiServiceInstanceSpec `json:"spec"`
}

// NewAPIServiceInstance creates an empty *APIServiceInstance
func NewAPIServiceInstance(name, scopeName string) *APIServiceInstance {
	return &APIServiceInstance{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _APIServiceInstanceGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: APIServiceInstanceScopes[0],
				},
			},
		},
	}
}

// APIServiceInstanceFromInstanceArray converts a []*ResourceInstance to a []*APIServiceInstance
func APIServiceInstanceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APIServiceInstance, error) {
	newArray := make([]*APIServiceInstance, 0)
	for _, item := range fromArray {
		res := &APIServiceInstance{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APIServiceInstance, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APIServiceInstance to a ResourceInstance
func (res *APIServiceInstance) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = APIServiceInstanceGVK()
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

// FromInstance converts a ResourceInstance to a APIServiceInstance
func (res *APIServiceInstance) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *APIServiceInstance) MarshalJSON() ([]byte, error) {
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
func (res *APIServiceInstance) UnmarshalJSON(data []byte) error {
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
func (res *APIServiceInstance) PluralName() string {
	return APIServiceInstanceResourceName
}
