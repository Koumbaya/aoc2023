package main

import (
	_ "embed"
	"fmt"
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

type choice struct {
	left  string
	right string
}

func pt1() {
	s := strings.Split(input, "\n")
	instructions := s[0]
	maps := make(map[string]choice, len(s)-2)
	for i := 2; i < len(s); i++ {
		loc, in, _ := strings.Cut(s[i], " = ")
		l, r, _ := strings.Cut(in, ", ")
		maps[loc] = choice{
			left:  l[1:],
			right: r[:len(r)-1],
		}
	}

	n := len(instructions)
	var i int
	target := "AAA"
	count := 0
	for x := 0; ; x++ {
		i = x % n
		if instructions[i] == 'L' {
			target = maps[target].left
		} else {
			target = maps[target].right
		}
		count++
		if target == "ZZZ" {
			break
		}
	}
	fmt.Println(count)
}

func pt2() {
	s := strings.Split(input, "\n")
	instructions := s[0]
	maps := make(map[string]choice, len(s)-2)
	for i := 2; i < len(s); i++ {
		loc, in, _ := strings.Cut(s[i], " = ")
		l, r, _ := strings.Cut(in, ", ")
		maps[loc] = choice{
			left:  l[1:],
			right: r[:len(r)-1],
		}
	}

	starts := make([]string, 0)
	for m := range maps {
		if m[2] == 'A' {
			starts = append(starts, m)
		}
	}

	allSteps := make([]int, 0)
	for _, start := range starts {
		cur := start
		steps := 0
		i := 0
		for !(cur[2] == 'Z') {
			t := instructions[i]
			if t == 'L' {
				cur = maps[cur].left
			} else {
				cur = maps[cur].right
			}
			steps++
			i = (i + 1) % len(instructions)
		}
		allSteps = append(allSteps, steps)
	}
	res := 1
	for i := 0; i < len(allSteps); i++ {
		res = lcm(res, allSteps[i])
	}
	fmt.Println(res)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
