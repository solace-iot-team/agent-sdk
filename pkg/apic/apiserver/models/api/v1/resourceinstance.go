package v1

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ResourceInstance API Server generic resource structure.
type ResourceInstance struct {
	ResourceMeta
	Spec         map[string]interface{} `json:"spec"`
	SubResources map[string]interface{} `json:"-"`
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

func (ri *ResourceInstance) UnmarshalJSON(data []byte) error {
	type BaseRI ResourceInstance

	if err := json.Unmarshal(data, &struct{ *BaseRI }{BaseRI: (*BaseRI)(ri)}); err != nil {
		return err
	}

	var allFields interface{}
	json.Unmarshal(data, &allFields)
	allFieldsMap := allFields.(map[string]interface{})

	baseFields := make(map[string]struct{})
	getBaseJsonFieldNames(reflect.TypeOf((*ResourceInstance)(nil)), baseFields)

	// Add all non-base fields into SubResources
	ri.SubResources = make(map[string]interface{})
	for k, v := range allFieldsMap {
		if _, ok := baseFields[k]; !ok {
			ri.SubResources[k] = v
		}
	}

	return nil
}

func getBaseJsonFieldNames(t reflect.Type, m map[string]struct{}) {
	// Return if not struct or pointer to struct.
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	// Iterate through fields collecting names in map.
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		jsonName, ok := getJsonFieldName(f)

		if ok {
			m[jsonName] = struct{}{}
		}

		// Recurse into anonymous fields.
		if f.Anonymous {
			getBaseJsonFieldNames(f.Type, m)
		}
	}
}

func getJsonFieldName(f reflect.StructField) (string, bool) {
	jsonName := strings.TrimSpace(strings.Split(f.Tag.Get("json"), ",")[0])
	if jsonName != "" && jsonName != "-" {
		return jsonName, true
	} else {
		return "", false
	}
}

type Interface interface {
	Meta
	AsInstance() (*ResourceInstance, error)
	FromInstance(from *ResourceInstance) error
}
