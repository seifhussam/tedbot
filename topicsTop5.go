// helloWorld project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type talkRes struct {
	speaker string
	talk    string
}

func fetchtopTalks(topic string) []talkRes {
	res := topTopicscrawler(strings.ToLower(topic))
	if len(res) == 0 || strings.Contains(res[0], "Looks like there werent any talks matching your search criteria Try using a more general term or searching with fewer filters") {
		return []talkRes{}
	}
	var talks = []talkRes{}
	for i := 0; i < len(res)-1; i += 2 {
		talks = append(talks, talkRes{res[i], res[i+1]})
	}
	// for _, e := range talks {
	// 	fmt.Println(e.speaker)
	// 	fmt.Println(e.talk)
	// }
	return talks
}
func topTopicscrawler(link string) []string {
	var res = []string{}
	resp, err := http.Get(link)
	var flag = false
	var x = 0
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
				if !strings.Contains(sth2.Data, "Under 6 minutes") && sth2.Data != "\n" && sth2.Data != " " && sth2.Data != "\t" && sth2.Data != "" {
					fmt.Println(sth2.Data)
					if strings.Contains(sth2.Data, "Showing all") {
						x = 41
					}
					if sth2.Data == "Informative" {
						flag = true
					} else if flag {
						reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
						if err != nil {
							log.Fatal(err)
						}

						if x%8 == 1 || x%8 == 2 {
							processedString := reg.ReplaceAllString(sth2.Data, "")
							//							fmt.Print("Line : ")
							//							fmt.Print(x)
							//							fmt.Print(" " + processedString)
							//							fmt.Println()
							res = append(res, processedString)
						}
						x++
						if x > 40 {
							flag = false
						}
					}
				}
			}
		}
	}
	return res
}
