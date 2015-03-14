package main

import "testing"

type LCSLengthTest struct {
	first  string
	second string
	count  int
}

func TestLCSLength(t *testing.T) {
	positives := []LCSLengthTest{
		LCSLengthTest{"abc", "abc", 3},
		LCSLengthTest{"onetwothreefour", "onetwothreefour", 15},
		LCSLengthTest{"one", "six", 0},
		LCSLengthTest{"abc", "defg", 0},
		LCSLengthTest{"hijk", "lmn", 0},
	}

	for i := 0; i < 2; i++ {
		test := positives[i]
		lcs := LCSLength(test.first, test.second)
		if lcs != test.count {
			t.Fatalf("%s, %s, expected=%d, actual=%d", test.first, test.second, test.count, lcs)
		}
	}
}
