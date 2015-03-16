package main

import (
	"fmt"
	"github.com/ctrlrsf/go-lcs"
	"os"
)

func printSize(s string) {
	fmt.Printf("\"%v\" length = %d\n", s, len(s))
}

func usage() {
	fmt.Printf("Usage: %s string1 string2\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args[1:]) < 2 {
		usage()
	}
	str1, str2 := os.Args[1], os.Args[2]

	length, sequence := lcs.LCS([]rune(str1), []rune(str2))
	fmt.Printf("Longest sequence: %s\n", string(sequence))
	fmt.Printf("Length: %d\n", length)
}
