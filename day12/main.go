package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	// First, parse the input
	grid := make(map[Point]int)
	var start Point
	var end Point
	var shortest *Point
	for i, line := range lines {
		r := []rune(line)
		for j, c := range r {
			if c == 'S' {
				start = Point{j, i}
				grid[start] = int('a') - 97
			} else if c == 'E' {
				end = Point{j, i}
				grid[end] = int('z') - 97
			} else {
				grid[Point{j, i}] = int(c) - 97
			}
		}
	}
	fmt.Println(grid)
	distances := make(map[Point]int)
	distances[end] = 0
	// set the end point to trace back
	dest := []Point{end}
	fmt.Println("len of queue", len(dest), dest)
	for len(dest) > 0 {
		curPt := dest[0]
		dest = dest[1:]
		fmt.Println(curPt, "dest", dest)
		up := Point{curPt.x - 1, curPt.y}
		down := Point{curPt.x + 1, curPt.y}
		left := Point{curPt.x, curPt.y - 1}
		right := Point{curPt.x, curPt.y + 1}
		possiblePaths := []Point{up, down, left, right}
		for _, path := range possiblePaths {
			_, okPath := grid[path]
			_, seenPath := distances[path]
			fmt.Println("distance,  move", distances[path], path)
			if okPath && !seenPath && grid[curPt] <= grid[path]+1 {
				// found next move, add distance count
				distances[path] = distances[curPt] + 1
				if grid[path] == 0 && shortest == nil {
					fmt.Println("shortest", path)
					shortest = &path
				}
				fmt.Println("distance, add to queue", distances[path], path)
				dest = append(dest, path)
			}
		}
	}

	fmt.Println("out sortest", *shortest)
	fmt.Printf("Part One Answer: %d\n", distances[start])
	fmt.Printf("Part two Answer: %d\n", distances[*shortest])

}
