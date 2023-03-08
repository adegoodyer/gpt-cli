package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	log.SetOutput(new(NullWriter))

	ClearScreen()
	PrintWelcome()

	ctx := context.Background()
	apiKey := loadAPIKey("API_KEY")
	client := newClient(apiKey)
	history := newHistory()

	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Printf("\x1b[%dm%s\x1b[0m", 34, "\nQuestion: ")

				// read input from stdin - line by line
				if !scanner.Scan() {
					break
				}

				question := scanner.Text()
				history.add(question)
				questionParam := validateQuestion(question)

				switch questionParam {
				case "q":
					quit = true
				case "":
					continue
				case "h":
					history.print()
				case "c":
					ClearScreen()
					PrintWelcome()
					continue
				default:
					getResponse(client, ctx, questionParam)
				}
			}
		},
	}
	log.Fatal(rootCmd.Execute())
}
