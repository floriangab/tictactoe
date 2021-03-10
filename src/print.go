package main

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintBoard(b [9]string) {
	ClearScreen()
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")
		}
	}
}

func PrintVictory(winner string) {
	if winner == "" {
		fmt.Println("it's a draw :(")
		return
	}
	fmt.Printf("%s is winner\n", winner)
}

func AskForPlay() int {
	var moveInt int
	fmt.Println("Enter Pos to play: ")
	fmt.Scan(&moveInt)
	return moveInt
}

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
