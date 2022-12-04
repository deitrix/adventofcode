package main

import (
	"fmt"
	"sort"

	. "github.com/deitrix/adventofcode/picnic"
)

func main() {
	elves := make([]int, 1)
	var mostIndex int
	for _, line := range Lines() {
		if line == "" {
			elves = append(elves, 0)
			continue
		}
		i := len(elves) - 1
		cals := Int(line)
		elves[i] += cals
		if len(elves) == 1 || elves[i] > elves[mostIndex] {
			mostIndex = i
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	var sum int
	for i := 0; i < 3; i++ {
		sum += elves[i]
	}

	fmt.Println(sum)
}
