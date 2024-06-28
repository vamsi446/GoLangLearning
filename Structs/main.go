package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type contactInfo struct{
	email string
	zipcode int
}

type Person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	// vamsi := Person{"Vamsi ", "Revuri"}
	// vamsi := Person{firstName: "Vamsi", lastName: "Revuri"}
	var vamsi Person
	vamsi.firstName = "Vamsi"
	vamsi.lastName = "Revuri"
	fmt.Println(vamsi)

	john := Person{
		firstName: "John",
		lastName: "Doe",
		contact: contactInfo{
			email: "john@gmail.com",
			zipcode: 94000,
		},
	}
	john.print()
	john.updateName("Jim")
	john.print()

	johnPointer := &john
	johnPointer.updateName("Jim")
	john.print()


	fmt.Println("**********************************************")
	mySlice := []string{"Hi", "There", "How", "are", "you"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)

	///// !!!!!!Important
	/*
	pass by values ---> int, float, string, bool, structs
	pass by reference ---> slices, maps, channels, pointers, functions
	*/

	fmt.Println("*************************************************")
	resp,err := http.Get("http://google.com")
	if err!=nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	fmt.Println(resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)

	links:= []string{
		"http://www.google.com",
		"http://facebook.com",
	}
	
	for _,link:= range links {
		checkLink(link)
	}
}

func checkLink(link string){
	_,err:=http.Get(link)
	if(err !=nil ){
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println()
}


func updateSlice(mySlice []string) {
	mySlice[0] = "Bye"
}

func (pointer *Person) updateName(newFirstName string){
	pointer.firstName = newFirstName
	pointer.print()
	//this doesn't update globally for the person 
	//Go - pass by value
}

func (p Person) print(){
	fmt.Printf("%+v\n", p)
}

