package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	waitTime  = 3 * time.Second
	execDigit = 5
	baseURL   = "https://git.io/"
	okFmt     = "%s OK!\n"
	ngFmt     = "%s NG!\n"
	locFmt    = "Location: %s\n"
)

var (
	n = kingpin.Flag("dry-run", "dry run mode.").Short('n').Bool()
	c = new(http.Client)
)

func init() {
	kingpin.Parse()
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error { return errors.New("not redirect") }
}

func main() {
	runes := []rune{'0'}

	for {
		uri := string(runes)

		checkURI(uri)

		runes = advanceRunes(runes)

		if len(runes) > execDigit {
			// exit
			break
		}

		time.Sleep(waitTime)
	}
}

func checkURI(uri string) {
	if *n {
		fmt.Printf(okFmt, uri)
		return
	}

	resp, err := c.Get(baseURL + uri)
	if err == nil {
		resp.Body.Close()
	}

	if resp.StatusCode == 404 {
		fmt.Printf(okFmt, uri)
	} else {
		fmt.Fprintf(os.Stderr, ngFmt, uri)
		fmt.Fprintf(os.Stderr, locFmt, resp.Header.Get("Location"))
	}
}

func advanceRunes(runes []rune) []rune {
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

	return runes
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
