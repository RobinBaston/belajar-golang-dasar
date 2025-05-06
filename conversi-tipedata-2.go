package main

import "fmt"

func main() {

	var name = "Robin Bastian"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
