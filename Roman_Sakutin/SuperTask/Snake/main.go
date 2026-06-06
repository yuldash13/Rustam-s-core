package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const size = 10

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

const (
	snakeHead = '8'
	snakeTail = '0'
	apple     = '*'
	emptySign = '-'
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	var (
		field [][]rune
		line  = strings.Repeat(string(emptySign), size)
	)

	for i := 0; i < size; i++ {
		field = append(field, []rune(line))
	}

	var (
		end              int
		head             = 0
		tail             = 2
		chunks           = make([][2]int, size*size)
		x, y             = size / 2, size / 2
		direction        = UP
		currentDirection = UP
		loop             = 5
		currentLoop      = 0
	)
	chunks[0][0], chunks[0][1] = x, y+1
	chunks[1][0], chunks[1][1] = x, y+2
	field[y][x] = snakeHead
	field[y+1][x] = snakeTail
	field[y+2][x] = snakeTail

	var finish = false
	go func() {
		for !finish {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					direction = UP
				case termbox.KeyArrowDown:
					direction = DOWN
				case termbox.KeyArrowLeft:
					direction = LEFT
				case termbox.KeyArrowRight:
					direction = RIGHT
				}
			default:
				break
			}
		}
	}()

	for {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		for _, str := range field {
			fmt.Println(string(str))
		}

		time.Sleep(time.Second / 2)

		field[y][x] = snakeTail
		head--
		if head < 0 {
			head = len(chunks) - 1
		}
		chunks[head][0], chunks[head][1] = x, y

		switch direction {
		case UP:
			if currentDirection == DOWN {
				y += 1
				break
			}
			y -= 1
			currentDirection = UP
		case DOWN:
			if currentDirection == UP {
				y -= 1
				break
			}
			y += 1
			currentDirection = DOWN
		case LEFT:
			if currentDirection == RIGHT {
				x += 1
				break
			}
			x -= 1
			currentDirection = LEFT
		case RIGHT:
			if currentDirection == LEFT {
				x -= 1
				break
			}
			x += 1
			currentDirection = RIGHT
		}

		if y < 0 {
			y = size - 1
		} else if y >= size {
			y = 0
		} else if x < 0 {
			x = size - 1
		} else if x >= size {
			x = 0
		}

		if field[y][x] == apple {
			tail++
		}

		if field[y][x] == snakeTail {
			finish = true
			cmd = exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("GAME OVER!")
			time.Sleep(time.Second * 3)
			return
		}

		field[y][x] = snakeHead

		end = tail - 1
		if end < 0 {
			end = len(chunks) - 1
		}

		tailX, tailY := chunks[end][0], chunks[end][1]
		field[tailY][tailX] = emptySign
		tail--

		if tail < 0 {
			tail = len(chunks) - 1
		}

		currentLoop++
		if currentLoop == loop {
			rand.Seed(time.Now().UnixNano())
			randomX := rand.Intn(size)
			randomY := rand.Intn(size)
			if field[randomY][randomX] == emptySign {
				field[randomY][randomX] = apple
			}
			currentLoop = 0
		}
	}
}
