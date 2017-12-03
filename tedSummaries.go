package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type talkdetail struct {
	summary  string
	critique string
}

// main function
func fetchSummary(keywords string, keywords1 string, keywords2 string) string {
	keywordsArr := strings.Split(keywords, "+")
	// fmt.Println(keywords)
	url, err := checkSummary(keywordsArr, (keywords))
	if err == nil {
		summary := crawlSummary(url)
		talkdetail := extractAttr(summary)
		return talkdetail.summary
	}
	// fmt.Println(err)
	return strings.Replace(strings.Replace(fetchAltSummary(keywords, keywords1, keywords2), "\"", "", -1), "\\", "", -1)
}

func extractAttr(strArr []string) talkdetail {
	var flag = false
	var res talkdetail
	for _, element := range strArr {
		if strings.EqualFold(element, "Summary") {
			flag = true
			continue
		}
		if strings.EqualFold(element, "Critique") || strings.EqualFold(element, "My Thoughts") {
			flag = false
			continue
		}
		if flag {
			res.summary += element
		}
	}
	return res
}

func putKeyWords(str []string) string {
	var res string
	for i, element := range str {
		if i == len(str)-1 {
			res += element
		} else {
			res += element + "+"
		}
	}
	return res
}
func checkSummary(keywordsArr []string, keywords string) (string, error) {
	var arr = tedSummarycrawler(TedSummaries + keywords)
	// fmt.Println(keywordsArr[0])
	// fmt.Println(keywordsArr[1])

	for _, e := range arr {
		if !strings.Contains(e, "?s=") && strings.Contains(e, strings.ToLower(keywordsArr[0])) && strings.Contains(e, strings.ToLower(keywordsArr[1])) && strings.Contains(e, strings.ToLower(keywordsArr[2])) {
			fmt.Println(e)
			fmt.Println("tedSummaries")
			return e, nil
		}
	}
	return "", fmt.Errorf("%s", SummaryError)

}
func tedSummarycrawler(link string) []string {
	var arr = []string{}
	//	var x int = 0

	resp, err := http.Get(link)

	if err == nil {
		defer resp.Body.Close()
		tokenized := html.NewTokenizer(resp.Body)
		for {
			sth := tokenized.Next()
			switch {
			case sth == html.ErrorToken:

				return arr
			case sth == html.StartTagToken:
				sth2 := tokenized.Token()
				isLink := sth2.Data == "a"
				if isLink {
					for _, a := range sth2.Attr {
						if a.Key == "href" {
							if a.Val != "" && len(a.Val) > 3 {
								var tempstr = "http"
								strSlice := a.Val[:4]
								if strSlice == tempstr {
									//fmt.Println("Link : " + a.Val)
								} else {
									//fmt.Println("Link : " + link + a.Val)
									a.Val = link + a.Val
								}
								arr = append(arr, a.Val)
							}
						}

					}
				}
			}
		}
	}
	return arr
}
func crawlSummary(link string) []string {
	var arr []string
	var f = false
	resp, err := http.Get(link)

	if err == nil {
		defer resp.Body.Close()
		tokenized := html.NewTokenizer(resp.Body)
		for {
			sth := tokenized.Next()
			switch {
			case sth == html.ErrorToken:

				return arr
			case sth == html.TextToken:
				sth2 := tokenized.Token()
				if sth2.Data != "" && sth2.Data != " " && sth2.Data != "\n" {
					if strings.Contains(sth2.Data, "Summary") {
						f = true
					} else if strings.Contains(sth2.Data, "\n") /*strings.Contains(sth2.Data, "Critique") || strings.Contains(sth2.Data, "My Thoughts")*/ {
						f = false
					}

					if f {
						arr = append(arr, sth2.Data)
					}
				}
			}
		}
	}
	return arr
}
