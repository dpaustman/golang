package main

import "fmt"

func demo(a int) int {
	if a < 5 {
		return 0
	} else {
		return 1
	}
}

func main() {
	fmt.Println(demo(4))
}
