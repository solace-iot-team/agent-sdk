/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// EnvironmentCompliancetasks Subresource that is only accessible by the backend. Used by compliance-controller to auto-start/cancel linting jobs based on changes made to \"spec.compliance\" properties detected via mutation hook. (management.v1alpha1.Environment)
type EnvironmentCompliancetasks struct {
	// GENERATE: The following code has been modified after code generation
	Design EnvironmentCompliancetasksLinting `json:"design,omitempty"`
	// GENERATE: The following code has been modified after code generation
	Security EnvironmentCompliancetasksLinting `json:"security,omitempty"`
}
