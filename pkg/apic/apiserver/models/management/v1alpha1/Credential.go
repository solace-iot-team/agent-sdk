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
	CredentialCtx log.ContextField = "credential"

	_CredentialGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Credential",
		},
		APIVersion: "v1alpha1",
	}

	CredentialScopes = []string{"Environment"}
)

const (
	CredentialResourceName              = "credentials"
	CredentialDataSubResourceName       = "data"
	CredentialPoliciesSubResourceName   = "policies"
	CredentialReferencesSubResourceName = "references"
	CredentialStateSubResourceName      = "state"
	CredentialStatusSubResourceName     = "status"
)

func CredentialGVK() apiv1.GroupVersionKind {
	return _CredentialGVK
}

func init() {
	apiv1.RegisterGVK(_CredentialGVK, CredentialScopes[0], CredentialResourceName)
	log.RegisterContextField(CredentialCtx)
}

// Credential Resource
type Credential struct {
	apiv1.ResourceMeta
	Data       interface{}          `json:"data"`
	Owner      *apiv1.Owner         `json:"owner"`
	Policies   CredentialPolicies   `json:"policies"`
	References CredentialReferences `json:"references"`
	Spec       CredentialSpec       `json:"spec"`
	State      CredentialState      `json:"state"`
	// Status     CredentialStatus     `json:"status"`
	Status *apiv1.ResourceStatus `json:"status"`
}

// NewCredential creates an empty *Credential
func NewCredential(name, scopeName string) *Credential {
	return &Credential{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _CredentialGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: CredentialScopes[0],
				},
			},
		},
	}
}

// CredentialFromInstanceArray converts a []*ResourceInstance to a []*Credential
func CredentialFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Credential, error) {
	newArray := make([]*Credential, 0)
	for _, item := range fromArray {
		res := &Credential{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Credential, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Credential to a ResourceInstance
func (res *Credential) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = CredentialGVK()
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

// FromInstance converts a ResourceInstance to a Credential
func (res *Credential) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *Credential) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["data"] = res.Data
	out["owner"] = res.Owner
	out["policies"] = res.Policies
	out["references"] = res.References
	out["spec"] = res.Spec
	out["state"] = res.State
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Credential) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Data
	if v, ok := aux.SubResources["data"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "data")
		err = json.Unmarshal(sr, &res.Data)
		if err != nil {
			return err
		}
	}

	// marshalling subresource Policies
	if v, ok := aux.SubResources["policies"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "policies")
		err = json.Unmarshal(sr, &res.Policies)
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

	// marshalling subresource Status
	if v, ok := aux.SubResources["status"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "status")
		// err = json.Unmarshal(sr, &res.Status)
		res.Status = &apiv1.ResourceStatus{}
		err = json.Unmarshal(sr, res.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *Credential) PluralName() string {
	return CredentialResourceName
}
