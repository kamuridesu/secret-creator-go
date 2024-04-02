package main

import (
	"encoding/base64"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Result *map[string]any

func ConvertValuesToBase64(contents *string) Result {
	byteContent := []byte(*contents)
	var data map[string]any
	if err := yaml.Unmarshal(byteContent, &data); err != nil {
		panic(err)
	}
	for key, value := range data {
		encodedValue := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", value)))
		data[key] = encodedValue
	}
	return &data
}

func PrintResult(data Result) {
	for key, value := range *data {
		fmt.Printf("%v: %v\n", key, value)
	}
}
