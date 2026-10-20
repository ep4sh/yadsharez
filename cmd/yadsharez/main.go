package main

import (
	"context"
	"log"
	"os"

	yadsh "github.com/ep4sh/yadsharez/pkg/yadsh"
)

func main() {
	args := os.Args[:]
	if len(args) < 2 || len(args) > 3 {
		yadsh.Help()
		return
	}

	filePath := os.Args[1]
	log.Printf("Filepath: %s\n", filePath)

	token, err := yadsh.GetOAuthToken()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = yadsh.Upload(ctx, filePath, token)
	if err != nil {
		log.Fatalf("error on Upload: %v\n", err)
	}
	err = yadsh.Download(ctx, filePath, token)
	if err != nil {
		log.Fatalf("error on Download: %v\n", err)
	}
}
