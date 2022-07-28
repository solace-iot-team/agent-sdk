/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AssetMappingStatusResource The resources that were impacted with the trigger of asset mapping. (catalog.v1alpha1.AssetMapping)
type AssetMappingStatusResource struct {
	AssetResource AssetMappingStatusResourceAssetResource `json:"assetResource,omitempty"`
	Stage         AssetMappingStatusResourceStage         `json:"stage,omitempty"`
}
