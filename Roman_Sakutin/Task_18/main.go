package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		skillHero, nameAttackBoss string

		hpHero       = 30
		manaHero     = 40
		attackHero1  = 4
		attackHero2  = 5
		attackHero3  = 8
		manaCostHero = 10
		heal         = 10

		hpBoss       = 50
		attackBoss1  = 6
		attackBoss2  = 3
		attackBoss3  = 8
		manaCostBoss = 2

		m = 0
		n = 0
	)

	for {
		var array, array1 []rune

		for i := 0; i < hpBoss/5; i++ {
			a := '('
			b := '`'
			c := '~'
			d := '`'
			e := ')'
			array = append(array, a)
			array = append(array, b)
			array = append(array, c)
			array = append(array, d)
			array = append(array, e)
		}

		for i := 0; i < hpHero/5; i++ {
			z := '('
			x := '"'
			w := '_'
			y := '"'
			u := ')'
			array1 = append(array1, z)
			array1 = append(array1, x)
			array1 = append(array1, w)
			array1 = append(array1, y)
			array1 = append(array1, u)
		}

		fmt.Println("\u001B[33mХод героя:\u001B[0m")
		fmt.Printf("У вас \u001B[31m%d\u001B[0m здоровья\nУ вас \u001B[34m%d\u001B[0m маны\n", hpHero, manaHero)
		fmt.Println("\u001B[31m" + string(array1) + "\u001B[0m")
		fmt.Println("===========================================================================")
		fmt.Println("\u001B[33mВаши\u001B[0m действия:\n\u001B[33mУдарМечом\u001B[0m  \u001B[33mОгненныйЗалп\u001B[0m  \u001B[33mОблитьМаслом\u001B[0m  \u001B[33mЛечение\u001B[0m")
		fmt.Println("===========================================================================")
		fmt.Scan(&skillHero)
		fmt.Println("===========================================================================")

		switch skillHero {
		case "УдарМечом":
			hpBoss -= attackHero1
			fmt.Printf("\u001B[33mВы\u001B[0m нанесли \u001B[35mБоссу\u001B[0m \u001B[36m%d\u001B[0m урона.\n", attackHero1)
			fmt.Println("===========================================================================")
		case "ОгненныйЗалп":
			manaHero -= manaCostHero

			if manaHero == 0 {
				fmt.Println("У \u001B[33mвас\u001B[0m недостаточно маны. \u001B[33mВы\u001B[0m пропускаете ход.")
				fmt.Println("===========================================================================")
				break
			} else {
				hpBoss -= attackHero2
				m += 1
				fmt.Printf("\u001B[33mВы\u001B[0m нанесли \u001B[35mБоссу\u001B[0m \u001B[36m%d\u001B[0m урона.\n", attackHero2)
				fmt.Println("===========================================================================")
			}
		case "ОблитьМаслом":
			if m == 1 {
				hpBoss -= attackHero3
				m += 1
				fmt.Printf("\u001B[33mВы\u001B[0m нанесли \u001B[35mБоссу\u001B[0m \u001B[36m%d\u001B[0m урона.\n", attackHero3)
				fmt.Println("===========================================================================")
			} else {
				fmt.Println("Сначала используйте огненный залп. \u001B[33mВы\u001B[0m пропускаете ход.")
				fmt.Println("===========================================================================")
			}

			if m == 2 {
				m = 0
			}
		case "Лечение":
			if n == 5 {
				fmt.Println("У \u001B[33mвас\u001B[0m закончились бинты. \u001B[33mВы\u001B[0m пропускаете ход.")
				fmt.Println("===========================================================================")
				break
			} else {
				hpHero += heal
				n += 1
				fmt.Printf("\u001B[33mВы\u001B[0m восстановили \u001B[32m%d\u001B[0m здоровья.\n", heal)
				fmt.Println("===========================================================================")
			}
		default:
			fmt.Println("\u001B[33mВы\u001B[0m споткнулись и пропускаете ход.")
			fmt.Println("===========================================================================")
			break
		}

		if hpBoss <= 0 {
			fmt.Println("\u001B[33mВы\u001B[0m победили. Наконец-то мир может вздохнуть с облегчением.")
			fmt.Println("===========================================================================")
			break
		}

		skillBoss := rand.Intn(4) + 1
		fmt.Printf("\u001B[35mХод Босса:\u001B[0m\n")
		fmt.Printf("У \u001B[35mБосса\u001B[0m \u001B[31m%d\u001B[0m здоровья\n", hpBoss)
		fmt.Println("\u001B[31m" + string(array) + "\u001B[0m")
		fmt.Println("===========================================================================")
		switch skillBoss {
		case 1:
			hpHero -= attackBoss1
			nameAttackBoss = "удар рукой"
			fmt.Printf("\u001B[35mБосс\u001B[0m использовал \u001B[35m%s\u001B[0m. \u001B[33mВы\u001B[0m получили \u001B[36m%d\u001B[0m урона.\n", nameAttackBoss, attackBoss1)
			fmt.Println("===========================================================================")
		case 2:
			hpHero -= attackBoss2
			manaHero -= manaCostBoss
			nameAttackBoss = "оскорбительный реп"
			fmt.Printf("\u001B[35mБосс\u001B[0m использовал \u001B[35m%s\u001B[0m. \u001B[33mВы\u001B[0m получили \u001B[36m%d\u001B[0m урона и \u001B[34m%d\u001B[0m маны.\n", nameAttackBoss, attackBoss2, manaCostBoss)
			fmt.Println("===========================================================================")
		case 3:
			hpHero -= attackBoss3
			nameAttackBoss = "пинок ногой"
			fmt.Printf("\u001B[35mБосс\u001B[0m использовал \u001B[35m%s\u001B[0m. \u001B[33mВы\u001B[0m получили \u001B[36m%d\u001B[0m урона.\n", nameAttackBoss, attackBoss3)
			fmt.Println("===========================================================================")
		case 4:
			fmt.Println("\u001B[35mБосс\u001B[0m использует танец из фортнайта.")
			fmt.Println("===========================================================================")
		}

		if hpHero <= 0 {
			fmt.Println("\u001B[33mВы\u001B[0m проиграли. К сожалению \u001B[33mГерой\u001B[0m погиб, мир никогда уже не будет прежним...")
			fmt.Println("===========================================================================")
			break
		}
	}
}