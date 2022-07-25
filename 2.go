package main

import "fmt"

func main() {
	// var nilai = 85

	// if nilai >= 90 {
	// 	fmt.Println("A")
	// } else if nilai >= 80 {
	// 	fmt.Println("B")
	// } else if nilai >= 70 {
	// 	fmt.Println("C")
	// } else if nilai >= 60 {
	// 	fmt.Println("D")
	// } else {
	// 	fmt.Println("E")
	// }
	//Jupri Eka Pratama

	for nilai := 0; nilai < 100; nilai++ {
		fmt.Print("Masukan Nilai : ")
		fmt.Scan(&nilai)
		if nilai >= 90 {
			fmt.Print(nilai, " Nilainya A \n")
		} else if nilai >= 80 {
			fmt.Print(nilai, " Nilainya B \n")
		} else if nilai >= 70 {
			fmt.Print(nilai, " Nilainya C \n")
		} else if nilai >= 60 {
			fmt.Print(nilai, " Nilainya D \n")
		} else {
			fmt.Print(nilai, " Nilainya E \n")
		}
	}
}
