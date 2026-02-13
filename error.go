package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
