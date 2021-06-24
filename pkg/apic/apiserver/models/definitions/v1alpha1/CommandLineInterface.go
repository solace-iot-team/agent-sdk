/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_CommandLineInterfaceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "CommandLineInterface",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	CommandLineInterfaceScope = "ResourceGroup"

	CommandLineInterfaceResourceName = "commandlines"
)

func CommandLineInterfaceGVK() apiv1.GroupVersionKind {
	return _CommandLineInterfaceGVK
}

func init() {
	apiv1.RegisterGVK(_CommandLineInterfaceGVK, CommandLineInterfaceScope, CommandLineInterfaceResourceName)
}

// CommandLineInterface Resource
type CommandLineInterface struct {
	apiv1.ResourceMeta

	Owner interface{} `json:"owner"`

	Spec CommandLineInterfaceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a CommandLineInterface
func (res *CommandLineInterface) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	err := json.Unmarshal(ri.GetRawResource(), res)
	return err
}

// CommandLineInterfaceFromInstanceArray converts a []*ResourceInstance to a []*CommandLineInterface
func CommandLineInterfaceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*CommandLineInterface, error) {
	newArray := make([]*CommandLineInterface, 0)
	for _, item := range fromArray {
		res := &CommandLineInterface{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*CommandLineInterface, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a CommandLineInterface to a ResourceInstance
func (res *CommandLineInterface) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = CommandLineInterfaceGVK()
	res.ResourceMeta = meta

	m, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	instance := apiv1.ResourceInstance{}
	err = json.Unmarshal(m, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
