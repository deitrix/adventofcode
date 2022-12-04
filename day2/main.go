package main

import (
	"fmt"
	"strings"

	. "github.com/deitrix/adventofcode/picnic"
)

func main() {
	var score uint
	for _, line := range Lines() {
		fields := strings.Fields(line)
		score += calculateScore(
			translateChoice(fields[0]),
			translateOutcome(fields[1]),
		)
	}
	fmt.Println(score)
}

func calculateScore(opponent, outcome uint) uint {
	return outcome + computeChoice(opponent, outcome)
}

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

const (
	Rock = iota + 1
	Paper
	Scissors
)

func translateChoice(s string) uint {
	switch s {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}
	return 0
}

func translateOutcome(s string) uint {
	switch s {
	case "X":
		return Loss
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	return 0
}

func computeChoice(opponent, outcome uint) uint {
	switch opponent {
	case Rock:
		switch outcome {
		case Loss:
			return Scissors
		case Draw:
			return Rock
		case Win:
			return Paper
		}
	case Paper:
		switch outcome {
		case Loss:
			return Rock
		case Draw:
			return Paper
		case Win:
			return Scissors
		}
	case Scissors:
		switch outcome {
		case Loss:
			return Paper
		case Draw:
			return Scissors
		case Win:
			return Rock
		}
	}
	return 0
}
