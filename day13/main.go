package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"strings"
)

var dividerPair = [][]any{{[]any{float64(2)}}, {[]any{float64(6)}}}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	packets := [][]any{}
	packet := strings.Split(string(input), "\n\n")
	sum := 0
	for index, line := range packet {
		part := strings.Split(line, "\n")
		var left, right []any
		json.Unmarshal([]byte(part[0]), &left)
		json.Unmarshal([]byte(part[1]), &right)
		pair := [2][]any{left, right}
		fmt.Println(pair)

		// number comparison
		result := compare(left, right, 0)

		if result <= 0 {
			sum += index + 1
			fmt.Println("sum", sum, "index", index)

		}
		packets = append(packets, pair[0], pair[1])

	}
	fmt.Println(sum)
	key := 0
	packets = append(packets, dividerPair...)

	sort.Slice(packets, func(i, j int) bool {
		left, right := packets[i], packets[j]
		return less(left, right)
	})

	for i, packet := range packets {
		if reflect.DeepEqual(packet, dividerPair[0]) {
			key = i + 1
		} else if reflect.DeepEqual(packet, dividerPair[1]) {
			key = key * (i + 1)
		}
	}
	fmt.Printf("part2: %v\n", key)

}
func less(left []any, right []any) bool {
	return compare(left, right, 0) <= 0
}
func compare(left []any, right []any, i int) (res int) {
	fmt.Println("left", left, "right", right, "i", i)
	if i >= len(left) && i >= len(right) {
		fmt.Println("both list ends")
		return 0
	} else if i >= len(left) {
		fmt.Println("left run out of items so correct")
		return -1
	} else if i >= len(right) {
		fmt.Println("right run out of items so not correct")
		return 1
	}

	fmt.Println("left[i]", left[i], "right[i]", right[i], "i", i)
	leftItem, rightItem := left[i], right[i]
	leftNum, isleftNum := leftItem.(float64)
	rightNum, isrightNum := rightItem.(float64)
	if isleftNum && isrightNum && leftNum == rightNum {
		// same number, continue next item
		fmt.Println("same num compare next")
		return compare(left, right, i+1)
	} else if isleftNum && isrightNum {
		return int(leftNum - rightNum)
	}

	// list comparison
	llist, isllist := leftItem.([]any)
	rlist, isrlist := rightItem.([]any)
	if !isllist {
		llist = []any{leftNum} // mixed, convert to list
	}
	if !isrlist {
		rlist = []any{rightNum} // mixed, convert to list
	}
	res = compare(llist, rlist, 0)
	if res == 0 { // tie, continue to next item
		fmt.Println("next item")
		return compare(left, right, i+1)
	}

	return res

}
