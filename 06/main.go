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
}

func pt1() {
	s := strings.Split(input, "\n")
	_, cut, _ := strings.Cut(s[0], ":")
	time := getNumbers(cut)
	_, cut, _ = strings.Cut(s[1], ":")
	distRecords := getNumbers(cut)

	fmt.Println(calc(time, distRecords))
}

func pt2() {
	s := strings.Split(input, "\n")
	_, cut, _ := strings.Cut(s[0], ":")
	time := getWholeNumbers(cut)
	_, cut, _ = strings.Cut(s[1], ":")
	distRecords := getWholeNumbers(cut)

	fmt.Println(calc(time, distRecords))
}

func calc(time, record []int) int {
	winPossible := make([]int, 0)
	for i, record := range record {
		tRace := time[i]
		winP := 0
		for tHold := 1; tHold < tRace; tHold++ { // 0 never wins
			dist := tHold * (tRace - tHold)
			if dist > record {
				winP++
			}
		}
		winPossible = append(winPossible, winP)
	}
	mult := winPossible[0]
	for i := 1; i < len(winPossible); i++ {
		mult *= winPossible[i]
	}
	return mult
}

func getWholeNumbers(s string) []int {
	res := make([]int, 0)
	s = strings.ReplaceAll(s, " ", "")
	splt := strings.Split(s, " ")
	for _, s2 := range splt {
		if val, err := strconv.Atoi(s2); err == nil {
			res = append(res, val)
		}
	}
	return res
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
