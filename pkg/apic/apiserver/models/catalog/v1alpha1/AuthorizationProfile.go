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
	AuthorizationProfileCtx log.ContextField = "authorizationProfile"

	_AuthorizationProfileGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AuthorizationProfile",
		},
		APIVersion: "v1alpha1",
	}

	AuthorizationProfileScopes = []string{""}
)

const (
	AuthorizationProfileResourceName              = "authprofiles"
	AuthorizationProfilePoliciesSubResourceName   = "policies"
	AuthorizationProfileReferencesSubResourceName = "references"
)

func AuthorizationProfileGVK() apiv1.GroupVersionKind {
	return _AuthorizationProfileGVK
}

func init() {
	apiv1.RegisterGVK(_AuthorizationProfileGVK, AuthorizationProfileScopes[0], AuthorizationProfileResourceName)
	log.RegisterContextField(AuthorizationProfileCtx)
}

// AuthorizationProfile Resource
type AuthorizationProfile struct {
	apiv1.ResourceMeta
	Owner      *apiv1.Owner                   `json:"owner"`
	Policies   AuthorizationProfilePolicies   `json:"policies"`
	References AuthorizationProfileReferences `json:"references"`
	Spec       AuthorizationProfileSpec       `json:"spec"`
}

// NewAuthorizationProfile creates an empty *AuthorizationProfile
func NewAuthorizationProfile(name string) *AuthorizationProfile {
	return &AuthorizationProfile{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _AuthorizationProfileGVK,
		},
	}
}

// AuthorizationProfileFromInstanceArray converts a []*ResourceInstance to a []*AuthorizationProfile
func AuthorizationProfileFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AuthorizationProfile, error) {
	newArray := make([]*AuthorizationProfile, 0)
	for _, item := range fromArray {
		res := &AuthorizationProfile{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AuthorizationProfile, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AuthorizationProfile to a ResourceInstance
func (res *AuthorizationProfile) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AuthorizationProfileGVK()
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

// FromInstance converts a ResourceInstance to a AuthorizationProfile
func (res *AuthorizationProfile) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *AuthorizationProfile) MarshalJSON() ([]byte, error) {
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
	out["policies"] = res.Policies
	out["references"] = res.References
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *AuthorizationProfile) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Policies
	if v, ok := aux.SubResources["policies"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "policies")
		err = json.Unmarshal(sr, &res.Policies)
		if err != nil {
			return err
		}
	}

	// marshalling subresource References
	if v, ok := aux.SubResources["references"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "references")
		err = json.Unmarshal(sr, &res.References)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *AuthorizationProfile) PluralName() string {
	return AuthorizationProfileResourceName
}
