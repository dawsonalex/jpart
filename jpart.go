// Package jpart is a utility for extracting partial data from JSON structures.
package jpart

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Path is a delimited list of elements used
// to select some part of a Structure.
type Path string

// Select returns the part of data that matches path. path is a period-delimited
// string of element names. If Select returns an error it will be of type Error.
func Select(path Path, data map[string]interface{}) (string, error) {
	pathStr := strings.TrimSpace(string(path))
	if pathStr == "" {
		output, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return "", makeError("error marshalling data", err)
		}
		return string(output), nil
	}
	pathParts := strings.Split(pathStr, ".")

	var currentElement interface{}
	currentElement = data
	for i, key := range pathParts {
		switch el := currentElement.(type) {
		case map[string]interface{}:
			val, ok := el[key]
			if !ok {
				return "", makeError(fmt.Sprintf("path %s doesn't exist, got to \"%s\"", path, strings.Join(pathParts[:i], ".")), nil)
			}
			currentElement = val
		default:
			// should be the last element in the path, otherwise we error
			if i != len(pathParts) {
				return "", makeError(fmt.Sprintf("path %s doesn't exist, got to \"%s\"", path, strings.Join(pathParts[:i], ".")), nil)
			}
			currentElement = el
		}
	}

	output, err := json.MarshalIndent(currentElement, "", "  ")
	if err != nil {
		return "", makeError("error marshalling data", err)
	}
	return string(output), nil
}

// Error represents an error in Matching a path
// to a Structure.
type jpartError struct {
	inner   error
	message string
}

// Error returns the package error message, and the inner error if there is one.
func (j jpartError) Error() string {
	if j.inner == nil {
		return fmt.Sprintf("jpart: %s", j.message)
	}
	return fmt.Sprintf("jpart: %s: %v", j.message, j.inner)
}

func makeError(message string, inner error) error {
	return jpartError{
		inner:   inner,
		message: message,
	}
}
