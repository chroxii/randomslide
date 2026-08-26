package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("RandomSlide starting...")
	if len(os.Args) != 2 {
		fmt.Println("Usage: randomslide <directory>")
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
}
