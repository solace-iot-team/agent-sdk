/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package definitions

// ResourceDefinitionSpecReferences  (definitions.v1alpha1.ResourceDefinition)
type ResourceDefinitionSpecReferences struct {
	// A list of resources that the current resources has references to.
	ToResources []ResourceDefinitionSpecReferencesToResources `json:"toResources,omitempty"`
	// A list of resources that the current resources is beging referenced from.
	FromResources []ResourceDefinitionSpecReferencesFromResources `json:"fromResources,omitempty"`
}
