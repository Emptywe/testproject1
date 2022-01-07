package main

import "testing"

func TestWholeStory(t *testing.T){

	seq := "1-Hello-2-World-3-Test-4-Run"
	expected := "Hello World Test Run"

	result := wholeStory(seq)

	if expected != result{
		t.Errorf("Incorrect whole story result. Expect: %v, got: %v", expected,result)
	}



}
