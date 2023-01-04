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
	K8SResourceCtx log.ContextField = "k8SResource"

	_K8SResourceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "K8SResource",
		},
		APIVersion: "v1alpha1",
	}

	K8SResourceScopes = []string{"K8SCluster"}
)

const K8SResourceResourceName = "k8sresources"

func K8SResourceGVK() apiv1.GroupVersionKind {
	return _K8SResourceGVK
}

func init() {
	apiv1.RegisterGVK(_K8SResourceGVK, K8SResourceScopes[0], K8SResourceResourceName)
	log.RegisterContextField(K8SResourceCtx)
}

// K8SResource Resource
type K8SResource struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner    `json:"owner"`
	Spec  K8SResourceSpec `json:"spec"`
}

// NewK8SResource creates an empty *K8SResource
func NewK8SResource(name, scopeName string) *K8SResource {
	return &K8SResource{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _K8SResourceGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: K8SResourceScopes[0],
				},
			},
		},
	}
}

// K8SResourceFromInstanceArray converts a []*ResourceInstance to a []*K8SResource
func K8SResourceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*K8SResource, error) {
	newArray := make([]*K8SResource, 0)
	for _, item := range fromArray {
		res := &K8SResource{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*K8SResource, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a K8SResource to a ResourceInstance
func (res *K8SResource) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = K8SResourceGVK()
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

// FromInstance converts a ResourceInstance to a K8SResource
func (res *K8SResource) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *K8SResource) MarshalJSON() ([]byte, error) {
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
func (res *K8SResource) UnmarshalJSON(data []byte) error {
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
func (res *K8SResource) PluralName() string {
	return K8SResourceResourceName
}
