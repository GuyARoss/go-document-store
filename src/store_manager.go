package main

import (
	"io/ioutil"
	"manager/store"
)

type SessionSettings type {
	Instances 	map[string]manger.Instance
}

func (session *SessionSettings) GetInstance(instanceName string) {
	return session.Instances[instanceName]
}

func Initialize() {
	instances := make(map[string]manager.Instance)

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}