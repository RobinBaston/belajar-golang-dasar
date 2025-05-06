package main

import "fmt"

func main() {

	//var name string

	name := "ROBIN BASTIAN" // : untuk ganti var hanya bisa digunakan dipertama saja
	fmt.Println(name)

	//var name = "Robin Bastian"
	//fmt.Println(name)

	name = "Robin Bastian"
	fmt.Println(name)

	name = "Robin Baston"
	fmt.Println(name)

	fmt.Println("MULTIPLE VARIABLE")

	var (
		firstName  = "Robin"
		middleName = "Bastian"
		lastName   = "Aritonang"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
