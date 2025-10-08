package main

import "fmt"

func main() {
	// 1. Operator Aritmatika
	numberOne()
	garis()

	// 2.Operator Relasi dan Logika
	numberTwo()
	garis()

	// 3. Perulangan for
	numberThree()
	garis()

	// 4. Perulangan dan Operator
	numberFour()
	garis()

	// 5. Array dan Perulangan
	numberFive()
}

func garis() {
	fmt.Println("===================================")
}

func numberOne() {
	// Deklarasi variabel
	nilai1 := 10
	nilai2 := 3

	// Menampilkan nilai awal
	fmt.Printf("Masukkan angka pertama: %d\n", nilai1)
	fmt.Printf("Masukkan angka kedua: %d\n", nilai2)
	// Penjumlahan
	hasilTambah := nilai1 + nilai2
	fmt.Printf("Hasil Penjumlahan: %d + %d = %d\n", nilai1, nilai2, hasilTambah)
	// Pengurangan
	hasilKurang := nilai1 - nilai2
	fmt.Printf("Hasil Pengurangan: %d - %d = %d\n", nilai1, nilai2, hasilKurang)
	// Perkalian
	hasilKali := nilai1 * nilai2
	fmt.Printf("Hasil Perkalian: %d * %d = %d\n", nilai1, nilai2, hasilKali)
	// Pembagian
	hasilBagi := nilai1 / nilai2
	fmt.Printf("Hasil Pembagian: %d / %d = %d\n", nilai1, nilai2, hasilBagi)
	// Sisa Bagi
	hasilMod := nilai1 % nilai2
	fmt.Printf("Hasil Sisa Bagi: %d %% %d = %d\n", nilai1, nilai2, hasilMod)
}

func numberTwo() {
	// Deklarasi variabel
	umur := 80;
	fmt.Printf("Masukkan umur Anda: %d\n", umur)
	if umur >= 0 && umur <= 12 {
		fmt.Println("Anda masih anak-anak.")
	} else if umur >= 13 && umur <= 19 {
		fmt.Println("Anda sudah remaja.")
	} else if umur >= 20 && umur <= 59 {
		fmt.Println("Anda sudah dewasa.")
	} else{
		fmt.Println("Anda sudah lansia.")
	} 
}

func numberThree(){
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func numberFour(){
	n := 10
	total := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	fmt.Println("Masukan nilai N:", n)
	fmt.Printf("Jumlah bilangan genap dari 1 sampai %d adalah: %d\n", n, total)
}

func numberFive(){
	Teman := [5]string{"Budi", "Siti", "Andi", "Lina", "Dika"}
	for index, nama := range Teman {
		fmt.Printf("Teman ke-%d: %s\n", index+1, nama)
	}
}
