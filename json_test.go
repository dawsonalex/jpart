package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestIntType(t *testing.T) {
	in := `{"a" : 1 }`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: NumberType,
		value:    float64(1),
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestFloatType(t *testing.T) {
	in := `{"a" : 1.1 }`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: NumberType,
		value:    1.1,
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestStringType(t *testing.T) {
	in := `{"a" : "Hello, World!"}`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: StringType,
		value:    "Hello, World!",
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestArrayType(t *testing.T) {
	in := `{"a" : [1, 2, 3]}`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: ArrayType,
		value: []interface{}{
			JsonElement{
				jsonType: NumberType,
				value:    float64(1),
			},
			JsonElement{
				jsonType: NumberType,
				value:    float64(2),
			},
			JsonElement{
				jsonType: NumberType,
				value:    float64(3),
			},
		},
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestNullType(t *testing.T) {
	in := `{"a" : null}`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: NullType,
		value:    nil,
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestBoolType(t *testing.T) {
	in := `{"a" : true}`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: BoolType,
		value:    true,
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestObjectType(t *testing.T) {
	in := `{"a" : {"b" : "c"}}`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err != nil {
		t.Errorf("error unmarhsalling json: %v", err)
	}

	expected := JsonElement{
		jsonType: ObjectType,
		value:    JsonElement{jsonType: StringType, value: "c"},
	}
	if !cmp.Equal(result, expected) {
		fmt.Printf("Got: %#v\n", result)
		fmt.Printf("Expected: %#v\n", expected)
		t.FailNow()
	}
}

func TestInvalidType(t *testing.T) {
	in := `{"a" : {"b" : invalid-value-here}}`

	var result JsonElement
	if err := json.Unmarshal([]byte(in), &result); err == nil {
		t.Errorf("Expected json.Unmarshal to error")
		t.FailNow()
	} else {
		if err, ok := err.(*JsonError); ok {
			t.Logf("json.Unmarshal failed with error:\n%s", err)
		}
	}
}
