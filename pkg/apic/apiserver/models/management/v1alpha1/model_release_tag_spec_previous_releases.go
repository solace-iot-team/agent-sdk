/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ReleaseTagSpecPreviousReleases  (management.v1alpha1.ReleaseTag)
type ReleaseTagSpecPreviousReleases struct {
	// Updates all prior non-archived releases to the deprecated state.
	UpdateState string `json:"updateState,omitempty"`
}
