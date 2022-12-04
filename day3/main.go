package main

import (
	"fmt"

	. "github.com/deitrix/adventofcode/picnic"
)

func main() {
	lines := Lines()
	var groupPriorities []int
	for i := 0; i < len(lines); i += 3 {
		masks := make([]int, 3)
		for j := i; j < i+3; j++ {
			for _, c := range lines[j] {
				masks[j-i] |= 1 << itemPriority(c)
			}
		}

		groupMask := masks[0] & masks[1] & masks[2]

		fmt.Printf("%053b\n", groupMask)

		var priority int
		for p := 1; p <= 52; p++ {
			priority += (groupMask >> p) & 1 * p
		}

		groupPriorities = append(groupPriorities, priority)
	}

	var total int
	for _, priority := range groupPriorities {
		total += priority
	}

	fmt.Println(total)
}

func itemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	}
	if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	}
	return 0
}
