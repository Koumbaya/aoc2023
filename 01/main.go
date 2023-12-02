package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

func main() {
	pt1()
	pt2()
	pt2Better()
}

func pt1() {
	firstDone := false
	first := 0
	last := 0
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		for _, r := range s {
			x, err := strconv.Atoi(string(r))
			if err == nil {
				if !firstDone {
					firstDone = true
					first = x
				}
				last = x
			}
		}
		firstDone = false
		sum += (first * 10) + last
	}
	fmt.Println(sum)
}

func pt2() {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		first, last := find(s)
		sum += (first * 10) + last
	}
	fmt.Println(sum)
}

func pt2Better() {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		first, last := findBetter(s)
		sum += (first * 10) + last
	}
	fmt.Println(sum)
}

func findBetter(s string) (int, int) {
	idxFirst := 100000
	idxLast := -1
	first := 0
	last := 0
	for word, digit := range spelledDigits {
		idx := strings.Index(s, word)
		if idx != -1 && idx < idxFirst {
			idxFirst = idx
			first = digit
		}

		idx = strings.LastIndex(s, word)
		if idx != -1 && idx > idxLast {
			idxLast = idx
			last = digit
		}
	}

	return first, last
}

func find(s string) (int, int) {
	calc := make([]int, len(s))
	word := ""
	digit := 0
	pos := 0
	for word, digit = range spelledDigits {
		for {
			idx := strings.Index(s[pos:], word)
			if idx == -1 {
				pos = 0
				break
			}
			actualIdx := idx + pos
			calc[actualIdx] = digit
			pos = actualIdx + len(word)
		}
	}

	firstFound := false
	first := 0
	last := 0
	for i := 0; i < len(calc); i++ {
		if calc[i] != 0 {
			if !firstFound {
				firstFound = true
				first = calc[i]
				last = calc[i]
			} else {
				last = calc[i]
			}
		}
	}

	return first, last
}

var spelledDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}
