package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

var (
	owner       string
	repo        string
	pullNumber  int
	hideDeleted bool
)

func main() {
	flag.StringVar(&owner, "owner", "", "Owner of the github account")
	flag.StringVar(&repo, "repo", "", "Repo name")
	flag.IntVar(&pullNumber, "pull-number", 0, "Pull request number")
	flag.BoolVar(&hideDeleted, "hide-deleted", true, "Hide deleted files ")
	flag.Parse()

	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opts := &github.ListOptions{
		Page:    1,
		PerPage: 100,
	}

	paths := []string{}
	// Run till we list all the files, NextPage is 0 on the Last Page
	for opts.Page != 0 {
		files, response, err := client.PullRequests.ListFiles(ctx, owner, repo, pullNumber, opts)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if hideDeleted && file.GetStatus() != "removed" {
				paths = append(paths, file.GetFilename())
			} else {
				paths = append(paths, file.GetFilename())
			}
		}

		opts.Page = response.NextPage
	}

	for _, item := range paths {
		fmt.Println(item)
	}
}
