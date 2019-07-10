package store

import (
	"os"
	"fmt"
	
	"log"
	"io/ioutil"
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
	pathname := fmt.Sprintf("temp/%s", name)
	os.MkdirAll("temp/", os.ModePerm)
	os.Create(pathname)

	err := ioutil.WriteFile(pathname, *stream, 0644)
	if err != nil {
		log.Fatal(err)
	}
}