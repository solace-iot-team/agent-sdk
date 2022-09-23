/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

import (
	// GENERATE: The following code has been modified after code generation
	//
	//	"time"
	time "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

// ConsumerInstanceStatusPhase  (management.v1alpha1.ConsumerInstance)
type ConsumerInstanceStatusPhase struct {
	Name string `json:"name"`
	// The criticality of the last update
	Level string `json:"level"`
	// Time of the current phase.
	TransitionTime time.Time `json:"transitionTime"`
	// Description of the phase.
	Message string `json:"message,omitempty"`
}
