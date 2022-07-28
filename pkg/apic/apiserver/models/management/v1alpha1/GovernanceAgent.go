/*
 * This file is automatically generated
 */

package management

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_GovernanceAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "GovernanceAgent",
		},
		APIVersion: "v1alpha1",
	}

	GovernanceAgentScopes = []string{"Environment"}
)

const GovernanceAgentResourceName = "governanceagents"

func GovernanceAgentGVK() apiv1.GroupVersionKind {
	return _GovernanceAgentGVK
}

func init() {
	apiv1.RegisterGVK(_GovernanceAgentGVK, GovernanceAgentScopes[0], GovernanceAgentResourceName)
}

// GovernanceAgent Resource
type GovernanceAgent struct {
	apiv1.ResourceMeta
	Agentconfigstatus GovernanceAgentAgentconfigstatus `json:"agentconfigstatus"`
	Owner             *apiv1.Owner                     `json:"owner"`
	Spec              GovernanceAgentSpec              `json:"spec"`
	Status            GovernanceAgentStatus            `json:"status"`
}

// NewGovernanceAgent creates an empty *GovernanceAgent
func NewGovernanceAgent(name, scopeName string) *GovernanceAgent {
	return &GovernanceAgent{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _GovernanceAgentGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: GovernanceAgentScopes[0],
				},
			},
		},
	}
}

// GovernanceAgentFromInstanceArray converts a []*ResourceInstance to a []*GovernanceAgent
func GovernanceAgentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*GovernanceAgent, error) {
	newArray := make([]*GovernanceAgent, 0)
	for _, item := range fromArray {
		res := &GovernanceAgent{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*GovernanceAgent, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a GovernanceAgent to a ResourceInstance
func (res *GovernanceAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = GovernanceAgentGVK()
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

// FromInstance converts a ResourceInstance to a GovernanceAgent
func (res *GovernanceAgent) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *GovernanceAgent) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["agentconfigstatus"] = res.Agentconfigstatus
	out["owner"] = res.Owner
	out["spec"] = res.Spec
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *GovernanceAgent) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Agentconfigstatus
	if v, ok := aux.SubResources["agentconfigstatus"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "agentconfigstatus")
		err = json.Unmarshal(sr, &res.Agentconfigstatus)
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
		err = json.Unmarshal(sr, &res.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *GovernanceAgent) PluralName() string {
	return GovernanceAgentResourceName
}
