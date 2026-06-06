package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "mfgah134517095aldrfgvh8h"
	fmt.Println(a)

	a = removeNumber(a)
	fmt.Println(a)
}

func removeNumber(str string) string {
	var newStr []rune
	for _, r := range str {
		if unicode.IsDigit(r) {
			continue
		}
		newStr = append(newStr, r)
	}
	return string(newStr)
}
