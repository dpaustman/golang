package main

import "fmt"
func printSlice(s []int) {
	fmt.Printf("%v,len =%d,caps =%d\n", s, len(s), cap(s))
}
func main() {
	var s []int
	for i :=0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2 + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8, 9}
	s2 := make([]int, 16)
	s3 := make([]int, 16, 48)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
	fmt.Println("Deleteing slice")
	fmt.Println(s2[:3])

	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
	printSlice(s2[:3])
	printSlice(s2)
	fmt.Println("Poping slice from front")

	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping slice from tail")
	tail := s2[len(s2) -1]
	s2 = s2[:len(s2) -1]
	fmt.Println(tail)
	printSlice(s2)
}
