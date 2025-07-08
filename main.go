package main

import (
	"fmt"
	"log"
	"os"

	githubutils "github.com/TextGames207/github-utils"
)

const API_URL = "https://api.github.com/users/%s/events"

func main() {
	if len(os.Args) < 2 {
		log.SetFlags(0)
		log.Fatal("Please type the username after the program!")
	}
	username := os.Args[1]

	dat, err := githubutils.GetGitHubData(username, API_URL)

	if err != nil {
		log.SetFlags(0)
		log.Fatal("Geting the info was an error: %w", err)
	}

	for _, event := range dat {
		fmt.Println("----------------------")
		event.Print()
		fmt.Println("----------------------")
	}
}
