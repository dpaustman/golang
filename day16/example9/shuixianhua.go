package main

import "fmt"

func isNumber(m int) bool {
	k := m/100
	v := m/10%10
	w := m%10
	return k*k*k + v*v*v + w*w*w == m
}

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	//fmt.Printf("%d %d", m, n)
	for i := m; i < n; i++ {
		if isNumber(i) == true {
			fmt.Println(i)
		}
	}
}
