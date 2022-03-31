/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetRequestApprovalState  (catalog.v1alpha1.AssetRequest)
type AssetRequestApprovalState struct {
	Name string `json:"name"`
	// Additional info on the state.
	Reason string `json:"reason,omitempty"`
}
