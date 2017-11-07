package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ramin0/chatbot"
)

// handle Name
func handleName(session chatbot.Session, message string) (string, error) {
	if strings.EqualFold(message, "tedbot") {
		return "", fmt.Errorf(SamenameString)
	}
	session["name"] = message
	session["phase"] = 1
	fmt.Println(session)
	return fmt.Sprintf(Outgoing[1][0]), nil
}

// handle RegularMessages
func handleChat(session chatbot.Session, message string) string {
	Phase = session["phase"].(int)
	R := "" + Check(message, session) + Respond()
	return R

}

func Check(M string, session chatbot.Session) string {
	M = strings.ToLower(M)
	M = strings.Trim(M, " ")

	if strings.EqualFold(M, "help") {
		return (Outgoing[Phase][2] + "<br>")
	}

	//M = M[0 : len(M)-2]

	for a, m := range Incoming[Phase] {
		Error = !strings.EqualFold(m, M)
		if Phase == 2 || Phase == 3 || Phase == 7 {
			Error = false
		}
		if !Error {
			switch Phase {
			case 1:
				if a == 0 {
					Phase = 2
					session["phase"] = 2
				}
				if a == 1 {
					Phase = 3
					session["phase"] = 3
				}
				if a == 2 {
					Phase = 1
					session["phase"] = 1
					return fetchTedinfo()
				}
				if a == 3 {
					Phase = 7
					session["phase"] = 7
				}
				break
			case 2:
				F := FindTopic(M)
				reg, err := regexp.Compile("[^a-zA-Z0-9]+")
				if err != nil {
					log.Fatal(err)
				}
				test := reg.ReplaceAllString(F, "")
				if test == "" {
					Error = true
					Phase = 2
					session["phase"] = 2
					return ""
				}
				Phase = 1
				session["phase"] = 1
				Error = false
				return F
				break
			case 3:
				F := searchSpeakername(M)
				reg, err := regexp.Compile("[^a-zA-Z0-9]+")
				if err != nil {
					log.Fatal(err)
				}
				test := reg.ReplaceAllString(F, "")
				if test == "" {
					Error = true
					Phase = 3
					session["phase"] = 3
					return ""
				}
				Phase = 1
				session["phase"] = 1
				Error = false
				return F
				break
			case 4:
				if a == 0 {
					Phase = 2
					session["phase"] = 2
				}
				if a == 1 {
					Phase = 1
					session["phase"] = 1
				}
				break
			case 5:
				if a == 0 {
					Phase = 3
					session["phase"] = 3
				}
				if a == 1 {
					Phase = 1
					session["phase"] = 1
				}
				break
			case 6:
				Phase = 1
				session["phase"] = 1
				Error = false
				break
			case 7:
				F := FindTalk(M)
				reg, err := regexp.Compile("[^a-zA-Z0-9]+")
				if err != nil {
					log.Fatal(err)
				}
				test := reg.ReplaceAllString(F, "")
				if test == "" {
					Error = true
					Phase = 7
					session["phase"] = 7
					return ""
				}
				Error = false
				Phase = 1
				session["phase"] = 1
				return F
				break
			case 8:
				if a == 0 {
					Phase = 7
					session["phase"] = 7
				}
				if a == 1 {
					Phase = 1
					session["phase"] = 1
				}
				break
			}
			break
		}
	}
	return " "
}

func Respond() string {
	if Error {
		return (Outgoing[Phase][1])
	}
	return (Outgoing[Phase][0])

}

func Init() {
	Phase = 1
	Error = false

	Incoming[1][0] = "topic"
	Incoming[1][1] = "speaker"
	Incoming[1][2] = "info"
	Incoming[1][3] = "talk"

	Incoming[4][0] = "yes"
	Incoming[4][1] = "no"

	Incoming[5][0] = "yes"
	Incoming[5][1] = "no"

	Incoming[8][0] = "yes"
	Incoming[8][1] = "no"

	//---------------------------------------------------------------------------------------------------
	Outgoing[1][0] = "Would you like to search for a Topic, a Speaker, Info or a Talk?"
	Outgoing[1][1] = "I don't understand, please choose Topic, Speaker, Info or Talk?"
	Outgoing[1][2] = "Please type Topic, Speaker, Info or a Talk"

	Outgoing[2][0] = "What topic should I get for you?"
	Outgoing[2][1] = "I couldn't find this topic, please tell me another one"
	Outgoing[2][2] = "Please type the name of the Topic you wish to view"

	Outgoing[3][0] = "Please tell me the name of the speaker?"
	Outgoing[3][1] = "I couldn't find a speaker with that name, please tell me another name"
	Outgoing[3][2] = "Please type a speaker name to view his/her talks"

	Outgoing[4][0] = "Here are some talks for this topic"
	Outgoing[4][1] = "Would you like to chosoe another topcic?"
	Outgoing[4][2] = "Please reply with a yes or a no"

	Outgoing[5][0] = "Here are some talks for this speaker"
	Outgoing[5][1] = "Would you like to choose another speaker?"
	Outgoing[5][2] = "Please reply with a yes or a no"

	Outgoing[7][0] = "What is the name of the talk you want?"
	Outgoing[7][1] = "I couldn't find this talk, please tell me another one"
	Outgoing[7][2] = "Please type the name of the talk you want"

	Outgoing[8][0] = "Here is the talk you asked for"
	Outgoing[8][1] = "Would you like to choose another talk?"
	Outgoing[8][2] = "Please reply with a yes or a no"

	//Outgoing[8][1] = "Sorry was that a yes or a no"

	//Outgoing[9][0] = "Here is the summary you asked for"
	//Outgoing[9][1] = "Would you like to choose another talk?"
	//Outgoing[9][2] = "Please reply with a yes or a no"

}

var Phase = 1
var s = ""
var Error = false

var Incoming = [10][5]string{}
var Outgoing = [10][3]string{}
