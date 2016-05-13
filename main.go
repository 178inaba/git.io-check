package main

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	log "github.com/Sirupsen/logrus"
)

const (
	execDigit = 5
	baseURL   = "https://git.io"
)

var (
	interval = kingpin.Flag("interval", "interval time.").Default((3 * time.Second).String()).Duration()
	n        = kingpin.Flag("dry-run", "dry run mode.").Short('n').Bool()
	c        = new(http.Client)
)

func init() {
	kingpin.Parse()
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error { return errors.New("not redirect") }
}

func main() {
	runes := []rune{'0'}

	for {
		checkPath("/" + string(runes))

		runes = advanceRunes(runes)

		if len(runes) > execDigit {
			// exit
			break
		}

		time.Sleep(*interval)
	}
}

func checkPath(path string) {
	if *n {
		dryRun(path)
		return
	}

	resp, err := c.Get(baseURL + path)
	if err == nil {
		resp.Body.Close()
	}

	if resp.StatusCode == 404 {
		okLog(path)
	} else {
		ngLog(path, resp.Header.Get("Location"))
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

func okLog(path string) {
	log.WithFields(log.Fields{"path": path}).Info("ok!")
}

func ngLog(path string, location string) {
	log.WithFields(log.Fields{"path": path, "location": location}).Error("ng.")
}

func dryRun(path string) {
	if rand.Intn(2) == 1 {
		okLog(path)
	} else {
		ngLog(path, "https://example.com/")
	}
}
