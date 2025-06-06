package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 3; i++ {
		time.Sleep(600 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	go display("Abhishek")
	display("Rohit")
}
