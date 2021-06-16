/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// GovernanceAgentStatus struct for GovernanceAgentStatus
type GovernanceAgentStatus struct {
	// Version name for the agent revision.
	Version string `json:"version,omitempty"`
	// Agent status:  * running - Passed all health checks.  Up and running  * stopped - Agent is not running  * failed - Failed health checks  * unhealthy - Agent is running with health check failure
	State string `json:"state,omitempty"`
	// A way to communicate details about the current status by the agent
	Message string `json:"message,omitempty"`
	// The last updated event timestamp provided by the agent
	LastActivityTime string `json:"lastActivityTime,omitempty"`
}
