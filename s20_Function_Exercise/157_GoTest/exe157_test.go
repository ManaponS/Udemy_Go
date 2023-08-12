package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(1, 5)
	if total != 10 {
		t.Errorf("Sum was in correct,got %d, want %d", total, 10)
	}
}
