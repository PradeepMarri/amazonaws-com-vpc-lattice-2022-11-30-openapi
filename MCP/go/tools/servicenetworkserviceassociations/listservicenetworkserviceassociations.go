package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-vpc-lattice/mcp-server/config"
	"github.com/amazon-vpc-lattice/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListservicenetworkserviceassociationsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["serviceIdentifier"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("serviceIdentifier=%v", val))
		}
		if val, ok := args["serviceNetworkIdentifier"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("serviceNetworkIdentifier=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/servicenetworkserviceassociations%s", cfg.BaseURL, queryString)
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
		var result models.ListServiceNetworkServiceAssociationsResponse
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

func CreateListservicenetworkserviceassociationsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_servicenetworkserviceassociations",
		mcp.WithDescription("<p>Lists the associations between the service network and the service. You can filter the list either by service or service network. You must provide either the service network identifier or the service identifier.</p> <p>Every association in Amazon VPC Lattice is given a unique Amazon Resource Name (ARN), such as when a service network is associated with a VPC or when a service is associated with a service network. If the association is for a resource that is shared with another account, the association will include the local account ID as the prefix in the ARN for each account the resource is shared with.</p>"),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of results to return.")),
		mcp.WithString("nextToken", mcp.Description("A pagination token for the next page of results.")),
		mcp.WithString("serviceIdentifier", mcp.Description("The ID or Amazon Resource Name (ARN) of the service.")),
		mcp.WithString("serviceNetworkIdentifier", mcp.Description("The ID or Amazon Resource Name (ARN) of the service network.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListservicenetworkserviceassociationsHandler(cfg),
	}
}
