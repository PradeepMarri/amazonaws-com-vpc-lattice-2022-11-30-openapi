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

func DeleteservicenetworkserviceassociationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		serviceNetworkServiceAssociationIdentifierVal, ok := args["serviceNetworkServiceAssociationIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serviceNetworkServiceAssociationIdentifier"), nil
		}
		serviceNetworkServiceAssociationIdentifier, ok := serviceNetworkServiceAssociationIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serviceNetworkServiceAssociationIdentifier"), nil
		}
		url := fmt.Sprintf("%s/servicenetworkserviceassociations/%s", cfg.BaseURL, serviceNetworkServiceAssociationIdentifier)
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
		var result models.DeleteServiceNetworkServiceAssociationResponse
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

func CreateDeleteservicenetworkserviceassociationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_servicenetworkserviceassociations_serviceNetworkServiceAssociationIdentifier",
		mcp.WithDescription("Deletes the association between a specified service and the specific service network. This request will fail if an association is still in progress."),
		mcp.WithString("serviceNetworkServiceAssociationIdentifier", mcp.Required(), mcp.Description("The ID or Amazon Resource Name (ARN) of the association.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteservicenetworkserviceassociationHandler(cfg),
	}
}
