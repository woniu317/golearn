package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p,%d\n", &i, i)
		}()
	}
	//fmt.Println("hello")
	time.Sleep(100000)
}
