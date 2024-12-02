package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const size = 50

func main() {
	var (
		face  = '>'
		car   = '='
		tail  = '-'
		line  = strings.Repeat(" ", size)
		line2 = []rune(line)
	)

	for index, index1, index2 := 0, 1, 2; ; {
		if index >= len(line2) || index1 >= len(line2) || index2 >= len(line2) {
			index = 0
			index1 = 1
			index2 = 2
		}

		line2[index2] = face
		line2[index1] = car
		line2[index] = tail

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("|" + string(line2) + "|")

		line2[index2] = ' '
		line2[index1] = ' '
		line2[index] = ' '

		index2++
		index1++
		index++

		time.Sleep(time.Second / 1000)
	}
}
