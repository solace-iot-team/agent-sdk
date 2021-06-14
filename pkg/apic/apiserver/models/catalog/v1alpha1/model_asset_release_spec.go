/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetReleaseSpec struct for AssetReleaseSpec
type AssetReleaseSpec struct {
	// Description of the asset release.
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	// version of the asset release.
	Version string `json:"version,omitempty"`
	Asset   string `json:"asset"`
}
