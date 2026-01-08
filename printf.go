package main

import "fmt"

func main() {
	x := "Test"
	const pi = 3.14787654
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", 2.1)
	fmt.Printf("%T \n", 1)
	fmt.Printf("%T \n", true)
	fmt.Printf("%t \n", false)
	fmt.Printf("%b \n", 80)
	fmt.Printf("%.3f \n", pi)
}
