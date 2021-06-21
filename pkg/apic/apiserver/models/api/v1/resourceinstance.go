package v1

// ResourceInstance API Server generic resource structure.
type ResourceInstance struct {
	ResourceMeta `mapstructure:",squash"`
	Spec         map[string]interface{} `json:"spec"`
	SubResources map[string]interface{} `mapstructure:",remain"`
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

type Interface interface {
	Meta
	AsInstance() (*ResourceInstance, error)
	FromInstance(from *ResourceInstance) error
}
