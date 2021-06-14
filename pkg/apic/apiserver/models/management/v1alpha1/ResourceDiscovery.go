/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceDiscoveryGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "ResourceDiscovery",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ResourceDiscoveryScope = "K8SCluster"

	ResourceDiscoveryResource = "resourcediscoveries"
)

func ResourceDiscoveryGVK() apiv1.GroupVersionKind {
	return _ResourceDiscoveryGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDiscoveryGVK, ResourceDiscoveryScope, ResourceDiscoveryResource)
}

// ResourceDiscovery Resource
type ResourceDiscovery struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec ResourceDiscoverySpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a ResourceDiscovery
func (res *ResourceDiscovery) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ResourceDiscoverySpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = ResourceDiscovery{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a ResourceDiscovery to a ResourceInstance
func (res *ResourceDiscovery) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = ResourceDiscoveryGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
