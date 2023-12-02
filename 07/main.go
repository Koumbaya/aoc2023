package main

import (
	"cmp"
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

type typeCard int

const (
	zero typeCard = iota
	joker
	two
	three
	four
	five
	six
	seven
	height
	nine
	tee // ???? what's T
	jack
	queen
	king
	ace
)

type kind int

const (
	highType kind = iota + 20
	pairType
	twoPairType
	threeKindType
	fullHouseType
	fourKindType
	fiveKindType
)

var relation = map[rune]typeCard{
	'2': two,
	'3': three,
	'4': four,
	'5': five,
	'6': six,
	'7': seven,
	'8': height,
	'9': nine,
	'A': ace,
	'K': king,
	'Q': queen,
	'J': jack,
	'T': tee,
}

type hand struct {
	debugStr string
	order    []typeCard
	bid      int
	kind     kind
}

func pt1() {
	in := strings.Split(input, "\n")
	ranks := make([]hand, len(in))
	for i, s := range in {
		order, bid, kind := parseCard(s)
		ranks[i] = hand{
			debugStr: s,
			order:    order,
			bid:      bid,
			kind:     kind,
		}
	}

	slices.SortFunc(ranks, func(a, b hand) int {
		if n := cmp.Compare(a.kind, b.kind); n != 0 {
			return n
		}
		for i := 0; i < len(a.order); i++ {
			if n := cmp.Compare(a.order[i], b.order[i]); n != 0 {
				return n
			}
		}
		return 0
	})
	sum := 0
	for i := 0; i < len(ranks); i++ {
		sum += (i + 1) * ranks[i].bid
	}
	fmt.Println(sum)
}

func pt2() {
	relation['J'] = joker // override J
	in := strings.Split(input, "\n")
	ranks := make([]hand, len(in))
	for i, s := range in {
		order, bid, k := parseCardJoker(s)
		ranks[i] = hand{
			debugStr: s,
			order:    order,
			bid:      bid,
			kind:     k,
		}
	}

	slices.SortFunc(ranks, func(a, b hand) int {
		if n := cmp.Compare(a.kind, b.kind); n != 0 {
			return n
		}
		for i := 0; i < len(a.order); i++ {
			if n := cmp.Compare(a.order[i], b.order[i]); n != 0 {
				return n
			}
		}
		return 0
	})
	sum := 0
	for i := 0; i < len(ranks); i++ {
		bid := ranks[i].bid
		mul := i + 1
		sum += bid * mul
	}
	fmt.Println(sum)
}

func parseCard(s string) (hand []typeCard, bid int, kindType kind) {
	h, b, _ := strings.Cut(s, " ")
	diff := make(map[typeCard]int)
	for _, r := range h {
		c := relation[r]
		hand = append(hand, c)
		diff[c]++
	}

	bid, _ = strconv.Atoi(b)
	pairs := 0
	threes := 0
	uniques := 0
	for _, nbCard := range diff {
		switch nbCard {
		case 1:
			uniques++
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			return hand, bid, fourKindType
		case 5:
			return hand, bid, fiveKindType
		}
	}

	switch {
	case pairs == 1 && threes == 1:
		return hand, bid, fullHouseType
	case pairs == 0 && threes == 1:
		return hand, bid, threeKindType
	case pairs == 2:
		return hand, bid, twoPairType
	case pairs == 1:
		return hand, bid, pairType
	case uniques == 5:
		return hand, bid, highType
	default:
		panic("wrong hand")
	}
}

func parseCardJoker(s string) (hand []typeCard, bid int, kindType kind) {
	h, b, _ := strings.Cut(s, " ")
	diff := make(map[typeCard]int)
	for _, r := range h {
		c := relation[r]
		hand = append(hand, c)
		diff[c]++
	}

	bid, _ = strconv.Atoi(b)

	nbJokers, exist := diff[joker]
	if exist {
		if nbJokers == 5 {
			return hand, bid, fiveKindType
		}

		delete(diff, joker)
		highestIdx := joker
		highest := 0
		for c, i := range diff {
			if i > highest {
				highest = i
				highestIdx = c
			}
		}
		diff[highestIdx] += nbJokers
	}

	// this whole block can probably be optimized
	pairs := 0
	threes := 0
	uniques := 0
	typeHand := highType
	for _, nbCard := range diff {
		switch nbCard {
		case 1:
			uniques++
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			return hand, bid, fourKindType
		case 5:
			return hand, bid, fiveKindType
		}
	}

	switch {
	case pairs == 1 && threes == 1:
		typeHand = fullHouseType
	case pairs == 0 && threes == 1:
		typeHand = threeKindType
	case pairs == 2:
		typeHand = twoPairType
	case pairs == 1:
		typeHand = pairType
	case uniques == 5:
		typeHand = highType
	default:
		panic("wrong hand")
	}

	return hand, bid, typeHand
}
