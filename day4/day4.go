package day4

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
	"github.com/jonasah/advent-of-code-2023/lib/sliceconv"
)

type Card struct {
	idx            int
	winningNumbers []int
	myNumbers      []int
}

func Part1(input string) int {
	cards := parseInput(input)

	points := 0
	for _, card := range cards {
		count := 0
		for _, n := range card.myNumbers {
			if slices.Contains(card.winningNumbers, n) {
				count++
			}
		}
		points += math.Pow(2, count-1)
	}

	return points
}

func Part2(input string) int {
	allCards := parseInput(input)

	var cards []Card
	cards = append(cards, allCards...)

	total := 0
	for len(cards) > 0 {
		total++
		card := cards[0]

		count := 0
		for _, n := range card.myNumbers {
			if slices.Contains(card.winningNumbers, n) {
				count++
			}
		}
		if count > 0 {
			cards = append(cards, allCards[card.idx+1:card.idx+1+count]...)
		}

		cards = cards[1:]
	}

	return total
}

func parseInput(input string) []Card {
	lines := common.GetLines(input)

	cards := make([]Card, 0, len(lines))
	for n, line := range lines {
		x := strings.Split(strings.Split(line, ": ")[1], " | ")

		cards = append(cards, Card{
			idx:            n,
			winningNumbers: sliceconv.Atoi(strings.Fields(x[0])),
			myNumbers:      sliceconv.Atoi(strings.Fields(x[1])),
		})
	}

	return cards
}
