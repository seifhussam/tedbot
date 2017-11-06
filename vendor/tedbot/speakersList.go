package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type speaker struct {
	firstName string
	lastName  string
	talks     []talk
}
type talk struct {
	talkName    string
	talkInfo    string
	talkKeyWord string
	nameKeyword string
}

var speakerNewName string

var speakersTalksList = make(map[string]speaker)

func fetchSpeakerTalkList() {
	localspeakersTalksList := speakerListcrawler(WikiLink)

	for name, element := range localspeakersTalksList {
		var sp = speaker{}
		var namearr = strings.Split(name, " ")
		var f = true
		for _, element1 := range namearr {
			if f {
				sp.firstName = strings.ToLower(element1)
				f = !f
			} else {
				sp.lastName += strings.ToLower(strings.Replace(element1, "!", "", -1)) + " "

			}
		}
		var talks = []talk{}
		talksarr := strings.Split(element, ",")
		index := 0
		for _, e := range talksarr {
			if index < len(talksarr)-1 {
				var t = talk{}
				t.talkName, t.talkInfo = decodeTalk(e)
				t.nameKeyword, t.talkKeyWord = generateTalkKeyWord(sp.firstName, t.talkName)
				talks = append(talks, t)
			}
			index++
		}
		sp.talks = talks
		speakersTalksList[name] = sp
	}
}

func generateTalkKeyWord(name string, talkstr string) (string, string) {
	var res = name + "+"
	var res1 = ""
	talkstrArr := strings.Split(talkstr, " ")
	for _, e := range talkstrArr {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(e, "")
		x := strings.Replace(processedString, " ", "", -1)
		if len(processedString) > 1 && len(x) > 1 {
			res += processedString + "+"
			res1 += processedString + "+"
		}
	}
	// fmt.Println(len(res1))
	// fmt.Println(res1)
	// if len(res1) == 0 {
	// 	//	res1 = " "
	// }
	return res1[:(len(res1) - 1)], res[:(len(res) - 1)]
}

func decodeTalk(talks string) (string, string) {

	var res1 = ""
	var res2 = ""
	var flag = true
	var temp = 0
	for i := 0; i < len(talks); i++ {
		if flag {
			if (i > 2 && talks[i] == '(') || i+1 == len(talks) {
				flag = false
				res1 = talks[:i]
				temp = i
			} else {

			}
		} else {
			if talks[i] == ')' {
				res2 = talks[temp:i]
				break
			}

		}
	}
	return res1, res2
}
func speakerListcrawler(url string) map[string]string {
	SpeakerTalk := make(map[string]string)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return nil
	}
	b := resp.Body
	defer b.Close()
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return SpeakerTalk
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data != "td" {
				continue
			}
			SLGetMyData(t, z, SpeakerTalk)
		}
	}
	//return SpeakerTalk
}

// SLGetMyData ..
func SLGetMyData(t html.Token, z *html.Tokenizer, SpeakerTalk map[string]string) {
	for i := 0; i < 5; i++ {
		tt := z.Next()
		if tt == html.TextToken {
			t = z.Token()
			//fmt.Println(t.Data)
			SLAddThisNow(SpeakerTalk, t.Data)
			//break
		}
	}
}

//SLAddThisNow ..
func SLAddThisNow(SpeakerTalk map[string]string, data string) {
	if data == "\n" || data == " " || strings.Contains(data, "[") || data == " " {
	} else {
		if strings.Contains(data, ")") {
			SpeakerTalk[speakerNewName] += data + ","
		} else {

			if strings.Contains(data, ",") {
				Stringo := strings.Split(data, ",")
				speakerNewName = Stringo[1][1:len(Stringo[1])-1] + Stringo[0]
			} else {
				speakerNewName = data
			}
		}
	}
}
