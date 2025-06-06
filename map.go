package main

import "fmt"

type vertex struct{
	Lat, Long float32
}

func main(){
	// ages := map[int]int{}
	// ages[234]=234
	// delete(ages, 234)
	// ages[23]=78
	// v, status := ages["Udit"]
	// fmt.Println(v, status)

	// for i, v := range ages {
	// 	fmt.Println(i, v)
	// }

	m := make(map[string]vertex)
	fmt.Printf("%d %v\n", len(m), m)

	m["Abhishek"] = vertex{23.2342, 23432.23}
	fmt.Printf("%d %v\n", len(m), m)

	fmt.Println(m["Abhishek"])
}