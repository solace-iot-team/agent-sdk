/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceDefinitionSpecReferencesToResources struct for ResourceDefinitionSpecReferencesToResources
type ResourceDefinitionSpecReferencesToResources struct {
	// Defines the kind of the resource.
	Group string `json:"group,omitempty"`
	// Defines the kind of the resource.
	Kind string `json:"kind,omitempty"`
	// Defines the scope kind of the resource. Omit for unscoped resources.
	ScopeKind string `json:"scopeKind,omitempty"`
	// The type of the reference.
	Types []string                             `json:"types,omitempty"`
	From  ResourceDefinitionSpecReferencesFrom `json:"from,omitempty"`
}
