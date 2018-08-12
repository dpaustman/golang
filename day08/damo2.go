package main

import "fmt"

func daemo(x int) int {
	switch x {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fallthrough
	case 5:
		fmt.Println(5)
	}
	return x
}

func main() {
	daemo(3)
	daemo(1)
	daemo(2)
	daemo(4)
}
