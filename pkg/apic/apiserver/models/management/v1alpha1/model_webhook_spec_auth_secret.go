/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// WebhookSpecAuthSecret  (management.v1alpha1.Webhook)
type WebhookSpecAuthSecret struct {
	// Secret name to be used as a reference. If the secret is removed, the reference gets removed and the webhook invocation will be done with no Authorization header.
	Name string `json:"name,omitempty"`
	// Key to be used from the referenced secret.
	Key string `json:"key,omitempty"`
}
