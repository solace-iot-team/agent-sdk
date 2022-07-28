/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// QuotaSpecLimitTypeLoose struct for QuotaSpecLimitTypeLoose
type QuotaSpecLimitTypeLoose struct {
	Type string `json:"type"`
	// The limit of the unit that is provided.
	// GENERATE: The following code has been modified after code generation
	Value    float64                         `json:"value"`
	Overages QuotaSpecLimitTypeLooseOverages `json:"overages"`
}
