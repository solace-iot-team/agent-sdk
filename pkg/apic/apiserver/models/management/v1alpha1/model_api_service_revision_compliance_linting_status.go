/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiServiceRevisionComplianceLintingStatus Provides the current linting job state and the final result.
type ApiServiceRevisionComplianceLintingStatus struct {
	// The state of the compliance job.
	State             string `json:"state,omitempty"`
	ApiSpecLintingJob string `json:"apiSpecLintingJob,omitempty"`
	// Ruleset logical name.
	Ruleset string `json:"ruleset,omitempty"`
	// File name of the APISpecLintingRuleset.
	RulesetFileName string                                          `json:"rulesetFileName,omitempty"`
	Result          ApiServiceRevisionComplianceLintingStatusResult `json:"result,omitempty"`
}
