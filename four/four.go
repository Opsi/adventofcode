package four

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	ID        int
	Scratched []int
	Winning   []int
}

func (c Card) WinningNumbers() int {
	sum := 0
	for _, nr := range c.Winning {
		if slices.Contains(c.Scratched, nr) {
			sum++
		}
	}
	return sum
}

func (c Card) Points() int {
	sum := 0
	for _, nr := range c.Winning {
		if !slices.Contains(c.Scratched, nr) {
			continue
		}
		if sum == 0 {
			sum = 1
		} else {
			sum *= 2
		}
	}
	return sum
}

var cardRegex = regexp.MustCompile(`Card\s*(\d+): (.+) \| (.+)`)

func parseCard(s string) (Card, error) {
	var card Card
	var err error
	matches := cardRegex.FindStringSubmatch(s)
	if len(matches) != 4 {
		return card, fmt.Errorf("invalid card: %s", s)
	}
	card.ID, err = strconv.Atoi(matches[1])
	if err != nil {
		return card, fmt.Errorf("invalid id: %v", err)
	}
	card.Scratched, err = parseNumbers(matches[2])
	if err != nil {
		return card, fmt.Errorf("invalid scratched: %v", err)
	}
	card.Winning, err = parseNumbers(matches[3])
	if err != nil {
		return card, fmt.Errorf("invalid winning: %v", err)
	}
	return card, nil
}

func parseNumbers(s string) ([]int, error) {
	subs := strings.Split(s, " ")
	numbers := make([]int, 0, len(subs))
	for _, sub := range subs {
		trimmed := strings.TrimSpace(sub)
		if trimmed == "" {
			continue
		}
		nr, err := strconv.Atoi(trimmed)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s': %v", trimmed, err)
		}
		numbers = append(numbers, nr)
	}
	return numbers, nil
}

func One(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		card, err := parseCard(line)
		if err != nil {
			return 0, fmt.Errorf("parse card: %v", err)
		}
		sum += card.Points()
	}
	return sum, nil
}

func Two(lines []string) (int, error) {
	cardValues := make([]int, len(lines))
	cardAmount := make([]int, len(lines))
	for i, line := range lines {
		card, err := parseCard(line)
		if err != nil {
			return 0, fmt.Errorf("parse card: %v", err)
		}
		cardValues[i] = card.WinningNumbers()
		cardAmount[i] = 1
	}
	total := 0
	for i := 0; i < len(cardValues); i++ {
		amount := cardAmount[i]
		total += amount
		value := cardValues[i]
		upperBound := min(len(cardValues), i+value+1)
		for j := i + 1; j < upperBound; j++ {
			cardAmount[j] += amount
		}
	}
	return total, nil
}
