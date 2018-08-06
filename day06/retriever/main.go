package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.deppon.com")

}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake deppon"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
