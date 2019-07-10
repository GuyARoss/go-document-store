package store

import (
	"manager/utilities"
)

// Instance type for the store
type Instance struct {
	stream 				*[]byte
	Name				string
}

// Update updates the instance stream with a new stream
func (instance *Instance) Update(newStream *[]byte) {
	instance.stream = newStream
	utilities.CopyToIo(instance.Name, newStream)
}

// Get retreives the current instances stream
func (instance *Instance) Get() *[]byte {
	return instance.stream
}

// CreateInstance creates a new instance
func CreateInstance(instanceName string) *Instance {
	stream := utilities.UpdateFromIo(instanceName)

	return &Instance{
		stream: stream,
		Name: instanceName,
	}
}