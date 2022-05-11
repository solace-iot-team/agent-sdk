/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ManagedApplicationMarketplaceResourceOwner Owner of the Application. (management.v1alpha1.ManagedApplication)
type ManagedApplicationMarketplaceResourceOwner struct {
	// The type of the owner.
	Type string `json:"type,omitempty"`
	// Id of the owner of the resource.
	Id           string                                                 `json:"id,omitempty"`
	Organization ManagedApplicationMarketplaceResourceOwnerOrganization `json:"organization"`
}