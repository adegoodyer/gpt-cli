package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func PrintWelcome() {
	fmt.Println("Welcome to GPT3!")
	fmt.Println("'q' to exit.")
	fmt.Println("'c' to clear screen")
	fmt.Println("'h' to view history.")
}
