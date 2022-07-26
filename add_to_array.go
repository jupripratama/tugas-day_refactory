package main

import "fmt"

func main() {
	my_slice := []string{
		"Meja", "Buku", "Topi", "Baju", "Kayu",
	}
	second_slice := []string{"Celana"}
	my_slice = append(my_slice, second_slice...)
	fmt.Println("Slice: ", my_slice)

	depan_slice2 := []string{"Handuk"}
	my_slice = append(depan_slice2, my_slice...)
	fmt.Println("Slice: ", my_slice)
}
