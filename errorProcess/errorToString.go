package main

import "fmt"

func New(str string) error {
	return &MyError{s: str}
}

type MyError struct {
	s string
}

func (myError MyError) String() string {
	fmt.Println("my String()")
	return ""
}
func (myError MyError) Error() string {
	fmt.Println("my Error()")
	return ""
}

func main() {
	er := New("hello")
	fmt.Println(er)
}
