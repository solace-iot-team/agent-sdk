/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiServiceCompliance  (management.v1alpha1.APIService)
type ApiServiceCompliance struct {
	// GENERATE: The following code has been modified after code generation
	Design ApiServiceComplianceLintingStatus `json:"design,omitempty"`
	// GENERATE: The following code has been modified after code generation
	Security ApiServiceComplianceLintingStatus `json:"security,omitempty"`
}
