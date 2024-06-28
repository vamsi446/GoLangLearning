package main

import (
"fmt" 
"strconv"
)

func main() {
	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr2)

	var cars = [...]string{"Tesla", "mercedes"}
	fmt.Println(cars)

	fmt.Println("cars[0] : ", cars[0])

	//initialize only specific elements
	arr3 := [5]int{1: 10, 2: 29}
	fmt.Println(arr3)

	fmt.Println(len(arr3))

	//slices are similar to arrays but more powerful
	//slices can grow and Strink
	mySlice := []int{}
	fmt.Println("len of myslice", len(mySlice))
	fmt.Println("cap of mySlice", cap(mySlice))

	//slice from an array
	mySlice2 := arr1[1:3]
	fmt.Println(mySlice2)
	mySlice3 := make([]int, 5, 10)
	fmt.Println(mySlice3)

	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println("element at index 2 of slice prices : ", prices[2])
	fmt.Println("cap and len of prices slice : ", cap(prices), len(prices))
	prices = append(prices, 60, 70, 80, 90)
	fmt.Println("updated prices slice : ", prices)
	fmt.Println("cap and len of prices slice : ", cap(prices), len(prices))

	pricesCopy := make([]int, len(prices))
	copy(pricesCopy, prices)

	fmt.Println(pricesCopy)

	conditions()
	loops()
	fmt.Println(sayHello("vamsi", 24))

	fmt.Println(" sum of 232 and 23232 : ", getSum(232, 23232) )


	var pers1 Person
	pers1.name = "vamsi"
	pers1.age = 24
	pers1.job = "unemployed"
	pers1.salary = 0

	fmt.Println("person details : ",pers1.name, pers1.age, pers1.job, pers1.salary);


	practiceMaps()

}

func conditions() {
	a, b := 1, 2
	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a equals b")
	} else {
		fmt.Println("a is less than b")
	}

	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

}

func loops() {
	fmt.Println("iterating through i ")
	for i := 0; i < 5; i++ {
		fmt.Println("i : ", i)
	}

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	//Range:
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}




	//format verbs
	/*
		%v	the value in a default format
		when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value
		%t	the word true or false
		%b	base 2
		%c	the character represented by the corresponding Unicode code point
		%d	base 10
		%o	base 8
		%O	base 8 with 0o prefix
		%q	a single-quoted character literal safely escaped with Go syntax.
		%x	base 16, with lower-case letters for a-f
		%X	base 16, with upper-case letters for A-F
		%U	Unicode format: U+1234; same as "U+%04X"
	*/
}

func sayHello(name string, age int) string{
	return "Hello: "+ name+","+strconv.Itoa(age) +" year old"
}

func getSum(x int, y int) (result int){
	result = x+y
	return 
}



//Structures
type Person struct{
	name string
	age int
	job string
	salary int

}

func practiceMaps(){
	var a = map[string]string{"brand": "Ford", "model": "Mustang"}
	fmt.Println(a)

	b := make(map[string]int)
	b["vamsi"] = 24;
	fmt.Println(b);

	delete(b,"vamsi")

	//check for existing key
	var cars= map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}
	val1, isPresent1 := cars["brand"]
	fmt.Println(val1, isPresent1)
	val2, _ := cars["model"]
	fmt.Println(val2)

	for k,v :=range cars{
		fmt.Println(k,v)
	}


}


///iterate over maps in specific order
func iterateMapsInSpecificOrder(){
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string             // defining the order
	b = append(b, "one", "two", "three", "four")
  
	for k, v := range a {        // loop with no order
	  fmt.Printf("%v : %v, ", k, v)
	}
  
	fmt.Println()
  
	for _, element := range b {  // loop with the defined order
	  fmt.Printf("%v : %v, ", element, a[element])
	}
}