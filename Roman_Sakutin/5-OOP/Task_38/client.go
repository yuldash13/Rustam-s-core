package main

import "math/rand"

type Client struct {
	name        string
	money       int
	basket      []*Product
	bag         []*Product
	superMarket *SuperMarket
}

func NewClient(name string, money int, superMarket *SuperMarket) *Client {
	return &Client{
		name:        name,
		money:       money,
		basket:      make([]*Product, 0),
		bag:         make([]*Product, 0),
		superMarket: superMarket,
	}
}

func (c *Client) JoinQueue() {
	c.superMarket.queue.Enqueue(c)
}

func (c *Client) TakeProduct() {
	product := c.superMarket.storage.GetRandomProduct()
	c.basket = append(c.basket, product)
}

func (c *Client) PayProducts() int {
	var cost int
	for _, r := range c.basket {
		cost += r.cost
	}
	if cost > c.money {
		id := rand.Intn(len(c.basket))
		c.basket = append(c.basket[:id], c.basket[id+1:]...)
		return c.PayProducts()
	}
	c.money -= cost
	c.SwitchProducts()
	return cost
}

func (c *Client) SwitchProducts() {
	c.bag = c.basket
	c.basket = make([]*Product, 0)
}
