package typewriter_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/j-delaney/typewriter"
	"github.com/stretchr/testify/assert"
)

func doTest(t *testing.T, tc testCase) {
	// Remove the blank starting newlines and prefix \t that goformatting leaves
	tc.expected = tc.expected[2:]
	expectedSplit := strings.Split(tc.expected, "\n")
	var buf bytes.Buffer
	for _, line := range expectedSplit {
		buf.WriteString(strings.TrimLeft(line, "\t") + "\n")
	}

	expected := buf.String()

	actual := typewriter.Run(tc.lines1, tc.lines2, tc.config)
	assert.Equal(t, expected, actual)
}

type testCase struct {
	name string

	lines1 []string
	lines2 []string
	config typewriter.Config

	expected string
}

var testCases = []testCase{
	{
		name: "simple",

		lines1: []string{"a", "b", "c"},
		lines2: []string{"d", "e", "f"},
		config: typewriter.Config{},

		expected: `
		ad
		be
		cf`,
	},
	{
		name: "different lengths",

		lines1: []string{"one", "two", "three", "four"},
		lines2: []string{"five", "six", "seven", "eight"},
		config: typewriter.Config{},

		expected: `
		one  five
		two  six
		threeseven
		four eight`,
	},
	{
		name: "column 1 longer",

		lines1: []string{"1", "100", "10"},
		lines2: []string{"2000", "2"},
		config: typewriter.Config{},

		expected: `
		1  2000
		1002
		10`,
	},
	{
		name: "column 2 longer",

		lines1: []string{"2000", "2"},
		lines2: []string{"1", "100", "10"},
		config: typewriter.Config{},

		expected: `
		20001
		2   100
		    10`,
	},
	{
		name: "padding",

		lines1: []string{"foo", "bars"},
		lines2: []string{"baz", "sizzle"},
		config: typewriter.Config{
			Padding: 3,
		},

		expected: `
		foo    baz
		bars   sizzle`,
	},
	{
		name: "separator",

		lines1: []string{"foo", "bars"},
		lines2: []string{"baz", "sizzle"},
		config: typewriter.Config{
			Padding:   3,
			Separator: "|",
		},

		expected: `
		foo    |baz
		bars   |sizzle`,
	},
}

func TestRun(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			doTest(t, tc)
		})
	}
}
