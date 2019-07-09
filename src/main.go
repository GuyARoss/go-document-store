package main

import (
	"net/http"
	"fmt"
	"manager/utilities"
	"io/ioutil"
)

func main() {
	session := Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/", session.storeHandler)

	fmt.Println("Server Starting on Port 4203")
	http.ListenAndServe(":4203", mux)
}

func (session *SessionSettings) storeHandler(w http.ResponseWriter, req *http.Request) {
	instanceName := utilities.GetParam(req, "id")

	// always need a reference as to what we are updating/ getting
	// will never implement get all
	if instanceName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch(req.Method) {
	case "GET":
		session.handleGet(w, instanceName)

		return
	case "POST":
		session.handlePost(w, req, instanceName)

		return
default: 
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (session *SessionSettings) handlePost(w http.ResponseWriter, req *http.Request, instanceName string) {
	instance := session.GetInstance(instanceName)
	
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	if instance == nil {
		session.CreateInstance(instanceName, &body)

		w.WriteHeader(http.StatusCreated)
		return
	}

	fmt.Print("passed!01")

	instance.Update(&body)
	w.WriteHeader(http.StatusOK)
}

func (session *SessionSettings) handleGet(w http.ResponseWriter, instanceName string) {
	instance := session.GetInstance(instanceName)
	if instance == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Write(*instance.Get())
	w.WriteHeader(http.StatusOK)
}