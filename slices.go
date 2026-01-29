package  main

import "fmt"

func main(){
	// with fixed array value
	evens := [6]int {2, 4, 6, 8, 10, 12}
	fmt.Println(evens[:4])
	fmt.Println(evens[1:3])
	fmt.Println(evens[3:])

	var s []int = evens[1:4]
	fmt.Println(s)

	evens[1] = 14
	fmt.Println(evens[1])


	// without setting the fixed array value
	odd := []int{1, 3, 5, 7}
	fmt.Println(odd)

	z := []struct {
		b bool
		i int
	}{
		{true, 3},
		{false, 7},
	}
	
	fmt.Println(z)

}