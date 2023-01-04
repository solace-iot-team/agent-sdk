/*
 * This file is automatically generated
 */

package management

import (
	"encoding/json"
	"fmt"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/util/log"
)

var (
	GraphQLDocumentCtx log.ContextField = "graphQLDocument"

	_GraphQLDocumentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "GraphQLDocument",
		},
		APIVersion: "v1alpha1",
	}

	GraphQLDocumentScopes = []string{"VirtualAPI", "VirtualAPIRelease"}
)

const GraphQLDocumentResourceName = "graphqldocuments"

func GraphQLDocumentGVK() apiv1.GroupVersionKind {
	return _GraphQLDocumentGVK
}

func init() {
	apiv1.RegisterGVK(_GraphQLDocumentGVK, GraphQLDocumentScopes[0], GraphQLDocumentResourceName)
	log.RegisterContextField(GraphQLDocumentCtx)
}

// GraphQLDocument Resource
type GraphQLDocument struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner `json:"owner"`
	Spec  interface{}  `json:"spec"`
}

// NewGraphQLDocument creates an empty *GraphQLDocument
func NewGraphQLDocument(name, scopeKind, scopeName string) (*GraphQLDocument, error) {
	validScope := false
	for _, s := range GraphQLDocumentScopes {
		if scopeKind == s {
			validScope = true
			break
		}
	}
	if !validScope {
		return nil, fmt.Errorf("scope '%s' not valid for GraphQLDocument kind", scopeKind)
	}

	return &GraphQLDocument{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _GraphQLDocumentGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: scopeKind,
				},
			},
		},
	}, nil
}

// GraphQLDocumentFromInstanceArray converts a []*ResourceInstance to a []*GraphQLDocument
func GraphQLDocumentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*GraphQLDocument, error) {
	newArray := make([]*GraphQLDocument, 0)
	for _, item := range fromArray {
		res := &GraphQLDocument{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*GraphQLDocument, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a GraphQLDocument to a ResourceInstance
func (res *GraphQLDocument) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = GraphQLDocumentGVK()
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

// FromInstance converts a ResourceInstance to a GraphQLDocument
func (res *GraphQLDocument) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *GraphQLDocument) MarshalJSON() ([]byte, error) {
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
func (res *GraphQLDocument) UnmarshalJSON(data []byte) error {
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
func (res *GraphQLDocument) PluralName() string {
	return GraphQLDocumentResourceName
}
