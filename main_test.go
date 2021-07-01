package main

import "testing"

func Test_addNumbers(t *testing.T) {
	result := 4
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
