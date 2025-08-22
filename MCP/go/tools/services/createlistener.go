package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-vpc-lattice/mcp-server/config"
	"github.com/amazon-vpc-lattice/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatelistenerHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		serviceIdentifierVal, ok := args["serviceIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serviceIdentifier"), nil
		}
		serviceIdentifier, ok := serviceIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serviceIdentifier"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/services/%s/listeners", cfg.BaseURL, serviceIdentifier)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateListenerResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatelistenerTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_services_serviceIdentifier_listeners",
		mcp.WithDescription("Creates a listener for a service. Before you start using your Amazon VPC Lattice service, you must add one or more listeners. A listener is a process that checks for connection requests to your services. For more information, see <a href="https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html">Listeners</a> in the <i>Amazon VPC Lattice User Guide</i>."),
		mcp.WithString("serviceIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the service.")),
		mcp.WithString("protocol", mcp.Required(), mcp.Description("Input parameter: The listener protocol HTTP or HTTPS.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: The tags for the listener.")),
		mcp.WithString("clientToken", mcp.Description("Input parameter: A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If you retry a request that completed successfully using the same client token and parameters, the retry succeeds without performing any actions. If the parameters aren't identical, the retry fails.")),
		mcp.WithObject("defaultAction", mcp.Required(), mcp.Description("Input parameter: Describes the action for a rule. Each rule must include exactly one of the following types of actions: <code>forward </code>or <code>fixed-response</code>, and it must be the last action to be performed.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the listener. A listener name must be unique within a service. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.")),
		mcp.WithNumber("port", mcp.Description("Input parameter: The listener port. You can specify a value from <code>1</code> to <code>65535</code>. For HTTP, the default is <code>80</code>. For HTTPS, the default is <code>443</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatelistenerHandler(cfg),
	}
}
