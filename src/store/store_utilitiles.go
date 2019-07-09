package store

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func updateFromIo(name string) *[]byte {
	data := &[]byte{}

	pname := fmt.Sprintf("/tmp/%s", name)
	if _, err := os.Stat(pname); os.IsNotExist(err) {
		os.Create(pname)
		return data
	}

	file, _ := os.Open(pname)
    defer file.Close()

	b, _ := ioutil.ReadAll(file)
	return &b
}

func copyToIo(name string, stream *[]byte) {
	pname := fmt.Sprintf("../tmp/%s", name)
	absFilePath, _ := filepath.Abs(pname)

	if _, err := os.Stat(absFilePath); os.IsNotExist(err) {
		os.Create(absFilePath)
	}

	fmt.Println(absFilePath)
	ioutil.WriteFile(absFilePath, *stream, 0644)
}