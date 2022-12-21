/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AssetRequestMarketplaceResource The Marketplace Access Request resource details. (catalog.v1alpha1.AssetRequest)
type AssetRequestMarketplaceResource struct {
	Metadata AssetRequestMarketplaceResourceMetadata `json:"metadata"`
	Owner    AssetRequestMarketplaceResourceOwner    `json:"owner,omitempty"`
}