package main

import (
	"fmt"
	"os"
)

func main() {
	args1 := os.Args

	fmt.Println(args1[1])
	fmt.Println(args1)
}
