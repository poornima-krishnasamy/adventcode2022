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
	lineNum := 1
	grp := make(map[int]struct{})
	common := make(map[int]struct{})
	for scanner.Scan() {
		lineStr := scanner.Text()

		for i := range lineStr {
			priority := int(lineStr[i])
			if priority >= 65 && priority < 91 {
				priority = priority - 65 + 27
			} else if priority >= 97 && priority < 123 {
				priority = priority - 97 + 1
			}

			if lineNum%3 == 0 {
				if _, ok := common[priority]; ok {
					total += priority
					// reset for next round of 3 lines
					grp = make(map[int]struct{})
					common = make(map[int]struct{})
					break
				}
			} else if lineNum%3 == 2 {
				if _, ok := grp[priority]; ok {
					common[priority] = struct{}{}
				}
			} else if lineNum%3 == 1 {
				grp[priority] = struct{}{}
			}
		}
		lineNum += 1
	}
	fmt.Println(total)
}
