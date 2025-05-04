package main

import "fmt"

type Gladiator2 struct {
	name        string
	hp          int
	defence     int
	damage      int
	damageCount int
}

func NewGladiator2(name string, hp int, defence int, damage int, damageCount int) *Gladiator2 {
	return &Gladiator2{
		name:        name,
		hp:          hp,
		defence:     defence,
		damage:      damage,
		damageCount: damageCount,
	}
}

func (g *Gladiator2) dealDamage(glad Gladiators) {
	var damage int
	if g.damageCount == 3 {
		damage = g.double()
		g.damageCount = 0
	}
	damage = g.punch()
	g.count()
	glad.takeDamage(damage)
}

func (g *Gladiator2) punch() int {
	return g.damage
}

func (g *Gladiator2) double() int {
	return g.damage * 2
}

func (g *Gladiator2) count() {
	g.damageCount += 1
}

func (g *Gladiator2) takeDamage(damage int) {
	damage -= g.defence
	g.hp -= damage
}

func (g *Gladiator2) showHp() int {
	return g.hp
}

func (g *Gladiator2) showStats() string {
	return fmt.Sprintf("%s:\n-hp:%d\n-defence:%d\n-damage:%d\n-damage count:%v\n", g.name, g.hp, g.defence, g.damage, g.damageCount)
}
