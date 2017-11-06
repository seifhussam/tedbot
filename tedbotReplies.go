package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ramin0/chatbot"
)

// handle Name
func handleName(session chatbot.Session, message string) (string, error) {
	if strings.EqualFold(message, "tedbot") {
		return "", fmt.Errorf(SamenameString)
	}
	session["name"] = message
	fmt.Println(session)
	return fmt.Sprintf(WelcomeString + " " + message + ", " + HelpString), nil
}

// handle RegularMessages
func handleChat(session chatbot.Session, message string) string {
	Check(message)
	return Respond()

}

func Check(M string) {
	M = strings.ToLower(M)
	M = strings.Trim(M, " ")
	M = M[0 : len(M)-2]

	for a, m := range Incoming[Phase] {
		Error = !strings.EqualFold(m, M)
		if Phase == 2 || Phase == 3 {
			Error = false
		}
		if !Error {
			switch Phase {
			case 1:
				if a == 0 {
					Phase = 2
				}
				if a == 1 {
					Phase = 3
				}
				break
			case 2:
				F := FindTopic(M)
				if F {
					Error = false
					Phase = 4
				} else {
					Error = true
					Phase = 4
				}
				break
			case 3:
				F := FindSpeaker(M)
				if F {
					Error = false
					Phase = 5
				} else {
					Error = true
					Phase = 5
				}
				break
			case 4:
				if a == 0 {
					Phase = 2
				}
				if a == 1 {
					fmt.Println("BYE 🙂")
					os.Exit(0)
				}
				break
			case 5:
				if a == 0 {
					Phase = 3
				}
				if a == 1 {
					fmt.Println("BYE 🙂")
					os.Exit(0)
				}
				break
			}
			break
		}
	}
}
func FindTopic(min string) bool {
	return true
}

func FindSpeaker(min string) bool {
	return true
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

	Incoming[4][0] = "yes"
	Incoming[4][1] = "no"

	Incoming[5][0] = "yes"
	Incoming[5][1] = "no"
	//---------------------------------------------------------------------------------------------------
	Outgoing[1][0] = "Hi stranger, I am TEDbot. Would you like to search for a topic or a sepaker?"
	Outgoing[1][1] = "I don't understand, please choose topic or speaker?"

	Outgoing[2][0] = "What topic should I get for you?"
	Outgoing[2][1] = "I couldn't find this topic, please tell me another one"

	Outgoing[3][0] = "Please tell me the name of the speaker?"
	Outgoing[3][1] = "I couldn't find a speaker with that name, please tell me another name"

	Outgoing[4][0] = "Here are some talks for this topic"
	Outgoing[4][1] = "Would you like to chosoe another topcic?"

	Outgoing[5][0] = "Here are some talks for this speaker"
	Outgoing[5][1] = "Would you like to choose another speaker?"
}

var Phase = 1
var s = ""
var Error = false

var Incoming = [6][2]string{}
var Outgoing = [6][2]string{}
