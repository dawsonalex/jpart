package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dawsonalex/jutil"
	"io/ioutil"
	"os"
)

var jsonPath string

func init() {
	flag.StringVar(&jsonPath,
		"p",
		"",
		"A path in the format 'first.second.third' which defines the data you want. Leave empty to see the whole JSON input.")
	flag.Usage = usage
}

// Input a valid JSON string, and return an arbitrary value.
func main() {
	flag.Parse()
	input := getStdIn()
	fmt.Printf("Got input: %v\n", string(input))

	var jsonIn jutil.JsonElement
	err := json.Unmarshal(input, &jsonIn)
	if err != nil {
		fmt.Printf("erro unmarshalling: %v\n", err)
	}

	//fmt.Printf("result: %v\n", jsonIn["a"])

	//fmt.Printf("result is: %v", findInMap("a", jsonIn))

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
Usage: jutil [-p <path>]
Options:
`)
	flag.PrintDefaults()
}
