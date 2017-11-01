package typewriter

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	colorOff     = "\033[0m"
	colorDefault = "\033[41m"
)

type Config struct {
	ShowLineNumbers bool

	// If true, marks the first spot where the two columns diverge.
	MarkFirstDifference bool

	// The escape sequence to use to mark a difference. If not specified
	// defaults to a red background.
	Marking string

	Padding   int
	Separator string

	LeftHeader  string
	RightHeader string
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
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

func rightZero(s string, desiredLength int, padChar string) (string, error) {
	if len(s) > desiredLength {
		return "", fmt.Errorf("length of string (%d) is greater than desired length (%d)", len(s), desiredLength)
	}

	padding := desiredLength - len(s)
	return s + strings.Repeat(padChar, padding), nil
}

func leftZero(s string, desiredLength int, padChar string) (string, error) {
	if len(s) > desiredLength {
		return "", fmt.Errorf("length of string (%d) is greater than desired length (%d)", len(s), desiredLength)
	}

	padding := desiredLength - len(s)
	return strings.Repeat(padChar, padding) + s, nil
}

func findDifference(s1, s2 string) (index int, found bool) {
	minLen := minInt(len(s1), len(s2))
	for i := 0; i < minLen; i++ {
		c1 := s1[i]
		c2 := s2[i]

		if c1 != c2 {
			return i, true
		}
	}

	if len(s1) > len(s2) {
		return len(s2), true
	} else if len(s2) > len(s1) {
		return len(s1), true
	}

	return -1, false
}

func Sprint(lines1, lines2 []string, config Config) (string, error) {
	var buf bytes.Buffer

	leftColumnWidth := maxWidth(lines1)
	if len(config.LeftHeader) > leftColumnWidth {
		leftColumnWidth = len(config.LeftHeader)
	}

	padding, err := rightZero("", config.Padding, " ")
	if err != nil {
		return "", err
	}

	if config.Marking == "" {
		config.Marking = colorDefault
	}

	// Find the max width we'll need for line numbers.
	maxLineNumber := maxInt(len(lines1), len(lines2))
	maxLineNumberWidth := len(strconv.Itoa(maxLineNumber)) + 2

	if config.LeftHeader != "" || config.RightHeader != "" {
		h1, err := rightZero(config.LeftHeader, leftColumnWidth, " ")
		if err != nil {
			return "", err
		}

		lineNumber := ""
		if config.ShowLineNumbers {
			lineNumber, err = leftZero(lineNumber, maxLineNumberWidth, " ")
			if err != nil {
				return "", err
			}
		}
		buf.WriteString(lineNumber + h1 + padding + config.Separator + config.RightHeader + "\n\n")
	}

	differenceFound := false
	for i := 0; i < maxInt(len(lines1), len(lines2)); i++ {
		s1 := getOrEmpty(lines1, i)
		s2 := getOrEmpty(lines2, i)

		extraWidth := 0
		if config.MarkFirstDifference && !differenceFound {
			diffIndex, found := findDifference(s1, s2)
			if found {
				differenceFound = true
				s1 = s1[:diffIndex] + config.Marking + s1[diffIndex:] + colorOff
				s2 = s2[:diffIndex] + config.Marking + s2[diffIndex:] + colorOff
				extraWidth = len(config.Marking) + len(colorOff)
			}
		}

		s1, err = rightZero(s1, leftColumnWidth+extraWidth, " ")
		if err != nil {
			return "", err
		}

		if config.ShowLineNumbers {
			lineNumber := strconv.Itoa(i+1) + ". "

			lineNumber, err = leftZero(lineNumber, maxLineNumberWidth, " ")
			if err != nil {
				return "", err
			}

			buf.WriteString(lineNumber)
		}

		buf.WriteString(s1 + padding + config.Separator + s2 + "\n")
	}

	return buf.String(), nil
}
