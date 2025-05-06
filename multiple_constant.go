package main

import "fmt"

func main() {

	const ( //constant tidak boleh berubah, variable boleh berubah
		firstName = "Robin"
		lastName  = "Bastian"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
