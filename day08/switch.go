package main

import "fmt"

func daemo(x int) int {
	switch {
	case 0 <= x && x <= 3:
		fmt.Println("0-3")
	case 4 <= x && x <= 6:
		fmt.Println("4-6")
	case 7 <= x && x <=9 :
		fmt.Println("7-9")
	}
	return x
}
func main() {
	daemo(5)
}
