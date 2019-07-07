package store

// type Instance interface {
// 	Overwrite()
// 	Retrieve()
// }

type Instance type {
	OutStream chan *[]byte
}

func Overwrite() {

}

func Retrieve() {

}

func store() (chan *[]byte, ) {
	stream := make(chan *[]byte)
	
	go func() {
		for {
			copyToIo()
			<-time.After(time.Second * 2) // copies to every ~2sec
		}
	}
}

func CreateInstance() chan *[]byte {
	return 
}