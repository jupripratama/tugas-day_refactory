package main

import "fmt"

func NChan(start, end int) chan int {
	outchan := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			outchan <- i
		}
		close(outchan)
	}()
	return outchan
}

func main() {
	for v := range NChan(4, 8) {
		fmt.Print(v)
	}
}
