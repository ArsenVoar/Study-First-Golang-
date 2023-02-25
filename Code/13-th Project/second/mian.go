package main

import (
	"context"
	"fmt"
)

func main() {
	Context()
}

func Context() {
	ctx := context.Background()
	fmt.Println(ctx)

	ToDo := context.TODO()
	fmt.Println(ToDo)
}
