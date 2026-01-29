package main

import "fmt"


type vetex struct {
	age, class int
}

func main(){
	ages := make(map[string]int)
	ages["Sangay"] = 30
	ages["Karma"] = 25
	fmt.Println(ages)

	class := map[string]int{
		"sonam" : 10,
		"wangdi" :12,
	}

	fmt.Println(class)
	// print the value by key 
	fmt.Println(ages["Sangay"])
	fmt.Println(class["sonam"])

	// using struct 

	m := make(map[string]vetex)
	m["age_class"] = vetex{
		12, 6,
	}
	fmt.Println(m["age_class"])

	o := map[string]vetex{
		"sangay" : vetex{
			14, 8,
		},
		"wangdi": vetex{
			19, 12,
		},
	}

	fmt.Println(o)


	// form the different 
	fmt.Println("Answere values")
	n := make(map[string]int)

	n["Answer"] = 42
	fmt.Println("The value:", n["Answer"])

	n["Answer"] = 48
	fmt.Println("The value:", n["Answer"])

	delete(n, "Answer")
	fmt.Println("The value:", n["Answer"])

	v, ok := n["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}