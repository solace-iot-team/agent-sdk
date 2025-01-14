/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// CategorySpec  (catalog.v1alpha1.Category)
type CategorySpec struct {
	// Markdown representing the category description.
	Description string `json:"description,omitempty"`
	// GENERATE: The following code has been modified after code generation
	Restriction interface{} `json:"restriction,omitempty"`
}
