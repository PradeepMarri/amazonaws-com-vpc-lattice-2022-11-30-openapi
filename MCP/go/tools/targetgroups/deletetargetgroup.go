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

func DeletetargetgroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		targetGroupIdentifierVal, ok := args["targetGroupIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: targetGroupIdentifier"), nil
		}
		targetGroupIdentifier, ok := targetGroupIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: targetGroupIdentifier"), nil
		}
		url := fmt.Sprintf("%s/targetgroups/%s", cfg.BaseURL, targetGroupIdentifier)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.DeleteTargetGroupResponse
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

func CreateDeletetargetgroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_targetgroups_targetGroupIdentifier",
		mcp.WithDescription("Deletes a target group. You can't delete a target group if it is used in a listener rule or if the target group creation is in progress."),
		mcp.WithString("targetGroupIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the target group.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletetargetgroupHandler(cfg),
	}
}
