package day25

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

func Part1(input string, wiresToCut []string) int {
	lines := common.GetLines(input)

	graph := map[string][]string{}

	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsAny(string(r), ": ") })
		from := fields[0]

		for _, to := range fields[1:] {
			if slices.Contains(wiresToCut, fmt.Sprintf("%s %s", from, to)) || slices.Contains(wiresToCut, fmt.Sprintf("%s %s", to, from)) {
				continue
			}

			graph[from] = append(graph[from], to)
			graph[to] = append(graph[to], from)
		}
	}

	return graphSize(graph, wiresToCut[0][:3]) * graphSize(graph, wiresToCut[0][4:])
}

func graphSize(graph map[string][]string, component string) int {
	toVisit := []string{component}
	var visited []string
	for len(toVisit) > 0 {
		component := toVisit[0]
		toVisit = toVisit[1:]

		if slices.Contains(visited, component) {
			continue
		}

		visited = append(visited, component)
		toVisit = append(toVisit, graph[component]...)
	}

	return len(visited)
}
