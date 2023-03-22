package main

import "fmt"

type Test interface {
	PointerCheck() string
}

type check struct {
	a int
}

func (t *check) PointerCheck() {
	fmt.Println(t.a)
}
func main() {
	var t *Test
	var c *check
	c = new(check)
	fmt.Println(t)
	fmt.Println(c)
}
