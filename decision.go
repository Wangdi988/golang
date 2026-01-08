package main

import "fmt"

func main() {

	age := 18

	if age > 18 {
		fmt.Println("Age is greater then 18")
	} else {
		fmt.Println("Age is lesser or equal to 18")
	}

	ages := 18

	switch ages {
	case 16:
		fmt.Println("Your are under age")
	case 18:
		fmt.Println("Your are allowed to marrige")
	default:
		fmt.Println("Are you alive")
	}

}
