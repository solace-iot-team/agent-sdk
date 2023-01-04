/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// SupportContactSpecOfficeHoursPeriods  (catalog.v1alpha1.SupportContact)
type SupportContactSpecOfficeHoursPeriods struct {
	// Time in 24hr ISO 8601 extended format (hh:mm). Valid values are 00:00-24:00.
	OpenTime string `json:"openTime"`
	OpenDay  string `json:"openDay"`
	// Time in 24hr ISO 8601 extended format (hh:mm). Valid values are 00:00-24:00.
	CloseTime string `json:"closeTime"`
	CloseDay  string `json:"closeDay"`
}
