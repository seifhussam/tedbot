// helloWorld project main.go
package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func fetchAltSummary(keyword string) string {

	x := TedSearch + keyword
	x1 := strings.Replace(keyword, "+", "_", -1)
	x2 := crawlALTSummaryLink(x, x1)
	summ := crawlAltSummary(x2)
	return summ

}
func crawlAltSummary(link string) string {
	var res = ""
	var flag = false
	resp, err := http.Get(link)
	if err == nil {
		defer resp.Body.Close()
		tokenized := html.NewTokenizer(resp.Body)
		for {
			sth := tokenized.Next()
			switch {
			case sth == html.ErrorToken:
				return res
			case sth == html.TextToken:
				sth2 := tokenized.Token()
				if flag {
					reg, err := regexp.Compile("[^a-zA-Z0-9]+")
					if err != nil {
						log.Fatal(err)
					}
					x := reg.ReplaceAllString(sth2.Data, "")
					if !strings.Contains(x, "\n") && x != "\n" && x != " " && x != "\t" && x != "" {
						flag = !flag
						res = sth2.Data
						break
					}
				}
				if strings.Contains(sth2.Data, "About the talk") {
					flag = true
				}
			}
		}
	}
	return res
}

func crawlALTSummaryLink(link string, name string) string {
	var res = ""
	//	var flag = false
	resp, err := http.Get(link)
	//	var flag = false
	//	var x = 0
	if err == nil {
		defer resp.Body.Close()
		tokenized := html.NewTokenizer(resp.Body)
		for {
			sth := tokenized.Next()
			switch {
			case sth == html.ErrorToken:
				return res
			case sth == html.StartTagToken:
				sth2 := tokenized.Token()
				if sth2.Data == "a" {
					for _, a := range sth2.Attr {
						if strings.Contains(a.Val, strings.ToLower(name)) && strings.HasPrefix(a.Val, "/talks") {
							res = "http://www.ted.com/" + a.Val
						}
					}
				}
			}
		}
	}
	return res
}
