package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		len := len(lineStr) / 2
		var priority byte
	NEXT:
		for i := 0; i < len; i++ {
			for j := len; j < len*2; j++ {
				if lineStr[i] == lineStr[j] {
					priority = lineStr[i]
					if priority >= 97 {
						total += ((int(priority) - 97) + 1)
					} else {
						total += (26 + (int(priority) - 65) + 1)
					}
					break NEXT
				}
			}

		}

	}
	fmt.Println(total)
}
