/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// ProductSpecAutoReleasePreviousReleases  (catalog.v1alpha1.Product)
type ProductSpecAutoReleasePreviousReleases struct {
	// Updates all prior non-archived releases to the deprecated state.
	UpdateState string `json:"updateState,omitempty"`
}
