/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// MeshDiscoverySpec  (management.v1alpha1.MeshDiscovery)
type MeshDiscoverySpec struct {
	// Target environment.
	EnvironmentRef string `json:"environmentRef"`
}
