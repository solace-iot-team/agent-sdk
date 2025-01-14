/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// ReleaseTagSpec  (catalog.v1alpha1.ReleaseTag)
type ReleaseTagSpec struct {
	// Description of the Release Tag.
	Description      string                         `json:"description,omitempty"`
	ReleaseType      string                         `json:"releaseType"`
	PreviousReleases ReleaseTagSpecPreviousReleases `json:"previousReleases,omitempty"`
}
