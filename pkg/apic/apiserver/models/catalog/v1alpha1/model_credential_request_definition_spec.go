/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// CredentialRequestDefinitionSpec  (catalog.v1alpha1.CredentialRequestDefinition)
type CredentialRequestDefinitionSpec struct {
	// JSON Schema draft \\#7 for describing the fields needed to provision Credentials of that type. (catalog.v1alpha1.CredentialRequestDefinition)
	Schema    map[string]interface{}                   `json:"schema"`
	Provision CredentialRequestDefinitionSpecProvision `json:"provision"`
}
