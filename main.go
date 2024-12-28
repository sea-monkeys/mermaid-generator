package main

import (
	"context"
	"time"

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

	var toolsLLM string
	if toolsLLM = os.Getenv("TOOLS_LLM"); toolsLLM == "" {
		//toolsLLM = "allenporter/xlam:1b"
		toolsLLM = "qwen2.5:1.5b"
	}

	url, _ := url.Parse(ollamaRawUrl)

	client := api.NewClient(url, http.DefaultClient)

	prompt := `Generate a mermaid graph from this: 
	- The MCP Client will be a simple HTTP client run by the Host GenAI application.
	- The MCP Client will make HTTP requests to the MCP Server to get 
	the list of tools and to make tool calls. 
	- The MCP Server will respond with the list of tools and the output of 
	the tool calls.
	Add clear explanation at the end with emoji.
	`

	// Prompt construction
	messages := []api.Message{
		{Role: "user", Content: prompt},
	}

	req := &api.ChatRequest{
		Model:    toolsLLM,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":   0.0,
			"repeat_last_n": 2,
		},
		Stream: &TRUE,
		//Format: json.RawMessage(`"json"`),
	}

	answer := ""
	err := client.Chat(ctx, req, func(resp api.ChatResponse) error {
		//fmt.Println("🖐️", resp.Message.ToolCalls)
		answer += resp.Message.Content
		fmt.Print(resp.Message.Content)
		//fmt.Println(answer)

		return nil
	})

	if err != nil {
		log.Fatalln("😡", err)
	}

	fmt.Println("***-----***")
	fmt.Println("🚀...", ollamaRawUrl, toolsLLM)

	// Create a channel for exit code
	exitCode := make(chan int)

	// Start the infinite loop in a goroutine
	go func() {
		for {
			//fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for exit code at the end of main
	code := <-exitCode
	os.Exit(code)

}
