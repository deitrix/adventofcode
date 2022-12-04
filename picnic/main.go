package picnic

import (
	"bufio"
	"os"
	"strconv"
)

func Lines() (lines []string) {
	f := Must1(os.Open("input.txt"))
	defer Must(f.Close())
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Int(s string) int {
	return Must1(strconv.Atoi(s))
}

func Uint(s string) uint {
	return uint(Must1(strconv.ParseUint(s, 10, 64)))
}

func Assert(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](v T, err error) T {
	Must(err)
	return v
}

func Must2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	return Must1(v1, err), v2
}
