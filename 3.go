package main

import "fmt"

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Print("Masukan Bilangan : ")
		fmt.Scan(&i)
		if i <= 10000 {
			if i%2 == 0 {
				fmt.Print(i, " Adalah Bilangan Genap \n ")
			} else {
				fmt.Print(i, " Adalah Bilangan Ganjil \n ")
			}
		} else {
			fmt.Print("Penyataan selesai, karena angka sudah lebih dari 10. Terima kasih \n")
		}
	}
}