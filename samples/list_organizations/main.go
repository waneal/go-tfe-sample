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

	config := &tfe.Config{
		Token:             token,
		RetryServerErrors: true,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	orgs, err := client.Organizations.List(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range orgs.Items {
		fmt.Println("Organization:", value.Name)
	}

}
