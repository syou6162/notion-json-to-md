package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jomei/notionapi"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: notion-to-md <block-id-or-url>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "  notion-to-md cec15681-9083-4e1f-a0ae-72d268507aab")
		fmt.Fprintln(os.Stderr, "  notion-to-md https://www.notion.so/10xall/By-name-cec1568190834e1fa0ae72d268507aab")
		os.Exit(1)
	}
	input := os.Args[1]

	// Extract block ID from URL or use directly
	blockID, err := extractBlockID(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Get NOTION_TOKEN
	token := os.Getenv("NOTION_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "Error: NOTION_TOKEN environment variable not set")
		os.Exit(1)
	}

	// Initialize Notion client
	client := notionapi.NewClient(notionapi.Token(token))

	// Fetch all blocks recursively
	ctx := context.Background()
	blocks, err := fetchAllBlocks(ctx, client, blockID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching blocks: %v\n", err)
		os.Exit(1)
	}

	// Convert to Markdown
	markdown := convert(blocks)
	fmt.Print(markdown)
}

// extractBlockID extracts block ID from a Notion URL or returns the input as-is if it's already an ID
func extractBlockID(input string) (notionapi.BlockID, error) {
	// If it's a URL, extract the ID
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		// Match 32-character hex string (without hyphens)
		re := regexp.MustCompile(`([a-f0-9]{32})`)
		matches := re.FindStringSubmatch(input)
		if len(matches) < 2 {
			return "", fmt.Errorf("invalid Notion URL: cannot extract block ID")
		}
		// Convert to UUID format with hyphens
		id := matches[1]
		formatted := fmt.Sprintf("%s-%s-%s-%s-%s",
			id[0:8], id[8:12], id[12:16], id[16:20], id[20:32])
		return notionapi.BlockID(formatted), nil
	}

	// Assume it's already a block ID
	return notionapi.BlockID(input), nil
}
