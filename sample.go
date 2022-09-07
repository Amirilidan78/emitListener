package emitListener

import (
	"fmt"
	"time"
)

func main() {

	// register listener
	go RegisterListener()

	// add listener to receiver message on specific route
	AddListener(Listener{
		Route: "route-name",
		Func: func(message Message) {
			fmt.Println(message)
		},
	})

	var i = 1

	for {
		i++
		// emit on route
		go Emit(Message{
			Route: "route-name",
			Arguments: map[string]any{
				"arg1": 0,
				"arg2": "string",
			},
		})
		time.Sleep(time.Second)
	}

	// you can clear listeners and register them again during running application
	ClearListeners()

}
