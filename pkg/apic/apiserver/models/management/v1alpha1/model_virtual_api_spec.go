/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// VirtualApiSpec The URL for a OAS specification.
type VirtualApiSpec struct {
	ApiType string `json:"apiType,omitempty"`
	Url     string `json:"url"`
}
