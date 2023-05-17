/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiSpecLintingJobSpec  (management.v1.APISpecLintingJob)
type ApiSpecLintingJobSpec struct {
	// Reference to Amplify Central APIServiceRevision
	ApiServiceRevision string `json:"apiServiceRevision"`
	// Reference to Amplify Central APISpecLintingRuleset
	ApiSpecLintingRuleset string `json:"apiSpecLintingRuleset"`
}
