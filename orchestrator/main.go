package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ollama/ollama/api"
)

func main() {
	ctx := context.Background()

	client, err := api.ClientFromEnvironment()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	req := &api.GenerateRequest{
		Model:  "mistral",
		Prompt: "Write a short poem about Go programming.",
	}
	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		if resp.Done {
			fmt.Println("\n[Generation complete]")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8090", nil)
}
