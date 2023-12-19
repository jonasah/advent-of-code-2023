package day19

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
	"golang.org/x/exp/maps"
)

type Category string
type Operator string

const (
	catX Category = "x"
	catM Category = "m"
	catA Category = "a"
	catS Category = "s"

	accepted = "A"
	rejected = "R"

	lessThan    = "<"
	greaterThan = ">"

	startWorkflow = "in"
)

type Rule struct {
	cat Category
	op  Operator
	cmp int
	dst string
}

func (r Rule) eval(part Part) bool {
	if r.op == lessThan {
		return part[r.cat] < r.cmp
	}

	if r.op == greaterThan {
		return part[r.cat] > r.cmp
	}

	return true
}

type Workflow struct {
	name  string
	rules []Rule
}

func (w Workflow) eval(part Part) string {
	idx := slices.IndexFunc(w.rules, func(r Rule) bool { return r.eval(part) })
	return w.rules[idx].dst
}

type Part map[Category]int

func NewPart(x, m, a, s int) Part {
	return map[Category]int{catX: x, catM: m, catA: a, catS: s}
}

func (p Part) totalRating() int {
	return math.Sum(maps.Values(p))
}

type Range struct {
	min, max int
}

func (r Range) length() int {
	return r.max - r.min + 1
}

type PartRange map[Category]Range

func NewPartRange(x, m, a, s Range) PartRange {
	return map[Category]Range{catX: x, catM: m, catA: a, catS: s}
}

func (pr PartRange) copy() PartRange {
	return NewPartRange(pr[catX], pr[catM], pr[catA], pr[catS])
}

func (pr PartRange) numCombinations() int {
	ranges := maps.Values(pr)
	c := 1
	for _, r := range ranges {
		c *= r.length()
	}
	return c
}

type Candidate struct {
	pr       PartRange
	workflow string
}

func Part1(input string) int {
	workflows, parts := parseInput(input)

	sum := 0
	for _, part := range parts {
		currWorkflow := startWorkflow
		for !(currWorkflow == accepted || currWorkflow == rejected) {
			currWorkflow = workflows[currWorkflow].eval(part)
		}

		if currWorkflow == accepted {
			sum += part.totalRating()
		}
	}

	return sum
}

func Part2(input string) int {
	workflows, _ := parseInput(input)

	defaultRange := Range{1, 4000}
	candidates := []Candidate{{NewPartRange(defaultRange, defaultRange, defaultRange, defaultRange), startWorkflow}}

	totalCombinations := 0
	for len(candidates) > 0 {
		c := candidates[0]
		candidates = candidates[1:]

		if c.workflow == accepted {
			totalCombinations += c.pr.numCombinations()
			continue
		}
		if c.workflow == rejected {
			continue
		}

		candidates = append(candidates, evalRange(c.pr, workflows[c.workflow])...)
	}

	return totalCombinations
}

var workflowRegExp = regexp.MustCompile(`(\w+){(.*)}`)
var ruleRegExp = regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`)
var partRegExp = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

func parseInput(input string) (map[string]Workflow, []Part) {

	content := strings.Split(input, "\n\n")
	workflowLines := common.GetLines(content[0])
	partLines := common.GetLines(content[1])

	workflows := make(map[string]Workflow, len(workflowLines))
	for _, line := range workflowLines {
		match := workflowRegExp.FindStringSubmatch(line)
		workflow := Workflow{name: match[1]}

		rules := strings.Split(match[2], ",")
		for _, r := range rules[:len(rules)-1] {
			m := ruleRegExp.FindStringSubmatch(r)
			cat := Category(m[1])
			op := Operator(m[2])
			cmp, _ := strconv.Atoi(m[3])
			dst := m[4]

			workflow.rules = append(workflow.rules, Rule{cat, op, cmp, dst})
		}

		workflow.rules = append(workflow.rules, Rule{dst: rules[len(rules)-1]})

		workflows[workflow.name] = workflow
	}

	var parts []Part
	for _, line := range partLines {
		match := partRegExp.FindStringSubmatch(line)
		x, _ := strconv.Atoi(match[1])
		m, _ := strconv.Atoi(match[2])
		a, _ := strconv.Atoi(match[3])
		s, _ := strconv.Atoi(match[4])

		parts = append(parts, NewPart(x, m, a, s))
	}

	return workflows, parts
}

func evalRange(pr PartRange, workflow Workflow) []Candidate {
	newCandidates := []Candidate{}
	for _, rule := range workflow.rules {
		var fulfilled PartRange
		fulfilled, pr = splitRange(pr, rule)

		newCandidates = append(newCandidates, Candidate{fulfilled, rule.dst})
	}

	return newCandidates
}

func splitRange(pr PartRange, rule Rule) (PartRange, PartRange) {
	if rule.op == lessThan {
		lower := pr.copy()
		lower[rule.cat] = Range{pr[rule.cat].min, rule.cmp - 1}

		upper := pr.copy()
		upper[rule.cat] = Range{rule.cmp, pr[rule.cat].max}

		return lower, upper
	}

	if rule.op == greaterThan {
		lower := pr.copy()
		lower[rule.cat] = Range{pr[rule.cat].min, rule.cmp}

		upper := pr.copy()
		upper[rule.cat] = Range{rule.cmp + 1, pr[rule.cat].max}

		return upper, lower
	}

	return pr, pr
}
