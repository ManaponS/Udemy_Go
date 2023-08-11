package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "Blue",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Tundra",
		doors: 2,
		color: "White",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.electric, v1.make, v1.model, v1.doors, v1.color)
	fmt.Println(v2.electric, v2.make, v2.model, v2.doors, v2.color)

	fmt.Println(v1.engine.electric, v1.electric)
	fmt.Println(v2.engine.electric, v2.electric)
}
