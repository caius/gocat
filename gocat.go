package main

import (
	"fmt"
	"io"
	"os"
)

var ARGV []string

func main() {
	ARGV = os.Args[1:]

	// If there are no arguments, then we read STDIN
	if len(ARGV) == 0 {
		ARGV = append(ARGV, "-")
	}

	// Iterate arguments to read filenames
	for _, filename := range ARGV {
		// - means read stdin as a special case, we just use /dev/stdin instead
		if filename == "-" {
			filename = "/dev/stdin"
		}

		// Time to try and read the file in question
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s: No such file or directory\n", filename)
			continue
		}

		// Copy our output across!
		io.Copy(os.Stdout, f)
	}

}
