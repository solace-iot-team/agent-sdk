/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// AmplifyJwtAuthRuleSpecFrom  (management.v1alpha1.AmplifyJWTAuthRule)
type AmplifyJwtAuthRuleSpecFrom struct {
	// Where to look for the token.
	In   string `json:"in,omitempty"`
	Name string `json:"name"`
	// The token prefix, e.g. \"Bearer \".
	Prefix string `json:"prefix,omitempty"`
}
