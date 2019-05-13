package main

import (
	"fmt"
	"roobo.com/golearn/visitUnexpectedParam/b"
)

type StatusCode string

func main() {
	s := b.Greet("world")
	fmt.Println(s)
	s = b.Hi("hi")
	fmt.Println(s)
	var sc StatusCode = "abc"
	fmt.Println(sc == "abc")
}
