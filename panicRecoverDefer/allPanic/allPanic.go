package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("panic,rec,", rec)
		}
	}()
	fmt.Println("Enter function main.")
	go caller3()
	caller1()
	fmt.Println("Exit function main.")

	time.Sleep(100000000)
	fmt.Println("sfs")
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2() {
	time.Sleep(100000)
	panic("new panic")
}

func caller3() {
	fmt.Println("Enter function caller3.")
	time.Sleep(10000000)
	fmt.Println("Exit function caller3.")
}
