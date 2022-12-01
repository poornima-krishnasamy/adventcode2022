package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var sortedCal = make([]int, len(calories))
	i := 0
	for _, cal := range calories {
		sortedCal[i] = cal
		i++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortedCal)))
	fmt.Println(sortedCal[0] + sortedCal[1] + sortedCal[2])

}
