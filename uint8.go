package main

import "fmt"

func main(){
	var a int8 = 127
	fmt.Println(a+1)//-128

	var b [3]int = [3]int{2,3,4}
	var c = b[:]

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}