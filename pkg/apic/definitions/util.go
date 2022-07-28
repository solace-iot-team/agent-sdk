package definitions

import (
	"strings"

	management "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// GetSubscriptionNameFromAccessRequest - Returns the subscription name from access request references
func GetSubscriptionNameFromAccessRequest(ar *management.AccessRequest) string {
	if ar == nil {
		return ""
	}

	subscriptionName := ""
	subsRefName := getSubscriptionRefName(ar)
	if subsRefName != "" {
		refElements := strings.Split(subsRefName, "/")
		if len(refElements) == 2 && refElements[0] == "catalog" {
			subscriptionName = refElements[1]
		}
	}
	return subscriptionName
}

func getSubscriptionRefName(ar *management.AccessRequest) string {
	for _, ref := range ar.References {
		switch arRef := ref.(type) {
		case map[string]interface{}:
			kind := arRef["kind"]
			if kind == Subscription {
				return arRef["name"].(string)
			}
		case management.AccessRequestReferencesSubscription:
			return getSubscriptionName(&arRef)
		case *management.AccessRequestReferencesSubscription:
			return getSubscriptionName(arRef)
		}
	}
	return ""
}

func getSubscriptionName(arRef *management.AccessRequestReferencesSubscription) string {
	if arRef.Kind == Subscription {
		return arRef.Name
	}
	return ""
}
