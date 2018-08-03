package main

import (
		"fmt"
	"strconv"
	"log"
)

func convertToBin(n int) string {
	result := ""
	if n < 0 {
		log.Printf("Decimal to binary error: the argument must be greater than zero")
		return ""
	}
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {

		lsb := n % 2
		result = strconv.Itoa(lsb) + result

	}
	return result
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(788939),
		convertToBin(0),
		convertToBin(-3),
		)
}