package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dispatcher := NewDispatcher("Yuldash")
	fmt.Println(dispatcher.MakeTrain("SPB - Dubai"))
	fmt.Println(dispatcher.MakeTrain("SPB - Tokio"))
	fmt.Println(dispatcher.MakeTrain("SPB - Uzbekistan"))
}
