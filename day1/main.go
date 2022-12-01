package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var calories = make(map[int]int)

	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elfNum := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		if lineStr != "" {
			cal, _ := strconv.Atoi(lineStr)
			calories[elfNum] += cal
		} else {
			elfNum++
		}
	}

	highVal := 0
	for _, cal := range calories {
		if cal > highVal {
			highVal = cal
		}
	}
	fmt.Println(highVal)

}
