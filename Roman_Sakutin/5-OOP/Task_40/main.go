package main

type Gladiators interface {
	dealDamage(glad Gladiators)
	takeDamage(damage int)
	showHp() int
	showStats() string
}

func main() {
	arena := NewArena()
	arena.showMenu()
	arena.initGladiators()
	arena.fight()
}
