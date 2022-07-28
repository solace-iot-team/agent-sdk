/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// VirtualServiceSpecBasiccredential The HTTP Basic credential.
type VirtualServiceSpecBasiccredential struct {
	Kind string `json:"kind"`
	// The name of the external secret containing the http basic credential.
	SecretName string `json:"secretName"`
}
