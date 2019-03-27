package main

import (
	"errors"
	"fmt"
)

type operate1 func(x, y int) int

var b int

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate1) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	op := func(x, y int) int {
		return x + y
	}
	x, y := 12, 23
	add := genCalculator(op)
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
