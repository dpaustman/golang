package main

import "fmt"

func reverse(str string) string {
	var result string
	for i :=0; i < len(str); i++ {
		result = result + fmt.Sprintf("%c", str[len(str) -i -1])
	}
	return result
}


func main() {
	str1 := "hello"
	str2 := "world!"
	str3 := str1 + " " + str2
	fmt.Println(str3)
	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)
	n := len(str3)
	fmt.Printf("len(str3)=%d\n", n)

	str5 := str3[3:]
	fmt.Println(str5)

	str6 := str3[3:6]

	fmt.Println(str6)
	str7 := reverse(str3)
	fmt.Printf(str7)


}
