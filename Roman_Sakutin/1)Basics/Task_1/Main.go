package main

import "fmt"

func main() {

	//minecraft coordinates
	var (
		x int32 = 1234
		y int32 = 56
		z int32 = 7890
	)

	//name
	var (
		name string = "Yuldash"
	)

	//body sizes
	var (
		growth   int32 = 181
		weight   int32 = 85
		footsize int32 = 42
	)

	//money on a card
	var (
		visa       float32 = 1567.52
		mastercard float32 = 45342.97
		world      float32 = 0.00
	)

	fmt.Println("/tp @s", x, y, z)
	fmt.Println("My name is", name)
	fmt.Println("My growth is", growth, ", weight:", weight, ", footsize:", footsize)
	fmt.Println("There are money on my cards: visa -", visa, ", mastercard -", mastercard, ", world -", world)
}
