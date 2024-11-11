package main

import "fmt"

func main() {

	var people int

	const waitingOnePeopleMinutes = 10
	const hourScale = 60

	var hour string

	var (
		hour1  = "час"
		hour2  = "часа"
		hour3  = "часов"
		minute = "минут"
	)

	fmt.Println("Сколько бабулек в очереди?")
	fmt.Scan(&people)

	totalWaitingTime := people * waitingOnePeopleMinutes
	hoursAll := totalWaitingTime / hourScale
	minutesRemaining := totalWaitingTime % hourScale

	if hoursAll == 1 {
		hour = hour1
	} else if hoursAll > 1 && hoursAll < 5 {
		hour = hour2
	} else if hoursAll >= 5 || hoursAll == 0 {
		hour = hour3
	}

	fmt.Printf("Вам придется ждать: %d %s, %d %s.", hoursAll, hour, minutesRemaining, minute)
}
