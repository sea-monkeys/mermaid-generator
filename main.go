package main

import (
	"context"

	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ollama/ollama/api"
)

var (
	FALSE = false
	TRUE  = true
)

func main() {
	ctx := context.Background()

	var ollamaRawUrl string
	if ollamaRawUrl = os.Getenv("OLLAMA_HOST"); ollamaRawUrl == "" {
		ollamaRawUrl = "http://localhost:11434"
	}

	model := "qwen2.5-coder:1.5b"
	systemInstructions := "You are a useful AI agent, expert with mermaid diagrams and source code generation."

	url, _ := url.Parse(ollamaRawUrl)

	fmt.Println("🤖", ollamaRawUrl, model)

	client := api.NewClient(url, http.DefaultClient)

	// Load the content of the prompt.txt file
	prompt, err := os.ReadFile("prompt.md")
	if err != nil {
		log.Fatalln("😡", err)
	}

	// Prompt construction
	messages := []api.Message{
		{Role: "system", Content: systemInstructions},
		{Role: "user", Content: string(prompt)},
	}

	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":   0.0,
			"repeat_last_n": 2,
		},
		Stream: &TRUE,
		//Format: json.RawMessage(`"json"`),
	}

	answer := ""
	err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
		answer += resp.Message.Content
		fmt.Print(resp.Message.Content)
		return nil
	})

	if err != nil {
		log.Fatalln("😡", err)
	}

	// generate a markdown file from the value of answer
	err = os.WriteFile("report.md", []byte(answer), 0644)
	if err != nil {
		log.Fatalln("😡", err)
	}

	//for {}
}
