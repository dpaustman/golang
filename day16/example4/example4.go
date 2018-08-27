package main

import "fmt"

func modify(a int) {
	a = 10
	return
}

func modify2(a *int) {
	*a = 200
}

func main() {
	var a int = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a = ", a)
	fmt.Println("b =", b)

	modify(a)
	fmt.Println(a)

	modify2(&a)
	fmt.Println(a)
}
