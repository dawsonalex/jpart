package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type JsonType int

type JsonElement struct {
	jsonType JsonType
	value    interface{}
}

type JsonError struct {
	Key     string
	Value   interface{}
	Wrapped error
}

const (
	NullType JsonType = iota
	ArrayType
	NumberType
	StringType
	BoolType
	ObjectType
)

func (e *JsonError) Error() string {
	return fmt.Sprintf("error parsing element with key %v: %v", e.Key, e.Wrapped)
}

func (j JsonElement) Type() JsonType {
	return j.jsonType
}

func (j JsonElement) Value() interface{} {
	return j.value
}
func (j JsonElement) Equal(alt JsonElement) bool {
	if j.jsonType != alt.jsonType {
		return false
	}

	if j.jsonType == ObjectType {
		if len(j.value.(map[string]interface{})) != len(alt.value.(map[string]interface{})) {
			return false
		}

		for k, v := range j.value.(map[string]interface{}) {
			if altValue, ok := alt.value.(map[string]interface{})[k]; ok {
				if !cmp.Equal(v, altValue) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}

	if j.jsonType == ArrayType {
		if len(j.value.([]interface{})) != len(alt.value.([]interface{})) {
			return false
		}

		for i, v := range j.value.([]interface{}) {
			altValue := alt.value.([]interface{})[i].(JsonElement)
			if !cmp.Equal(v, altValue) {
				return false
			}
		}
		return true
	}

	if cmp.Equal(j.value, alt.value) {
		return true
	}
	return false
}

func (j *JsonElement) UnmarshalJSON(data []byte) error {
	var jsonValue map[string]interface{}
	if err := json.Unmarshal(data, &jsonValue); err != nil {
		fmt.Printf("error unmarshalling %v, %v", data, err)
	}

	for k, v := range jsonValue {
		err := buildJson(v, j)
		if err != nil {
			return &JsonError{
				Key:     k,
				Value:   v,
				Wrapped: err,
			}
		}
	}

	return nil
}

func buildJson(val interface{}, j *JsonElement) error {
	switch val.(type) {
	case nil:
		j.jsonType = NullType
		j.value = nil
	case int, float32, float64:
		j.jsonType = NumberType
		j.value = val.(float64)
	case string:
		j.jsonType = StringType
		j.value = val
	case bool:
		j.jsonType = BoolType
		j.value = val
	case []interface{}:
		j.jsonType = ArrayType
		jsonArray, err := buildJsonArray(val.([]interface{}))
		if err != nil {
			return err
		}
		j.value = jsonArray
	case map[string]interface{}:
		j.jsonType = ObjectType
		jsonObject, err := buildJsonObject(val.(map[string]interface{}))
		if err != nil {
			return err
		}
		j.value = jsonObject
	default:
		return fmt.Errorf("cannot determine type for: %v", val)
	}
	return nil
}

func buildJsonObject(val map[string]interface{}) (JsonElement, error) {
	// TODO: Need to decide on how value is mapped to key for objects. Should match array type.
	var obj JsonElement
	for k, v := range val {
		var el JsonElement
		err := buildJson(v, &el)
		if err != nil {
			return obj, fmt.Errorf("cannot build JSON for key %v in object %v", k, v)
		}
		obj.value = el
	}
	return obj, nil
}

func buildJsonArray(in []interface{}) ([]interface{}, error) {
	a := make([]interface{}, 0)
	for i, v := range in {
		var arrayElement JsonElement
		err := buildJson(v, &arrayElement)
		if err != nil {
			return a, fmt.Errorf("cannot build JSON for index %d in array %v", i, in)
		}
		a = append(a, arrayElement)
	}
	return a, nil
}
