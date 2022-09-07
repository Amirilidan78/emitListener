package emitListener

func Emit(message Message) {
	getChannel() <- message
}
