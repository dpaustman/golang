package main

import "fmt"

func grade(score int) string {
	g := ""
	switch  {
	case score < 0 || score > 100:
		fmt.Printf("wrong score: %d", score)
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	fmt.Println(grade(110))
	fmt.Println(grade(50))
	fmt.Println(grade(61))
	fmt.Println(grade(83))
	fmt.Println(grade(99))

}
