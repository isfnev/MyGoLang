package main

import "fmt"

type person struct{
	name string
	age int
}

type integer int

type book struct {
	name string
}

func main(){
	a := person{name:"Abhishek", age:234}
	a.display()

	b := integer(234)
	fmt.Println(b.square())

	c := book{"Rajat"}
	fmt.Println("Cname : ", c.name)
	c.changeName("Vivek")
	fmt.Println("Cname : ", c.name)
}

func (b *book) changeName(s string){
	b.name=s
}

func (p person) display(){
	fmt.Printf("Name : %v\n", p.name)
	fmt.Printf("Age : %v\n", p.age)
}

func (i integer) square() integer{
	return i*i
}
