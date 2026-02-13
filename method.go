package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

// Method
func (v Vertex) add() int {
	return v.x + v.y
}

// Function
func div(d Vertex) float64 {
	return float64(d.x) / float64(d.y)
}

// pointer reciver
func (v *Vertex) sub() int{
	return v.x - v.y;
}

func main() {
	v := Vertex{2, 9}
	fmt.Println(v.add()) // 11

	d := Vertex{8, 2}
	fmt.Println(div(d)) // 4

	x := &Vertex{20, 4}
	fmt.Println(x.sub())
}