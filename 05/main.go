package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

func main() {
	pt1()
	pt2()
}

type correspondence struct {
	name    string
	dests   []int
	sources []int
	ranges  []int
}

func (c *correspondence) dest(source int) int {
	for i := 0; i < len(c.sources); i++ {
		if source >= c.sources[i] && source <= c.sources[i]+c.ranges[i] {
			delta := source - c.sources[i]
			return c.dests[i] + delta
		}
	}
	return source
}

type ranged struct {
	start int
	rang  int
	end   int
}

func pt1() {
	almanac := strings.Split(input, "\n")
	seeds := getNumbers(almanac[0][7:])

	corres := make([]correspondence, 0)
	_ = corres

	for line := 2; line < len(almanac); line++ {
		if len(almanac[line]) == 0 {
			continue
		}

		c := correspondence{
			sources: make([]int, 0),
			dests:   make([]int, 0),
			ranges:  make([]int, 0),
		}

		if almanac[line][len(almanac[line])-1] == ':' {
			c.name = almanac[line]
		}

		line++
		for ; len(almanac[line]) != 0; line++ {
			nbs := getNumbers(almanac[line])
			c.dests = append(c.dests, nbs[0])
			c.sources = append(c.sources, nbs[1]) // inverted ?
			c.ranges = append(c.ranges, nbs[2])
		}

		corres = append(corres, c)
	}

	lowest := math.MaxInt32
	for _, seed := range seeds {
		location := seed
		for _, c := range corres {
			location = c.dest(location)
		}
		if location < lowest {
			lowest = location
		}
	}
	fmt.Println(lowest)
}

// unfortunately pt2 is bruteforce, but very light on ram (3,2MB!), only CPU is heavily used.
func pt2() {
	t := time.Now()
	almanac := strings.Split(input, "\n")
	seeds := getSeedRange(almanac[0][7:])

	corres := make([]correspondence, 0)
	_ = corres

	for line := 2; line < len(almanac); line++ {
		if len(almanac[line]) == 0 {
			continue
		}

		c := correspondence{
			sources: make([]int, 0),
			dests:   make([]int, 0),
			ranges:  make([]int, 0),
		}

		if almanac[line][len(almanac[line])-1] == ':' {
			c.name = almanac[line]
		}

		line++
		for ; len(almanac[line]) != 0; line++ {
			nbs := getNumbers(almanac[line])
			c.dests = append(c.dests, nbs[0])
			c.sources = append(c.sources, nbs[1]) // inverted ?
			c.ranges = append(c.ranges, nbs[2]-1)
		}

		corres = append(corres, c)
	}

	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	minOfRange := make([]int, 0)

	getMinForRange := func(start, end int) {
		minRg := math.MaxInt32
		var loc int
		var s = start
		for ; s < end; s++ {
			loc = s
			for c := 0; c < len(corres); c++ {
				loc = corres[c].dest(loc)
			}
			if loc < minRg {
				minRg = loc
			}
		}

		m.Lock()
		minOfRange = append(minOfRange, minRg)
		m.Unlock()

		wg.Done()
	}

	seedsDivided := divideSeeds(seeds, 25) // optimum for input, does around 25s to compute after this point
	for i := 0; i < len(seedsDivided); i++ {
		wg.Add(1)
		go getMinForRange(seedsDivided[i].start, seedsDivided[i].end)
	}
	wg.Wait()

	fmt.Println(slices.Min(minOfRange))
	fmt.Printf("seconds elapsed %f\n", time.Since(t).Seconds())
}

func divideSeeds(ranges []ranged, divide int) []ranged {
	var dividedRanges []ranged

	for _, r := range ranges {
		interval := r.rang / divide
		for i := 0; i < divide; i++ {
			newStart := r.start + i*interval
			newRange := interval
			if i == divide-1 {
				newRange = r.end - newStart + 1
			}
			dividedRanges = append(dividedRanges, ranged{
				start: newStart,
				rang:  newRange,
				end:   newStart + newRange - 1,
			})
		}
	}

	return dividedRanges
}

func getSeedRange(s string) []ranged {
	s = strings.ReplaceAll(s, "  ", " ")
	splt := strings.Split(s, " ")
	res := make([]ranged, len(splt)/2)
	for i := 0; i < len(splt); i += 2 {
		start, _ := strconv.Atoi(splt[i])
		rang, _ := strconv.Atoi(splt[i+1])
		res[i/2].start = start
		res[i/2].rang = rang
		res[i/2].end = start + rang - 1
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
