package main

import (
	"fmt"
	"math/rand"
)

type Gladiator4 struct {
	name    string
	hp      int
	defence int
	damage  int
	mana    int
}

func NewGladiator4(name string, hp int, defence int, damage int, mana int) *Gladiator4 {
	return &Gladiator4{
		name:    name,
		hp:      hp,
		defence: defence,
		damage:  damage,
		mana:    mana,
	}
}

func (g *Gladiator4) dealDamage(glad Gladiators) {
	i := rand.Intn(2)
	var damage int
	if i == 0 {
		damage = g.punch()
	} else if i == 1 {
		damage = g.fireball()
	}
	glad.takeDamage(damage)
}

func (g *Gladiator4) punch() int {
	return g.damage
}

func (g *Gladiator4) fireball() int {
	if g.mana > 0 {
		g.mana -= 5
		return g.damage + 3
	}
	return g.damage
}

func (g *Gladiator4) takeDamage(damage int) {
	damage -= g.defence
	g.hp -= damage
}

func (g *Gladiator4) showHp() int {
	return g.hp
}

func (g *Gladiator4) showStats() string {
	return fmt.Sprintf("%s:\n-hp:%d\n-defence:%d\n-damage:%d\n-mana:%d\n", g.name, g.hp, g.defence, g.damage, g.mana)
}
