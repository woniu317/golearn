package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	fmt.Println("start...")
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		//cΩ® h <- 5
	}()
	go func(chan int) {
		a := <-ch
		fmt.Println(a)
		a = <-ch
		fmt.Println("--------")
		fmt.Println(a)
	}(ch)
	go chanRange(ch)
	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("over!!!")
}

func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
	fmt.Println("chan range over!")
}
