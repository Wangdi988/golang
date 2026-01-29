package main

import "fmt"

func main(){
	var arr [3]string
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[0], arr[1])

	even := [4]int{2, 4, 6, 8}
	fmt.Println(even)
	fmt.Println(even[2], even[3])
	
}