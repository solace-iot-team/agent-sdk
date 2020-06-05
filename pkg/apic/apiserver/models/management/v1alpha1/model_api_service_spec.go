/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ApiServiceSpec struct for ApiServiceSpec
type ApiServiceSpec struct {
	// The description of the api service.
	Description string             `json:"description,omitempty"`
	Icon        ApiServiceSpecIcon `json:"icon,omitempty"`
}
