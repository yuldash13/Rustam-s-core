package main

import "fmt"

func main() {
	var (
		a string
		b = "exit"
	)

	for {
		fmt.Scan(&a)
		if a == b {
			return
		}
	}
}
