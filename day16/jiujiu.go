package main

import "fmt"

func jiujiu(n int) {
	for i := 1; i < n; i++ {
		for j :=1; j <= i; j++ {
			fmt.Printf("%d x %d = %d \t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	jiujiu(10)
}