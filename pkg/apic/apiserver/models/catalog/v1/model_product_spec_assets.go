/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// ProductSpecAssets  (catalog.v1.Product)
type ProductSpecAssets struct {
	Name    string             `json:"name"`
	Filters ProductSpecFilters `json:"filters,omitempty"`
}
