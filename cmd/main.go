package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dawsonalex/jpart"
	"io/ioutil"

	"os"
	"strings"
)

type arrayValue []string

var jsonPaths arrayValue
var displayPath bool

var Version string

func (f arrayValue) String() string {
	return "[" + strings.Join(f, ", ") + "]"
}

func (f *arrayValue) Set(s string) error {
	*f = append(*f, s)
	return nil
}

func init() {
	flag.Var(&jsonPaths,
		"p",
		"A path in the format 'first.second.third' which defines the data you want. Leave empty to see the whole JSON input. jutil will output a value for each -p if there are multiple.")
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
		fmt.Printf("jutil: error parsing JSON, %v\n", err)
		os.Exit(1)
	}

	// just format and output the input if there is not a path.
	if len(jsonPaths) == 0 {
		output, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Printf("jutil: error outputting JSON, %v", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
		os.Exit(0)
	}
	if len(jsonPaths) > 1 {
		displayPath = true
	}

	for _, path := range jsonPaths {
		output, err := jpart.Select(jpart.Path(path), data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if displayPath {
			headingLine := strings.Repeat("-", len(path)+2)
			fmt.Printf("%s\n %s\n%s\n", headingLine, path, headingLine)
		}
		fmt.Println(output)
		fmt.Println()
	}
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
jutil %s
=======
Usage: jutil [-p <path>] [-v]
Options:
`, Version)
	flag.PrintDefaults()
}
