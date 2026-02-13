package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b;
}

func sub(a, b int) (int, int) {
	return a - b, a + b;
}

func main(){
	addtion := add(2, 6)
	fmt.Println(addtion)

	v, r:= sub(8,2)
	fmt.Println(v, r)

}