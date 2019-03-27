package main

import "fmt"

func sui() int8 {
	fmt.Println("sui")
	return 1
}
func main() {
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[2] {
	case sui(), 1 + 9:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
