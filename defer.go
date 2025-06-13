package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	value := func(a, b int) (v int) {
		v = a * b
		return
	}(23, 34)

	fmt.Println("23 * 34 : ", value)
	defer fmt.Println(234)
	defer fmt.Println("I am Abhishek")
	fmt.Println(34534)

	createFile("docker.txt", "name.txt")
}

func createFile(srcName, dstName string) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		log.Fatal(err)
	}

	defer dst.Close()

	written, err := io.Copy(dst, src)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Written : ", written)
}
