package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var surName, name, thirdName, group string
	var maxLen int

	fmt.Println("Введите ФИО через пробел:")
	fmt.Scan(&surName, &name, &thirdName)

	fmt.Println("Введите название группы:")
	fmt.Scan(&group)

	labaStr := "| Лабораторная работа №1 |"
	nameStr := "| " + surName + " " + name + " " + thirdName + " |"
	groupStr := "| Выполнил: ст. гр. " + group + " |"

	if utf8.RuneCountInString(nameStr) >= utf8.RuneCountInString(groupStr) {
		maxLen = utf8.RuneCountInString(nameStr)

	} else {
		maxLen = utf8.RuneCountInString(groupStr)
	}

	spaceFull := maxLen - utf8.RuneCountInString(labaStr)

	var array []rune

	for i := 0; i < spaceFull; i++ {
		space := ' '
		array = append(array, space)
	}

	spaceFullStr := string(array)

	labaStrFull := "| Лабораторная работа №1" + spaceFullStr + " |"

	spaceFull1 := maxLen - utf8.RuneCountInString(groupStr)

	var array1 []rune

	for i := 0; i < spaceFull1; i++ {
		space := ' '
		array1 = append(array1, space)
	}

	spaceFullStr1 := string(array1)

	groupStrFull := "| Выполнил: ст. гр. " + group + spaceFullStr1 + " |"

	spaceFull2 := maxLen - utf8.RuneCountInString(nameStr)

	var array2 []rune

	for i := 0; i < spaceFull2; i++ {
		space := ' '
		array2 = append(array2, space)
	}

	spaceFullStr2 := string(array2)

	nameStrFull := "| " + surName + " " + name + " " + thirdName + spaceFullStr2 + " |"

	var array3 []rune

	for j := 0; j < maxLen; j++ {
		roof := '-'
		array3 = append(array3, roof)
	}

	roofFull := string(array3)
	floor := roofFull

	fmt.Printf("%s\n%s\n%s\n%s\n%s", roofFull, labaStrFull, nameStrFull, groupStrFull, floor)
}
