package store

type Instance struct {
	stream 				*[]byte
	Name				string
}

func (instance *Instance) Update(newStream *[]byte) {
	instance.stream = newStream
	copyToIo(instance.Name, newStream)
}

func (instance *Instance) Get() *[]byte {
	return instance.stream
}

func CreateInstance(instanceName string) *Instance {
	stream := updateFromIo(instanceName)

	return &Instance{
		stream: stream,
		Name: instanceName,
	}
}