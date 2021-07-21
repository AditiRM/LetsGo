package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	me := person{name: "Aditi", age: 20}
	fmt.Println(me)
} 
