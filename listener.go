package main

var listeners []Listener

type Listener struct {
	Route string
	Func  func(message Message)
}

func AddListener(listener Listener) {
	listeners = append(listeners, listener)
}

func ClearListeners() {
	listeners = []Listener{}
}

func RegisterListener() {
	for {
		message := <-getChannel()
		for _, listener := range listeners {
			if listener.Route == message.Route {
				go listener.Func(message)
			}
		}
	}
}
