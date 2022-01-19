/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductPlanSpec struct for ProductPlanSpec
type ProductPlanSpec struct {
	Product string `json:"product"`
	// description of the Plan.
	Description string `json:"description,omitempty"`
	// The type of the Plan.
	Type string `json:"type,omitempty"`
	// Defines all features supported by the Plan.
	Features     []ProductPlanSpecFeatures   `json:"features,omitempty"`
	Subscription ProductPlanSpecSubscription `json:"subscription,omitempty"`
}