package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(11)

	var degree = 0

	for two := 1; randomNum >= two; degree++ {
		two *= 2
	}

	fmt.Printf("%d<2^%d", randomNum, degree)
}
