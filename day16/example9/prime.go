package main

import "fmt"

func prime(n int)  {
	j := 0
	for i :=1; i <=n ; i++ {
		if i % 2 != 0 {

			j++
			fmt.Println(i)
		}
	}
	fmt.Println(j)
}

func main() {
	prime(100)
}
