/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// QuotaSpecFixedPricingType struct for QuotaSpecFixedPricingType
type QuotaSpecFixedPricingType struct {
	Type     string `json:"type"`
	Interval string `json:"interval"`
	// GENERATE: The following code has been modified after code generation
	Limit interface{} `json:"limit"`
}
