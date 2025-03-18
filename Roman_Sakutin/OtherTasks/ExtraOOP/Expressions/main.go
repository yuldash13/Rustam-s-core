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
			exp1 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).string()

			exp2 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).string()
			fmt.Printf("%v = %v\n", exp1, exp2)

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

			exp3 := NewMultiply(
				NewConst(2),
				NewVariable("x"),
			).string()

			exp4 := NewMultiply(
				NewVariable("x"),
				NewConst(2),
			).string()
			fmt.Printf("%v = %v\n", exp3, exp4)

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

			exp5 := NewAdd(
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
			).string()

			exp6 := NewSubtract(
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
				NewMultiply(
					NewConst(2),
					NewVariable("x"),
				),
			).string()
			fmt.Printf("%v = %v\n", exp5, exp6)

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
