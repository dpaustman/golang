package main

import "fmt"

func sum()  {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	}

func main() {
	sum()
}