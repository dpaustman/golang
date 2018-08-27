package main

import "fmt"

func shuixian(n int) {
	for i :=100; i <= n; i++ {
		k := i/100
		v := i/10%10
		w := i%10
		if (k*k*k + v*v*v + w*w*w) == i {
			fmt.Println(i)
		}
	}
}

func main() {
	shuixian(999)
}
