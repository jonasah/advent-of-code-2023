package day22

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
	"github.com/jonasah/advent-of-code-2023/lib/sliceconv"
)

type Block struct {
	i int

	zMin, zMax int

	xyArea []struct{ x, y int }
}

func NewBlock(i, xMin, xMax, yMin, yMax, zMin, zMax int) Block {
	b := Block{i, zMin, zMax, nil}

	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			b.xyArea = append(b.xyArea, struct{ x, y int }{x, y})
		}
	}

	return b
}

func (b *Block) setZMin(zMin int) {
	d := b.zMax - b.zMin
	b.zMin = zMin
	b.zMax = zMin + d
}

func (b Block) equals(other Block) bool {
	return b.i == other.i && b.zMin == other.zMin
}

func Part1(input string) int {
	numSafe := 0
	doIt(input, func(numChanged int) {
		if numChanged == 0 {
			numSafe++
		}
	})

	return numSafe
}

func Part2(input string) int {
	totalNumChanged := 0
	doIt(input, func(numChanged int) { totalNumChanged += numChanged })

	return totalNumChanged
}

func doIt(input string, cb func(numChanged int)) {
	blocks := parseInput(input)

	fall(blocks)

	blockMap := make(map[int]Block, len(blocks))
	for _, block := range blocks {
		blockMap[block.i] = block
	}

	for _, blockToRemove := range blocks {
		var newBlocks []Block
		newBlocks = append(newBlocks, blocks...)
		newBlocks = slices.DeleteFunc(newBlocks, func(b Block) bool { return b.equals(blockToRemove) })

		fall(newBlocks)

		numChanged := 0
		for _, b := range newBlocks {
			if !b.equals(blockMap[b.i]) {
				numChanged++
			}
		}

		cb(numChanged)
	}
}

func parseInput(input string) []Block {
	lines := common.GetLines(input)

	blocks := make([]Block, 0, len(lines))
	for i, line := range lines {
		s := strings.Split(line, "~")
		p0 := sliceconv.Atoi(strings.Split(s[0], ","))
		p1 := sliceconv.Atoi(strings.Split(s[1], ","))

		blocks = append(blocks, NewBlock(
			i,
			math.Min(p0[0], p1[0]),
			math.Max(p0[0], p1[0]),
			math.Min(p0[1], p1[1]),
			math.Max(p0[1], p1[1]),
			math.Min(p0[2], p1[2]),
			math.Max(p0[2], p1[2]),
		))
	}

	return blocks
}

func fall(blocks []Block) {
	slices.SortFunc(blocks, func(a, b Block) int { return a.zMin - b.zMin })

	heightMap := [10][10]int{}

	for i := range blocks {
		h := 0
		for _, p := range blocks[i].xyArea {
			h = math.Max(h, heightMap[p.x][p.y])
		}

		if blocks[i].zMin-h > 1 {
			blocks[i].setZMin(h + 1)
		}

		for _, p := range blocks[i].xyArea {
			heightMap[p.x][p.y] = blocks[i].zMax
		}
	}
}
