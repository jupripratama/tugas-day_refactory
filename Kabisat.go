package main

import "fmt"

//Jupri Eka Pratama
func main() {
	var tahun = 0
	fmt.Print("Masukan Tahun : ")
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
