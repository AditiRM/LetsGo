package main

import "fmt"

func main() {
	var a = 1;
	fmt.Println(a);

	// The below is not allowed as we already used 'a' as a int
	// a = "not allowrd"
	// Println(a)

	// in go we don't need to mention the type as shown above, though we can link
	var test string = "hello"
	fmt.Println(test)

	// The simplest for declaring as well as initialising is using ":=" this notaion
	simple := "How to teach your dog QUANTUM PHYSICS!!"
	fmt.Println(simple);

	// if not initalized corressponds to its zero values 
	var x int
	var y string
	fmt.Println(x) // 0
	fmt.Println("Emptinessss ", y) // ""
}
