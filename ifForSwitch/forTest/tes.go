package main

import "fmt"

func main() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		fmt.Println(i)
	}
	fmt.Println()
}
