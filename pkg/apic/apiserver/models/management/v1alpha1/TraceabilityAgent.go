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
)

const (
	TraceabilityAgentScope = "Environment"

	TraceabilityAgentResource = "traceabilityagents"
)

func TraceabilityAgentGVK() apiv1.GroupVersionKind {
	return _TraceabilityAgentGVK
}

func init() {
	apiv1.RegisterGVK(_TraceabilityAgentGVK, TraceabilityAgentScope, TraceabilityAgentResource)
}

// TraceabilityAgent Resource
type TraceabilityAgent struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec TraceabilityAgentSpec `json:"spec"`

	Status TraceabilityAgentStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a TraceabilityAgent
func (res *TraceabilityAgent) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &TraceabilityAgentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = TraceabilityAgent{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a TraceabilityAgent to a ResourceInstance
func (res *TraceabilityAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = TraceabilityAgentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
