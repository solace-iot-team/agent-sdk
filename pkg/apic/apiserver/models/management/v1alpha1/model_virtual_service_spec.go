/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// VirtualServiceSpec  (management.v1alpha1.VirtualService)
type VirtualServiceSpec struct {
	// The path prefix to match. Example /api/v1
	Prefix string `json:"prefix"`
	// The list of headers and values to match. These are ANDed together. (management.v1alpha1.VirtualService)
	HeaderMatch map[string]string `json:"headerMatch,omitempty"`
	// The hosts that the VirtualService will match.
	VirtualHosts []string `json:"virtualHosts,omitempty"`
	Cors         string   `json:"cors,omitempty"`
	// GENERATE: The following code has been modified after code generation
	Auth    interface{} `json:"auth"`
	Lambdas []string    `json:"lambdas,omitempty"`
	// GENERATE: The following code has been modified after code generation
	Route []VirtualServiceSpecServiceRouting `json:"route"`
}
