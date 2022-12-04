package main

import (
	"fmt"
	"strings"

	. "github.com/deitrix/aoc/picnic"
)

func main() {
	var count int
	for _, line := range Lines() {
		pair := strings.Split(line, ",")
		amin, amax := parseRange(pair[0])
		bmin, bmax := parseRange(pair[1])
		if amin <= bmax && amax >= bmin {
			count++
		}
	}
	fmt.Println(count)
}

func parseRange(rng string) (min, max uint) {
	parts := strings.Split(rng, "-")
	return Uint(parts[0]), Uint(parts[1])
}
