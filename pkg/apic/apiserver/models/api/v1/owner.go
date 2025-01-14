package v1

import (
	"encoding/json"
	"fmt"
)

// OwnerType -
type OwnerType uint

const (
	TeamOwner OwnerType = iota
)

var ownerTypeToString = map[OwnerType]string{
	TeamOwner: "team",
}

var ownerTypeFromString = map[string]OwnerType{
	"team": TeamOwner,
}

// Owner is the owner of a resource
type Owner struct {
	Type OwnerType `json:"type,omitempty"`
	ID   string    `json:"id"`
}

// SetType sets the type of the owner
func (o *Owner) SetType(t OwnerType) {
	o.Type = t
}

// SetID sets the id of the owner
func (o *Owner) SetID(id string) {
	o.ID = id
}

// MarshalJSON marshals the owner to JSON
func (o *Owner) MarshalJSON() ([]byte, error) {
	var t string
	var ok bool
	if t, ok = ownerTypeToString[o.Type]; !ok {
		return nil, fmt.Errorf("unknown owner type %d", o.Type)
	}

	aux := struct {
		Type string `json:"type,omitempty"`
		ID   string `json:"id"`
	}{}

	aux.Type = t
	aux.ID = o.ID

	return json.Marshal(aux)
}

// UnmarshalJSON unmarshalls the owner from JSON to convert the owner type to a string
func (o *Owner) UnmarshalJSON(bytes []byte) error {
	aux := struct {
		Type string `json:"type,omitempty"`
		ID   string `json:"id"`
	}{}

	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	ownerString, ok := ownerTypeFromString[aux.Type]
	if !ok {
		return fmt.Errorf("unknown owner type %s", aux.Type)
	}
	o.Type = ownerString
	o.ID = aux.ID

	return nil
}
