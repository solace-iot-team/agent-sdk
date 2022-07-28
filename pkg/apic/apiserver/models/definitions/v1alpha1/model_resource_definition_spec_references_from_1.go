/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package definitions

// ResourceDefinitionSpecReferencesFrom1  (definitions.v1alpha1.ResourceDefinition)
type ResourceDefinitionSpecReferencesFrom1 struct {
	// Defines the subResource referencing this resource. Omit for non subResource references.
	SubResourceName string `json:"subResourceName,omitempty"`
}
