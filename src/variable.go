package main

import "fmt"

func main(){
	var name string
	name = "Muzammil"

	var address string = "Jakarta"
	var gender = "male"

	email := "email@email.com"

	var x, y, z int = 1, 2, 3
	var (
		firstName = "Muzammil"
		lastName = "nm"
	  )
	a, b := "Go", "Language"

	fmt.Println(name);
	fmt.Println(address)
	fmt.Println(gender)
	fmt.Println(email)
	fmt.Println(x, y, z)
	fmt.Println(firstName, lastName)
	fmt.Println(a, b)
}