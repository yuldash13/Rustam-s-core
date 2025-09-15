package main

import "fmt"

type Gladiator3 struct {
	name    string
	hp      int
	defence int
	damage  int
	rage    int
}

func NewGladiator3(name string, hp int, defence int, damage int, rage int) *Gladiator3 {
	return &Gladiator3{
		name:    name,
		hp:      hp,
		defence: defence,
		damage:  damage,
		rage:    rage,
	}
}

func (g *Gladiator3) dealDamage(glad Gladiators) {
	if g.rage == 3 {
		g.heal()
		g.rage = 0
	}
	damage := g.punch()
	g.count()
	glad.takeDamage(damage)
}

func (g *Gladiator3) punch() int {
	return g.damage
}

func (g *Gladiator3) heal() {
	g.hp += 6
}

func (g *Gladiator3) count() {
	g.rage += 1
}

func (g *Gladiator3) takeDamage(damage int) {
	damage -= g.defence
	g.hp -= damage
}

func (g *Gladiator3) showHp() int {
	return g.hp
}

func (g *Gladiator3) showStats() string {
	return fmt.Sprintf("%s:\n-hp:%d\n-defence:%d\n-damage:%d\n-rage:%d\n", g.name, g.hp, g.defence, g.damage, g.rage)
}
