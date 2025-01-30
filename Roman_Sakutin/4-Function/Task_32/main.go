package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var (
		menu string
		y, x = 5, 3
		n    int
	)
	field := [][]string{
		{"-", "-", "-", "-", "-", "-", "-"},
		{"|", " ", " ", "*", " ", " ", "|"},
		{"|", "*", "-", "-", "-", "*", "|"},
		{"|", " ", "-", "-", "-", " ", "|"},
		{"|", " ", "-", "-", "-", " ", "|"},
		{"|", "*", " ", "0", "|", "*", "|"},
		{"-", "-", "-", "-", "-", "-", "-"},
	}
	for {
		fmt.Println("Меню:\n1)Карта\n2)Перемещение\n3)Выход")
		fmt.Scan(&menu)
		switch menu {
		case "Карта":
			clearConsol()
			showMap(field)
		case "Перемещение":
			clearConsol()
		walking:
			for {
				var direction string
				showMap(field)
				fmt.Println("Куда вы хотите пойти:\n1)Вверх\n2)Влево\n3)Вниз\n4)Вправо\n5)Выход")
				fmt.Scan(&direction)
				clearConsol()
				switch direction {
				case "Вверх":
					y--
					if field[y][x] == "-" {
						fmt.Println("Вы не можете ходить сквозь стены")
						y++
						break
					} else if field[y][x] == "*" {
						n++
					}
					field[y][x] = "0"
					field[y+1][x] = " "
					showMap(field)
				case "Влево":
					x--
					if field[y][x] == "-" || field[y][x] == "|" {
						fmt.Println("Вы не можете ходить сквозь стены")
						x++
						break
					} else if field[y][x] == "*" {
						n++
					}
					field[y][x] = "0"
					field[y][x+1] = " "
					showMap(field)
				case "Вниз":
					y++
					if field[y][x] == "-" || field[y][x] == "|" {
						fmt.Println("Вы не можете ходить сквозь стены")
						y--
						break
					} else if field[y][x] == "*" {
						n++
					}
					field[y][x] = "0"
					field[y-1][x] = " "
					showMap(field)
				case "Вправо":
					x++
					if field[y][x] == "-" || field[y][x] == "|" {
						fmt.Println("Вы не можете ходить сквозь стены")
						x--
						break
					} else if field[y][x] == "*" {
						n++
					}
					field[y][x] = "0"
					field[y][x-1] = " "
					showMap(field)
				case "Выход":
					clearConsol()
					break walking
				default:
					clearConsol()
					fmt.Println("Вы ввели команду неверно.")
				}
				if n == 1 {
					fmt.Println("1/5")
				} else if n == 2 {
					fmt.Println("2/5")
				} else if n == 3 {
					fmt.Println("3/5")
				} else if n == 4 {
					fmt.Println("4/5")
				} else if n == 5 {
					fmt.Println("5/5")
					fmt.Println("Вы собрали все звезды, поздравляем!")
					return

				}
			}
		case "Выход":
			clearConsol()
			return
		default:
			clearConsol()
			fmt.Println("Вы ввели команду неверно.")
		}

	}
}

func showMap(field [][]string) {
	for i := 0; i < len(field); i++ {
		fmt.Println(field[i])
	}
	fmt.Println("Обозначения:\n1)0 - вы\n2)* - звезда\n3)-,| - стены")
}

func clearConsol() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
