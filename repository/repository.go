package repository

import (
	"strings"
)

type Product struct {
	Name     string
	Category string
	Price    float64
	IsActive bool
}

type Specification interface {
	IsSatisfiedBy(p Product) bool
}

type Repository interface {
	FindBy(s Specification) []Product
}

type ProductRepository struct {
	products []Product
}

func (p *ProductRepository) FindBy(s Specification) []Product {
	var result []Product
	for _, prod := range p.products {
		if s.IsSatisfiedBy(prod) {
			result = append(result, prod)
		}
	}
	return result
}

func New() Repository {
	return &ProductRepository{
		products: []Product{
			{Name: "Nike Running Shoes", Category: "Footwear", Price: 120.00, IsActive: true},
			{Name: "Adidas Running Shoes", Category: "Footwear", Price: 160.00, IsActive: true},
			{Name: "Adidas Shoes", Category: "Footwear", Price: 100.00, IsActive: true},
			{Name: "Adidas Sneakers", Category: "Footwear", Price: 110.00, IsActive: true},
			{Name: "Leather Boots", Category: "Footwear", Price: 150.00, IsActive: false},
			{Name: "Slim Fit T-Shirt", Category: "Clothing", Price: 25.00, IsActive: true},
			{Name: "Denim Jeans", Category: "Clothing", Price: 60.00, IsActive: true},
			{Name: "Hoodie Jacket", Category: "Clothing", Price: 45.00, IsActive: false},
			{Name: "Bluetooth Headphones", Category: "Electronics", Price: 89.99, IsActive: true},
			{Name: "Smartphone Case", Category: "Electronics", Price: 15.00, IsActive: true},
			{Name: "4K Monitor", Category: "Electronics", Price: 320.00, IsActive: false},
			{Name: "Wireless Mouse", Category: "Electronics", Price: 29.99, IsActive: true},
			{Name: "Mechanical Keyboard", Category: "Electronics", Price: 99.99, IsActive: true},
			{Name: "Gaming Chair", Category: "Furniture", Price: 220.00, IsActive: true},
			{Name: "Office Desk", Category: "Furniture", Price: 180.00, IsActive: true},
			{Name: "Bookshelf", Category: "Furniture", Price: 85.00, IsActive: false},
			{Name: "LED Lamp", Category: "Home", Price: 35.00, IsActive: true},
			{Name: "Coffee Maker", Category: "Home", Price: 70.00, IsActive: true},
			{Name: "Blender", Category: "Home", Price: 55.00, IsActive: false},
			{Name: "Water Bottle", Category: "Accessories", Price: 12.00, IsActive: true},
			{Name: "Sunglasses", Category: "Accessories", Price: 40.00, IsActive: true},
			{Name: "Backpack", Category: "Accessories", Price: 50.00, IsActive: false},
			{Name: "Perfume", Category: "Beauty", Price: 65.00, IsActive: true},
			{Name: "Lipstick", Category: "Beauty", Price: 20.00, IsActive: true},
			{Name: "Foundation", Category: "Beauty", Price: 30.00, IsActive: false},
			{Name: "Yoga Mat", Category: "Sports", Price: 25.00, IsActive: true},
			{Name: "Dumbbell Set", Category: "Sports", Price: 90.00, IsActive: true},
			{Name: "Tennis Racket", Category: "Sports", Price: 75.00, IsActive: false},
			{Name: "Cookbook", Category: "Books", Price: 18.00, IsActive: true},
			{Name: "Science Fiction Novel", Category: "Books", Price: 22.00, IsActive: true},
			{Name: "Notebook", Category: "Stationery", Price: 5.00, IsActive: true},
			{Name: "Pen Pack", Category: "Stationery", Price: 3.50, IsActive: true},
		},
	}
}

type PriceBelowSpecification struct {
	Value float64
}

func (p PriceBelowSpecification) IsSatisfiedBy(pro Product) bool {
	return pro.Price < p.Value
}

type NameContainsSpecification struct {
	Name string
}

func (n NameContainsSpecification) IsSatisfiedBy(pro Product) bool {
	return strings.Contains(strings.ToLower(pro.Name), strings.ToLower(n.Name))
}

type AndSpecification struct {
	Left, Right Specification
}

func (a AndSpecification) IsSatisfiedBy(pro Product) bool {
	return a.Left.IsSatisfiedBy(pro) && a.Right.IsSatisfiedBy(pro)
}
