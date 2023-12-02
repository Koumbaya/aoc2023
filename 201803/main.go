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
	matrice := make([][]int, 1000)
	for i := range matrice {
		matrice[i] = make([]int, 1000)
	}
	for _, s := range strings.Split(input, "\n") {
		_, s, _ = strings.Cut(s, "@ ")
		locStr, sizeStr, _ := strings.Cut(s, ": ")
		xStr, yStr, _ := strings.Cut(locStr, ",")
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)
		lStr, wStr, _ := strings.Cut(sizeStr, "x")
		l, _ := strconv.Atoi(lStr)
		w, _ := strconv.Atoi(wStr)
		for i := y; i < w+y; i++ {
			for j := x; j < l+x; j++ {
				if i == 1000 || j == 1000 {
					fmt.Println("a")
				}
				matrice[j][i]++
			}
		}
	}
	sum := 0
	for y := 0; y < len(matrice); y++ {
		for x := 0; x < len(matrice[0]); x++ {
			if matrice[y][x] > 1 {
				sum++
			}
		}
	}
	//fmt.Println(count)
	fmt.Println(sum)
}

func pt2() {
	type submatrix struct {
		id    string
		count int
	}
	matrice := make([][]submatrix, 1000)
	for i := range matrice {
		matrice[i] = make([]submatrix, 1000)
	}
	claims := make(map[string]struct{})
	for _, s := range strings.Split(input, "\n") {
		claim, s, _ := strings.Cut(s, "@ ")
		claims[claim] = struct{}{}
		locStr, sizeStr, _ := strings.Cut(s, ": ")
		xStr, yStr, _ := strings.Cut(locStr, ",")
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)
		lStr, wStr, _ := strings.Cut(sizeStr, "x")
		l, _ := strconv.Atoi(lStr)
		w, _ := strconv.Atoi(wStr)

		for i := y; i < w+y; i++ {
			for j := x; j < l+x; j++ {
				matrice[j][i].count++
				if matrice[j][i].id != "" {
					delete(claims, matrice[j][i].id)
					delete(claims, claim)
				}
				matrice[j][i].id = claim
			}
		}
	}
	fmt.Println(claims)
}
