package main

import "fmt"

func main() {
	anon1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Jansen",
		friends: map[string]int{
			"Jib":   24,
			"Mango": 35,
		},
		favDrinks: []string{
			"Mineral Water",
			"Est",
		},
	}

	for k, v := range anon1.friends {
		fmt.Println(anon1.first, "- friends - ", k, v)
	}

	for k, v := range anon1.favDrinks {
		fmt.Println(anon1.first, "- friends - ", k, v)
	}
}
