/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// ProductPlanSpecBilling Paid plan billing details. (catalog.v1alpha1.ProductPlan)
type ProductPlanSpecBilling struct {
	Currency string `json:"currency"`
	// The base price for the plan.
	Price float64 `json:"price,omitempty"`
	// The billing cycle type.
	Cycle    string `json:"cycle,omitempty"`
	Interval string `json:"interval"`
}
