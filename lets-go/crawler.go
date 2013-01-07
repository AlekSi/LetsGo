package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

// GET OMIT
func get(u string) *http.Response {
	res, err := http.Get(u)
	if err != nil {
		log.Print(err)
	}
	return res
}

// PARSE OMIT
var re = regexp.MustCompile(`href="(.+?)"`)

func parse(res *http.Response) (urls []string) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return
	}
	res.Body.Close()

	matches := re.FindAllSubmatch(body, -1)

	// PARSE 1 OMIT
	for _, match := range matches {
		m := string(match[1])
		u, err := url.Parse(m)
		if err != nil {
			continue
		}
		if u.Scheme == "" {
			u.Scheme = res.Request.URL.Scheme
		}
		if u.Host == "" {
			u.Host = res.Request.URL.Host
		}
		if u.Scheme == "http" || u.Scheme == "https" {
			urls = append(urls, u.String())
		}
	}
	return
}

// CRAWLER OMIT
var (
	urls      = make(chan string, 1000)
	responses = make(chan *http.Response, 1000)
)

func crawler() {
	for u := range urls {
		res := get(u)
		if res != nil {
			fmt.Printf("%d - %s\n", res.StatusCode, res.Request.URL)
			responses <- res
		}
	}
}

// PARSER OMIT
func parser() {
	m := make(map[string]bool)
	for res := range responses {
		for _, u := range parse(res) {
			if !m[u] {
				m[u] = true
				urls <- u
			}
		}
	}
}

// MAIN OMIT
func main() {
	log.SetFlags(0)
	for i := 0; i < 2; i++ {
		go crawler()
	}

	urls <- "http://railsclub.ru/"
	parser()
}
