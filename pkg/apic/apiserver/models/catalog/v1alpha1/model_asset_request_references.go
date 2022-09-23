/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AssetRequestReferences  (catalog.v1alpha1.AssetRequest)
type AssetRequestReferences struct {
	AssetRequestDefinition string `json:"assetRequestDefinition,omitempty"`
	// Reference to Release that got created from this asset request.
	AssetRelease string `json:"assetRelease,omitempty"`
	// Reference to Release that got created from this asset request.
	Asset                        string   `json:"asset,omitempty"`
	CredentialRequestDefinitions []string `json:"credentialRequestDefinitions,omitempty"`
	// Reference to Access Request resource
	AccessRequest string `json:"accessRequest,omitempty"`
}
