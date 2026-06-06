package main

import (
	"fmt"
	"math/rand"
)

type Gladiator1 struct {
	name         string
	hp           int
	defence      int
	damage       int
	doubleDamage bool
}

func NewGladiator1(name string, hp int, defence int, damage int, doubleDamage bool) *Gladiator1 {
	return &Gladiator1{
		name:         name,
		hp:           hp,
		defence:      defence,
		damage:       damage,
		doubleDamage: doubleDamage,
	}
}

func (g *Gladiator1) dealDamage(glad Gladiators) {
	i := rand.Intn(2)
	var damage int
	if i == 0 {
		damage = g.punch()
	} else if i == 1 {
		damage = g.double()
	}
	glad.takeDamage(damage)
}

func (g *Gladiator1) punch() int {
	return g.damage
}

func (g *Gladiator1) double() int {
	i := rand.Intn(2)
	if i == 0 {
		return g.damage
	}
	return g.damage * 2
}

func (g *Gladiator1) takeDamage(damage int) {
	damage -= g.defence
	g.hp -= damage
}

func (g *Gladiator1) showHp() int {
	return g.hp
}

func (g *Gladiator1) showStats() string {
	return fmt.Sprintf("%s:\n-hp:%d\n-defence:%d\n-damage:%d\n-chance to double damage:%v\n", g.name, g.hp, g.defence, g.damage, g.doubleDamage)
}
