/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AccessControlListSpecAccessLevelScopedKind struct for AccessControlListSpecAccessLevelScopedKind
type AccessControlListSpecAccessLevelScopedKind struct {
	// Resource level at which access is being granted.
	Level string `json:"level,omitempty"`
	// The kind of scoped resources to provide access to or \"*\" for all kinds.
	Kind string `json:"kind"`
}
