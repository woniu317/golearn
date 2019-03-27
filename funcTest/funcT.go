package main

import "fmt"

type Printer func(content string, id int) (n int, err error)
type Printer1 func(string, int) (int, error)
type Printer2 func(int, string) (int, error)

func main() {
	var printer Printer
	printer = echo
	printer("printer", 0)
	var printer1 Printer1
	printer1 = echo
	printer1("printer1", 1)
	//var printer2 Printer2
	//printer2 = echo
}

func echo(content string, id int) (n int, err error) {
	fmt.Println(content, id)
	return 0, nil
}
