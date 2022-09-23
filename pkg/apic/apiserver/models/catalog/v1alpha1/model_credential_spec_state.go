/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// CredentialSpecState Desired state of the Credential. (catalog.v1alpha1.Credential)
type CredentialSpecState struct {
	Name string `json:"name"`
	// Additional info on the state.
	Reason string `json:"reason,omitempty"`
	// Additional info on the state.
	Rotate bool `json:"rotate,omitempty"`
}
