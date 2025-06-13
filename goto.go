package main

import "fmt"

func main() {
	var x int
	goto label
	x = 4
	fmt.Println(x)
label:
	x = 3
	fmt.Println(x)
}
