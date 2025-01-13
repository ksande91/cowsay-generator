package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	// Get the OpenAI API key from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY not set")
		return
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))

	prompt := "Generate a motivational or humorous quote for a terminal startup cow."

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	// Get the response content
	response := chatCompletion.Choices[0].Message.Content

	// Prepare the cowsay command
	cmd := exec.Command("cowsay")

	// Pipe the response to the cowsay input
	cmd.Stdin = bytes.NewReader([]byte(response))

	// Capture the output of the cowsay command
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running cowsay:", err)
		return
	}

	// Print the cowsay output
	fmt.Println(string(output))
}

