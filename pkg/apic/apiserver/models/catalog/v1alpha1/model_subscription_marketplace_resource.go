/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// SubscriptionMarketplaceResource The Marketplace Subscription resource details. (catalog.v1alpha1.Subscription)
type SubscriptionMarketplaceResource struct {
	Metadata SubscriptionMarketplaceResourceMetadata `json:"metadata"`
	Owner    SubscriptionMarketplaceResourceOwner    `json:"owner,omitempty"`
}
