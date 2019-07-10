package store

import (
	"io/ioutil"
)

// Store type that holds the instances
type Store struct {
	Instances 	map[string]*Instance
}

// GetInstance retrieves an instance from the instance name
func (session *Store) GetInstance(instanceName string) *Instance {
	return session.Instances[instanceName]
}

// CreateInstance creates an instance 
func (session *Store) CreateInstance(instanceName string, data *[]byte) {
	createdInstance := CreateInstance(instanceName)
	createdInstance.Update(data)

	session.Instances[instanceName] = createdInstance
}

// Init creates an instance of the store
func Init() *Store {
	instances := make(map[string]*Instance)

	files, _ := ioutil.ReadDir("./tmp")
	for _, f := range files {
		instances[f.Name()] = CreateInstance(f.Name())
	}

	return &Store{
		Instances: instances,
	}
}