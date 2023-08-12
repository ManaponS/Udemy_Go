package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect got: %d, want: %d.", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(5, 4)
	if total != 1 {
		t.Errorf("Sum was incorrect got: %d, want: %d.", total, 1)
	}
}

func TestDoMathAdd(t *testing.T) {
	totalAdd := doMath(5, 7, add)
	if totalAdd != 12 {
		log.Fatalf("Sum was incorrect got: %d, want: %d.", totalAdd, 12)
	}
}
