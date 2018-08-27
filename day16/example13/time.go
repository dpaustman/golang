package main

import (
	"time"
	"fmt"
)

func test() {
	time.Sleep(time.Millisecond*100)
}
func main() {
	now := time.Now()
	fmt.Println(now.Format("2018/08/27 16:56:05"))
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost: %s us\n", (end - start)/1000)
}
