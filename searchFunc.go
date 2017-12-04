package main

import "strings"

func FindTopic(min string) string {

	for k, e := range talkTopicsList {
		if strings.Contains(strings.ToLower(k), strings.ToLower(min)) || strings.Contains(strings.ToLower(min), strings.ToLower(k)) {
			min = e
			break
		}
	}
	arr := fetchtopTalks(min)
	res := ""
	for _, e := range arr {
		res = res + e.speaker + " : " + e.talk + "<br>"
	}
	return "Here are the top talks in this topic. <br>" + res
}
func FindTalk(min string) string {
	res := ""
	x := 0
	for _, e := range speakersTalksList {
		for _, e1 := range e.talks {
			if strings.Contains(strings.ToLower(e1.talkName), strings.ToLower(min)) || strings.Contains(strings.ToLower(min), strings.ToLower(e1.talkName)) {
				res = res + strings.Title(e.firstName) + " " + strings.Title(e.lastName) + " : " + e1.talkName
				res = res + " \nVideo Link : " + fetchVideoLink(e1.talkName)
				res = res + " \nTalk Summary : " + fetchSummary(e1.talkKeyWord, e1.nameKeyword, e1.speakerKeyword) /* fetchAltSummary(e1.nameKeyword)*/ + "<br>"
				x++
				break
			}
		}
		if x > 5 {
			break
		}
	}
	if x > 1 {
		res = "I found a couple of talks that matches <br>" + res
	}
	return res
}

func searchSpeakername(min string) string {
	res := ""
	x := 0
	for k, e := range speakersTalksList {
		if strings.Contains(strings.ToLower(min), strings.ToLower(k)) || strings.Contains(strings.ToLower(k), strings.ToLower(min)) {
			for _, e1 := range e.talks {
				res = res + strings.Title(e.firstName) + " " + strings.Title(e.lastName) + " : " + e1.talkName
				res = res + " \nVideo Link : " + fetchVideoLink(e1.talkName)
				res = res + " \nTalk Summary : " + fetchSummary(e1.talkKeyWord, e1.nameKeyword, e1.speakerKeyword) /* fetchAltSummary(e1.nameKeyword)*/ + "<br>"
				x++
			}
			if x > 5 {
				break
			}
		}
	}
	if x > 1 {
		res = "I found a couple of talks that matches <br>" + res
	}
	return res
}
