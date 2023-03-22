package main

import (
	"fmt"
	"sync"
)

func main() {
	chans := make(chan int, 10)
	wg := sync.WaitGroup{}
	//for i := 0; i < 10; i++ {
	wg.Add(2)
	//go sayHello(chans, &wg)
	//go increment(chans, &wg)
	go func(ch chan int) {

	}(chans)
	go func(ch chan int) {
		//counter := <-ch
		//counter++
		//chans <- counter
		//wg.Done()
		for i := 0; i < 10; i++ {
			chans <- i
		}
		close(chans)
	}(chans)
	//}
	wg.Wait()
}

func sayHello(chans chan int, group *sync.WaitGroup) {
	counter := <-chans
	fmt.Println("Hello ", counter)
	group.Done()
}

func increment(chans chan int, group *sync.WaitGroup) {
	counter := <-chans
	counter++
	chans <- counter
	group.Done()
}
