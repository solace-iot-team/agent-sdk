/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// ProductState the model 'ProductState'
type ProductState string

// List of ProductState
const (
	// GENERATE: The following code has been modified after code generation
	ProductStateDRAFT ProductState = "draft"
	// GENERATE: The following code has been modified after code generation
	ProductStateACTIVE ProductState = "active"
	// GENERATE: The following code has been modified after code generation
	ProductStateDEPRECATED ProductState = "deprecated"
	// GENERATE: The following code has been modified after code generation
	ProductStateARCHIVED ProductState = "archived"
)
