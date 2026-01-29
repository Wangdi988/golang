package main

import "fmt"

type FieldColl struct {
	x int
	y int
}


type User struct {
    ID     int
    Name   string
    Email  string
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

	u := User{
		ID:   1,
		Name: "Sangay",
		Email: "example@gmail.com",
	}

	fmt.Println(u)

	us := User{
		1,
		"Sangay  wangdi",
		"wangdi@gmail.com",
	}
	
	fmt.Println(us)
}
