package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

// Load .env file and return matching key
func loadAPIKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func newClient(apiKey string) gpt3.Client {
	client := gpt3.NewClient(apiKey)
	return client
}

func validateQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}

// ctx is context used to control lifecycle of request
//
// gpt3.TextDavinci003Engine is the engine used to generate response
//   - https://platform.openai.com/docs/models/overview
//
// gpt3.CompletionRequest is a struct literal that contains the prompt and other parameters
//
// a callback function then handles the response
//   - each completion response (there can be one or more) is passed to the callback function
//   - allows program to specify what should happen, without needing to block/wait for event to occur
//
// https://platform.openai.com/docs/api-reference/completions/create
func getResponse(client gpt3.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		// tokens are tokenized words, broken down into syllables
		MaxTokens: gpt3.IntPtr(3000),
		// sampling temperature to use (between 0 and 2)
		// 0 is more focused/deterministic, 2 is more random
		// also consider using top_p instead
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		// prints first choice returned in response
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}
	fmt.Printf("\n\n")
}
