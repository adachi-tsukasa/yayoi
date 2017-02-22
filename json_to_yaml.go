package yayoi

import (
	"fmt"
	"os"

	"github.com/ghodss/yaml"
)

func JsonToYaml(inputBytes []byte) []byte {
	buffer, err := yaml.JSONToYAML(inputBytes)

	if err != nil {
		fmt.Println("convert error")
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	return buffer
}
