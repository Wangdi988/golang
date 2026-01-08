package main

import "fmt"

func main() {
	fmt.Println("for loop")

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	sum := 1
	fmt.Println("For another loop")
	for sum < 30 {
		sum += 1
	}
	fmt.Println(sum)

	fmt.Println("while loop for GO")
	// while loop for GO
	x := 1
	for x < 1000 {
		x += x
	}
	fmt.Println(x)
}
