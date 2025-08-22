package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-vpc-lattice/mcp-server/config"
	"github.com/amazon-vpc-lattice/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetruleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/services/%s/listeners/%s/rules/%s", cfg.BaseURL, listenerIdentifier, ruleIdentifier, serviceIdentifier)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetRuleResponse
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

func CreateGetruleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_services_serviceIdentifier_listeners_listenerIdentifier_rules_ruleIdentifier",
		mcp.WithDescription("Retrieves information about listener rules. You can also retrieve information about the default listener rule. For more information, see <a href="https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules">Listener rules</a> in the <i>Amazon VPC Lattice User Guide</i>."),
		mcp.WithString("listenerIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the listener.")),
		mcp.WithString("ruleIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the listener rule.")),
		mcp.WithString("serviceIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the service.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetruleHandler(cfg),
	}
}
