package day4

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
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
		card := Card{idx: n, winningNumbers: make([]int, 0), myNumbers: make([]int, 0)}
		x := strings.Split(strings.Split(line, ": ")[1], " | ")

		w := strings.Fields(x[0])
		for _, v := range w {
			n, _ := strconv.Atoi(v)
			card.winningNumbers = append(card.winningNumbers, n)
		}

		m := strings.Fields(x[1])
		for _, v := range m {
			n, _ := strconv.Atoi(v)
			card.myNumbers = append(card.myNumbers, n)
		}

		cards = append(cards, card)
	}

	return cards
}
