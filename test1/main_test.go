package main

import "testing"

func TestCalc(t *testing.T) {
	if calc(4, 4) != 8 {
		t.Errorf("wanted 8, got %d", calc(4, 4))
	}
}

func TestTableCalc(t *testing.T) {
	var tests = []struct {
		input1, input2, expected int
	} {
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
	}
	for caseCount, test := range tests {
		if res := calc(test.input1, test.input2); res != test.expected {
			t.Errorf("Test %d : Inputs (%d, %d) ==> Output (%d) != Expected (%d)", caseCount, test.input1, test.input2, res, test.expected)
		}
	}
}