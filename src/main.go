package main

import (
	"net/http"
)

func main() {
	mux := http.newMuxServer()

	mux.HandleFunc("/", storeHandler)

	http.ListenAndServe(":4203", mux)
}

func storeHandler(w http.ResponseWriter, req *http.Request) {
		
}