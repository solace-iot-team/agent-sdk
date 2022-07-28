/*
 * This file is automatically generated
 */

package management

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ConsumerSubscriptionDefinitionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "ConsumerSubscriptionDefinition",
		},
		APIVersion: "v1alpha1",
	}

	ConsumerSubscriptionDefinitionScopes = []string{"Environment"}
)

const ConsumerSubscriptionDefinitionResourceName = "consumersubscriptiondefs"

func ConsumerSubscriptionDefinitionGVK() apiv1.GroupVersionKind {
	return _ConsumerSubscriptionDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_ConsumerSubscriptionDefinitionGVK, ConsumerSubscriptionDefinitionScopes[0], ConsumerSubscriptionDefinitionResourceName)
}

// ConsumerSubscriptionDefinition Resource
type ConsumerSubscriptionDefinition struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner                       `json:"owner"`
	Spec  ConsumerSubscriptionDefinitionSpec `json:"spec"`
}

// NewConsumerSubscriptionDefinition creates an empty *ConsumerSubscriptionDefinition
func NewConsumerSubscriptionDefinition(name, scopeName string) *ConsumerSubscriptionDefinition {
	return &ConsumerSubscriptionDefinition{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ConsumerSubscriptionDefinitionGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: ConsumerSubscriptionDefinitionScopes[0],
				},
			},
		},
	}
}

// ConsumerSubscriptionDefinitionFromInstanceArray converts a []*ResourceInstance to a []*ConsumerSubscriptionDefinition
func ConsumerSubscriptionDefinitionFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ConsumerSubscriptionDefinition, error) {
	newArray := make([]*ConsumerSubscriptionDefinition, 0)
	for _, item := range fromArray {
		res := &ConsumerSubscriptionDefinition{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ConsumerSubscriptionDefinition, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ConsumerSubscriptionDefinition to a ResourceInstance
func (res *ConsumerSubscriptionDefinition) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ConsumerSubscriptionDefinitionGVK()
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

// FromInstance converts a ResourceInstance to a ConsumerSubscriptionDefinition
func (res *ConsumerSubscriptionDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ConsumerSubscriptionDefinition) MarshalJSON() ([]byte, error) {
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
func (res *ConsumerSubscriptionDefinition) UnmarshalJSON(data []byte) error {
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
func (res *ConsumerSubscriptionDefinition) PluralName() string {
	return ConsumerSubscriptionDefinitionResourceName
}
