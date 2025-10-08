package main

import (
	"fmt"
	"meating_2/lines"
)

func main() {

	// Array
	lines.Line("Bagian Array")
	exampleArray()

	// Slice
	lines.Line("Bagian Slice")
	exampleSlice()

	// Map
	lines.Line("Bagian Map")
	exampleMap()

	// Struct
	lines.Line("Bagian Struct")
	exampleStruct()

	// Pointer
	lines.Line("Bagian Pointer")
	examplePointer()

}

func exampleArray() {
	var names [3]string
	names[0] = "Renaldy"
	names[1] = "Galih"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println("Jumlah Array :", len(names))
}

func exampleSlice() {
	var buah = []string{"jeruk", "nanas", "pepaya"}
	fmt.Println(buah)
	fmt.Println("Jumlah Slice :", len(buah))
	fmt.Println("Kapasitas Slice :", cap(buah))
	buah = append(buah, "mangga")
	fmt.Println(buah)
	fmt.Println("Jumlah Slice :", len(buah))
	fmt.Println("Kapasitas Slice :", cap(buah))
}

func exampleMap() {
	var person = map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}
	fmt.Println(person)
	fmt.Println("Name :", person["name"])
	fmt.Println("Address :", person["address"])
	fmt.Println("Jumlah Map :", len(person))
	person["title"] = "Programmer"
	fmt.Println("New Map :", person)
	fmt.Println("Jumlah Map :", len(person))
}

func exampleStruct() {
	// Instansiasi awal dari struct
	type Person struct {
		nama              string
		umur              int
		status_pernikahan bool
	}

	// variabel bertipe struct
	var person1 Person
	person1.nama = "Eko"
	person1.umur = 30
	person1.status_pernikahan = false

	// menampilkan data struct
	fmt.Println("Nama :", person1.nama)
	fmt.Println("Umur :", person1.umur)
	fmt.Println("Status Pernikahan :", person1.status_pernikahan)
}

func examplePointer() {

	alamat1 := "Subang"
	alamat2 := &alamat1 // pointer
	alamat1 = "Bandung"
	fmt.Println("Alamat 1 :", alamat1)
	fmt.Println("Alamat 2 :", *alamat2) // dereference
}
