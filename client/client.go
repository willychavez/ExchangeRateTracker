package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Bid string `json:"bid"`
}

func main() {
	// Start the client
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Call the server
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON response
	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	// Save the quote to a file
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Write to the file in the desired format
	content := fmt.Sprintf("Dólar:{%s}\n", quote.Bid)
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("quote saved to cotacao.txt")
}
