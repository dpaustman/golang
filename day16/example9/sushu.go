package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var m int
	var n int
	fmt.Scanf("%d %d",&n, &m)
	fmt.Printf("%d %d\n", n, m)
	for j := n; j < m; j++ {
		if isPrime(j) == true {
			fmt.Printf("%d\n", j)
			continue
		}
	}
}
