package main

import "fmt"

type Player struct {
	ID       int
	nickName string
	level    int
	banned   bool
}

func NewPlayer(id int, name string, level int, banned bool) *Player {
	return &Player{
		ID:       id,
		nickName: name,
		level:    level,
		banned:   banned,
	}
}

func (p *Player) BanPlayer() {
	p.banned = true
}

func (p *Player) UnPlayer() {
	p.banned = false
}

func (p *Player) ToString() string {
	return fmt.Sprintf("%v:\n1) ID:%v\n2) NickName:%v\n3) Level:%v\n4) Banned:%v\n", p.nickName, p.ID, p.nickName, p.level, p.banned)
}

func (p *Player) ShowBan() bool {
	return p.banned
}
