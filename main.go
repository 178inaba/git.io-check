package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	waitTime  = 3 * time.Second
	execDigit = 5
)

func main() {
	c := new(http.Client)
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error { return errors.New("not redirect") }

	runes := []rune{'0'}

	for {
		var uri string
		for _, v := range runes {
			uri += string(v)
		}

		resp, err := c.Get("https://git.io/" + uri)
		if err == nil {
			resp.Body.Close()
		}

		if resp.StatusCode == 404 {
			fmt.Println(uri + " OK!")
		} else {
			fmt.Fprintln(os.Stderr, uri+" NG")
			fmt.Fprintln(os.Stderr, "Location:", resp.Header.Get("Location"))
		}

		if runes[len(runes)-1] == 'Z' {
			// carry over
			addFlg := true
			for i := len(runes) - 1; i > -1; i-- {
				beforeRune := runes[i]
				runes[i] = getNextRune(runes[i])
				if beforeRune != 'Z' {
					addFlg = false
					break
				}
			}

			if addFlg {
				runes = append(runes, '0')
			}
		} else {
			runes[len(runes)-1] = getNextRune(runes[len(runes)-1])
		}

		if len(runes) > execDigit {
			// exit
			break
		}

		time.Sleep(waitTime)
	}
}

func getNextRune(r rune) rune {
	// 0 -> 9, a -> z, A -> Z
	if r == '9' {
		return 'a'
	} else if r == 'z' {
		return 'A'
	} else if r == 'Z' {
		return '0'
	}

	r++
	return r
}
