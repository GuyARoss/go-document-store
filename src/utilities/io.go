package utilities

import (
	"os"
	"fmt"
	
	"log"
	"io/ioutil"
)

// UpdateFromIo reads the current bytes from the temp file
func UpdateFromIo(name string) *[]byte {
	pname := fmt.Sprintf("tmp/%s", name)

	file, _ := os.Open(pname)
    defer file.Close()

	b, _ := ioutil.ReadAll(file)
	return &b
}

// CopyToIo writes a byte stream to io
func CopyToIo(name string, stream *[]byte) {
	pathname := fmt.Sprintf("tmp/%s", name)
	os.MkdirAll("tmp/", os.ModePerm)
	os.Create(pathname)

	err := ioutil.WriteFile(pathname, *stream, 0644)
	if err != nil {
		log.Fatal(err)
	}
}