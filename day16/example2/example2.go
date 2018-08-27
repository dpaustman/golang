package main

import (
	"time"
	"fmt"
)

const (
	man = 1
	woman = 2
)

func test() {
	for {
		second := time.Now().Unix()
		if (second%woman == 0) {
			fmt.Println("woman")
		} else {
			fmt.Println("man")
		}
	}
	time.Sleep(time.Second)
}

func main() {
	test()
}

