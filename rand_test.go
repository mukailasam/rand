package main

import "testing"

func TestRand(t *testing.T) {
	input := 16
	leastexpectedOutput := 16
	mostExpectedout := 100000

	r := &RandStrct{}
	r.NumberofBytes = input

	rval, _ := RandomBytes(r)
	if rval < leastexpectedOutput {
		t.Fatal("Testing failed...")
	} else if rval > mostExpectedout {
		t.Fatal("Testing failed...")
	}

}
