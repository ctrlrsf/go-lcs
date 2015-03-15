package main

import "testing"

type LCSTest struct {
	first    string
	second   string
	length   int
	sequence string
}

func TestLCS(t *testing.T) {
	positives := []LCSTest{
		LCSTest{"abc", "abc", 3, "abc"},
		LCSTest{"onetwothreefour", "onetwothreefour", 15, "onetwothreefour"},
		LCSTest{"one", "six", 0, ""},
		LCSTest{"abc", "defg", 0, ""},
		LCSTest{"hijk", "lmn", 0, ""},
	}

	for i := 0; i < 2; i++ {
		test := positives[i]
		length, sequence := LCS([]rune(test.first), []rune(test.second))
		if length != test.length {
			t.Fatalf("bad length: %s, %s, expected=%d, actual=%d", test.first, test.second, test.length, length)
		}
		if test.sequence != string(sequence) {
			t.Fatalf("bad sequence: %s, %s, expected=%d, actual=%d",
				test.first, test.second, test.sequence, sequence)
		}
	}
}
