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

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func pt1() {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		game, valid := parseGame(s)
		if valid {
			sum += game
		}
	}
	fmt.Println(sum)
}

func pt2() {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		splt := strings.Split(s, ":")
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, set := range strings.Split(splt[1], ";") {
			colors := strings.Split(set, ",")
			for _, color := range colors {
				switch {
				case strings.Contains(color, "red"):
					color = strings.ReplaceAll(strings.ReplaceAll(color, "red", ""), " ", "")
					r := mustAtoi(color)
					if r > minRed {
						minRed = r
					}
				case strings.Contains(color, "green"):
					color = strings.ReplaceAll(strings.ReplaceAll(color, "green", ""), " ", "")
					g := mustAtoi(color)
					if g > minGreen {
						minGreen = g
					}
				case strings.Contains(color, "blue"):
					color = strings.ReplaceAll(strings.ReplaceAll(color, "blue", ""), " ", "")
					b := mustAtoi(color)
					if b > minBlue {
						minBlue = b
					}
				}
			}
		}
		sum += (minRed * minGreen * minBlue)
	}
	fmt.Println(sum)
}

func parseGame(s string) (game int, valid bool) {
	splt := strings.Split(s, ":")
	game = mustAtoi(strings.Replace(splt[0], "Game ", "", 1))
	splt = strings.Split(strings.ReplaceAll(splt[1], ";", ","), ",")
	for _, val := range splt {
		switch {
		case strings.Contains(val, "red"):
			val = strings.ReplaceAll(strings.ReplaceAll(val, "red", ""), " ", "")
			r := mustAtoi(val)
			if r > maxRed {
				return game, false
			}
		case strings.Contains(val, "green"):
			val = strings.ReplaceAll(strings.ReplaceAll(val, "green", ""), " ", "")
			g := mustAtoi(val)
			if g > maxGreen {
				return game, false
			}
		case strings.Contains(val, "blue"):
			val = strings.ReplaceAll(strings.ReplaceAll(val, "blue", ""), " ", "")
			b := mustAtoi(val)
			if b > maxBlue {
				return game, false
			}
		}
	}
	return game, true
}

func mustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
