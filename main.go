package main

import (
	"fmt"
	repositoryV2 "github.com/danilobandeira29/specification-pattern/repository/v2"
)

func main() {
	repo := repositoryV2.New()
	spec := repositoryV2.NameContains("Shoes").
		And(repositoryV2.PriceBelow(121.00))
	products := repo.FindBy(spec)
	for _, p := range products {
		fmt.Println(p)
	}
}
