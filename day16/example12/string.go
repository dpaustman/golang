package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	str := "   hello world abc \t \n   "
	result := strings.Replace(str, "hello", "nihao", 1)
	fmt.Println("string.replace---", result)

	count := strings.Count(str, "o")
	fmt.Println("string.count---", count)

	repeat := strings.Repeat(str, 3)
	fmt.Println("string.repeate", repeat)

	upper := strings.ToUpper(str)
	fmt.Println("string.upper", upper)

	lower := strings.ToLower(upper)
	fmt.Println("string.lower", lower)

	space := strings.TrimSpace(str)
	fmt.Println("string.trimspace", space)

	left := strings.TrimLeft(str, "\n\r")
	fmt.Println("string.leftspace", left)

	right := strings.TrimRight(str, "\n\r")
	fmt.Println("string.right", right)

	trim := strings.Trim(str, "l")
	fmt.Println("string.trim", trim)

	splitresult := strings.Fields(str)
	for i := 0; i < len(splitresult); i++ {
		fmt.Println(splitresult[i])
	}

	splitResult := strings.Split(str, "o")
	for i :=0; i <len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str3 := strings.Join(splitResult, "o")
	fmt.Println("string.join", str3)

	str4 := strconv.Itoa(1000)
	fmt.Println("string.itoa", str4)

	number, err := strconv.Atoi(str4)
	if err != nil {
		fmt.Println("can't to convert to int", err)
		return
	} else {
		fmt.Println("string.Atoi", number)
	}
}
