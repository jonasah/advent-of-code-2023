package day2

import (
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

type Color string

type CubeSet map[Color]int

type Game struct {
	id    int
	cubes []CubeSet
}

const (
	red   Color = "red"
	green Color = "green"
	blue  Color = "blue"
)

func Part1(input string) int {
	games := parseGames(input)

	const maxRed = 12
	const maxGreen = 13
	const maxBlue = 14

	res := 0

outer:
	for _, g := range games {
		for _, cs := range g.cubes {
			if cs[red] > maxRed || cs[green] > maxGreen || cs[blue] > maxBlue {
				continue outer
			}
		}

		res += g.id
	}

	return res
}

func Part2(input string) int {
	games := parseGames(input)

	res := 0
	for _, g := range games {
		min := CubeSet{}

		for _, cs := range g.cubes {
			min[red] = math.Max(min[red], cs[red])
			min[green] = math.Max(min[green], cs[green])
			min[blue] = math.Max(min[blue], cs[blue])
		}

		res += min[red] * min[green] * min[blue]
	}

	return res
}

func parseGames(input string) []Game {
	lines := common.GetLines(input)

	games := make([]Game, 0, len(input))
	for _, v := range lines {
		games = append(games, parseGame(v))
	}

	return games
}

func parseGame(line string) Game {
	x := strings.Split(line, ": ")

	gameId, _ := strconv.Atoi(strings.Split(x[0], " ")[1])
	sets := strings.Split(x[1], "; ")

	game := Game{id: gameId, cubes: make([]CubeSet, 0)}
	for _, s := range sets {
		c := strings.Split(s, ", ")

		set := CubeSet{}
		for _, v := range c {
			a := strings.Split(v, " ")
			amount, _ := strconv.Atoi(a[0])
			color := Color(a[1])

			set[color] = amount
		}

		game.cubes = append(game.cubes, set)
	}

	return game
}
