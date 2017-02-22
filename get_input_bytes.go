package yayoi

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetInputBytes(source, str string) []byte {
	var inputBytes []byte

	switch source {
	case "ARGV":
	case "FILE":
		var err error
		inputBytes, err = ioutil.ReadFile(str)

		if err != nil {
			fmt.Println("read error")
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
	case "STDIN":

	default:
		fmt.Println("Unknown")
		os.Exit(1)
	}

	return inputBytes
}
