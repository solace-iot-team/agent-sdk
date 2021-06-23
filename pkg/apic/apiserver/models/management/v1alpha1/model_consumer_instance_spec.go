/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ConsumerInstanceSpec struct for ConsumerInstanceSpec
type ConsumerInstanceSpec struct {
	// Maps to the name of the Catalog Item. If not provided, the resource title will be used.
	Name string `json:"name,omitempty"`
	// The name of an APIServiceInstance resource that specifies where the API is deployed.
	ApiServiceInstance string `json:"apiServiceInstance,omitempty"`
	// Name of the team that owns the Catalog Item. If not provided, the Default team will be used.
	OwningTeam string `json:"owningTeam,omitempty"`
	// Maps to the description of the Catalog Item. Defaults to the API service description.
	Description string `json:"description,omitempty"`
	// The visibility of the Catalog Item:  * PUBLIC - If Catalog Item is in PUBLISHED state, it will be visible to the entire organization.  * RESTRICTED - If Catalog Item is in PUBLISHED state, it will be visible to the owning team and teams part of the Catalog Item Access Control List.
	Visibility string `json:"visibility,omitempty"`
	// Version name for the revision.
	Version string `json:"version,omitempty"`
	// Catalog Item state:  * UNPUBLISHED - Only developers in the owning team will be able to access the Catalog Item.  * PUBLISHED - Developers and Consumers in the owning team and any additional team in the ACL will be able to access the Catalog Item.
	State string `json:"state,omitempty"`
	// A way to communicate the status of the service to consumers. Examples: DRAFT, BETA, GA
	Status string `json:"status,omitempty"`
	// List of tags to be set on the Catalog Item. Max allowed tags for the Catalog Item is 80.
	Tags []string `json:"tags,omitempty"`
	// GENERATE: The following code has been modified after code generation
	// 	Icon ConsumerInstanceSpecIcon `json:"icon,omitempty"`
	Icon *ConsumerInstanceSpecIcon `json:"icon,omitempty"`
	// Markdown documentation for the Catalog Item documentation.
	Documentation              string                                         `json:"documentation,omitempty"`
	UnstructuredDataProperties ConsumerInstanceSpecUnstructuredDataProperties `json:"unstructuredDataProperties,omitempty"`
	AdditionalDataProperties   ConsumerInstanceSpecAdditionalDataProperties   `json:"additionalDataProperties,omitempty"`
	Subscription               ConsumerInstanceSpecSubscription               `json:"subscription,omitempty"`
	Categories                 []string                                       `json:"categories,omitempty"`
}
