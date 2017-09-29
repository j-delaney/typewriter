# Typewriter

[![Build Status](https://travis-ci.org/j-delaney/typewriter.svg?branch=travis)](https://travis-ci.org/j-delaney/typewriter)

A library and command-line tool for displaying 2 columns of text.
The original purpose of this library is to improve failed test output when comparing two multiline strings.
Go from this:

![](/readme_imgs/before.png?raw=true "Before")

To this:
![](/readme_imgs/after.png?raw=true "After")

## Installing The CLI

If you just want the `side_by_side` command-line tool you should run this command:
```
go get -u github.com/j-delaney/typewriter/cmd/side_by_side
```

This will put the `side_by_side` tool in your `~/go/bin` directory.

## Usage & Example

First you'll want to get the library with

```
go get -u github.com/j-delaney/typewriter
```

Example of a function for printing two slices of strings side-by-side:

```go
func PrintSideBySide(s1, s2 []string) {
	var output string
	output = typewriter.Run(s1, s2, Config{
		ShowLineNumbers:     true,
		Separator:           "|",
		MarkFirstDifference: true,
	})

	fmt.Print(output)
}
```

This would print something in the form of:

```
1. abc |abcd
2. defg|ef
3. hi  |ghi
4. j   |j
```
