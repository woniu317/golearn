package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var capacity = 10
var producerNum = 5
var consumerNum = 3

var lock sync.RWMutex
var senCond = sync.NewCond(&lock)

//var recvCond = sync.NewCond(lock.RLocker())
var recvCond = sync.NewCond(&lock)

func main() {

	rand.Seed(time.Now().UnixNano())

	quit := make(chan bool)
	product := make(chan int, capacity)

	producer(product)
	consumer(product)

	<-quit
}

func producer(out chan<- int) {
	for i := 0; i < producerNum; i++ {
		go func(nu int) {
			for {
				lock.Lock()
				for len(out) == capacity {
					fmt.Println("capacity is full,stop produce")
					senCond.Wait()
				}
				num := rand.Intn(100)
				out <- num
				fmt.Printf("Produce %d produce: num %d\n", nu, num)
				lock.Unlock()
				recvCond.Signal()
				time.Sleep(time.Second)
			}
		}(i)

	}
}

func consumer(in <-chan int) {
	for i := 0; i < consumerNum; i++ {
		go func(nu int) {
			for {
				//lock.RLock()
				lock.Lock()
				for len(in) == 0 {
					fmt.Println("capacity is empty,stop consume")
					recvCond.Wait()
				}
				num := <-in
				fmt.Printf("Goroutine %d: consume num %d\n", nu, num)
				lock.Unlock()
				senCond.Signal()
				time.Sleep(time.Second)
			}
		}(i)
	}
}
