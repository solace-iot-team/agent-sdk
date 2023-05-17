/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ApiSpecLintingJobResultLocation The location of the linting result. (management.v1.APISpecLintingJob)
type ApiSpecLintingJobResultLocation struct {
	// The 1-based line number in the API Specification File.
	Line int32 `json:"line"`
	// The 1-based character number in the API Specification File.
	Character int32 `json:"character"`
}
