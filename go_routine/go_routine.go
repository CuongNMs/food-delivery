package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increament()
	}
	wg.Wait()
}

func increament() {
	counter++
	m.Unlock()
	wg.Done()
}

func sayHello() {
	fmt.Println("Hello ", counter)
	m.RUnlock()
	wg.Done()
}

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})
