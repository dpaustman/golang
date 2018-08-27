package main

import (
	"os"
	"fmt"
)

func main() {
	var goos string = os.Getenv("GOOS")
	goos := os.Getenv()
	fmt.Printf("The operating system is %s\n", goos)

	//var path string = os.Getenv("PATH")
	path := os.Getenv("PATH")
	fmt.Printf("The environment variable is %s\n", path)
}
