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

func UpdateruleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		ruleIdentifierVal, ok := args["ruleIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ruleIdentifier"), nil
		}
		ruleIdentifier, ok := ruleIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ruleIdentifier"), nil
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
		url := fmt.Sprintf("%s/services/%s/listeners/%s/rules/%s", cfg.BaseURL, listenerIdentifier, ruleIdentifier, serviceIdentifier)
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
		var result models.UpdateRuleResponse
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

func CreateUpdateruleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_services_serviceIdentifier_listeners_listenerIdentifier_rules_ruleIdentifier",
		mcp.WithDescription("Updates a rule for the listener. You can't modify a default listener rule. To modify a default listener rule, use <code>UpdateListener</code>."),
		mcp.WithString("listenerIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the listener.")),
		mcp.WithString("ruleIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the rule.")),
		mcp.WithString("serviceIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the service.")),
		mcp.WithNumber("priority", mcp.Description("Input parameter: The rule priority. A listener can't have multiple rules with the same priority.")),
		mcp.WithObject("action", mcp.Description("Input parameter: Describes the action for a rule. Each rule must include exactly one of the following types of actions: <code>forward </code>or <code>fixed-response</code>, and it must be the last action to be performed.")),
		mcp.WithObject("match", mcp.Description("Input parameter: Describes a rule match.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateruleHandler(cfg),
	}
}
