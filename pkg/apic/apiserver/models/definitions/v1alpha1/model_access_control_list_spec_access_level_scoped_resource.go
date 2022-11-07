/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package definitions

// AccessControlListSpecAccessLevelScopedResource struct for AccessControlListSpecAccessLevelScopedResource
type AccessControlListSpecAccessLevelScopedResource struct {
	// Resource level at which access is being granted.
	Level string `json:"level,omitempty"`
	// The kind of scoped resources to provide access.
	Kind string `json:"kind"`
	// The name of the scoped resource to provide access to.
	Name string `json:"name"`
	// Set true to allow users to delete the referenced scoped resource.
	AllowDelete bool `json:"allowDelete,omitempty"`
	// Set true to allow users to update the referenced scoped resource.
	AllowWrite bool `json:"allowWrite,omitempty"`
}
