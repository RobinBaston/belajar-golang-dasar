package main

import "fmt"

func main() {

	type NoKTP string

	var ktpRobin NoKTP = "1111111111"

	var contoh string = "2222222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpRobin)
	fmt.Println(contohKTP)

}
