/*
 * This file is automatically generated
 */

package v1alpha1

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
)

const (
	GovernanceAgentScope = "Environment"

	GovernanceAgentResource = "governanceagents"
)

func GovernanceAgentGVK() apiv1.GroupVersionKind {
	return _GovernanceAgentGVK
}

func init() {
	apiv1.RegisterGVK(_GovernanceAgentGVK, GovernanceAgentScope, GovernanceAgentResource)
}

// GovernanceAgent Resource
type GovernanceAgent struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec GovernanceAgentSpec `json:"spec"`

	Status GovernanceAgentStatus `json:"status"`

	RuntimeConfig map[string]interface{}
}

// FromInstance converts a ResourceInstance to a GovernanceAgent
func (res *GovernanceAgent) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &GovernanceAgentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = GovernanceAgent{ResourceMeta: ri.ResourceMeta, Spec: *spec, RuntimeConfig: ri.SubResources["runtimeconfig"].(map[string]interface{})}
	return err
}

// AsInstance converts a GovernanceAgent to a ResourceInstance
func (res *GovernanceAgent) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = GovernanceAgentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
