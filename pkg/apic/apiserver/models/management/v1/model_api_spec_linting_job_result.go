/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiSpecLintingJobResult  (management.v1.APISpecLintingJob)
type ApiSpecLintingJobResult struct {
	// The API Specification Linting Result details.
	Details []ApiSpecLintingJobResultDetails `json:"details"`
	Summary ApiSpecLintingJobResultSummary   `json:"summary"`
	// Reference to the APISpecLintingRuleset revision
	ApiSpecLintingRulesetRevision int32 `json:"apiSpecLintingRulesetRevision,omitempty"`
	// Set the value to true if the linting result details count has reached the threshold
	DetailsThresholdExceeded bool `json:"detailsThresholdExceeded,omitempty"`
}
