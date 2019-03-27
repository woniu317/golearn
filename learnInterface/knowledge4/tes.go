package main

import "fmt"

type Animal interface {
	Id() int
	//Name() string
}

//type Dog interface {
//	Animal
//	Pet
//}

type mys struct {
}

func (m mys) Id() int {
	fmt.Println("hellow")
	return 0
}

type Pet interface {
	Name() string
	Category() string
	Animal
}

func main() {
	mm := mys{}
	//mm.Id()
	var ani Animal
	ani = mm
	ani.Id()
}
