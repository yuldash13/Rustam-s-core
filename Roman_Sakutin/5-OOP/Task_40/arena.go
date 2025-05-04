package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Arena struct {
	gladiator1 Gladiators
	gladiator2 Gladiators
}

func NewArena() *Arena {
	return &Arena{
		gladiator1: nil,
		gladiator2: nil,
	}
}

func (a *Arena) showMenu() {
	fmt.Println("Variants of gladiators:\n")
	fmt.Println("1)gladiator №1:\n-hp:20\n-defence:2\n-damage:4\n-chance to double damage\n")
	fmt.Println("2)gladiator №2:\n-hp:20\n-defence:2\n-damage:4\n-after 3 attacks, deal double damage\n")
	fmt.Println("3)gladiator №3:\n-hp:20\n-defence:2\n-damage:4\n-after rage is full, using heal\n")
	fmt.Println("4)gladiator №4:\n-hp:20\n-defence:2\n-damage:4\n-mana: 15\n-can use fireball\n")
	fmt.Println("5)gladiator №5:\n-hp:20\n-defence:2\n-damage:4\n-chance to dodge attack\n")
}

func (a *Arena) initGladiator1() *Gladiator1 {
	return NewGladiator1("gladiator1", 20, 2, 4, true)
}

func (a *Arena) initGladiator2() *Gladiator2 {
	return NewGladiator2("gladiator2", 20, 2, 4, 0)
}

func (a *Arena) initGladiator3() *Gladiator3 {
	return NewGladiator3("gladiator3", 20, 2, 4, 0)
}

func (a *Arena) initGladiator4() *Gladiator4 {
	return NewGladiator4("gladiator4", 20, 2, 4, 20)
}

func (a *Arena) initGladiator5() *Gladiator5 {
	return NewGladiator5("gladiator5", 20, 2, 4, true)
}

func (a *Arena) initGladiators() {
	var gladiators []Gladiators
	for i := 0; i < 2; i++ {
		fmt.Printf("Choose gladiator №%d\n", i+1)
		var answer int
		fmt.Scan(&answer)
		if answer == 1 {
			gladiators = append(gladiators, a.initGladiator1())
		} else if answer == 2 {
			gladiators = append(gladiators, a.initGladiator2())
		} else if answer == 3 {
			gladiators = append(gladiators, a.initGladiator3())
		} else if answer == 4 {
			gladiators = append(gladiators, a.initGladiator4())
		} else if answer == 5 {
			gladiators = append(gladiators, a.initGladiator5())
		}
	}
	a.gladiator1 = gladiators[0]
	a.gladiator2 = gladiators[1]
}

func (a *Arena) showStats() string {
	return fmt.Sprintf(a.gladiator1.showStats()) + fmt.Sprintln() + fmt.Sprintf(a.gladiator2.showStats())
}

func (a *Arena) fight() {
	for {
		a.gladiator1.dealDamage(a.gladiator2)
		if a.gladiator2.showHp() <= 0 {
			fmt.Println("Gladiator №1 wins")
			return
		}
		a.gladiator2.dealDamage(a.gladiator1)
		if a.gladiator1.showHp() <= 0 {
			fmt.Println("Gladiator №2 wins")
			return
		}
		clearConsole()
		fmt.Println(a.showStats())
		time.Sleep(time.Second * 2)
	}
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
