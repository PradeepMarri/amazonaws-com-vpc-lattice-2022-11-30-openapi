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

func BatchupdateruleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		listenerIdentifierVal, ok := args["listenerIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: listenerIdentifier"), nil
		}
		listenerIdentifier, ok := listenerIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: listenerIdentifier"), nil
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
		url := fmt.Sprintf("%s/services/%s/listeners/%s/rules", cfg.BaseURL, listenerIdentifier, serviceIdentifier)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.BatchUpdateRuleResponse
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

func CreateBatchupdateruleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_services_serviceIdentifier_listeners_listenerIdentifier_rules",
		mcp.WithDescription("Updates the listener rules in a batch. You can use this operation to change the priority of listener rules. This can be useful when bulk updating or swapping rule priority. "),
		mcp.WithString("listenerIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the listener.")),
		mcp.WithString("serviceIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the service.")),
		mcp.WithArray("rules", mcp.Required(), mcp.Description("Input parameter: The rules for the specified listener.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BatchupdateruleHandler(cfg),
	}
}
