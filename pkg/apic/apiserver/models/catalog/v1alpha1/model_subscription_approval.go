/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SubscriptionApproval  (catalog.v1alpha1.Subscription)
type SubscriptionApproval struct {
	State string `json:"state"`
	// Reason for the state.
	Reason string `json:"reason,omitempty"`
}
