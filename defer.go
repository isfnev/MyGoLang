package main

import "fmt"

func main(){
	value := func(a, b int) (v int){
		v = a*b
		return
	}(23, 34)
	
	fmt.Println("23 * 34 : ", value)
	defer fmt.Println(234)
	defer fmt.Println("I am Abhishek")
	fmt.Println(34534)
}