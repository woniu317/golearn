package main

import "fmt"

func fn1(m map[int]int) {
	m = make(map[int]int)
}

func main() {
	var m map[int]int = make(map[int]int)
	fn1(m)
	m[0] = 1
	fmt.Println(m == nil)
}
