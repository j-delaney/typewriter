# Typewriter

[![Build Status](https://travis-ci.org/j-delaney/typewriter.svg?branch=travis)](https://travis-ci.org/j-delaney/typewriter)

A library and command-line tool for displaying 2 columns of text.
The original purpose of this library is to improve failed test output when comparing two multiline strings.
Go from this:

![](/readme_imgs/before.png?raw=true "Before")

To this:
![](/readme_imgs/after.png?raw=true "After")

## Library Usage & Example

First you'll want to get the library with

```
go get -u github.com/j-delaney/typewriter
```

Example of a function for printing two slices of strings side-by-side:

```go
func PrintSideBySide(s1, s2 []string) {
  var output string
  var err error

  output, err = typewriter.Sprint(s1, s2, Config{
    ShowLineNumbers:     true,
    Separator:           "|",
    MarkFirstDifference: true,
  })

  if err != nil {
    panic(err)    	  
  }
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

## Command-Line Installation and Usage

If you just want the `side_by_side` command-line tool you should run this command:

```
go get -u github.com/j-delaney/typewriter/cmd/side_by_side
```

This will put the `side_by_side` tool in your `~/go/bin` directory.

If `~/go/bin` is in your `PATH` then you can just run `side_by_side` from your shell. Otherwise you'll need to run `~/go/bin/side_by_side`.

It expects two file paths as arguments. For example, if you want to see foo.txt and bar.txt side-by-side run `side_by_side path/to/foo.txt path/to/bar.txt`.

Run `side_by_side -h` to see all options.