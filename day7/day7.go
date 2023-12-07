package day7

import (
	"slices"
	"strconv"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

type HandType int

const (
	highCard HandType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind

	aceValue           = 14
	kingValue          = 13
	queenValue         = 12
	jokerValue         = 11
	tenValue           = 10
	jokerWildcardValue = 1
)

type Hand struct {
	cards []int
	_type HandType
	bid   int
}

func (h Hand) compare(other Hand) int {
	if h._type > other._type {
		return -1
	}
	if h._type < other._type {
		return 1
	}

	for i, c := range h.cards {
		if c > other.cards[i] {
			return -1
		}
		if c < other.cards[i] {
			return 1
		}
	}

	return 0
}

func Part1(input string) int {
	return calculateWinnings(parseInput(input, false))
}

func Part2(input string) int {
	return calculateWinnings(parseInput(input, true))
}

func calculateWinnings(hands []Hand) int {
	slices.SortFunc(hands, func(h1, h2 Hand) int { return h2.compare(h1) })

	winnings := 0
	for i, h := range hands {
		rank := i + 1
		winnings += h.bid * rank
	}

	return winnings
}

func parseInput(input string, jokerIsWildcard bool) []Hand {
	lines := common.GetLines(input)

	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		cardsStr := line[:5]
		bid, _ := strconv.Atoi(line[6:])

		cards := make([]int, 0, 5)
		for _, c := range cardsStr {
			switch c {
			case 'T':
				cards = append(cards, tenValue)
			case 'J':
				if jokerIsWildcard {
					cards = append(cards, jokerWildcardValue)
				} else {
					cards = append(cards, jokerValue)
				}
			case 'Q':
				cards = append(cards, queenValue)
			case 'K':
				cards = append(cards, kingValue)
			case 'A':
				cards = append(cards, aceValue)
			default:
				cards = append(cards, int(c-'0'))
			}
		}

		hands = append(hands, Hand{cards, getType(cards, jokerIsWildcard), bid})
	}

	return hands
}

func getType(cards []int, jokerIsWildcard bool) HandType {
	cardCounts := make([]int, 15)
	numJokers := 0
	for _, c := range cards {
		if jokerIsWildcard && c == jokerWildcardValue {
			numJokers++
		} else {
			cardCounts[c]++
		}
	}

	slices.SortFunc(cardCounts, func(a, b int) int { return b - a })

	if jokerIsWildcard {
		cardCounts[0] += numJokers
	}

	if cardCounts[0] == 5 {
		return fiveOfAKind
	}

	if cardCounts[0] == 4 {
		return fourOfAKind
	}

	if cardCounts[0] == 3 {
		if cardCounts[1] == 2 {
			return fullHouse
		}

		return threeOfAKind
	}

	if cardCounts[0] == 2 {
		if cardCounts[1] == 2 {
			return twoPair
		}

		return onePair
	}

	return highCard
}
