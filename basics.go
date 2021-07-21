package main

import "fmt"

func main() {
	//Looping 
	for i := 1; i <= 10; i++{
		fmt.Println(i)
	}

	//if-else
	x := 3;
	if (x == 1){
		fmt.Println("Inside if")
	}else if (x == 2){
		//SPECIAL NOTE : else if should be in the same line as in ending if brace!
		fmt.Println("Inside else if")
	}else {
		fmt.Println("Inside else")
	}

	//Arrays
	var test[3] int
	test[1] = 73
	fmt.Println(len(test))
	fmt.Println(test[0])	//0

	//SLICES : Arrays + more features
	slice1 := make([]int, 3)
	slice1[0] = 3
	fmt.Println(slice1[0]) //3
	//append
	slice2 := append(slice1, 4)
	fmt.Println(slice2[3]) //4
	//copy
	slice3 := make([]int, 4)
	copy(slice3, slice2)
	fmt.Println(slice3[3]) //4
	fmt.Println(slice3)

	//MAPS
	map1 := make(map[string]int)
	map2 := map[string]int{"Holaa" : 1, "Hii": 2}

	map1["Holaa"] = 1
	map1["Hii"]  = 2

	fmt.Println(map2) //map[Holaa:1 Hii:2]

	delete(map1, "Holaa")
	fmt.Println(map1) //map[Hii:2]

	//Range : Iterate over arrays & maps
	nums := [3]int{1, 2, 3}
	for ind, value := range nums {
		fmt.Println("index", ind, "value", value)
	}

	for key, value := range map2 {
		fmt.Println("key", key, "value", value)
	}

}
