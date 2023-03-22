package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchAPI(url string) string {
	time.Sleep(time.Duration(rand.Int31n(1e3)) * time.Millisecond)
	return "data: " + url
}

func main() {
	responseChan := make(chan string)
	var result []string
	go func() {
		responseChan <- fetchAPI("http://google.com")
	}()

	go func() {
		responseChan <- fetchAPI("http://youtube.com")
	}()

	go func() {
		responseChan <- fetchAPI("http://facebook.com")
	}()
	for i := 0; i < 3; i++ {
		result = append(result, <-responseChan)
	}
	fmt.Println(result)
}
