/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// AssetLatestrelease Provides newest non-archived release and version. Will be unassigned if no releases exist. (catalog.v1alpha1.Asset)
type AssetLatestrelease struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}