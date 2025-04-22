package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	choreoservice "github.com/rajithacharith/choreo-mcp/internal/choreo_service"
)

var orgID string
var orgHandler string
var token = os.Getenv("TOKEN")

func main() {
	if token == "" {
		fmt.Println("TOKEN environment variables must be set")
		os.Exit(1)
	}

	organizations, err := choreoservice.GetOrganizations(token)
	if err != nil {
		fmt.Printf("Error getting organizations: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Organizations: %v\n", organizations)
	orgID = organizations[0].ID
	orgHandler = organizations[0].Handle

	s := server.NewMCPServer(
		"Choreo Management MCP",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	getProjectsTool := mcp.NewTool("get_projects",
		mcp.WithDescription("Get projects from Choreo organization"),
	)
	s.AddTool(getProjectsTool, getProjects)

	getComponentsTool := mcp.NewTool("get_components",
		mcp.WithDescription("Get components from Choreo project"),
		mcp.WithString("projectId",
			mcp.Required(),
			mcp.Description("UUID of the selected project."),
		),
	)
	s.AddTool(getComponentsTool, getComponents)

	getEnvironmentsTool := mcp.NewTool("get_environments",
		mcp.WithDescription("Get environments from Choreo organization"),
	)
	s.AddTool(getEnvironmentsTool, getEnvironments)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func getProjects(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	projects, err := choreoservice.GetProjects(orgID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", orgID, token, err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", projects)), nil
}

func getComponents(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	projectID, ok := request.Params.Arguments["projectId"].(string)
	if !ok {
		return nil, fmt.Errorf("projectId is required")
	}

	components, err := choreoservice.GetComponents(orgHandler, projectID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get components: %w", orgID, token, err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", components)), nil
}

func getEnvironments(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	environments, err := choreoservice.GetEnvironments(orgID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get environments: %w", orgID, token, err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", environments)), nil
}
