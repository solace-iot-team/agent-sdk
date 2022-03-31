/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetRequestStatus  (catalog.v1alpha1.AssetRequest)
type AssetRequestStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	Reasons []AssetRequestStatusReasons `json:"reasons,omitempty"`
}
