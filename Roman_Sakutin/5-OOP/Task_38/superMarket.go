package main

type SuperMarket struct {
	name    string
	money   int
	storage *DataBase
	queue   *Queue
}

func NewSuperMarket(name string, storage *DataBase, queue *Queue) *SuperMarket {
	return &SuperMarket{
		name:    name,
		money:   0,
		storage: storage,
		queue:   queue,
	}
}
