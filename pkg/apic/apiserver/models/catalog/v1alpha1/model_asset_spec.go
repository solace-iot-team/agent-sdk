/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetSpec struct for AssetSpec
type AssetSpec struct {
	// description of the asset.
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}
