package main

import "math/rand"

type DataBase struct {
	products map[int]*Product
}

func NewDataBase() *DataBase {
	return &DataBase{products: make(map[int]*Product)}
}

func (DB *DataBase) AddProduct(product *Product) {
	DB.products[product.ID] = product
}

func (DB *DataBase) InitDataBase() {
	DB.AddProduct(NewProduct(0, "Chocolate", 80))
	DB.AddProduct(NewProduct(1, "Chips", 60))
	DB.AddProduct(NewProduct(2, "Coke", 100))
	DB.AddProduct(NewProduct(3, "Cookies", 40))
	DB.AddProduct(NewProduct(4, "Cake", 200))
}

func (DB *DataBase) GetRandomProduct() *Product {
	id := rand.Intn(len(DB.products))
	return DB.products[id]
}
