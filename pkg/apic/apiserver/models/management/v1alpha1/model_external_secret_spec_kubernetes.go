/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// ExternalSecretSpecKubernetes K8s Secret Information.
type ExternalSecretSpecKubernetes struct {
	Provider string `json:"provider"`
	// Name of K8s Secret.
	Name string `json:"name"`
	// Namespace which contains the K8s Secret.  If none is provided, the secret will be sourced from the namespace into which Amplify Gateway is installed.
	Namespace string `json:"namespace,omitempty"`
}
