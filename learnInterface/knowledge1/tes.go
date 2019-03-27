package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	SayHello(string) string
}
type MyStruct struct{}

func (my MyStruct) SayHello(string) string {
	fmt.Printf("hello")
	return ""
}
func main() {
	var myI string
	fmt.Println(reflect.TypeOf(myI))
	var mys MyStruct
	fmt.Println(reflect.TypeOf(mys))
	//myStruct := MyStruct{}
	//myI = myStruct
	fmt.Println(reflect.TypeOf(myI))
	//myI.SayHello("")
}
