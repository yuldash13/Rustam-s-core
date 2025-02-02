package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "to be or not to be that is the question"
	str2 := "корабли лавировали лавировали да не вылавировали"
	countSimilar(str1)
	countSimilar(str2)
}

func countSimilar(str string) {
	arr := strings.Split(str, " ")
	strMap := make(map[string]int)
	for _, r := range arr {
		strMap[r] += 1
	}
	fmt.Println(strMap)
}
