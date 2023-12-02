package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	freq := 0
	freqs := make(map[int]struct{})
	for {
		for _, s := range strings.Split(input, "\n") {
			val, _ := strconv.Atoi(s[1:])
			if s[0] == '+' {
				freq += val
			} else {
				freq -= val
			}
			if _, exist := freqs[freq]; exist {
				fmt.Printf("first cycle %d\n", freq)
				return
			} else {
				freqs[freq] = struct{}{}
			}
		}
		fmt.Printf("end frequency %d\n", freq)
	}
}
