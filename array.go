package main

import "fmt"

func main() {

	println("deklarasi array dasar")
	var name [3]string

	name[0] = "ROBIN"
	name[1] = "BASTIAN"
	name[2] = "ARITONANG"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	println("membuat array langsung saat deklarasi variable")
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	println("cara mencari panjang array")
	var value = [...]int{

		90, 80, 70, 60, 100,
	}

	fmt.Println(len(value))
}
