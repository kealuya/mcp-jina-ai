package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// fetchURLContent fetches content from a URL using Jina.ai's reader service
func fetchURLContentFromJina(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("request failed: empty URL")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "", fmt.Errorf("request failed: invalid URL format")
	}

	// Create a new resty client with timeout
	restyClient := resty.New().SetTimeout(30 * time.Second)

	// Make GET request to jina.ai with the provided URL
	resp, err := restyClient.R().Get("https://r.jina.ai/" + url)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}

	// Return the response content
	return string(resp.Body()), nil
}

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"mcp-jina-ai",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// Add a calculator tool
	mcpJinaAi := mcp.NewTool("mcp-jina-ai",
		mcp.WithDescription("Read the URL content and convert it to a large model-friendly markdown format for return"),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("URL is required, it is the webpage you want to summarize, for example https://example.com/a"),
		),
	)

	// Add the calculator handler
	s.AddTool(mcpJinaAi, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		url := request.Params.Arguments["url"].(string)

		content, err := fetchURLContentFromJina(url)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(content), nil
	})

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
