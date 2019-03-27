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
	kingdom        int
	scientificName string
	AnimalCategory
}

func (ac Animal) categoryName() {
	fmt.Println("Animal categoryName ...")
}

func main() {
	category := AnimalCategory{species: "cat", kingdom: "kingdom"}
	animal := Animal{scientificName: "American Shorthair", AnimalCategory: category, kingdom: 16}
	animal.categoryName()
	fmt.Println(animal.kingdom)

}
