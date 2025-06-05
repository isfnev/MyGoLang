package main

import "fmt"

func main(){
	var sc = make([]string, 5)
	fmt.Println(len(sc), cap(sc))

	sc = append(sc, "Abhishek")
	for i, v:= range sc{
		fmt.Println(i, v)
	}
	fmt.Println()
}