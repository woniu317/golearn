package main

import (
	"fmt"
	"time"
)

func main() {
	var maxRoutineNum = 5
	ch := make(chan struct{}, maxRoutineNum)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		go func(v int) {
			fmt.Println(v)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	fmt.Println("hello")
	time.Sleep(100000)
}
