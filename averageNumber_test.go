package main

import "testing"

func TestAverageNumber(t *testing.T) {

	seq := "1-Hello-2-World-3-Test-4-Run"
	expected := 2

	result := averageNumber(seq)

	if expected != result {
		t.Errorf("Incorrect average number result. Expect: %v, got: %v", expected, result)
	}

}
