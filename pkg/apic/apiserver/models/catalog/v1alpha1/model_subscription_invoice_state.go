/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// SubscriptionInvoiceState  (catalog.v1alpha1.SubscriptionInvoice)
type SubscriptionInvoiceState struct {
	Name string `json:"name,omitempty"`
	// Additional info on the state.
	Reason string `json:"reason,omitempty"`
}
