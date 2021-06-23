package v1

import "encoding/json"

// ResourceInstance API Server generic resource structure.
type ResourceInstance struct {
	ResourceMeta
	// GENERATE: The following code has been modified after code generation
	//  	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`
	// Resource instance specs.
	Spec map[string]interface{} `json:"spec"`

	SubResources map[string]json.RawMessage `json:"-"`
}

// AsInstance -
func (ri *ResourceInstance) AsInstance() (*ResourceInstance, error) {
	return ri, nil
}

// FromInstance -
func (ri *ResourceInstance) FromInstance(from *ResourceInstance) error {
	*ri = *from

	return nil
}

//UnmarshalJSON - custom unmarshaler for ResourceInstance struct to handle subResources
func (ri *ResourceInstance) UnmarshalJSON(data []byte) error {
	type Alias ResourceInstance // Create an intermediate type to unmarshal the base attributes

	if err := json.Unmarshal(data, &struct{ *Alias }{Alias: (*Alias)(ri)}); err != nil {
		return err
	}

	var allFields map[string]json.RawMessage
	json.Unmarshal(data, &allFields)
	ri.SubResources = make(map[string]json.RawMessage)
	for key, value := range allFields {
		if key != "owner" && key != "spec" {
			ri.SubResources[key] = value
		}
	}
	return nil
}

//Interface -
type Interface interface {
	Meta
	AsInstance() (*ResourceInstance, error)
	FromInstance(from *ResourceInstance) error
}
