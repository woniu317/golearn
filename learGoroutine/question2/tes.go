package main

import (
	"fmt"
	//"time"
	"sync"
	"time"
)

func main() {
	method1()
	method2()
	method3()
}

func method1() {
	num := 10
	for i := 0; i < num; i++ {
		go func() {
			fmt.Println("method1", i)
		}()
	}

	time.Sleep(time.Millisecond * 500)
	fmt.Println(">>>method1 is done")
}

func method2() {
	num := 10
	sign := make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func() {
			fmt.Println("method2", i)
			sign <- struct{}{}
		}()
	}
	for j := 0; j < num; j++ {
		<-sign

	}
	fmt.Println(">>>method2 is done")
}

func method3() {
	num := 10
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			fmt.Println("method3", i)
		}()
	}
	wg.Wait()
	fmt.Println(">>>method3 is done")
}
