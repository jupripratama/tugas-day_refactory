package main

import "fmt"

func main() {
	var nilai = 85

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 80 {
		fmt.Println("B")
	} else if nilai >= 70 {
		fmt.Println("C")
	} else if nilai >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
