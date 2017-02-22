package yayoi

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/ghodss/yaml"
)

func YamlToJason(inputBytes []byte, minifyFlg bool) []byte {
	buffer, err := yaml.YAMLToJSON(inputBytes)

	if err != nil {
		fmt.Println("convert error")
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	if minifyFlg {
		return buffer
	} else {
		jsonData, err := simplejson.NewJson(buffer)

		if err != nil {
			fmt.Println("parse error")
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		prettyByte, err := jsonData.EncodePretty()

		if err != nil {
			fmt.Println("encode error")
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		return prettyByte
	}
}
