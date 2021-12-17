package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var jsonPath string
var displayPath bool

func init() {
	flag.StringVar(&jsonPath,
		"p",
		"",
		"A path in the format 'first.second.third' which defines the data you want. Leave empty to see the whole JSON input.")
	flag.BoolVar(&displayPath,
		"v",
		false,
		"Verbose mode displays the path that the element was found on the line before the value output.")
	flag.Usage = usage
}

// Input a valid JSON string, and return an arbitrary value.
func main() {
	flag.Parse()
	input := getStdIn()

	var data map[string]interface{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		fmt.Printf("jutil: error parsing JSON, %v", err)
		os.Exit(1)
	}

	// just format and output the input if there is not path.
	if jsonPath == "" {
		output, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Printf("jutil: error outputting JSON, %v", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
		os.Exit(0)
	}

	keys := strings.Split(jsonPath, ".")

	var currentElement interface{}
	currentElement = data
	for i, key := range keys {
		switch el := currentElement.(type) {
		case map[string]interface{}:
			val, ok := el[key]
			if !ok {
				fmt.Printf("jutil: JSON structure doesn't match path %s, got to %s\n", jsonPath, strings.Join(keys[:i], "."))
				os.Exit(1)
			}
			currentElement = val
		default:
			// should be the last element in the path, otherwise we error
			if i != len(keys)-1 {
				fmt.Printf("jutil: JSON structure doesn't match path %s, got to %s\n", jsonPath, strings.Join(keys[:i], "."))
				os.Exit(1)
			}
			currentElement = el
		}
	}

	output, err := json.MarshalIndent(currentElement, "", "  ")
	if err != nil {
		fmt.Printf("jutil: error outputting JSON, %v", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
	os.Exit(0)
}

func getStdIn() []byte {
	in := make([]byte, 0)
	reader := bufio.NewReader(os.Stdin)
	stdIn, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("error reading from stdin: %v\n", err)
		os.Exit(1)
	}
	in = stdIn

	return in
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `=======
jutil
=======
Usage: jutil [-p <path>] [-v]
Options:
`)
	flag.PrintDefaults()
}
