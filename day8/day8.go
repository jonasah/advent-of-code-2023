package day8

import (
	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

type Node struct {
	name, left, right string
}

func Part1(input string) int {
	lines := common.GetLines(input)

	instructions := lines[0]
	nodes := parseNodes(lines[2:])

	return traverse(instructions, nodes, "AAA", func(node Node) bool { return node.name == "ZZZ" })
}

func Part2(input string) int {
	lines := common.GetLines(input)

	instructions := lines[0]
	nodes := parseNodes(lines[2:])

	startNodes := make([]Node, 0)
	for _, node := range nodes {
		if node.name[len(node.name)-1] == 'A' {
			startNodes = append(startNodes, node)
		}
	}

	steps := make([]int, 0)
	for _, node := range startNodes {
		s := traverse(instructions, nodes, node.name, func(node Node) bool {
			return node.name[len(node.name)-1] == 'Z'
		})
		steps = append(steps, s)
	}

	return math.Lcm(steps[0], steps[1], steps[2:]...)
}

func traverse(instructions string, nodes map[string]Node, startNode string, isDone func(node Node) bool) int {
	steps := 0
	currentNode := nodes[startNode]
	currentInstruction := 0
	for !isDone(currentNode) {
		steps++

		if instructions[currentInstruction] == 'L' {
			currentNode = nodes[currentNode.left]
		} else {
			currentNode = nodes[currentNode.right]
		}

		currentInstruction++
		if currentInstruction == len(instructions) {
			currentInstruction = 0
		}
	}

	return steps
}

func parseNodes(lines []string) map[string]Node {
	nodes := make(map[string]Node, len(lines))
	for _, line := range lines {
		name := line[:3]
		left := line[7:10]
		right := line[12:15]
		nodes[name] = Node{name, left, right}
	}

	return nodes
}
