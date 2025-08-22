package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetServiceNetworkServiceAssociationRequest represents the GetServiceNetworkServiceAssociationRequest schema from the OpenAPI specification
type GetServiceNetworkServiceAssociationRequest struct {
}

// ServiceNetworkSummary represents the ServiceNetworkSummary schema from the OpenAPI specification
type ServiceNetworkSummary struct {
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Numberofassociatedservices interface{} `json:"numberOfAssociatedServices,omitempty"`
	Numberofassociatedvpcs interface{} `json:"numberOfAssociatedVPCs,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// AccessLogSubscriptionSummary represents the AccessLogSubscriptionSummary schema from the OpenAPI specification
type AccessLogSubscriptionSummary struct {
	Id interface{} `json:"id"`
	Lastupdatedat interface{} `json:"lastUpdatedAt"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourceid interface{} `json:"resourceId"`
	Arn interface{} `json:"arn"`
	Createdat interface{} `json:"createdAt"`
	Destinationarn interface{} `json:"destinationArn"`
}

// CreateRuleRequest represents the CreateRuleRequest schema from the OpenAPI specification
type CreateRuleRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Match interface{} `json:"match"`
	Name interface{} `json:"name"`
	Priority interface{} `json:"priority"`
	Tags interface{} `json:"tags,omitempty"`
	Action interface{} `json:"action"`
}

// RuleMatch represents the RuleMatch schema from the OpenAPI specification
type RuleMatch struct {
	Httpmatch interface{} `json:"httpMatch,omitempty"`
}

// DeleteServiceNetworkRequest represents the DeleteServiceNetworkRequest schema from the OpenAPI specification
type DeleteServiceNetworkRequest struct {
}

// ServiceSummary represents the ServiceSummary schema from the OpenAPI specification
type ServiceSummary struct {
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
	Dnsentry interface{} `json:"dnsEntry,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// BatchUpdateRuleRequest represents the BatchUpdateRuleRequest schema from the OpenAPI specification
type BatchUpdateRuleRequest struct {
	Rules interface{} `json:"rules"`
}

// CreateServiceNetworkVpcAssociationRequest represents the CreateServiceNetworkVpcAssociationRequest schema from the OpenAPI specification
type CreateServiceNetworkVpcAssociationRequest struct {
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Servicenetworkidentifier interface{} `json:"serviceNetworkIdentifier"`
	Tags interface{} `json:"tags,omitempty"`
	Vpcidentifier interface{} `json:"vpcIdentifier"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// UpdateRuleRequest represents the UpdateRuleRequest schema from the OpenAPI specification
type UpdateRuleRequest struct {
	Priority interface{} `json:"priority,omitempty"`
	Action interface{} `json:"action,omitempty"`
	Match interface{} `json:"match,omitempty"`
}

// RuleUpdateSuccess represents the RuleUpdateSuccess schema from the OpenAPI specification
type RuleUpdateSuccess struct {
	Id interface{} `json:"id,omitempty"`
	Isdefault interface{} `json:"isDefault,omitempty"`
	Match interface{} `json:"match,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Action interface{} `json:"action,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// GetServiceNetworkVpcAssociationResponse represents the GetServiceNetworkVpcAssociationResponse schema from the OpenAPI specification
type GetServiceNetworkVpcAssociationResponse struct {
	Servicenetworkarn interface{} `json:"serviceNetworkArn,omitempty"`
	Servicenetworkname interface{} `json:"serviceNetworkName,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Servicenetworkid interface{} `json:"serviceNetworkId,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
}

// TargetGroupSummary represents the TargetGroupSummary schema from the OpenAPI specification
type TargetGroupSummary struct {
	Name interface{} `json:"name,omitempty"`
	Port interface{} `json:"port,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Protocol interface{} `json:"protocol,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Vpcidentifier interface{} `json:"vpcIdentifier,omitempty"`
	Ipaddresstype interface{} `json:"ipAddressType,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Servicearns interface{} `json:"serviceArns,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// RuleUpdateFailure represents the RuleUpdateFailure schema from the OpenAPI specification
type RuleUpdateFailure struct {
	Failurecode interface{} `json:"failureCode,omitempty"`
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Ruleidentifier interface{} `json:"ruleIdentifier,omitempty"`
}

// ListServicesResponse represents the ListServicesResponse schema from the OpenAPI specification
type ListServicesResponse struct {
	Items interface{} `json:"items,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListTargetGroupsResponse represents the ListTargetGroupsResponse schema from the OpenAPI specification
type ListTargetGroupsResponse struct {
	Items interface{} `json:"items,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListListenersRequest represents the ListListenersRequest schema from the OpenAPI specification
type ListListenersRequest struct {
}

// DeleteServiceRequest represents the DeleteServiceRequest schema from the OpenAPI specification
type DeleteServiceRequest struct {
}

// GetServiceNetworkVpcAssociationRequest represents the GetServiceNetworkVpcAssociationRequest schema from the OpenAPI specification
type GetServiceNetworkVpcAssociationRequest struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// DeleteAccessLogSubscriptionResponse represents the DeleteAccessLogSubscriptionResponse schema from the OpenAPI specification
type DeleteAccessLogSubscriptionResponse struct {
}

// TargetGroupConfig represents the TargetGroupConfig schema from the OpenAPI specification
type TargetGroupConfig struct {
	Healthcheck interface{} `json:"healthCheck,omitempty"`
	Ipaddresstype interface{} `json:"ipAddressType,omitempty"`
	Port interface{} `json:"port"`
	Protocol interface{} `json:"protocol"`
	Protocolversion interface{} `json:"protocolVersion,omitempty"`
	Vpcidentifier interface{} `json:"vpcIdentifier"`
}

// RuleUpdate represents the RuleUpdate schema from the OpenAPI specification
type RuleUpdate struct {
	Action interface{} `json:"action,omitempty"`
	Match interface{} `json:"match,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Ruleidentifier interface{} `json:"ruleIdentifier"`
}

// DeleteServiceNetworkVpcAssociationResponse represents the DeleteServiceNetworkVpcAssociationResponse schema from the OpenAPI specification
type DeleteServiceNetworkVpcAssociationResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// RegisterTargetsResponse represents the RegisterTargetsResponse schema from the OpenAPI specification
type RegisterTargetsResponse struct {
	Successful interface{} `json:"successful,omitempty"`
	Unsuccessful interface{} `json:"unsuccessful,omitempty"`
}

// GetServiceNetworkResponse represents the GetServiceNetworkResponse schema from the OpenAPI specification
type GetServiceNetworkResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Numberofassociatedservices interface{} `json:"numberOfAssociatedServices,omitempty"`
	Numberofassociatedvpcs interface{} `json:"numberOfAssociatedVPCs,omitempty"`
}

// GetServiceNetworkServiceAssociationResponse represents the GetServiceNetworkServiceAssociationResponse schema from the OpenAPI specification
type GetServiceNetworkServiceAssociationResponse struct {
	Status interface{} `json:"status,omitempty"`
	Servicenetworkarn interface{} `json:"serviceNetworkArn,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Servicenetworkname interface{} `json:"serviceNetworkName,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Dnsentry interface{} `json:"dnsEntry,omitempty"`
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
	Servicenetworkid interface{} `json:"serviceNetworkId,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Servicearn interface{} `json:"serviceArn,omitempty"`
	Serviceid interface{} `json:"serviceId,omitempty"`
	Servicename interface{} `json:"serviceName,omitempty"`
}

// HealthCheckConfig represents the HealthCheckConfig schema from the OpenAPI specification
type HealthCheckConfig struct {
	Matcher interface{} `json:"matcher,omitempty"`
	Protocol interface{} `json:"protocol,omitempty"`
	Healthythresholdcount interface{} `json:"healthyThresholdCount,omitempty"`
	Unhealthythresholdcount interface{} `json:"unhealthyThresholdCount,omitempty"`
	Healthcheckintervalseconds interface{} `json:"healthCheckIntervalSeconds,omitempty"`
	Path interface{} `json:"path,omitempty"`
	Protocolversion interface{} `json:"protocolVersion,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Healthchecktimeoutseconds interface{} `json:"healthCheckTimeoutSeconds,omitempty"`
	Port interface{} `json:"port,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// HttpMatch represents the HttpMatch schema from the OpenAPI specification
type HttpMatch struct {
	Headermatches interface{} `json:"headerMatches,omitempty"`
	Method interface{} `json:"method,omitempty"`
	Pathmatch interface{} `json:"pathMatch,omitempty"`
}

// ListServiceNetworkServiceAssociationsRequest represents the ListServiceNetworkServiceAssociationsRequest schema from the OpenAPI specification
type ListServiceNetworkServiceAssociationsRequest struct {
}

// GetTargetGroupResponse represents the GetTargetGroupResponse schema from the OpenAPI specification
type GetTargetGroupResponse struct {
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Config interface{} `json:"config,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Servicearns interface{} `json:"serviceArns,omitempty"`
}

// ServiceNetworkServiceAssociationSummary represents the ServiceNetworkServiceAssociationSummary schema from the OpenAPI specification
type ServiceNetworkServiceAssociationSummary struct {
	Servicenetworkarn interface{} `json:"serviceNetworkArn,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
	Dnsentry interface{} `json:"dnsEntry,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Serviceid interface{} `json:"serviceId,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Servicename interface{} `json:"serviceName,omitempty"`
	Servicenetworkid interface{} `json:"serviceNetworkId,omitempty"`
	Servicenetworkname interface{} `json:"serviceNetworkName,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Servicearn interface{} `json:"serviceArn,omitempty"`
}

// CreateServiceNetworkRequest represents the CreateServiceNetworkRequest schema from the OpenAPI specification
type CreateServiceNetworkRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
}

// ListTargetGroupsRequest represents the ListTargetGroupsRequest schema from the OpenAPI specification
type ListTargetGroupsRequest struct {
}

// ServiceNetworkVpcAssociationSummary represents the ServiceNetworkVpcAssociationSummary schema from the OpenAPI specification
type ServiceNetworkVpcAssociationSummary struct {
	Servicenetworkid interface{} `json:"serviceNetworkId,omitempty"`
	Servicenetworkname interface{} `json:"serviceNetworkName,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Servicenetworkarn interface{} `json:"serviceNetworkArn,omitempty"`
}

// DeleteResourcePolicyResponse represents the DeleteResourcePolicyResponse schema from the OpenAPI specification
type DeleteResourcePolicyResponse struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// GetResourcePolicyRequest represents the GetResourcePolicyRequest schema from the OpenAPI specification
type GetResourcePolicyRequest struct {
}

// CreateServiceNetworkServiceAssociationRequest represents the CreateServiceNetworkServiceAssociationRequest schema from the OpenAPI specification
type CreateServiceNetworkServiceAssociationRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Serviceidentifier interface{} `json:"serviceIdentifier"`
	Servicenetworkidentifier interface{} `json:"serviceNetworkIdentifier"`
	Tags interface{} `json:"tags,omitempty"`
}

// UpdateServiceNetworkRequest represents the UpdateServiceNetworkRequest schema from the OpenAPI specification
type UpdateServiceNetworkRequest struct {
	Authtype interface{} `json:"authType"`
}

// PathMatchType represents the PathMatchType schema from the OpenAPI specification
type PathMatchType struct {
	Exact interface{} `json:"exact,omitempty"`
	Prefix interface{} `json:"prefix,omitempty"`
}

// GetRuleRequest represents the GetRuleRequest schema from the OpenAPI specification
type GetRuleRequest struct {
}

// ListAccessLogSubscriptionsResponse represents the ListAccessLogSubscriptionsResponse schema from the OpenAPI specification
type ListAccessLogSubscriptionsResponse struct {
	Items interface{} `json:"items"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateServiceNetworkResponse represents the CreateServiceNetworkResponse schema from the OpenAPI specification
type CreateServiceNetworkResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// PutAuthPolicyResponse represents the PutAuthPolicyResponse schema from the OpenAPI specification
type PutAuthPolicyResponse struct {
	State interface{} `json:"state,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
}

// UpdateAccessLogSubscriptionResponse represents the UpdateAccessLogSubscriptionResponse schema from the OpenAPI specification
type UpdateAccessLogSubscriptionResponse struct {
	Id interface{} `json:"id"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourceid interface{} `json:"resourceId"`
	Arn interface{} `json:"arn"`
	Destinationarn interface{} `json:"destinationArn"`
}

// DeleteRuleResponse represents the DeleteRuleResponse schema from the OpenAPI specification
type DeleteRuleResponse struct {
}

// CreateRuleResponse represents the CreateRuleResponse schema from the OpenAPI specification
type CreateRuleResponse struct {
	Name interface{} `json:"name,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Action interface{} `json:"action,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Match interface{} `json:"match,omitempty"`
}

// PutAuthPolicyRequest represents the PutAuthPolicyRequest schema from the OpenAPI specification
type PutAuthPolicyRequest struct {
	Policy interface{} `json:"policy"`
}

// CreateServiceNetworkVpcAssociationResponse represents the CreateServiceNetworkVpcAssociationResponse schema from the OpenAPI specification
type CreateServiceNetworkVpcAssociationResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ListRulesRequest represents the ListRulesRequest schema from the OpenAPI specification
type ListRulesRequest struct {
}

// CreateServiceRequest represents the CreateServiceRequest schema from the OpenAPI specification
type CreateServiceRequest struct {
	Customdomainname interface{} `json:"customDomainName,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Certificatearn interface{} `json:"certificateArn,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// DeleteResourcePolicyRequest represents the DeleteResourcePolicyRequest schema from the OpenAPI specification
type DeleteResourcePolicyRequest struct {
}

// RegisterTargetsRequest represents the RegisterTargetsRequest schema from the OpenAPI specification
type RegisterTargetsRequest struct {
	Targets interface{} `json:"targets"`
}

// DeleteAuthPolicyResponse represents the DeleteAuthPolicyResponse schema from the OpenAPI specification
type DeleteAuthPolicyResponse struct {
}

// ListRulesResponse represents the ListRulesResponse schema from the OpenAPI specification
type ListRulesResponse struct {
	Items interface{} `json:"items"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateRuleResponse represents the UpdateRuleResponse schema from the OpenAPI specification
type UpdateRuleResponse struct {
	Match interface{} `json:"match,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Action interface{} `json:"action,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Isdefault interface{} `json:"isDefault,omitempty"`
}

// CreateServiceNetworkServiceAssociationResponse represents the CreateServiceNetworkServiceAssociationResponse schema from the OpenAPI specification
type CreateServiceNetworkServiceAssociationResponse struct {
	Dnsentry interface{} `json:"dnsEntry,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
}

// ListTargetsRequest represents the ListTargetsRequest schema from the OpenAPI specification
type ListTargetsRequest struct {
	Targets interface{} `json:"targets,omitempty"`
}

// UpdateServiceNetworkResponse represents the UpdateServiceNetworkResponse schema from the OpenAPI specification
type UpdateServiceNetworkResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ListTargetsResponse represents the ListTargetsResponse schema from the OpenAPI specification
type ListTargetsResponse struct {
	Items interface{} `json:"items"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// Matcher represents the Matcher schema from the OpenAPI specification
type Matcher struct {
	Httpcode interface{} `json:"httpCode,omitempty"`
}

// CreateTargetGroupRequest represents the CreateTargetGroupRequest schema from the OpenAPI specification
type CreateTargetGroupRequest struct {
	Config interface{} `json:"config,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	TypeField interface{} `json:"type"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// ListAccessLogSubscriptionsRequest represents the ListAccessLogSubscriptionsRequest schema from the OpenAPI specification
type ListAccessLogSubscriptionsRequest struct {
}

// UpdateServiceRequest represents the UpdateServiceRequest schema from the OpenAPI specification
type UpdateServiceRequest struct {
	Authtype interface{} `json:"authType,omitempty"`
	Certificatearn interface{} `json:"certificateArn,omitempty"`
}

// DeleteServiceResponse represents the DeleteServiceResponse schema from the OpenAPI specification
type DeleteServiceResponse struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// GetListenerRequest represents the GetListenerRequest schema from the OpenAPI specification
type GetListenerRequest struct {
}

// RuleAction represents the RuleAction schema from the OpenAPI specification
type RuleAction struct {
	Fixedresponse interface{} `json:"fixedResponse,omitempty"`
	Forward interface{} `json:"forward,omitempty"`
}

// CreateAccessLogSubscriptionRequest represents the CreateAccessLogSubscriptionRequest schema from the OpenAPI specification
type CreateAccessLogSubscriptionRequest struct {
	Resourceidentifier interface{} `json:"resourceIdentifier"`
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Destinationarn interface{} `json:"destinationArn"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// PutResourcePolicyRequest represents the PutResourcePolicyRequest schema from the OpenAPI specification
type PutResourcePolicyRequest struct {
	Policy interface{} `json:"policy"`
}

// UpdateTargetGroupResponse represents the UpdateTargetGroupResponse schema from the OpenAPI specification
type UpdateTargetGroupResponse struct {
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Config interface{} `json:"config,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// GetServiceNetworkRequest represents the GetServiceNetworkRequest schema from the OpenAPI specification
type GetServiceNetworkRequest struct {
}

// GetTargetGroupRequest represents the GetTargetGroupRequest schema from the OpenAPI specification
type GetTargetGroupRequest struct {
}

// ListServiceNetworkVpcAssociationsResponse represents the ListServiceNetworkVpcAssociationsResponse schema from the OpenAPI specification
type ListServiceNetworkVpcAssociationsResponse struct {
	Items interface{} `json:"items"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// GetRuleResponse represents the GetRuleResponse schema from the OpenAPI specification
type GetRuleResponse struct {
	Id interface{} `json:"id,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Isdefault interface{} `json:"isDefault,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Match interface{} `json:"match,omitempty"`
	Action interface{} `json:"action,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// TargetFailure represents the TargetFailure schema from the OpenAPI specification
type TargetFailure struct {
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Port interface{} `json:"port,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
}

// CreateTargetGroupResponse represents the CreateTargetGroupResponse schema from the OpenAPI specification
type CreateTargetGroupResponse struct {
	Status interface{} `json:"status,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Config interface{} `json:"config,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// GetServiceResponse represents the GetServiceResponse schema from the OpenAPI specification
type GetServiceResponse struct {
	Certificatearn interface{} `json:"certificateArn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Failurecode interface{} `json:"failureCode,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
	Dnsentry interface{} `json:"dnsEntry,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// DeleteServiceNetworkResponse represents the DeleteServiceNetworkResponse schema from the OpenAPI specification
type DeleteServiceNetworkResponse struct {
}

// UpdateServiceNetworkVpcAssociationResponse represents the UpdateServiceNetworkVpcAssociationResponse schema from the OpenAPI specification
type UpdateServiceNetworkVpcAssociationResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UpdateListenerRequest represents the UpdateListenerRequest schema from the OpenAPI specification
type UpdateListenerRequest struct {
	Defaultaction interface{} `json:"defaultAction"`
}

// WeightedTargetGroup represents the WeightedTargetGroup schema from the OpenAPI specification
type WeightedTargetGroup struct {
	Targetgroupidentifier interface{} `json:"targetGroupIdentifier"`
	Weight interface{} `json:"weight,omitempty"`
}

// ForwardAction represents the ForwardAction schema from the OpenAPI specification
type ForwardAction struct {
	Targetgroups interface{} `json:"targetGroups"`
}

// GetListenerResponse represents the GetListenerResponse schema from the OpenAPI specification
type GetListenerResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Defaultaction interface{} `json:"defaultAction,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Port interface{} `json:"port,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Servicearn interface{} `json:"serviceArn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Protocol interface{} `json:"protocol,omitempty"`
	Serviceid interface{} `json:"serviceId,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// FixedResponseAction represents the FixedResponseAction schema from the OpenAPI specification
type FixedResponseAction struct {
	Statuscode interface{} `json:"statusCode"`
}

// CreateListenerResponse represents the CreateListenerResponse schema from the OpenAPI specification
type CreateListenerResponse struct {
	Name interface{} `json:"name,omitempty"`
	Port interface{} `json:"port,omitempty"`
	Protocol interface{} `json:"protocol,omitempty"`
	Servicearn interface{} `json:"serviceArn,omitempty"`
	Serviceid interface{} `json:"serviceId,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Defaultaction interface{} `json:"defaultAction,omitempty"`
	Id interface{} `json:"id,omitempty"`
}

// GetAuthPolicyResponse represents the GetAuthPolicyResponse schema from the OpenAPI specification
type GetAuthPolicyResponse struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
	State interface{} `json:"state,omitempty"`
}

// ListServiceNetworkServiceAssociationsResponse represents the ListServiceNetworkServiceAssociationsResponse schema from the OpenAPI specification
type ListServiceNetworkServiceAssociationsResponse struct {
	Items interface{} `json:"items"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// PutResourcePolicyResponse represents the PutResourcePolicyResponse schema from the OpenAPI specification
type PutResourcePolicyResponse struct {
}

// CreateListenerRequest represents the CreateListenerRequest schema from the OpenAPI specification
type CreateListenerRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Defaultaction interface{} `json:"defaultAction"`
	Name interface{} `json:"name"`
	Port interface{} `json:"port,omitempty"`
	Protocol interface{} `json:"protocol"`
}

// DeleteTargetGroupResponse represents the DeleteTargetGroupResponse schema from the OpenAPI specification
type DeleteTargetGroupResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ListServiceNetworkVpcAssociationsRequest represents the ListServiceNetworkVpcAssociationsRequest schema from the OpenAPI specification
type ListServiceNetworkVpcAssociationsRequest struct {
}

// DeregisterTargetsResponse represents the DeregisterTargetsResponse schema from the OpenAPI specification
type DeregisterTargetsResponse struct {
	Successful interface{} `json:"successful,omitempty"`
	Unsuccessful interface{} `json:"unsuccessful,omitempty"`
}

// ListServicesRequest represents the ListServicesRequest schema from the OpenAPI specification
type ListServicesRequest struct {
}

// UpdateServiceNetworkVpcAssociationRequest represents the UpdateServiceNetworkVpcAssociationRequest schema from the OpenAPI specification
type UpdateServiceNetworkVpcAssociationRequest struct {
	Securitygroupids interface{} `json:"securityGroupIds"`
}

// DeleteServiceNetworkServiceAssociationRequest represents the DeleteServiceNetworkServiceAssociationRequest schema from the OpenAPI specification
type DeleteServiceNetworkServiceAssociationRequest struct {
}

// UpdateAccessLogSubscriptionRequest represents the UpdateAccessLogSubscriptionRequest schema from the OpenAPI specification
type UpdateAccessLogSubscriptionRequest struct {
	Destinationarn interface{} `json:"destinationArn"`
}

// GetServiceRequest represents the GetServiceRequest schema from the OpenAPI specification
type GetServiceRequest struct {
}

// DnsEntry represents the DnsEntry schema from the OpenAPI specification
type DnsEntry struct {
	Domainname interface{} `json:"domainName,omitempty"`
	Hostedzoneid interface{} `json:"hostedZoneId,omitempty"`
}

// ListListenersResponse represents the ListListenersResponse schema from the OpenAPI specification
type ListListenersResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Items interface{} `json:"items"`
}

// UpdateTargetGroupRequest represents the UpdateTargetGroupRequest schema from the OpenAPI specification
type UpdateTargetGroupRequest struct {
	Healthcheck interface{} `json:"healthCheck"`
}

// DeleteRuleRequest represents the DeleteRuleRequest schema from the OpenAPI specification
type DeleteRuleRequest struct {
}

// HeaderMatch represents the HeaderMatch schema from the OpenAPI specification
type HeaderMatch struct {
	Casesensitive interface{} `json:"caseSensitive,omitempty"`
	Match interface{} `json:"match"`
	Name interface{} `json:"name"`
}

// DeleteTargetGroupRequest represents the DeleteTargetGroupRequest schema from the OpenAPI specification
type DeleteTargetGroupRequest struct {
}

// DeleteServiceNetworkServiceAssociationResponse represents the DeleteServiceNetworkServiceAssociationResponse schema from the OpenAPI specification
type DeleteServiceNetworkServiceAssociationResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// DeregisterTargetsRequest represents the DeregisterTargetsRequest schema from the OpenAPI specification
type DeregisterTargetsRequest struct {
	Targets interface{} `json:"targets"`
}

// GetAuthPolicyRequest represents the GetAuthPolicyRequest schema from the OpenAPI specification
type GetAuthPolicyRequest struct {
}

// RuleSummary represents the RuleSummary schema from the OpenAPI specification
type RuleSummary struct {
	Id interface{} `json:"id,omitempty"`
	Isdefault interface{} `json:"isDefault,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// DeleteListenerRequest represents the DeleteListenerRequest schema from the OpenAPI specification
type DeleteListenerRequest struct {
}

// DeleteServiceNetworkVpcAssociationRequest represents the DeleteServiceNetworkVpcAssociationRequest schema from the OpenAPI specification
type DeleteServiceNetworkVpcAssociationRequest struct {
}

// UpdateListenerResponse represents the UpdateListenerResponse schema from the OpenAPI specification
type UpdateListenerResponse struct {
	Protocol interface{} `json:"protocol,omitempty"`
	Servicearn interface{} `json:"serviceArn,omitempty"`
	Serviceid interface{} `json:"serviceId,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Defaultaction interface{} `json:"defaultAction,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Port interface{} `json:"port,omitempty"`
}

// HeaderMatchType represents the HeaderMatchType schema from the OpenAPI specification
type HeaderMatchType struct {
	Contains interface{} `json:"contains,omitempty"`
	Exact interface{} `json:"exact,omitempty"`
	Prefix interface{} `json:"prefix,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// GetAccessLogSubscriptionRequest represents the GetAccessLogSubscriptionRequest schema from the OpenAPI specification
type GetAccessLogSubscriptionRequest struct {
}

// PathMatch represents the PathMatch schema from the OpenAPI specification
type PathMatch struct {
	Casesensitive interface{} `json:"caseSensitive,omitempty"`
	Match interface{} `json:"match"`
}

// BatchUpdateRuleResponse represents the BatchUpdateRuleResponse schema from the OpenAPI specification
type BatchUpdateRuleResponse struct {
	Successful interface{} `json:"successful,omitempty"`
	Unsuccessful interface{} `json:"unsuccessful,omitempty"`
}

// ListenerSummary represents the ListenerSummary schema from the OpenAPI specification
type ListenerSummary struct {
	Name interface{} `json:"name,omitempty"`
	Port interface{} `json:"port,omitempty"`
	Protocol interface{} `json:"protocol,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
}

// TargetSummary represents the TargetSummary schema from the OpenAPI specification
type TargetSummary struct {
	Status interface{} `json:"status,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Port interface{} `json:"port,omitempty"`
	Reasoncode interface{} `json:"reasonCode,omitempty"`
}

// Target represents the Target schema from the OpenAPI specification
type Target struct {
	Id interface{} `json:"id"`
	Port interface{} `json:"port,omitempty"`
}

// UpdateServiceResponse represents the UpdateServiceResponse schema from the OpenAPI specification
type UpdateServiceResponse struct {
	Authtype interface{} `json:"authType,omitempty"`
	Certificatearn interface{} `json:"certificateArn,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// CreateServiceResponse represents the CreateServiceResponse schema from the OpenAPI specification
type CreateServiceResponse struct {
	Dnsentry interface{} `json:"dnsEntry,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Certificatearn interface{} `json:"certificateArn,omitempty"`
	Customdomainname interface{} `json:"customDomainName,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// GetAccessLogSubscriptionResponse represents the GetAccessLogSubscriptionResponse schema from the OpenAPI specification
type GetAccessLogSubscriptionResponse struct {
	Createdat interface{} `json:"createdAt"`
	Destinationarn interface{} `json:"destinationArn"`
	Id interface{} `json:"id"`
	Lastupdatedat interface{} `json:"lastUpdatedAt"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourceid interface{} `json:"resourceId"`
	Arn interface{} `json:"arn"`
}

// ListServiceNetworksRequest represents the ListServiceNetworksRequest schema from the OpenAPI specification
type ListServiceNetworksRequest struct {
}

// GetResourcePolicyResponse represents the GetResourcePolicyResponse schema from the OpenAPI specification
type GetResourcePolicyResponse struct {
	Policy interface{} `json:"policy,omitempty"`
}

// DeleteListenerResponse represents the DeleteListenerResponse schema from the OpenAPI specification
type DeleteListenerResponse struct {
}

// CreateAccessLogSubscriptionResponse represents the CreateAccessLogSubscriptionResponse schema from the OpenAPI specification
type CreateAccessLogSubscriptionResponse struct {
	Id interface{} `json:"id"`
	Resourcearn interface{} `json:"resourceArn"`
	Resourceid interface{} `json:"resourceId"`
	Arn interface{} `json:"arn"`
	Destinationarn interface{} `json:"destinationArn"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// DeleteAccessLogSubscriptionRequest represents the DeleteAccessLogSubscriptionRequest schema from the OpenAPI specification
type DeleteAccessLogSubscriptionRequest struct {
}

// DeleteAuthPolicyRequest represents the DeleteAuthPolicyRequest schema from the OpenAPI specification
type DeleteAuthPolicyRequest struct {
}

// ListServiceNetworksResponse represents the ListServiceNetworksResponse schema from the OpenAPI specification
type ListServiceNetworksResponse struct {
	Items interface{} `json:"items"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}
