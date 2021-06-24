/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_K8SClusterGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "K8SCluster",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	K8SClusterScope = ""

	K8SClusterResourceName = "k8sclusters"
)

func K8SClusterGVK() apiv1.GroupVersionKind {
	return _K8SClusterGVK
}

func init() {
	apiv1.RegisterGVK(_K8SClusterGVK, K8SClusterScope, K8SClusterResourceName)
}

// K8SCluster Resource
type K8SCluster struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec K8SClusterSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a K8SCluster
func (res *K8SCluster) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	err := json.Unmarshal(ri.GetRawResource(), res)
	return err
}

// K8SClusterFromInstanceArray converts a []*ResourceInstance to a []*K8SCluster
func K8SClusterFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*K8SCluster, error) {
	newArray := make([]*K8SCluster, 0)
	for _, item := range fromArray {
		res := &K8SCluster{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*K8SCluster, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a K8SCluster to a ResourceInstance
func (res *K8SCluster) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = K8SClusterGVK()
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
