package management

// GENERATE: All of the code below was replaced after code generation

import (
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

// DiscoveryAgentSpecConfig Represents the discovery agent config (management.v1alpha1.DiscoveryAgent)
type DiscoveryAgentSpecConfig struct {
	// Defines the tag based filter expression to be evaluated for discovering the API from the API Gateway. The filter value is a conditional expression that can use logical operators to compare two value. The conditional expression must have \"tag\" as the prefix/selector in the symbol name. For e.g. ``` tag.SOME_TAG == \"somevalue\" ``` The expression can be a simple condition as shown above or compound condition in which more than one simple conditions are evaluated using logical operator. For e.g. ``` tag.SOME_TAG == \"somevalue\" || tag.ANOTHER_TAG != \"some_other_value\" ``` In addition to logical expression, the filter can hold call based expressions. Below are the list of supported call expressions #### Exists Exists call can be made to evaluate if the tag name exists in the list of tags on API. This call expression can be used as unary expression For e.g. ``` tag.SOME_TAG.Exists() ``` #### Any Any call can be made in a simple expression to evaluate if the tag with any name has specified value or not in the list of tags on the API. For e.g. ``` tag.Any() == \"Tag with some value\" || tag.Any() != \"Tag with other value\" ``` #### Contains Contains call can be made in a simple expression to evaluate if the the specified tag contains specified argument as value. This call expression requires string argument that will be used to perform lookup in tag value For e.g. ``` tag.Contains(\"somevalue\") ``` #### MatchRegEx MatchRegEx call can be used for evaluating the specified tag value to match specified regular expression. This call expression requires a regular expression as the argument. For e.g. ``` tag.MatchRegEx(\"(some){1}\") ``` 
	Filter         string `json:"filter,omitempty"`
	// The list of tags to be added to the API service resource that the agent publishes to Amplify Central
	AdditionalTags []string `json:"additionalTags,omitempty"`
	// The list of tags to exclude from the API service resource that the agent publishes to Amplify Central
	IgnoreTags     []string `json:"ignoreTags,omitempty"`
	// Name of the team that owns the catalog item created by agent. If not provided, the default team will be used.
	OwningTeam     string `json:"owningTeam,omitempty"`
	Owner          *apiv1.Owner `json:"owner,omitempty"`
}
