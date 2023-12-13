package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 0, "Day to generate code for")
	flag.Parse()

	dayStr := strconv.Itoa(day)

	if flag.NFlag() == 0 {
		fmt.Fprintln(os.Stderr, "Missing required 'day' argument")
		os.Exit(1)
	}
	if day < 1 || day > 25 {
		fmt.Fprintln(os.Stderr, "Invalid day: "+dayStr)
		os.Exit(1)
	}

	dirName := "day" + dayStr

	implFilePath := path.Join(dirName, "day"+dayStr+".go")
	implContent := getImplContent(dayStr)

	testFilePath := path.Join(dirName, "day"+dayStr+"_test.go")
	testContent := getTestContent(dayStr)

	inputFilePath := path.Join(dirName, "input.txt")

	if !pathExists(dirName) {
		os.Mkdir(dirName, 0777)
	}
	if !pathExists(implFilePath) {
		os.WriteFile(implFilePath, []byte(implContent), 0666)
	}
	if !pathExists(testFilePath) {
		os.WriteFile(testFilePath, []byte(testContent), 0666)
	}
	if !pathExists(inputFilePath) {
		os.Create(inputFilePath)
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getImplContent(day string) string {
	return `package day` + day + `

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}
`
}

func getTestContent(day string) string {
	return `package day` + day + `

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, -1, Part1(testInput))
	// require.Equal(t, -1, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, -1, Part2(testInput))
	// require.Equal(t, -1, Part2(realInput))
}

var realInput = common.GetInput(` + day + `)

const testInput = ` + "``" + `
`
}
