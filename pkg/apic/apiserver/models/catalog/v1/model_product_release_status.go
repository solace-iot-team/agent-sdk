/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	// GENERATE: The following code has been modified after code generation
	// 	"time"
	time "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

// ProductReleaseStatus struct for ProductReleaseStatus
type ProductReleaseStatus struct {
	Type string `json:"type"`
	// Details of the error.
	Detail string `json:"detail"`
	// Time when the update occurred.
	Timestamp time.Time                         `json:"timestamp"`
	Meta      map[string]map[string]interface{} `json:"meta,omitempty"`
}
