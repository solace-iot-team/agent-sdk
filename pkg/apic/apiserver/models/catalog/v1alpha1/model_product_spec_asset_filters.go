/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductSpecAssetFilters Filters what AssetRelease the product points to
type ProductSpecAssetFilters struct {
	// The stages that are included in the Product.
	Stages []string `json:"stages,omitempty"`
	// The Asset Release version to use. Examples:   - 1.0.1 for a specific asset release version   - 1.* for all minor and patch versions of version 1   - 1.2.* for all the patch version for version 1.2
	Version string `json:"version,omitempty"`
}