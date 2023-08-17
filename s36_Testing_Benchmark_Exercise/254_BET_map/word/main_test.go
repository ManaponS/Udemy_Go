package word

import (
	"fmt"
	"testing"

	"github.com/ManaponS/Udemy_Go/s36_Testing_Benchmark_Exercise/254_BET_map/quote"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Expected", 3, "Got", n)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Expected", k, "=", 1, "Got", v)
			}
		case "two":
			if v != 2 {
				t.Error("Expected", k, "=", 2, "Got", v)
			}
		case "three":
			if v != 3 {
				t.Error("Expected", k, "=", 3, "Got", v)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
