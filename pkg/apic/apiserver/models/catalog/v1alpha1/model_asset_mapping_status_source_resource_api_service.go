/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AssetMappingStatusSourceResourceApiService  (catalog.v1alpha1.AssetMapping)
type AssetMappingStatusSourceResourceApiService struct {
	Ref string `json:"ref,omitempty"`
	// GENERATE: The following code has been modified after code generation
	OperationType AssetMappingStatusOperationType `json:"operationType,omitempty"`
}
