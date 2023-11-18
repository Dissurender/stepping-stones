package main

import (
	"fmt"
	"os"
)

func main() {
	props := os.Args[1:]

	fmt.Println(props)

	os.Exit(0)
}
