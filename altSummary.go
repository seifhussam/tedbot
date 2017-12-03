// helloWorld project main.go
package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func fetchAltSummary(keyword1 string, keyword2 string, keyword string) string {

	x := TedSearch + keyword2
	x1 := strings.Replace(keyword, "+", "_", -1)
	x3 := strings.Replace(keyword1, "+", "_", -1)
	x4 := strings.Replace(keyword2, "+", "_", -1)
	fmt.Println("x:" + x)
	// fmt.Println("x1:" + x1)
	x2 := crawlALTSummaryLink(x, x1, x3, x4)
	// fmt.Println("alt summary")
	// fmt.Println("x2:" + x2)
	summ := crawlAltSummary(x2)
	return summ

}
func crawlAltSummary(link string) string {
	var res = ""
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

				if strings.Contains(sth2.Data, "__INITIAL_DATA__") {
					res = parse("description", sth2.Data)
				}

			}
		}
	}
	return "Summary not availble"
}
func parse(attr string, json string) string {

	arr := strings.Split(json, "\":")
	var flag = false
	var res = ""
	for _, e := range arr {
		if flag {
			arr := strings.Split(e, ",")
			for index, e := range arr {
				if index < len(arr)-1 {
					res += e
				}

			}
			return res[1:]
		}
		if strings.Contains(e, attr) {
			flag = true
		}
	}
	return "Summary not found"
}
func crawlALTSummaryLink(link string, name string, keyword1 string, keyword2 string) string {
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
						// if strings.HasPrefix(a.Val, "/talks") {
						// 	fmt.Println("Compare")
						// 	fmt.Println(a.Val)
						// 	fmt.Println(keyword1)
						// 	fmt.Println(name)
						// 	fmt.Println("-------------------------")
						// }

						if (strings.Contains(a.Val, strings.ToLower(keyword2)) || strings.Contains(a.Val, strings.ToLower(keyword1)) || strings.Contains(a.Val, strings.ToLower(name))) && strings.HasPrefix(a.Val, "/talks") {
							res = "http://www.ted.com" + a.Val
						}
					}
				}
			}
		}
	}
	return res
}
