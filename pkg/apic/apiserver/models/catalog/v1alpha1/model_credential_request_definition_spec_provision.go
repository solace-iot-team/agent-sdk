/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// CredentialRequestDefinitionSpecProvision  (catalog.v1alpha1.CredentialRequestDefinition)
type CredentialRequestDefinitionSpecProvision struct {
	// JSON Schema draft \\#7 for describing the generated credentials format. (catalog.v1alpha1.CredentialRequestDefinition)
	Schema   map[string]interface{}                           `json:"schema"`
	Policies CredentialRequestDefinitionSpecProvisionPolicies `json:"policies,omitempty"`
}
