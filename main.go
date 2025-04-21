package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	choreoservice "github.com/rajithacharith/choreo-mcp/internal/choreo_service"
)

func main() {
	s := server.NewMCPServer(
		"Choreo Management MCP",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	getProjectsTool := mcp.NewTool("get_projects",
		mcp.WithDescription("Get projects from Choreo organization"),
		mcp.WithString("orgId",
			mcp.Required(),
			mcp.Description("Organization ID"),
		),
		mcp.WithString("token",
			mcp.Required(),
			mcp.Description("Bearer token for authentication"),
		),
	)
	s.AddTool(getProjectsTool, getProjects)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func getProjects(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgID, ok := request.Params.Arguments["orgId"].(string)
	if !ok {
		return nil, errors.New("orgId must be an string")
	}

	token, ok := request.Params.Arguments["token"].(string)
	if !ok {
		return nil, errors.New("token must be a string")
	}

	projects, err := choreoservice.GetProjects(orgID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", projects)), nil
}
