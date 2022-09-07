package emitListener

import (
	"sync"
)

var lock = &sync.Mutex{}

var channel chan Message

func createChannel() {
	lock.Lock()
	defer lock.Unlock()
	if channel == nil {
		channel = make(chan Message)
	}
}

func getChannel() chan Message {
	if channel == nil {
		createChannel()
	}
	return channel
}
