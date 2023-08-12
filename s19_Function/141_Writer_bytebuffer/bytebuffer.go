package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gopher")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("Reset")
	fmt.Println(b.String())

	b.Write([]byte("Byte"))
	fmt.Println(b.String())
}
