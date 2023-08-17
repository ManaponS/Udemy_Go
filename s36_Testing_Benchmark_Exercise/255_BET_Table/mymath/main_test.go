package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}
	tests := []test{
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{2, 4, 6, 8, 10, 12}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	//output:
	//4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}
