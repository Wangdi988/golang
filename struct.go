package main

import "fmt"

type FieldColl struct {
	x int
	y int
}

func main() {
	fmt.Println(FieldColl{2, 5})
	b := FieldColl{20, 30}
	b.x = 10
	fmt.Println(b.x)

	d := &b
	d.x = 11

	fmt.Println(b)
	fmt.Println(*d)
}
