package lcs

import "math"

// LCS returns the length and the longest common subsequence of two rune arrays
// http://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func LCS(str1, str2 []rune) (int, []rune) {
	var matrix [][]int
	var matrixSeq [][][]rune

	str1Len := len(str1)
	str2Len := len(str2)

	matrix = make([][]int, str1Len+1)
	matrixSeq = make([][][]rune, str1Len+1)

	for row := range matrix {
		matrix[row] = make([]int, str2Len+1)
	}

	for row := range matrixSeq {
		matrixSeq[row] = make([][]rune, str2Len+1)
	}

	// Loop is not needed as int and []rune are already default
	// to zero value.
	// Algorithm depends on first column of matrix being zero
	// values, so leaving here for clarity.
	for i := 0; i < str2Len+1; i++ {
		matrix[0][i] = 0
		matrixSeq[0][i] = make([]rune, 0)
	}

	// Loop is not needed as int and []rune are already default
	// to zero value.
	// Algorithm depends on first row of matrix being zero
	// values, so leaving here for clarity.
	for j := 0; j < str1Len+1; j++ {
		matrix[j][0] = 0
		matrixSeq[j][0] = make([]rune, 0)
	}

	for i := 1; i <= str1Len; i++ {
		for j := 1; j <= str2Len; j++ {
			if str1[i-1] == str2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
				matrixSeq[i][j] = append(matrixSeq[i-1][j-1], str1[i-1])
			} else {
				leftLength := matrix[i][j-1]
				topLength := matrix[i-1][j]
				matrix[i][j] = int(math.Max(float64(leftLength), float64(topLength)))
				if leftLength > topLength {
					matrixSeq[i][j] = append(matrixSeq[i][j-1])
				} else {
					matrixSeq[i][j] = append(matrixSeq[i-1][j])
				}
			}
		}
	}

	return matrix[str1Len][str2Len], matrixSeq[str1Len][str2Len]
}
