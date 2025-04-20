package main

import (
	"fmt"
	"github.com/danilobandeira29/specification-pattern/repository"
)

func main() {
	repo := repository.New()
	nameContains := repository.NameContainsSpecification{
		Name: "Shoes",
	}
	priceBelow := repository.PriceBelowSpecification{Value: 121.00}
	and := repository.AndSpecification{
		Left:  nameContains,
		Right: priceBelow,
	}
	products := repo.FindBy(and)
	for _, p := range products {
		fmt.Println(p)
	}
}
