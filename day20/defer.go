package main

import "fmt"

func f() {
	for i :=1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

func a() {
	i := 0
	fmt.Println(i)
	i++
	return
}

func main() {
	f()
	a()
}
