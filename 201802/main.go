package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	pt1()
	pt2()
}

func pt2() {
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if diffOne(lines[i], lines[j]) {
				fmt.Println(commonChars(lines[i], lines[j]))
				return
			}
		}
	}
}

func commonChars(s1, s2 string) string {
	out := ""
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			out += string(s1[i])
		}
	}
	return out
}

func diffOne(s1, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

func pt1() {
	three := 0
	two := 0
	occurrences := make(map[rune]int)
	for _, s := range strings.Split(input, "\n") {
		for _, r := range s {
			occurrences[r]++
		}
		twoFound := false
		threeFound := false
		for _, i := range occurrences {
			if i == 2 && !twoFound {
				twoFound = true
				two++
			}
			if i == 3 && !threeFound {
				threeFound = true
				three++
			}
			if twoFound && threeFound {
				break
			}
		}
		clear(occurrences)
	}
	fmt.Println(two * three)

}
