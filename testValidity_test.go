package main

import (
	"github.com/Emptywe/testproject1/randomExamples"
	"testing"
)

func TestTestValidity(t *testing.T){
	seq1 := randomExamples.Generate(true)
	seq2 := randomExamples.Generate(false)

	expected1 := true
	expected2 := false

	result1 := testValidity(seq1)
	result2 := testValidity(seq2)

	if expected1 != result1{
		t.Errorf("Incorrect validation result. Expect: %v, got: %v", expected1,result1)
	}
	if expected2 != result2{
		t.Errorf("Incorrect validation result. Expect: %v, got: %v", expected2,result2)
	}

}
