package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func addCounter() {
	for count := 0; count < 2; count++ {
		defer func() {
			fmt.Println("<<<")
			lock.Unlock()
		}()
		fmt.Println(">>>")
		//do someThing
		lock.Lock()
	}
}

//func addCounter() {
//	for count := 0; count < 2; count++ {
//		doSomeThing()
//	}
//}

func doSomeThing() {
	lock.Lock()
	defer lock.Unlock()
}
func main() {
	addCounter()
}
