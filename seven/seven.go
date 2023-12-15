package seven

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Cards string

type Hand struct {
	Cards Cards
	Bid   int
}

func (c Cards) typeCount() map[rune]int {
	count := make(map[rune]int)
	for _, card := range c {
		count[card]++
	}
	return count
}

func (c Cards) TypeOne() int {
	count := c.typeCount()
	return typeValue(count)
}

func (c Cards) TypeTwo() int {
	count := c.typeCount()
	jokerCount, ok := count['J']
	if !ok {
		// there are no jokers
		return typeValue(count)
	}
	delete(count, 'J')
	if len(count) == 0 {
		// there are only jokers
		return FiveOfAKind
	}
	var maxKey rune
	for k, v := range count {
		if v > count[maxKey] {
			maxKey = k
		}
	}
	// add jokers to the most common card
	count[maxKey] += jokerCount
	return typeValue(count)
}

func typeValue(count map[rune]int) int {
	switch {
	case len(count) == 1:
		return FiveOfAKind
	case len(count) == 2:
		for _, v := range count {
			switch v {
			case 1, 4:
				return FourOfAKind
			case 2, 3:
				return FullHouse
			}
		}
		panic("invalid cards")
	case len(count) == 3:
		for _, v := range count {
			switch v {
			case 2:
				return TwoPair
			case 3:
				return ThreeOfAKind
			}
		}
		panic("invalid cards")
	case len(count) == 4:
		return OnePair
	case len(count) == 5:
		return HighCard
	default:
		panic("invalid card count")
	}
}

func cardValueOne(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func cardValueTwo(card rune) int {
	if card == 'J' {
		return 0
	}
	return cardValueOne(card)
}

func parseInput(lines []string) ([]Hand, error) {
	handRegex := regexp.MustCompile(`^([AKQJT98765432]{5}) (\d+)$`)
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		matches := handRegex.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, fmt.Errorf("invalid hand '%s'", line)
		}
		bid, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("parse bid: %v", err)
		}
		hands[i] = Hand{
			Cards: Cards(matches[1]),
			Bid:   bid,
		}
	}
	return hands, nil
}

func One(lines []string) (int, error) {
	hands, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Cards.TypeOne() != b.Cards.TypeOne() {
			return a.Cards.TypeOne() - b.Cards.TypeOne()
		}
		for i, ca := range a.Cards {
			cb := rune(b.Cards[i])
			if cardValueOne(ca) != cardValueOne(cb) {
				return cardValueOne(ca) - cardValueOne(cb)
			}
		}
		return 0
	})
	sum := 0
	for i, hand := range hands {
		sum += hand.Bid * (i + 1)
	}
	return sum, nil
}

func Two(lines []string) (int, error) {
	hands, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Cards.TypeTwo() != b.Cards.TypeTwo() {
			return a.Cards.TypeTwo() - b.Cards.TypeTwo()
		}
		for i, ca := range a.Cards {
			cb := rune(b.Cards[i])
			if cardValueTwo(ca) != cardValueTwo(cb) {
				return cardValueTwo(ca) - cardValueTwo(cb)
			}
		}
		return 0
	})
	sum := 0
	for i, hand := range hands {
		sum += hand.Bid * (i + 1)
	}
	return sum, nil
}
