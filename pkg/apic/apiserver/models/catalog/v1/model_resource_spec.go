/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package catalog

// ResourceSpec  (catalog.v1.Resource)
type ResourceSpec struct {
	// The type of the resource, example: pdf
	FileType string `json:"fileType"`
	// The content type
	ContentType string `json:"contentType"`
	// GENERATE: The following code has been modified after code generation
	Data interface{} `json:"data"`
}
