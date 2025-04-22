# Choreo MCP Server

The Choreo MCP Server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction)
server that provides seamless integration with Choreo APIs, enabling advanced
automation and interaction capabilities for developers and tools.

## Use Cases

- Automating Choreo workflows and processes.
- Extracting and analyzing data from Choreo Organization.
- Building AI powered tools and applications that interact with Choreo ecosystem.


## Installation

### Usage with VS Code

For quick installation, use one of the one-click install buttons at the top of this README. Once you complete that flow, toggle Agent mode (located by the Copilot Chat text input) and the server will start.

For manual installation, add the following JSON block to your User Settings (JSON) file in VS Code. You can do this by pressing `Ctrl + Shift + P` and typing `Preferences: Open User Settings (JSON)`.

Optionally, you can add it to a file called `.vscode/mcp.json` in your workspace. This will allow you to share the configuration with others.

> Note that the `mcp` key is not needed in the `.vscode/mcp.json` file.

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "choreo_token",
        "description": "Choreo Personal Access Token",
        "password": true
      }
    ],
    "servers": {
      "choreo": {
        "command": "/path/to/choreo-mcp-server",
        "env": {
          "TOKEN": "${input:choreo_token}"
        }
      }
    }
  }
}
```

More about using MCP server tools in VS Code's [agent mode documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

### Usage with Claude Desktop

```json
{
  "mcpServers": {
    "choreo": {
      "command": "/path/to/choreo-mcp-server",
      "env": {
        "TOKEN": "<YOUR_TOKEN>"
      }
    }
  }
}
```

### Build from source

You can use `go build` to build the binary in the
`main` directory, and use the `main stdio` command with the `GITHUB_PERSONAL_ACCESS_TOKEN` environment variable set to your token. To specify the output location of the build, use the `-o` flag. You should configure your server to use the built executable as its `command`. For example:

```JSON
{
  "mcp": {
    "servers": {
      "github": {
        "command": "/path/to/choreo-mcp-server",
        "args": ["stdio"],
        "env": {
          "TOKEN": "<YOUR_TOKEN>"
        }
      }
    }
  }
}
```

## Tools

### Projects

- **get_projects** - Get details of the projects in default choreo organization
  - No parameters required

### Components

- **get_components** - Gets the components in a specific project

## License

This project is licensed under the terms of the MIT open source license. Please refer to [MIT](./LICENSE) for the full terms.