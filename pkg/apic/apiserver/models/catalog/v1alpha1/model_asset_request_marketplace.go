/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AssetRequestMarketplace Details about the marketplace Access Request. (catalog.v1alpha1.AssetRequest)
type AssetRequestMarketplace struct {
	// The name of the Marketplace.
	Name     string                          `json:"name"`
	Resource AssetRequestMarketplaceResource `json:"resource"`
}
