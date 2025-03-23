package main

import (
	"fmt"
)

type Expression interface {
	evaluate(x int) int
	string() string
	equals(Expression) bool
}

func main() {
	fmt.Println("Уравнение:\nx^2 - 2x + 1")
	fmt.Println()

	for {
		var menu string
		fmt.Println("Меню:\n1)Покажи\n2)Вычисли\n3)Проверь\n4)Выход")
		fmt.Println()
		fmt.Scan(&menu)

		switch menu {
		case "Покажи":
			exp := NewAdd(
				NewSubtract(
					NewMultiply(
						NewVariable("x"),
						NewVariable("x"),
					),
					NewMultiply(
						NewConst(2),
						NewVariable("x"),
					),
				),
				NewConst(1),
			).string()
			fmt.Println(exp)
			fmt.Println()
		case "Вычисли":
			exp := NewAdd(
				NewSubtract(
					NewMultiply(
						NewVariable("x"),
						NewVariable("x"),
					),
					NewMultiply(
						NewConst(2),
						NewVariable("x"),
					),
				),
				NewConst(1),
			).evaluate(5)
			fmt.Println(exp)
			fmt.Println()
		case "Проверь":
			exp01 := NewConst(1).string()
			exp02 := NewVariable("x").string()
			fmt.Printf("%v = %v\n", exp01, exp02)
			eq0 := NewConst(1).equals(NewVariable("x"))
			fmt.Println(eq0)
			fmt.Println()

			exp11 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).string()

			exp12 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).string()
			fmt.Printf("%v = %v\n", exp11, exp12)

			eq1 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).equals(NewMultiply(
				NewConst(2),
				NewVariable("x"),
			),
			)
			fmt.Println(eq1)
			fmt.Println()

			exp21 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).string()

			exp22 := NewMultiply(
				NewVariable("x"),
				NewConst(2),
			).string()
			fmt.Printf("%v = %v\n", exp21, exp22)

			eq2 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).equals(NewMultiply(
				NewVariable("x"),
				NewConst(2),
			),
			)
			fmt.Println(eq2)
			fmt.Println()

			exp31 := NewAdd(
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
			).string()

			exp32 := NewSubtract(
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
			).string()
			fmt.Printf("%v = %v\n", exp31, exp32)

			eq3 := NewAdd(
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
			).equals(NewSubtract(
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
			),
			)
			fmt.Println(eq3)
			fmt.Println()
		case "Выход":
			return
		default:
			fmt.Println("Вы ввели что-то не то.")
			fmt.Println()
		}
	}
}
