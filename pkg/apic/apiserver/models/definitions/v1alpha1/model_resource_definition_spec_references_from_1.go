/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceDefinitionSpecReferencesFrom1 struct for ResourceDefinitionSpecReferencesFrom1
type ResourceDefinitionSpecReferencesFrom1 struct {
	// Defines the subResource referencing this resource. Omit for non subResource references.
	SubResourceName string `json:"subResourceName,omitempty"`
}