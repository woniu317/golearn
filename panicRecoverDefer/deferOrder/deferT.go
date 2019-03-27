package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	fmt.Println("main...")
}
