package typewriter

import (
	"bytes"
	"fmt"
)

type Config struct {
	ShowLineNumbers          bool
	MarkFirstDifference      bool
	MarkFirstLineDifferences bool

	Padding   int
	Separator string
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the length of the longest string.
func maxWidth(lines []string) int {
	max := 0
	for _, line := range lines {
		max = maxInt(len(line), max)
	}
	return max
}

func getOrEmpty(strings []string, index int) string {
	if index >= len(strings) {
		return ""
	}
	return strings[index]
}

func rightZero(s string, desiredLength int, padChar string) string {
	if len(s) > desiredLength {
		panic(fmt.Sprintf("Length of string (%d) is greater than desired length (%d)", len(s), desiredLength))
	}

	// TODO: Optimize
	padding := desiredLength - len(s)
	for i := 0; i < padding; i++ {
		s = s + padChar
	}

	return s
}

func Run(lines1, lines2 []string, config Config) string {
	var buf bytes.Buffer

	leftColumnWidth := maxWidth(lines1)

	padding := rightZero("", config.Padding, " ")

	for i := 0; i < maxInt(len(lines1), len(lines2)); i++ {
		s1 := getOrEmpty(lines1, i)
		s2 := getOrEmpty(lines2, i)

		s1 = rightZero(s1, leftColumnWidth, " ")

		buf.WriteString(s1 + padding + config.Separator + s2 + "\n")
	}

	return buf.String()
}
