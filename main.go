package main

import (
	"github.com/google/go-github/github"
	"fmt"
)

func main() {

	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List("willnorris", nil)

	if err != nil {

	}

	fmt.Println(orgs)
}