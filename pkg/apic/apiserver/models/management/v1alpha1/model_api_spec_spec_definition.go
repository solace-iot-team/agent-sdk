/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiSpecSpecDefinition  (management.v1alpha1.APISpec)
type ApiSpecSpecDefinition struct {
	// The type of the api specification.
	Type string `json:"type,omitempty"`
	// Base64 encoded value of the api specification.
	Value string `json:"value,omitempty"`
	// Consistent hash of the value.
	Hash string `json:"hash,omitempty"`
}
