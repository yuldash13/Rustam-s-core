package main

import (
	"fmt"
	"math/rand"
)

type Gladiator5 struct {
	name    string
	hp      int
	defence int
	damage  int
	dodge   bool
}

func NewGladiator5(name string, hp int, defence int, damage int, dodge bool) *Gladiator5 {
	return &Gladiator5{
		name:    name,
		hp:      hp,
		defence: defence,
		damage:  damage,
		dodge:   dodge,
	}
}

func (g *Gladiator5) dealDamage(glad Gladiators) {
	damage := g.punch()
	glad.takeDamage(damage)
}

func (g *Gladiator5) punch() int {
	return g.damage
}

func (g *Gladiator5) roll() {
	g.hp -= 0
}

func (g *Gladiator5) takeDamage(damage int) {
	i := rand.Intn(2)
	if i == 1 {
		g.roll()
		return
	}
	damage -= g.defence
	g.hp -= damage
}

func (g *Gladiator5) showHp() int {
	return g.hp
}

func (g *Gladiator5) showStats() string {
	return fmt.Sprintf("%s:\n-hp:%d\n-defence:%d\n-damage:%d\n-chance to dodge:%v\n", g.name, g.hp, g.defence, g.damage, g.dodge)
}
