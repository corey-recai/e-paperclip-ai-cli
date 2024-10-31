package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func greeting(firstStart *bool) {
	fmt.Println("Axial prosperity!")
	fmt.Println("Paperclip OS assistive AI REPL. Type 'quit' to exit")
	*firstStart = false
}

func prompt(model *genai.GenerativeModel, ctx *context.Context) {
	var query string
	fmt.Println(">")
	_, err := fmt.Scanln(&query)
	if err != nil {
		log.Fatal(err)
	}

	if query == "quit" {
		return
	}

	resp, err := model.GenerateContent(*ctx, genai.Text(query))
	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
	}

	printResponse(resp)
	prompt(model, ctx)
}

func main() {
	ctx := context.Background()

	firstStart := true

	data, err := os.ReadFile(".api.env")
	if err != nil {
		log.Fatal(err)
	}
	result := strings.Split(string(data[:]), "=")
	os.Setenv(result[0], result[1])

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	if firstStart == true {
		greeting(&firstStart)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt(model, &ctx)

}
