package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	c := new(http.Client)
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error { return errors.New("not redirect") }

	for i := 0; i < 1000; i++ {
		res, _ := c.Get(fmt.Sprintf("http://git.io/%d", i))
		if res.StatusCode == 404 {
			fmt.Println(i)
		}
	}
}
