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

	pathFromUser := os.Args[1]

	info, err := os.Stat(pathFromUser)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: path does not exist")
		} else {
			fmt.Println("An error occurred:", err)
		}
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Println("Error: path is not a directory")
		os.Exit(1)
	}
	fmt.Println("Media directory:", pathFromUser)
}
