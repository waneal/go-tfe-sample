package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/hashicorp/go-tfe"
)

func main() {

	// create a new client
	token := os.Getenv("TF_API_TOKEN")
	config := &tfe.Config{
		Token:             token,
		RetryServerErrors: true,
	}
	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// parse flags
	var (
		teamId = flag.String("team-id", "", "Team id")
		username = flag.String("username", "", "User name")
	)
	flag.Parse()

	println(*teamId, *username)

	// define options
	options := tfe.TeamMemberAddOptions{
		Usernames: []string{*username},
	}

	// add member to team
	err = client.TeamMembers.Add(context.Background(), *teamId, options)
	if err != nil {
		log.Fatal(err)
	}

}
