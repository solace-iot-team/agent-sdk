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
	DiscoveryAgentCtx log.ContextField = "discoveryAgent"

	_DiscoveryAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "DiscoveryAgent",
		},
		APIVersion: "v1alpha1",
	}

	DiscoveryAgentScopes = []string{"Environment"}
)

const (
	DiscoveryAgentResourceName             = "discoveryagents"
	DiscoveryAgentDataplaneSubResourceName = "dataplane"
	DiscoveryAgentStatusSubResourceName    = "status"
)

func DiscoveryAgentGVK() apiv1.GroupVersionKind {
	return _DiscoveryAgentGVK
}

func init() {
	apiv1.RegisterGVK(_DiscoveryAgentGVK, DiscoveryAgentScopes[0], DiscoveryAgentResourceName)
	log.RegisterContextField(DiscoveryAgentCtx)
}

// DiscoveryAgent Resource
type DiscoveryAgent struct {
	apiv1.ResourceMeta
	Dataplane DiscoveryAgentDataplane `json:"dataplane"`
	Owner     *apiv1.Owner            `json:"owner"`
	Spec      DiscoveryAgentSpec      `json:"spec"`
	Status    DiscoveryAgentStatus    `json:"status"`
}

// NewDiscoveryAgent creates an empty *DiscoveryAgent
func NewDiscoveryAgent(name, scopeName string) *DiscoveryAgent {
	return &DiscoveryAgent{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _DiscoveryAgentGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: DiscoveryAgentScopes[0],
				},
			},
		},
	}
}

// DiscoveryAgentFromInstanceArray converts a []*ResourceInstance to a []*DiscoveryAgent
func DiscoveryAgentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*DiscoveryAgent, error) {
	newArray := make([]*DiscoveryAgent, 0)
	for _, item := range fromArray {
		res := &DiscoveryAgent{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*DiscoveryAgent, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a DiscoveryAgent to a ResourceInstance
func (res *DiscoveryAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = DiscoveryAgentGVK()
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

// FromInstance converts a ResourceInstance to a DiscoveryAgent
func (res *DiscoveryAgent) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *DiscoveryAgent) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["dataplane"] = res.Dataplane
	out["owner"] = res.Owner
	out["spec"] = res.Spec
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *DiscoveryAgent) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Dataplane
	if v, ok := aux.SubResources["dataplane"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "dataplane")
		err = json.Unmarshal(sr, &res.Dataplane)
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
func (res *DiscoveryAgent) PluralName() string {
	return DiscoveryAgentResourceName
}
