package main

import (
	"fmt"
	"sync"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) getName(mutex sync.Mutex) {
	fmt.Println("get name...")
}

func (s Student) getAge(ch chan<- int) {
	ch <- 1
}
func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}

func main() {
	var s *Student
	StudentRegister("Jim", 18)
	fmt.Println(s.Name)
}
