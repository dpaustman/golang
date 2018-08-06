package main

import (
	"fmt"
	"retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.deppon.com")

}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake deppon"}
	fmt.Println(download(r))
}
