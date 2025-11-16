package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Read JSON from stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	// Parse JSON
	var response NotionResponse
	if err := json.Unmarshal(data, &response); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Convert to Markdown
	markdown := Convert(response)

	// Write to stdout
	fmt.Print(markdown)
}
