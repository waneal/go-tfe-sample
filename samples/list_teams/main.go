package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-tfe"
)

func main() {

	token := os.Getenv("TF_API_TOKEN")
	organizationId := os.Getenv("TF_ORGANIZATION_ID")

	config := &tfe.Config{
		Token:             token,
		RetryServerErrors: true,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	teams, err := client.Teams.List(context.Background(), organizationId, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range teams.Items {
		fmt.Println(value.ID)
	}

}
