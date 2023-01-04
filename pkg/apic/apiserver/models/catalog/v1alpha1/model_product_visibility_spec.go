/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// ProductVisibilitySpec  (catalog.v1alpha1.ProductVisibility)
type ProductVisibilitySpec struct {
	// Defines where the visibility settings apply.
	Products []ProductVisibilitySpecProducts `json:"products"`
	// Determines if the list of subjects should be excluded from the product visibility.
	Exclude bool `json:"exclude,omitempty"`
	// GENERATE: The following code has been modified after code generation
	Subjects []interface{} `json:"subjects,omitempty"`
}
