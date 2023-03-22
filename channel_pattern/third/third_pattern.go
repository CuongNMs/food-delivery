package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchAPI(url string) string {
	time.Sleep(time.Duration(rand.Intn(10e3)) * time.Millisecond)
	return "data: " + url
}
func queryFirst(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(s string) { c <- fetchAPI(s) }(url)
	}
	return c
}

func main() {
	result := queryFirst("a", "b", "c", "d", "e", "f", "g")
	fmt.Println(<-result)
}
