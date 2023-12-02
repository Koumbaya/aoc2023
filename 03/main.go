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

type grid struct {
	g []string
}

type pos struct {
	y int
	x int
}

func (g *grid) hasNeighborSymbol(p pos) bool {
	n := g.neighbors(p)
	for _, p2 := range n {
		if g.isSymbol(p2) {
			return true
		}
	}
	return false
}

func (g *grid) hasNeighborGear(p pos) (pos, bool) {
	n := g.neighbors(p)
	for _, p2 := range n {
		if g.valid(p2) && rune(g.g[p2.y][p2.x]) == '*' {
			return p2, true
		}
	}
	return pos{}, false
}

func (g *grid) getPartNumber(p pos) int {
	str := ""
	for {
		if g.isDigit(pos{y: p.y, x: p.x}) {
			str += string(g.g[p.y][p.x])
			p.x++
		} else {
			break
		}
	}
	val, _ := strconv.Atoi(str)
	return val
}

func (g *grid) getPartAndGearPos(p pos) (int, pos, bool) {
	str := ""
	gear := pos{}
	found := false
	for {
		if g.isDigit(p) {
			str += string(g.g[p.y][p.x])

			if posGear, gearFound := g.hasNeighborGear(p); gearFound {
				gear = posGear
				found = true
			}

			p.x++
		} else {
			break
		}
	}
	val, _ := strconv.Atoi(str)
	return val, gear, found
}

func (g *grid) getPartHasSymbol(p pos) (int, bool) {
	symbol := false
	str := ""
	for {
		if g.isDigit(p) {
			str += string(g.g[p.y][p.x])

			if g.hasNeighborSymbol(p) {
				symbol = true
			}

			p.x++
		} else {
			break
		}
	}
	val, _ := strconv.Atoi(str)
	return val, symbol
}

func (g *grid) isDigit(p pos) bool {
	if !g.valid(p) {
		return false
	}
	r := rune(g.g[p.y][p.x])
	return '0' <= r && r <= '9'
}

func (g *grid) isSymbol(p pos) bool {
	if !g.valid(p) {
		return false
	}
	r := rune(g.g[p.y][p.x])
	return r != '.' && (r < '0' || r > '9')
}

func (g *grid) neighbors(p pos) []pos {
	dirs := make([]pos, 8)
	dirs[0] = pos{x: p.x - 1, y: p.y - 1}
	dirs[1] = pos{x: p.x, y: p.y - 1}
	dirs[2] = pos{x: p.x + 1, y: p.y - 1}
	dirs[3] = pos{x: p.x - 1, y: p.y}
	dirs[4] = pos{x: p.x + 1, y: p.y}
	dirs[5] = pos{x: p.x - 1, y: p.y + 1}
	dirs[6] = pos{x: p.x, y: p.y + 1}
	dirs[7] = pos{x: p.x + 1, y: p.y + 1}
	return dirs
}

func (g *grid) valid(p pos) bool {
	if p.x < 0 || p.y < 0 || p.y >= len(g.g) || p.x >= len((g.g)[0]) {
		return false
	}
	return true
}

func pt1() {
	plan := grid{
		g: strings.Split(input, "\n"),
	}
	parts := make([]int, 0)

	processing := false
	for y := 0; y < len(plan.g); y++ {
		for x := 0; x < len(plan.g[0]); x++ {
			if !plan.isDigit(pos{x: x, y: y}) {
				processing = false
				continue
			}
			if !processing {
				processing = true
				partNumber, hasSymbol := plan.getPartHasSymbol(pos{x: x, y: y})
				if hasSymbol {
					parts = append(parts, partNumber)
				}
			}
		}
	}
	sum := 0
	for _, p := range parts {
		sum += p
	}
	fmt.Println(sum)
}

func pt2() {
	plan := grid{
		g: strings.Split(input, "\n"),
	}

	type gearPts struct {
		part1 int
		part2 int
	}

	gears := make(map[pos]gearPts)

	processing := false
	for y := 0; y < len(plan.g); y++ {
		for x := 0; x < len(plan.g[0]); x++ {
			if !plan.isDigit(pos{x: x, y: y}) {
				processing = false
				continue
			}
			if !processing {
				processing = true
				partNumber, gearPos, found := plan.getPartAndGearPos(pos{x: x, y: y})
				if gear, exist := gears[gearPos]; exist && found {
					gear.part2 = partNumber
					gears[gearPos] = gear
				} else if !exist && found {
					gears[gearPos] = gearPts{part1: partNumber}
				}
			}
		}
	}
	sum := 0
	for _, pts := range gears {
		sum += pts.part1 * pts.part2
	}
	fmt.Println(sum)
}
