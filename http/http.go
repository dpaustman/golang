package http

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	resp, err := http.Get("http://www.deppon.com")
	if err != nil {
		panic(err)
	}

	s, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
