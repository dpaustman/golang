package main

import "fmt"

func concat(a string, args...string) (result string) {
	result = a
	for i := 0; i < len(args); i++ {
		result += args[i]
	}
	return
}

func main() {
	res := concat("hello" + " " + "world")
	fmt.Println(res)
}
