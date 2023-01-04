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
	APIServiceCtx log.ContextField = "apiService"

	_APIServiceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APIService",
		},
		APIVersion: "v1alpha1",
	}

	APIServiceScopes = []string{"Environment"}
)

const APIServiceResourceName = "apiservices"

func APIServiceGVK() apiv1.GroupVersionKind {
	return _APIServiceGVK
}

func init() {
	apiv1.RegisterGVK(_APIServiceGVK, APIServiceScopes[0], APIServiceResourceName)
	log.RegisterContextField(APIServiceCtx)
}

// APIService Resource
type APIService struct {
	apiv1.ResourceMeta
	Details ApiServiceDetails `json:"details"`
	Owner   *apiv1.Owner      `json:"owner"`
	Spec    ApiServiceSpec    `json:"spec"`
	// Status  ApiServiceStatus  `json:"status"`
	Status *apiv1.ResourceStatus `json:"status"`
}

// NewAPIService creates an empty *APIService
func NewAPIService(name, scopeName string) *APIService {
	return &APIService{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _APIServiceGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: APIServiceScopes[0],
				},
			},
		},
	}
}

// APIServiceFromInstanceArray converts a []*ResourceInstance to a []*APIService
func APIServiceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APIService, error) {
	newArray := make([]*APIService, 0)
	for _, item := range fromArray {
		res := &APIService{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APIService, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APIService to a ResourceInstance
func (res *APIService) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = APIServiceGVK()
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

// FromInstance converts a ResourceInstance to a APIService
func (res *APIService) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *APIService) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["details"] = res.Details
	out["owner"] = res.Owner
	out["spec"] = res.Spec
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *APIService) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Details
	if v, ok := aux.SubResources["details"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "details")
		err = json.Unmarshal(sr, &res.Details)
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
func (res *APIService) PluralName() string {
	return APIServiceResourceName
}
