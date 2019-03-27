package main

import "fmt"

type AnimalCategory struct {
	kingdom string // 界。
	species string // 种。
}

func (ac AnimalCategory) categoryName() {
	fmt.Println("AnimalCategory categoryName...")
}

type Animal struct {
	scientificName string
	AnimalCategory
}

func main() {
	category := AnimalCategory{species: "cat", kingdom: "kingdom"}
	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	animal.categoryName()

}
