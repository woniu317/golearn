package main

import "fmt"

func main() {
	arrayRange()
	//sliceRange()
}

func arrayRange() {
	numbers2 := [...]int{1, 1, 1, 1, 1, 1}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		fmt.Println(&e)
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()
}

func sliceRange() {
	numbers2 := []int{1, 1, 1, 1, 1, 1}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()
}
