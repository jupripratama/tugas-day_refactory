package main

import "fmt"

type People struct {
	id   int
	name string
	age  int
}

func main() {
	people := []People{
		{1, "Udin", 12},
		{2, "Wati", 51},
		{3, "Budi", 34},
		{4, "Agus", 16},
		{5, "Sari", 19},
		{6, "Ririn", 21},
	}

	var mapData []People
	for _, v := range people {
		if v.age < 21 {
			mapData = append(mapData, v)
		}
	}
	fmt.Print(mapData)
}
