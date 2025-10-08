package main

import "fmt"

func main() {

	// Variable
	// var nama string = "Renaldy"
	// umur := 23
	// fmt.Println("Halo, nama saya", nama, "umur saya", umur, "tahun")

	// Tipe Data
	// var (
	// 	namaDepan string = "Renaldy"
	// 	namaBelakang string = "Pratama"
	// 	umur int = 23
	// 	tinggi float32 = 175.5
	// 	isMarried bool = false
	// )
	// fmt.Println("Nama saya", namaDepan, namaBelakang)
	// fmt.Println("Umur saya", umur, "tahun")
	// fmt.Println("Tinggi saya", tinggi, "cm")
	// fmt.Println("Sudah menikah?", isMarried)

	// Konstanta
	// const  phi = 3.14
	// fmt.Println("Konstanta phi =", phi)

	// Percabangan
	// nilai := 100
	// if nilai >= 80 {
	// 	fmt.Println("Selamat, Anda mendapatkan nilai A")
	// } else if nilai >= 70 {
	// 	fmt.Println("Anda mendapatkan nilai B")
	// } else if nilai >= 60 {
	// 	fmt.Println("Anda mendapatkan nilai C")
	// } else if nilai >= 50 {
	// 	fmt.Println("Anda mendapatkan nilai D")
	// } else {
	// 	fmt.Println("Anda mendapatkan nilai E")
	// }

	// switch {
	// case nilai >= 80:
	// 	fmt.Println("Selamat, Anda mendapatkan nilai A")
	// case nilai >= 70:
	// 	fmt.Println("Anda mendapatkan nilai B")
	// case nilai >= 60:
	// 	fmt.Println("Anda mendapatkan nilai C")
	// case nilai >= 50:
	// 	fmt.Println("Anda mendapatkan nilai D")
	// default:
	// 	fmt.Println("Anda mendapatkan nilai E")
	// }

	// Perulangan
	// Cara 1
	// for i := 1; i <= 10; i++{
	// 	fmt.Println("Perulangan ke", i)
	// }

	// Cara 2
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Perulangan ke", i)
	// 	i++
	// }

	for i := 1; i <= 10; i++{
		for j := i; j <= 10; j++{
			fmt.Print(j, "")
		}
		fmt.Println()
	} 

}
