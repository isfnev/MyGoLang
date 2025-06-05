package main

import "fmt"

func main() {
	// var a uint8 = 256   Gives Error
	var a uint8 = 255
	fmt.Printf("%T\n", a)
	fmt.Printf("%d\n", a+1)// if a is 255 as the range of uint is 0 to 255, adding 1 turn it to 0

	var b int8 = 127
	fmt.Println(b+127)
	fmt.Println(a+b) // adding both the types with same types doesn't give error
}
