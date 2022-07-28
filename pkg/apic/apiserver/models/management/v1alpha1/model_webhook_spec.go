/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

// WebhookSpec  (management.v1alpha1.Webhook)
type WebhookSpec struct {
	Auth    WebhookSpecAuth `json:"auth,omitempty"`
	Enabled bool            `json:"enabled,omitempty"`
	Url     string          `json:"url"`
	// A list of headers that will be sent as par of the http call to the webhook endpoint. (management.v1alpha1.Webhook)
	Headers map[string]string `json:"headers,omitempty"`
}
