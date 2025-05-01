package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dataBase := NewDataBase()
	dataBase.InitDataBase()

	queue := NewQueue(10, 0, 0)

	superMarket := NewSuperMarket("Magnet", dataBase, queue)

	clients := []*Client{
		NewClient("Yuldash", 600, superMarket),
		NewClient("Rustam", 500, superMarket),
		NewClient("Nazar", 400, superMarket),
	}

	for _, r := range clients {
		quantityProdacts := 5
		for i := 0; i < quantityProdacts; i++ {
			r.TakeProduct()
		}
		r.JoinQueue()
	}

	seller := NewSeller("Samir", superMarket)

	for i := 0; i < len(clients); i++ {
		seller.NextClient()
	}

	fmt.Printf("Super Market got %d money\n", superMarket.money)
	for _, r := range clients {
		fmt.Printf("Client %s have %d money and got:\n", r.name, r.money)
		for _, product := range r.bag {
			fmt.Printf("-%s\n", product.name)
		}
	}
}
