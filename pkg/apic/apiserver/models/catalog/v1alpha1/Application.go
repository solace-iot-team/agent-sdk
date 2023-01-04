/*
 * This file is automatically generated
 */

package catalog

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

var (
	ApplicationCtx log.ContextField = "application"

	_ApplicationGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Application",
		},
		APIVersion: "v1alpha1",
	}

	ApplicationScopes = []string{""}
)

const ApplicationResourceName = "applications"

func ApplicationGVK() apiv1.GroupVersionKind {
	return _ApplicationGVK
}

func init() {
	apiv1.RegisterGVK(_ApplicationGVK, ApplicationScopes[0], ApplicationResourceName)
	log.RegisterContextField(ApplicationCtx)
}

// Application Resource
type Application struct {
	apiv1.ResourceMeta
	Marketplace ApplicationMarketplace `json:"marketplace"`
	Owner       *apiv1.Owner           `json:"owner"`
	Spec        ApplicationSpec        `json:"spec"`
}

// NewApplication creates an empty *Application
func NewApplication(name string) *Application {
	return &Application{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ApplicationGVK,
		},
	}
}

// ApplicationFromInstanceArray converts a []*ResourceInstance to a []*Application
func ApplicationFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Application, error) {
	newArray := make([]*Application, 0)
	for _, item := range fromArray {
		res := &Application{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Application, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Application to a ResourceInstance
func (res *Application) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ApplicationGVK()
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

// FromInstance converts a ResourceInstance to a Application
func (res *Application) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *Application) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["marketplace"] = res.Marketplace
	out["owner"] = res.Owner
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *Application) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Marketplace
	if v, ok := aux.SubResources["marketplace"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "marketplace")
		err = json.Unmarshal(sr, &res.Marketplace)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *Application) PluralName() string {
	return ApplicationResourceName
}
