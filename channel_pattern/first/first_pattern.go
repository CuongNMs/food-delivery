package main

import (
	"fmt"
	"time"
)

func startSender(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 5; i++ {
			c <- (name + " hello")
			time.Sleep(time.Second)
		}
	}()
	return c
}
func main() {
	sender1 := startSender("Ti")
	sender2 := startSender("Teo")
	for i := 1; i <= 5; i++ {
		select {
		case msgTi := <-sender1:
			fmt.Println(msgTi)
		case msgTeo := <-sender2:
			fmt.Println(msgTeo)
		}

	}
}
