package store

type Instance type {
	stream 			chan *[]byte
	Name				string
	Update			interface{}
	Get					interface{}
}

func (instance *Instance) update(newStream *[]byte) {
	instance.stream <- newStream

	go copyToIo(instance.Name, *[]byte)
}

func (instance *Instance) get() *[]byte {
	return <-instance.stream
}

func store(string instanceName) chan *[]byte {
	stream := make(chan *[]byte)
	
	go func() {
		for {
			updateFromIo(instanceName) // in the case some other utility updates it.
			<-time.After(time.Second * 2) 
		}
	}
}

func CreateInstance(string instanceName) *Instance {
	stream := store(instanceName)

	return &Instance{
		stream: stream,
		Name: instanceName,
		Update: update,
		Get: get,
	}
}