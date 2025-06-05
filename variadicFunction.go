package main

import "fmt"

func main(){
	fmt.Println(sum(1,2,3))
	fmt.Println(sum(4,5))
	fmt.Println(sum())
}

func sum(nums ...int) (total int) {
	for _, v := range nums {
		total += v
	}
	return
}