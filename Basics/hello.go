package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		fmt.Println(index, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
