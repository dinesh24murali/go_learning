package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args1 := os.Args

	fmt.Println(args1)
	readFile(args1[1])
}

func readFile(fileName string) {
	file, _ := os.Open(fileName) // For read access.
	data := make([]byte, 1000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count: %d: %q\n", count, data[:count])
}
