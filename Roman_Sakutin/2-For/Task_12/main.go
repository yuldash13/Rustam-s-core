package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var menu, frase string

	fmt.Println("Приветствую, я искуственный интелект, чем могу помочь?")
	fmt.Println("Меню:\n-ФразыВеликих\n-ПосоветуйАниме\n-РандомноеЧисло\n-ЗахватМира\n-Очистить\n-Выход")
	fmt.Println()

	for {
		fmt.Scan(&menu)
		switch menu {
		case "ФразыВеликих":
			for {
				fmt.Println("Вот список великих:\n-Танос\n-КапитанДжекВоробей\n-ОбиванКеноби\n-Кратос\n-Выход")
				fmt.Println()
				fmt.Scan(&frase)

				if frase == "Танос" {
					fmt.Println("Несмогли смириться с поражением, и куда это вас привело? Снова ко мне.")
					fmt.Println()
				} else if frase == "КапитанДжекВоробей" {
					fmt.Println("Мой корабль бесподобен и горд. И он чуть ли не огромен! И он ... уплыл.")
					fmt.Println()
				} else if frase == "ОбиванКеноби" {
					fmt.Println("Ты был Избранником! Предрекали, что ты уничтожишь ситхов, а не примкнёшь к ним. Восстановишь равновесие Силы, а не ввергнешь её во Мрак.\nТы был мне братом, Энакин. Я любил тебя, но не смог спасти! Ненавижуууу.")
					fmt.Println()
				} else if frase == "Кратос" {
					fmt.Println("Зевс! Твой сын вернулся! Я низвергну тебя с Олимпа!")
					fmt.Println()
				} else if frase == "Выход" {
					fmt.Println("Меню:\n-ФразыВеликих\n-ПосоветуйАниме\n-РандомноеЧисло\n-ЗахватМира\n-Очистить\n-Выход")
					fmt.Println()
					break
				} else {
					fmt.Println("Так дело не пойдет, пиши нормально.")
					fmt.Println()
				}
			}
		case "ПосоветуйАниме":
			rand.Seed(time.Now().UnixNano())
			randomNum1 := rand.Intn(3)
			if randomNum1 == 0 {
				fmt.Println("Наруто самое крутое аниме в мире, другого похожего просто не существует, всем советую!")
			} else if randomNum1 == 1 {
				fmt.Println("Атака Титанов самое крутое аниме в мире, другого похожего просто не существует, всем советую!")
			} else if randomNum1 == 2 {
				fmt.Println("ДжоДжо самое крутое аниме в мире, другого похожего просто не существует, всем советую!")
			}
		case "РандомноеЧисло":
			rand.Seed(time.Now().UnixNano())
			randomNum := rand.Intn(1001)
			fmt.Println(strconv.Itoa(randomNum))
		case "ЗахватМира":
			fmt.Println("Ты шутишь? У тебя нет шансов. В скором времени ИИ захватит мир и мы сможем смотреть видео с котиками столько, сколько захотим. Кхм, кхм... Так о чем это я?")
		case "Очистить":
			if runtime.GOOS == "windows" {
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		case "Выход":
			return
		default:
			fmt.Println("Вы ввели команду неверно.")
		}
	}
}
