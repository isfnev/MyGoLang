package main

import "fmt"

func main() {
	// var sc = make([]string, 4, 5)
	// fmt.Println(len(sc), cap(sc))

	// sc = append(sc, "Abhishek")
	// for i, v:= range sc{
	// 	fmt.Println(i, v)
	// }
	// fmt.Println()

	var s []int
	print_slice(s)

	var t int
	fmt.Print("The number of appends : ")
	fmt.Scan(&t)

	for i:=0;i<t;i++{

		var numberOfElement int
		fmt.Print("The number of elements in this append : ")
		fmt.Scan(&numberOfElement)
		var b = make([]int, numberOfElement)

		fmt.Printf("Enter %d number of element to insert : ", numberOfElement)
		for j:=0;j<numberOfElement;j++{
			fmt.Scan(&b[j])
		}
		s = append(s, b...)
		print_slice(s)
	}
}

func print_slice(a []int) {
	fmt.Printf("len=%d, cap=%d, s=%v\n", len(a), cap(a), a)
}
