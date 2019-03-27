package main

import "fmt"

type AnimalCategory struct {
	kingdom string // 界。
	species string // 种。
}

func (ac AnimalCategory) categoryName() {
	fmt.Println("AnimalCategory categoryName...")
}

type AnimalCategory1 struct {
	kingdom string // 界。
	species string // 种。
}

func (ac AnimalCategory1) categoryName() {
	fmt.Println("AnimalCategory categoryName...")
}

type Animal struct {
	scientificName string
	AnimalCategory
	AnimalCategory1
}

func (ac Animal) categoryName() {
	fmt.Println("Animal categoryName...")
}
func main() {
	category := AnimalCategory{species: "cat", kingdom: "kingdom"}
	category1 := AnimalCategory1{species: "cat1", kingdom: "kingdom1"}
	animal := Animal{
		scientificName:  "American Shorthair",
		AnimalCategory:  category,
		AnimalCategory1: category1,
	}
	fmt.Println(animal)
	animal.categoryName()
}
