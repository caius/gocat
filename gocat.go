package main

import (
	"fmt"
	"io"
	"os"
)

var ARGV []string

func main() {
	ARGV = os.Args[1:]

	if len(ARGV) >= 1 {
		for _, filename := range ARGV {
			if filename == "--" {
				continue
			}

			// - means read stdin as a special case
			if filename == "-" {
				filename = "/dev/stdin"
			}

			// Otherwise we're after a file itself
			f, err := os.Open(filename)
			if err != nil {
				fmt.Printf("gocat: %s: No such file or directory\n", filename)
				continue
			}

			// Copy our output across!
			io.Copy(os.Stdout, f)
		}
	}
}
