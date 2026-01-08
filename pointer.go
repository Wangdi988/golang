package main

import "fmt"

func main() {
	x := 5
	change_value(&x)
	fmt.Println(x)

	v := 9
	y := &v
	*y = 45

	fmt.Println(*y)

	a, b := 20, 40

	c := &a
	*c = 30

	fmt.Println(a, b)
}

func change_value(x *int) {
	*x = 7

}
