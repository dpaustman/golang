package main

import "fmt"

func loop(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(loop(10))
	fmt.Println(loop(100))
}
