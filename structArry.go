package main

import "fmt"

type person struct {
	height float32
	name   string
	age    int
}

func main() {
	arr := [3]person{
		{height: 234.234, name: "absdkj", age: 23},
		{height: 24.4, name: "abkj", age: 24},
		{height: 2.34, name: "sdkj", age: 34},
	}

	fmt.Print(arr)
}
