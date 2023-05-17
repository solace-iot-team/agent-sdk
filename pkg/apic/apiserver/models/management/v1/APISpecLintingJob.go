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
	APISpecLintingJobCtx log.ContextField = "apiSpecLintingJob"

	_APISpecLintingJobGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APISpecLintingJob",
		},
		APIVersion: "v1",
	}

	APISpecLintingJobScopes = []string{"Environment"}
)

const APISpecLintingJobResourceName = "apispeclintingjobs"

func APISpecLintingJobGVK() apiv1.GroupVersionKind {
	return _APISpecLintingJobGVK
}

func init() {
	apiv1.RegisterGVK(_APISpecLintingJobGVK, APISpecLintingJobScopes[0], APISpecLintingJobResourceName)
	log.RegisterContextField(APISpecLintingJobCtx)
}

// APISpecLintingJob Resource
type APISpecLintingJob struct {
	apiv1.ResourceMeta
	Archived interface{}             `json:"archived"`
	Owner    *apiv1.Owner            `json:"owner"`
	Result   ApiSpecLintingJobResult `json:"result"`
	Spec     ApiSpecLintingJobSpec   `json:"spec"`
	State    ApiSpecLintingJobState  `json:"state"`
}

// NewAPISpecLintingJob creates an empty *APISpecLintingJob
func NewAPISpecLintingJob(name, scopeName string) *APISpecLintingJob {
	return &APISpecLintingJob{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _APISpecLintingJobGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: APISpecLintingJobScopes[0],
				},
			},
		},
	}
}

// APISpecLintingJobFromInstanceArray converts a []*ResourceInstance to a []*APISpecLintingJob
func APISpecLintingJobFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APISpecLintingJob, error) {
	newArray := make([]*APISpecLintingJob, 0)
	for _, item := range fromArray {
		res := &APISpecLintingJob{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APISpecLintingJob, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APISpecLintingJob to a ResourceInstance
func (res *APISpecLintingJob) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = APISpecLintingJobGVK()
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

// FromInstance converts a ResourceInstance to a APISpecLintingJob
func (res *APISpecLintingJob) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *APISpecLintingJob) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["archived"] = res.Archived
	out["owner"] = res.Owner
	out["result"] = res.Result
	out["spec"] = res.Spec
	out["state"] = res.State

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *APISpecLintingJob) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Archived
	if v, ok := aux.SubResources["archived"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "archived")
		err = json.Unmarshal(sr, &res.Archived)
		if err != nil {
			return err
		}
	}

	// marshalling subresource Result
	if v, ok := aux.SubResources["result"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "result")
		err = json.Unmarshal(sr, &res.Result)
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

	return nil
}

// PluralName returns the plural name of the resource
func (res *APISpecLintingJob) PluralName() string {
	return APISpecLintingJobResourceName
}
