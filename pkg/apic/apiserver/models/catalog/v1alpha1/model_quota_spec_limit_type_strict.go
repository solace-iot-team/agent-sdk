/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// QuotaSpecLimitTypeStrict struct for QuotaSpecLimitTypeStrict
type QuotaSpecLimitTypeStrict struct {
	Type string `json:"type"`
	// The limit of the unit that is provided.
	// GENERATE: The following code has been modified after code generation
	Value float64 `json:"value"`
}
