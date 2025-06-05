package main

import "fmt"

func main(){
	x := 5
	y := 10

	fmt.Printf("Before: x = %d and y = %d\n", x, y)
	multiply(&x, &y)
	fmt.Printf("After: x = %d and y = %d\n", x, y)
}

func multiply(a *int, b* int) {	
	*a *= *b
}