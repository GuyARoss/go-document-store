package main

import (
	"fmt"
	"io/ioutil"
	"manager/store"
	"path/filepath"
)

type SessionSettings struct {
	Instances 	map[string]*store.Instance
}

func (session *SessionSettings) GetInstance(instanceName string) *store.Instance {
	return session.Instances[instanceName]
}

func (session *SessionSettings) CreateInstance(instanceName string, data *[]byte) {
	createdInstance := store.CreateInstance(instanceName)
	fmt.Print("passed!02")
	createdInstance.Update(data)

	session.Instances[instanceName] = createdInstance
}

func Init() *SessionSettings {
	instances := make(map[string]*store.Instance)

	absFilePath, _ := filepath.Abs("../temp/")
	files, _ := ioutil.ReadDir(absFilePath)

	for _, f := range files {
		instances[f.Name()] = store.CreateInstance(f.Name())
	}

	return &SessionSettings{
		Instances: instances,
	}
}