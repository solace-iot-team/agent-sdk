/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package definitions

// CommandLineInterfaceSpecNames  (definitions.v1alpha1.CommandLineInterface)
type CommandLineInterfaceSpecNames struct {
	// Defines the name used to access resources in this group. Also provided as default in the autocomplete for listing commands.
	Plural string `json:"plural"`
	// Defines the name used to access a resource in this group. Also provided as default in the autocomplete for single resource access commands.
	Singular string `json:"singular"`
	// Defines the short names that the cli can use to fetch a resource in the group.
	ShortNames []string `json:"shortNames"`
	// Defines the short names alias that the cli can use to fetch a resource in the group.
	ShortNamesAlias []string `json:"shortNamesAlias,omitempty"`
}
