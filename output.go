package yayoi

import (
	"fmt"
	"io/ioutil"
)

func Output(outputBytes []byte, dest string) {
	if dest == "STDOUT" {
		fmt.Println(string(outputBytes))
	} else {
		ioutil.WriteFile(dest, outputBytes, 0644)
	}
	return
}
