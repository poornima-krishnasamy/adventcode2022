package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Println(lineStr)
		s := strings.Split(lineStr, " ")
		switch s[0] {
		//Rock
		case "A":
			fmt.Println(s[0], s[1])
			// Lose = scissors
			if s[1] == "X" {
				score = score + 3 + 0
			}
			// Draw - Rock
			if s[1] == "Y" {
				score = score + 1 + 3
			}
			// Win - Paper
			if s[1] == "Z" {
				score = score + 2 + 6
			}
			// Paper
		case "B":
			fmt.Println(s[0], s[1])
			// Lose - Rock
			if s[1] == "X" {
				score = score + 1 + 0
			}
			// Draw - Paper
			if s[1] == "Y" {
				score = score + 2 + 3
			}
			// Win = scissors
			if s[1] == "Z" {
				score = score + 3 + 6
			}
			// Scissors
		case "C":
			fmt.Println(s[0], s[1])
			// Lose - Paper
			if s[1] == "X" {
				score = score + 2 + 0
			}
			// Draw - Scissors
			if s[1] == "Y" {
				score = score + 3 + 3
			}
			// Win = Rock
			if s[1] == "Z" {
				score = score + 1 + 6
			}
		}
		fmt.Println(score)
	}
	fmt.Println("total score", score)
}
