package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Eric","Last":"Blunc","Age":45}]`

	bs := []byte(s)
	fmt.Printf("%T %v \n", s, s)
	fmt.Printf("%T %v \n", bs, bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------Unmarshal-----")
	fmt.Println(people)
	for i, v := range people {
		fmt.Printf("Index %v = %v \n", i, v)
	}

}
