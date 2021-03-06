package main

import (
	"flag"
	"os"

	"fmt"
	"io/ioutil"
	"strings"

	"github.com/j-delaney/typewriter"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: side_by_side [options] leftFile rightFile\nOptions:\n")
		flag.PrintDefaults()
	}

	padding := flag.Int("padding", 5, "The padding between the two columns")
	separator := flag.String("separator", "", "Character to separate the two columns")
	markDifference := flag.Bool("diff", false, "Mark the first difference found")
	lineNumbers := flag.Bool("linenums", false, "Show line numbers")
	showHeader := flag.Bool("header", false, "Use the filenames as headers")

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(2)
	}

	filePath1 := flag.Arg(0)
	filePath2 := flag.Arg(1)

	bytes1, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read %v: %v", filePath1, err)
		os.Exit(1)
	}

	bytes2, err := ioutil.ReadFile(filePath2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read %v: %v", filePath2, err)
		os.Exit(1)
	}

	lines1 := strings.Split(string(bytes1), "\n")
	lines2 := strings.Split(string(bytes2), "\n")

	config := typewriter.Config{
		Padding:   *padding,
		Separator: *separator,

		MarkFirstDifference: *markDifference,
		ShowLineNumbers:     *lineNumbers,
	}

	if *showHeader {
		config.LeftHeader = filePath1
		config.RightHeader = filePath2
	}

	s, err := typewriter.Sprint(lines1, lines2, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, `Encountered an unknown error: %v
		Please report this to https://github.com/j-delaney/typewriter/issues/new`, err)
		os.Exit(1)
	}

	fmt.Print(s)
}
