/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SecretSpec  (catalog.v1alpha1.Secret)
type SecretSpec struct {
	// Key value pairs. The value will be stored encrypted. (catalog.v1alpha1.Secret)
	Data map[string]string `json:"data,omitempty"`
}