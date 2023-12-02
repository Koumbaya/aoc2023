package main

import (
	_ "embed"
	"fmt"
	"slices"
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
}

func pt1() {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		_, s, _ := strings.Cut(s, ": ")
		w, g, _ := strings.Cut(s, " | ")
		ws := getNumbers(w)
		gs := getNumbers(g)
		count := 0
		for _, candidate := range ws {
			if slices.Contains(gs, candidate) {
				if count == 0 {
					count++
				} else {
					count = count * 2
				}
			}
		}
		sum += count
	}
	fmt.Println(sum)
}

func pt2() {
	in := strings.Split(input, "\n")
	cardsWin := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		cardsWin[i] = 1
	}
	total := 0
	for i, s := range in {
		_, s, _ := strings.Cut(s, ": ")
		w, g, _ := strings.Cut(s, " | ")
		ws := getNumbers(w)
		gs := getNumbers(g)
		wins := 0
		for _, candidate := range ws {
			if slices.Contains(gs, candidate) {
				wins++
			}
		}
		if wins == 0 {
			continue
		}
		for j := i + 1; j <= i+wins; j++ {
			cardsWin[j] += cardsWin[i]
		}
	}
	for _, v := range cardsWin {
		total += v
	}
	fmt.Println(total)
}

func getNumbers(s string) []int {
	res := make([]int, 0)
	s = strings.ReplaceAll(s, "  ", " ")
	splt := strings.Split(s, " ")
	for _, s2 := range splt {
		if val, err := strconv.Atoi(s2); err == nil {
			res = append(res, val)
		}
	}
	return res
}
