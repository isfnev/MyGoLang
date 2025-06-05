package main

import "fmt"

func main(){
	ages := map[int]int{}
	ages[234]=234
	delete(ages, 234)
	ages[23]=78
	// v, status := ages["Udit"]
	// fmt.Println(v, status)

	for i, v := range ages {
		fmt.Println(i, v)
	}
}