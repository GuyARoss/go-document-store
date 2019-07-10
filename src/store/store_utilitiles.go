package store

import (
	"os"
	"fmt"
	
	"log"
	"io/ioutil"
)

func updateFromIo(name string) *[]byte {
	pname := fmt.Sprintf("tmp/%s", name)

	file, _ := os.Open(pname)
    defer file.Close()

	b, _ := ioutil.ReadAll(file)
	fmt.Println(string(b))
	return &b
}

func copyToIo(name string, stream *[]byte) {
	pathname := fmt.Sprintf("tmp/%s", name)
	os.MkdirAll("tmp/", os.ModePerm)
	os.Create(pathname)

	err := ioutil.WriteFile(pathname, *stream, 0644)
	if err != nil {
		log.Fatal(err)
	}
}