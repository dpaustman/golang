package main

import "fmt"

func add(a int, args...int) int {
	var sum int = a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println(add(10))
	fmt.Println(add(1, 3, 5, 7, 9))
}
