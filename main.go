package main

import (
	"fmt"
	"math"
	"os"
)

func printSize(s string) {
	fmt.Printf("\"%v\" length = %d\n", s, len(s))
}

func usage() {
	fmt.Printf("Usage: %s string1 string2", os.Args[0])
	os.Exit(1)
}

var str1, str2 string
var matrix [][]int

func main() {
	if len(os.Args[1:]) < 2 {
		usage()
	}
	str1, str2 = os.Args[1], os.Args[2]

	lcsLength := LCSLength(str1, str2)
	fmt.Printf("%d\n", lcsLength)
}

// LCSLength returns the length of the longest common subsequence of two strings
// http://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func LCSLength(str1, str2 string) int {
	str1Len := len(str1)
	str2Len := len(str2)

	matrix = make([][]int, str1Len+1)

	for row := range matrix {
		matrix[row] = make([]int, str2Len+1)
	}

	for i := 0; i < str2Len+1; i++ {
		matrix[0][i] = 0
	}

	for j := 0; j < str1Len+1; j++ {
		matrix[j][0] = 0
	}

	for i := 1; i <= str1Len; i++ {
		for j := 1; j <= str2Len; j++ {
			if str1[i-1] == str2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = int(math.Max(float64(matrix[i][j-1]), float64(matrix[i-1][j])))
			}
		}
	}

	return matrix[str1Len][str2Len]
}
