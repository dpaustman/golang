package main

import (
	"fmt"
	"strconv"
)

func main() {
	 var str string
	 var result = 0
	 fmt.Scanf("%s", &str)
	 for i := 0; i < len(str); i++ {
	 	num := int(str[i] - '0')
	 	result += num*num*num
	 }

	 number, err := strconv.Atoi(str)
	 if err != nil {
	 	 fmt.Printf("can't convert %s to int", str)
		 return
	 }
	if result == number {
		fmt.Printf("%d is shuixianhuashu", number)
	} else {
		fmt.Printf("%d is not shuixianhuashu", number)
	}
}
