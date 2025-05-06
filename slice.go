package main

import (
	"fmt"
)

func main() {

	names := [...]string{"ROBIN", "BASTIAN", "ARITONANG", "FLORENTINA", "CHRISTABEL", "LOUIS"}

	println("ambil data array dari array ke 4, dari 6 data")
	slices1 := names[4:6]
	fmt.Println(slices1)

	println("ambil data array pertama, dari 3 data")
	slice2 := names[:3]
	fmt.Println(slice2)

	println("ambil data array ke 3, dari semua data")
	slice3 := names[3:]
	fmt.Println(slice3)

	println("ambil data array pertama dari semua data")
	slice4 := names[:]
	fmt.Println(slice4)
}
