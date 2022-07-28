/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// AssetMappingTemplateSpec  (management.v1alpha1.AssetMappingTemplate)
type AssetMappingTemplateSpec struct {
	// A list of filters for the API Service resource on which the template applies.
	Filters []AssetMappingTemplateSpecFilters `json:"filters,omitempty"`
}
