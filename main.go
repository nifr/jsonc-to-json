package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/akshaybharambe14/go-jsonc"
	"github.com/iancoleman/orderedmap"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var jsonFilePath string
	var prettyPrint bool
	var validateJson bool
	flag.StringVar(&jsonFilePath, "file", "", "JSONC file path to convert to JSON (use \"-\" to read from stdin)")
	flag.BoolVar(&validateJson, "validate", false, "validate decoded JSON result")
	flag.BoolVar(&prettyPrint, "pretty", false, "pretty print result with 2 space indentation (orders keys alphabetically)")
	flag.Parse()

	if len(jsonFilePath) == 0 {
		flag.Usage()
		return
	}

	var jsoncBody []byte
	if jsonFilePath == "-" {
		jsoncBody = ReadBodyFromStdin()
	} else {
		jsoncBody = ReadBodyFromFile(jsonFilePath)
	}

	jsoncDecoder := jsonc.NewDecoder(bytes.NewBuffer([]byte(jsoncBody)))
	jsonBody, err := io.ReadAll(jsoncDecoder)
	if err != nil {
		log.Fatalf("error while decoding JSONC: ", err)
	}

	if validateJson {
		if !(json.Valid(jsonBody)) {
			log.Fatalf("error after decoding JSONC: result is not valid JSON!")
		}
	}

	if prettyPrint {
		// note: Unmarshal re-orders the JSON keys alphabetically
		// var jsonBodyMap map[string]any
		// note: we use an ordered map to work around the issue
		// see: https://github.com/golang/go/issues/27179
		// see: https://stackoverflow.com/a/18668885
		jsonBodyMap := orderedmap.New()
		err = json.Unmarshal([]byte(jsonBody), &jsonBodyMap)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		jsonBody, err = json.MarshalIndent(jsonBodyMap, "", "  ")
	}

	fmt.Println(string(jsonBody))
}

func ReadBodyFromStdin() []byte {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatalf("error while reading from stdin: stdin is not open")
	}

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error while reading from stdin: ", err)
	}
	return stdin
}

func ReadBodyFromFile(jsonFilePath string) []byte {
	body, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	return body
}
