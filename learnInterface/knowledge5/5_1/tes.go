package main

import (
	"fmt"
)

type Dog struct {
	name string // 名字。
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}
