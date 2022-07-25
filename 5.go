package main

import "fmt"

func main() {
	// var usia = 0
	// fmt.Print("Masukan Bilangan : ")
	// fmt.Scan(&usia)
	// if usia >= 21 {
	// 	fmt.Print("Ini FIlM Untuk DEWASA. \n ")
	// } else if usia >= 13 {
	// 	fmt.Print("Ini FIlM Untuk REMAJA. \n ")
	// } else if usia >= 9 {
	// 	fmt.Print("Ini FIlM Untuk BIMBINGAN ORANG TUA. \n ")
	// } else {
	// 	fmt.Print("Ini FIlM Untuk SEMUA USIA. \n ")
	// }

	//Jupri Eka Pratama

	for usia := 0; usia < 100; usia++ {
		fmt.Print("Masukan Usia Anda : ")
		fmt.Scan(&usia)
		if usia <= 100 {
			if usia >= 21 {
				fmt.Print("FIlM ini Untuk DEWASA. \n")
			} else if usia >= 13 {
				fmt.Print("FIlM ini Untuk REMAJA. \n ")
			} else if usia >= 9 {
				fmt.Print("FIlM ini Untuk BIMBINGAN ORANG TUA. \n ")
			} else {
				fmt.Print("FIlM ini Untuk SEMUA USIA. \n ")
			}
		}
	}

}
