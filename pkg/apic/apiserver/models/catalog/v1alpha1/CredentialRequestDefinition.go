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
	CredentialRequestDefinitionCtx log.ContextField = "credentialRequestDefinition"

	_CredentialRequestDefinitionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "CredentialRequestDefinition",
		},
		APIVersion: "v1alpha1",
	}

	CredentialRequestDefinitionScopes = []string{"AuthorizationProfile"}
)

const CredentialRequestDefinitionResourceName = "credentialrequestdefinitions"

func CredentialRequestDefinitionGVK() apiv1.GroupVersionKind {
	return _CredentialRequestDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_CredentialRequestDefinitionGVK, CredentialRequestDefinitionScopes[0], CredentialRequestDefinitionResourceName)
	log.RegisterContextField(CredentialRequestDefinitionCtx)
}

// CredentialRequestDefinition Resource
type CredentialRequestDefinition struct {
	apiv1.ResourceMeta
	Owner      *apiv1.Owner                          `json:"owner"`
	References CredentialRequestDefinitionReferences `json:"references"`
	Spec       CredentialRequestDefinitionSpec       `json:"spec"`
	Webhooks   interface{}                           `json:"webhooks"`
}

// NewCredentialRequestDefinition creates an empty *CredentialRequestDefinition
func NewCredentialRequestDefinition(name, scopeName string) *CredentialRequestDefinition {
	return &CredentialRequestDefinition{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _CredentialRequestDefinitionGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: CredentialRequestDefinitionScopes[0],
				},
			},
		},
	}
}

// CredentialRequestDefinitionFromInstanceArray converts a []*ResourceInstance to a []*CredentialRequestDefinition
func CredentialRequestDefinitionFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*CredentialRequestDefinition, error) {
	newArray := make([]*CredentialRequestDefinition, 0)
	for _, item := range fromArray {
		res := &CredentialRequestDefinition{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*CredentialRequestDefinition, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a CredentialRequestDefinition to a ResourceInstance
func (res *CredentialRequestDefinition) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = CredentialRequestDefinitionGVK()
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

// FromInstance converts a ResourceInstance to a CredentialRequestDefinition
func (res *CredentialRequestDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *CredentialRequestDefinition) MarshalJSON() ([]byte, error) {
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
	out["references"] = res.References
	out["spec"] = res.Spec
	out["webhooks"] = res.Webhooks

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *CredentialRequestDefinition) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Webhooks
	if v, ok := aux.SubResources["webhooks"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "webhooks")
		err = json.Unmarshal(sr, &res.Webhooks)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *CredentialRequestDefinition) PluralName() string {
	return CredentialRequestDefinitionResourceName
}
