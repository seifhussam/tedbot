package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ramin0/chatbot"
	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
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

// handle Video

func fetchVideoLink(talk string) string {
	var (
		query      = flag.String("query", talk, "Search term")
		maxResults = flag.Int64("max-results", 1, "Max YouTube results")
	)

	const developerKey = "AIzaSyAnHiHotrP8zetr9MYdJwNxAcaXBdOYrL4"
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}
	thisid := ""
	for id := range videos {
		thisid = id
	}
	return Youtubelink + thisid
}
