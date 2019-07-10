package main

import (
	"fmt"
	"io/ioutil"
	"manager/store"
)

type SessionSettings struct {
	Instances 	map[string]*store.Instance
}

func (session *SessionSettings) GetInstance(instanceName string) *store.Instance {
	return session.Instances[instanceName]
}

func (session *SessionSettings) CreateInstance(instanceName string, data *[]byte) {
	createdInstance := store.CreateInstance(instanceName)
	createdInstance.Update(data)

	session.Instances[instanceName] = createdInstance
}

func Init() *SessionSettings {
	instances := make(map[string]*store.Instance)

	files, _ := ioutil.ReadDir("./tmp")
	for _, f := range files {
		fmt.Println(f.Name())
		instances[f.Name()] = store.CreateInstance(f.Name())
	}

	return &SessionSettings{
		Instances: instances,
	}
}