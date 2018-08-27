package main

import "fmt"

func main() {
	str1 := "hello world!"
	str2 := `
	床前明月光
	疑是地上霜
	举头望明月
	我是郭德纲`

	b := 'c'
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Printf("%c\n", b)
	fmt.Println(b)
}
