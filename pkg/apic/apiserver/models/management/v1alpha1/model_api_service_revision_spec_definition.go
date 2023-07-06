/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiServiceRevisionSpecDefinition The api specification details. (management.v1alpha1.APIServiceRevision)
type ApiServiceRevisionSpecDefinition struct {
	// The type of the api specification.
	Type string `json:"type,omitempty"`
	// Base64 encoded value of the api specification.
	Value string `json:"value,omitempty"`
	// content-type of the spec.
	ContentType string `json:"contentType,omitempty"`
	// The version of the api specification. Will be extracted from \"value\" if not assigned.
	Version string `json:"version,omitempty"`
}
