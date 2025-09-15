package main

type Product struct {
	ID   int
	name string
	cost int
}

func NewProduct(id int, name string, cost int) *Product {
	return &Product{
		ID:   id,
		name: name,
		cost: cost,
	}
}
