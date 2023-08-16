package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx, _ = context.WithCancel(ctx)
	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context Type:\t%T\n", ctx)
}
