package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Se inicializa cond
	c := sync.NewCond(&sync.Mutex{})    // <1>
	queue := make([]interface{}, 0, 10) // <2>

	removeFromQueue := func(delay time.Duration) {
		//Se agrega un delay para llenar la queue a proposito
		time.Sleep(delay)
		c.L.Lock()        // <8>
		queue = queue[1:] // <9>
		fmt.Println("Removed from queue")
		c.L.Unlock() // <10>
		c.Signal()   // <11>
	}

	//Lockeo Appendeo hasta que la cola tiene longitud 2.
	//Mientras eso no ocurro itero haciendo lock y unlock para no tener condiciones de carrera con removeFromQueue
	//Que comparten la queue
	// Cuando la queue tiene longitud 2, se suspende el thread y se espera hasta que alguien le de signal
	for i := 0; i < 10; i++ {
		c.L.Lock()            // <3>
		for len(queue) == 2 { // <4>
			c.Wait() // <5>
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // <6>
		c.L.Unlock()                        // <7>
	}
}
