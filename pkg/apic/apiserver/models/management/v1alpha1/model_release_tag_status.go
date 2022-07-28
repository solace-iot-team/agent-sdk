/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ReleaseTagStatus  (management.v1alpha1.ReleaseTag)
type ReleaseTagStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	// GENERATE: The following code has been modified after code generation
	Reasons []interface{} `json:"reasons,omitempty"`
}
