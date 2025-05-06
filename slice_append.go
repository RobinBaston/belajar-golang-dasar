package main

import "fmt"

func main() {

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	dayslice1 := days[5:] //Sabtu , Minggu
	fmt.Println(dayslice1)

	//merubah data array dari slice diatas
	dayslice1[0] = "Sabtu Baru"
	dayslice1[1] = "Minggu Baru"
	fmt.Println(dayslice1)
	fmt.Println(days) // caek array aslinya sudah berubah

	println("=== APPEND ===")
	dayslice2 := append(dayslice1, "Libur Baru")
	//anggapan bahwa membuat array baru
	//daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(dayslice1)
	fmt.Println(dayslice2)
	fmt.Println(days)
}
