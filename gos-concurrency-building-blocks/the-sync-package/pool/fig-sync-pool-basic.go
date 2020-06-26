package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	// Crea una instancia
	myPool.Get()             // <1>
	// Crea una instancia
	instance := myPool.Get() // <1>
	// Volves a dejar disponible la instancia
	myPool.Put(instance)     // <2>
	// Utilizas la que dejaste disponible
	myPool.Get()             // <3>
}
