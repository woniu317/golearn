package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
	id   int
}

func (dog Dog) Name() string {
	return dog.name
}

//
//func (dog Dog) String() string {
//	fmt.Print("hello dog")
//	return ""
//}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	//pointerTest()
	notpointerTest()
}

func pointerTest() {
	var dog *Dog = nil
	var pet Pet
	fmt.Println(pet == nil)
	fmt.Println(dog == nil)
	fmt.Println(pet, reflect.TypeOf(pet)) //<nil> <nil>
	pet = dog
	fmt.Println(pet, reflect.TypeOf(pet)) //<nil> *main.Dog
	fmt.Println(pet == nil)
}

type II *int

func notpointerTest() {
	var i II = nil
	fmt.Println(i)
	var dog Dog
	var pet Pet
	fmt.Println(dog, reflect.TypeOf(dog))
	fmt.Println(pet, reflect.TypeOf(pet)) //<nil> <nil>
	pet = dog
	fmt.Println(pet, reflect.TypeOf(pet)) //{} *main.Dog
}
