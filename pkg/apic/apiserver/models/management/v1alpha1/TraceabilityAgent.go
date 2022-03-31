/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_TraceabilityAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "TraceabilityAgent",
		},
		APIVersion: "v1alpha1",
	}

	TraceabilityAgentScopes = []string{"Environment"}
)

const TraceabilityAgentResourceName = "traceabilityagents"

func TraceabilityAgentGVK() apiv1.GroupVersionKind {
	return _TraceabilityAgentGVK
}

func init() {
	apiv1.RegisterGVK(_TraceabilityAgentGVK, TraceabilityAgentScopes[0], TraceabilityAgentResourceName)
}

// TraceabilityAgent Resource
type TraceabilityAgent struct {
	apiv1.ResourceMeta
	Owner  *apiv1.Owner            `json:"owner"`
	Spec   TraceabilityAgentSpec   `json:"spec"`
	Status TraceabilityAgentStatus `json:"status"`
}

// NewTraceabilityAgent creates an empty *TraceabilityAgent
func NewTraceabilityAgent(name, scopeName string) *TraceabilityAgent {
	return &TraceabilityAgent{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _TraceabilityAgentGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: TraceabilityAgentScopes[0],
				},
			},
		},
	}
}

// TraceabilityAgentFromInstanceArray converts a []*ResourceInstance to a []*TraceabilityAgent
func TraceabilityAgentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*TraceabilityAgent, error) {
	newArray := make([]*TraceabilityAgent, 0)
	for _, item := range fromArray {
		res := &TraceabilityAgent{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*TraceabilityAgent, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a TraceabilityAgent to a ResourceInstance
func (res *TraceabilityAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = TraceabilityAgentGVK()
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

// FromInstance converts a ResourceInstance to a TraceabilityAgent
func (res *TraceabilityAgent) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *TraceabilityAgent) MarshalJSON() ([]byte, error) {
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
func (res *TraceabilityAgent) UnmarshalJSON(data []byte) error {
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
func (res *TraceabilityAgent) PluralName() string {
	return TraceabilityAgentResourceName
}
