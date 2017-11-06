package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ramin0/chatbot"
	// Autoload environment variables in .env
	_ "github.com/joho/godotenv/autoload"
)

func chatbotProcess(session chatbot.Session, message string) (string, error) {
	if session["name"] == nil {
		return handleName(session, message)
	}
	return fmt.Sprintf("I got your name"), nil

}

//var speakersTalksList map[string]string

func main() {

	// Uncomment the following lines to customize the chatbot
	chatbot.WelcomeMessage = HelloString
	chatbot.ProcessFunc(chatbotProcess)
	//go getLink()
	// Use the PORT environment variable

	// fill speakers
	fetchSpeakerTalkList()

	// for _, element := range speakersTalksList {
	// 	// fmt.Print("First Name : " + element.firstName)
	// 	// fmt.Print("	Last Name : " + element.lastName)
	// 	for _, e := range element.talks {
	// 		// fmt.Print(" Talk : " + e.talkName)
	// 		// fmt.Print(" info : " + e.talkInfo)
	// 		// fmt.Print(" Namekey : " + e.nameKeyword)
	// 		// fmt.Print(" key : " + e.talkKeyWord)
	// 		// fmt.Println()
	// 	}

	// }

	//fill topics
	fetchTopics()

	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}
