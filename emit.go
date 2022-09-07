package main

func Emit(message Message) {
	getChannel() <- message
}
