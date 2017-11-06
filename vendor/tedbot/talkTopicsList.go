package main

import (
	"fmt"
	"net/http"
	//"os"
	"strings"

	"golang.org/x/net/html"
)

var talkTopicsList = make(map[string]string)

//TopicsCounter ..
var TopicsCounter = 0

//TTcrawler ..
func TTcrawler(url string) []string {
	TalkTopics := make([]string, 500)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return TalkTopics
	}
	b := resp.Body
	defer b.Close()
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return TalkTopics
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data != "li" {
				continue
			}
			for _, a := range t.Attr {
				if a.Key == "class" {
					if a.Val == "d:b" {
						TTGetMyData(t, z, TalkTopics)
					}
				}
			}

		}
	}
	// return TalkTopics
}

//TTGetMyData ..
func TTGetMyData(t html.Token, z *html.Tokenizer, TalkTopics []string) {
	for i := 0; i < 5; i++ {
		tt := z.Next()
		if tt == html.TextToken {
			t = z.Token()
			//fmt.Println(t.Data)
			TTReplaceAllSpaces(t.Data, TalkTopics)
		}
	}
}

//TTReplaceAllSpaces ..
func TTReplaceAllSpaces(data string, TalkTopics []string) {
	var data2 = strings.Replace(data, "\t", "", -1)
	data2 = strings.Replace(data2, "\n", "", -1)
	if data2 == "\n" || len(data2) == 0 || len(data2) == 1 {

	} else {
		TTAddThisNow(TalkTopics, data2)
	}
}

//TTAddThisNow ..
func TTAddThisNow(TalkTopics []string, data string) {
	TalkTopics[TopicsCounter] = data
	TopicsCounter++
}

func fetchTopics() {

	TalkTopics := TTcrawler(TopicsURL)

	for i := 0; i < len(TalkTopics); i++ {
		if TalkTopics[i] != "" {
			talkTopicsList[TalkTopics[i]] = TopicsURL1 + strings.Replace(strings.ToLower(TalkTopics[i]), " ", "+", -1)
		}
	}

}
