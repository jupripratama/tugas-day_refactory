package main

import "fmt"

func main() {
	var tahun = 0
	fmt.Print("Masukan Bilangan : ")
	fmt.Scan(&tahun)
	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				fmt.Print("Tahun Kabisat \n")
			} else {
				fmt.Print("Bukan Tahun Kabisaat \n")
			}
		} else {
			fmt.Print("Tahun Kabisat \n")
		}
	} else {
		fmt.Print("Bukan Tahun Kabisat \n")
	}
}
