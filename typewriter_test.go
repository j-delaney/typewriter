package typewriter

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

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

	actual := Run(tc.lines1, tc.lines2, tc.config)

	if expected != actual {
		if tc.config.MarkFirstDifference {
			t.Log("Not using typewriter to show difference because color characters mess up output")
		} else {
			t.Log("\n" + Run(strings.Split(expected, "\n"), strings.Split(actual, "\n"), Config{
				MarkFirstDifference: true,
				Separator:           "‖",
				Padding:             3,

				LeftHeader:  "expected",
				RightHeader: "actual",
			}))
		}
	}
	assert.Equal(t, expected, actual)
}

type testCase struct {
	name string

	lines1 []string
	lines2 []string
	config Config

	expected string
}

var testCases = []testCase{
	{
		name: "simple",

		lines1: []string{"a", "b", "c"},
		lines2: []string{"d", "e", "f"},
		config: Config{},

		expected: `
		ad
		be
		cf`,
	},
	{
		name: "different lengths",

		lines1: []string{"one", "two", "three", "four"},
		lines2: []string{"five", "six", "seven", "eight"},
		config: Config{},

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
		config: Config{},

		expected: `
		1  2000
		1002
		10 `,
	},
	{
		name: "column 2 longer",

		lines1: []string{"2000", "2"},
		lines2: []string{"1", "100", "10"},
		config: Config{},

		expected: `
		20001
		2   100
		    10`,
	},
	{
		name: "padding",

		lines1: []string{"foo", "bars"},
		lines2: []string{"baz", "sizzle"},
		config: Config{
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
		config: Config{
			Padding:   3,
			Separator: "|",
		},

		expected: `
		foo    |baz
		bars   |sizzle`,
	},
	{
		name: "line numbers",

		lines1: []string{"foo", "bars"},
		lines2: []string{"baz", "sizzle", "bat"},
		config: Config{
			Padding:         3,
			Separator:       "|",
			ShowLineNumbers: true,
		},

		expected: `
		1. foo    |baz
		2. bars   |sizzle
		3.        |bat`,
	},
	{
		name: "lots of line numbers",

		lines1: []string{"a", "a"},
		lines2: []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
		config: Config{
			Padding:         3,
			Separator:       "|",
			ShowLineNumbers: true,
		},

		expected: `
		 1. a   |a
		 2. a   |a
		 3.     |a
		 4.     |a
		 5.     |a
		 6.     |a
		 7.     |a
		 8.     |a
		 9.     |a
		10.     |a
		11.     |a
		12.     |a`,
	},
	{
		name: "mark no difference",

		lines1: []string{"a", "a"},
		lines2: []string{"a", "a"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: `
		a|a
		a|a`,
	},
	{
		name: "mark simple difference",

		lines1: []string{"a", "a"},
		lines2: []string{"a", "b"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: fmt.Sprintf(`
		a|a
		%sa%s|%sb%s`, colorDefault, colorOff, colorDefault, colorOff),
	},
	{
		name: "mark difference",

		lines1: []string{"abc", "def", "ghi"},
		lines2: []string{"abc", "dff", "ghi"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: fmt.Sprintf(`
		abc|abc
		d%sef%s|d%sff%s
		ghi|ghi`, colorDefault, colorOff, colorDefault, colorOff),
	},
	{
		name: "mark left column missing line difference",

		lines1: []string{"abc", "def"},
		lines2: []string{"abc", "def", "ghi"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: fmt.Sprintf(`
		abc|abc
		def|def
		%s%s   |%sghi%s`, colorDefault, colorOff, colorDefault, colorOff),
	},
	{
		name: "mark right column missing line difference",

		lines1: []string{"abc", "def", "ghi"},
		lines2: []string{"abc", "def"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: fmt.Sprintf(`
		abc|abc
		def|def
		%sghi%s|%s%s`, colorDefault, colorOff, colorDefault, colorOff),
	},
	{
		name: "mark difference left column longer",

		lines1: []string{"abc", "defg", "hij"},
		lines2: []string{"abc", "def", "hij"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: fmt.Sprintf(`
		abc |abc
		def%sg%s|def%s%s
		hij |hij`, colorDefault, colorOff, colorDefault, colorOff),
	},
	{
		name: "mark difference right column longer",

		lines1: []string{"abc", "def", "hij"},
		lines2: []string{"abc", "defg", "hij"},
		config: Config{
			Separator:           "|",
			MarkFirstDifference: true,
		},

		expected: fmt.Sprintf(`
		abc|abc
		def%s%s|def%sg%s
		hij|hij`, colorDefault, colorOff, colorDefault, colorOff),
	},
	{
		name: "simple headers",

		lines1: []string{"abc", "def", "hij"},
		lines2: []string{"abc", "def", "hij"},
		config: Config{
			Separator:   "|",
			LeftHeader:  "l",
			RightHeader: "r",
		},

		expected: `
		l  |r

		abc|abc
		def|def
		hij|hij`,
	},
	{
		name: "long headers",

		lines1: []string{"abc", "def", "hij"},
		lines2: []string{"abc", "def", "hij"},
		config: Config{
			Separator:   "|",
			LeftHeader:  "left header",
			RightHeader: "right header",
		},

		expected: `
		left header|right header

		abc        |abc
		def        |def
		hij        |hij`,
	},
	{
		name: "left header only",

		lines1: []string{"abc", "def", "hij"},
		lines2: []string{"abc", "def", "hij"},
		config: Config{
			Separator:  "|",
			LeftHeader: "left header",
		},

		expected: `
		left header|

		abc        |abc
		def        |def
		hij        |hij`,
	},
	{
		name: "right header only",

		lines1: []string{"abc", "def", "hij"},
		lines2: []string{"abc", "def", "hij"},
		config: Config{
			Separator:   "|",
			RightHeader: "right header",
		},

		expected: `
		   |right header

		abc|abc
		def|def
		hij|hij`,
	},
}

func TestRun(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			doTest(t, tc)
		})
	}
}

var (
	simpleLine1 = []string{"abc", "defg", "hijk", "l", "mnopq", "rstuv", "wx", "y", "z"}
	simpleLine2 = []string{"abcde", "f", "g", "hijkl", "mn", "opqrstuv", "wx", "yz"}

	smallCodeSnippet1 = []string{
		"func getOrEmpty(strings []string, index int) string {",
		"	if index >= len(strings) {",
		"		return \"\"",
		"	}",
		"	return strings[index]",
		"}",
	}
	smallCodeSnippet2 = []string{
		"func getOrEmpty(strings []string, index int) string {",
		"	if len(strings) <= index {",
		"		return \"\"",
		"	}",
		"	return strings[index]",
		"}",
	}
)

var benchmarks = []testCase{
	{
		name: "base",

		config: Config{},
	},
	{
		name: "difference",

		config: Config{
			MarkFirstDifference: true,
		},
	},
	{
		name: "line numbers",

		config: Config{
			ShowLineNumbers: true,
		},
	},
	{
		name: "padding",

		config: Config{
			Padding: 50,
		},
	},
	{
		name: "separator",

		config: Config{
			Separator: "|",
		},
	},
	{
		name: "headers",

		config: Config{
			LeftHeader:  "Left Header",
			RightHeader: "RightHeader",
		},
	},
}

func BenchmarkRun(b *testing.B) {
	for _, tc := range benchmarks {
		tc := tc

		tc.lines1 = simpleLine1
		tc.lines2 = simpleLine2

		b.Run("Simple/" + tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Run(tc.lines1, tc.lines2, tc.config)
			}
		})
	}

	for _, tc := range benchmarks {
		tc := tc

		tc.lines1 = smallCodeSnippet1
		tc.lines2 = smallCodeSnippet2
		b.Run("SmallCodeSnippet/" + tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Run(tc.lines1, tc.lines2, tc.config)
			}
		})
	}
}
