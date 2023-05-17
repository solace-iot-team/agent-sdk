/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// EnvironmentSpecCompliance Compliance for the Environment. (management.v1alpha1.Environment)
type EnvironmentSpecCompliance struct {
	// Reference to Amplify Central design Ruleset
	Design string `json:"design,omitempty"`
	// Reference to Amplify Central security Ruleset
	Security string `json:"security,omitempty"`
}
