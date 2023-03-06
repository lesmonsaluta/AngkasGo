package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// key is the string, the value is the int
	counts := make(map[string]int)

	// defining how input works
	input := bufio.NewScanner(os.Stdin)

	// while there is an input (input.Scan() returns True), the value
	// of input.Text() (that specific line that we are counting)
	// should be increased

	// the string is the key, adding another same string increases the value by 1
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	// range iterates over every element in counts
	// counts is (String (key), Int (value)), hence
	// line is the key, n is the value
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
