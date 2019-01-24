package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createVirtualNetworkClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientAcceptLocalPeeringToken(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "AcceptLocalPeeringToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AcceptLocalPeeringToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "AcceptLocalPeeringToken", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "AcceptLocalPeeringToken")
	assert.NoError(t, err)

	type AcceptLocalPeeringTokenRequestInfo struct {
		ContainerId string
		Request     core.AcceptLocalPeeringTokenRequest
	}

	var requests []AcceptLocalPeeringTokenRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.AcceptLocalPeeringToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientAttachServiceId(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "AttachServiceId")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AttachServiceId is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "AttachServiceId", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "AttachServiceId")
	assert.NoError(t, err)

	type AttachServiceIdRequestInfo struct {
		ContainerId string
		Request     core.AttachServiceIdRequest
	}

	var requests []AttachServiceIdRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.AttachServiceId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientBulkAddVirtualCircuitPublicPrefixes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "BulkAddVirtualCircuitPublicPrefixes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkAddVirtualCircuitPublicPrefixes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "BulkAddVirtualCircuitPublicPrefixes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "BulkAddVirtualCircuitPublicPrefixes")
	assert.NoError(t, err)

	type BulkAddVirtualCircuitPublicPrefixesRequestInfo struct {
		ContainerId string
		Request     core.BulkAddVirtualCircuitPublicPrefixesRequest
	}

	var requests []BulkAddVirtualCircuitPublicPrefixesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			var response interface{}
			err := c.BulkAddVirtualCircuitPublicPrefixes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientBulkDeleteVirtualCircuitPublicPrefixes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "BulkDeleteVirtualCircuitPublicPrefixes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkDeleteVirtualCircuitPublicPrefixes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "BulkDeleteVirtualCircuitPublicPrefixes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "BulkDeleteVirtualCircuitPublicPrefixes")
	assert.NoError(t, err)

	type BulkDeleteVirtualCircuitPublicPrefixesRequestInfo struct {
		ContainerId string
		Request     core.BulkDeleteVirtualCircuitPublicPrefixesRequest
	}

	var requests []BulkDeleteVirtualCircuitPublicPrefixesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			var response interface{}
			err := c.BulkDeleteVirtualCircuitPublicPrefixes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientConnectLocalPeeringConnections(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ConnectLocalPeeringConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ConnectLocalPeeringConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ConnectLocalPeeringConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ConnectLocalPeeringConnections")
	assert.NoError(t, err)

	type ConnectLocalPeeringConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ConnectLocalPeeringConnectionsRequest
	}

	var requests []ConnectLocalPeeringConnectionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ConnectLocalPeeringConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientConnectLocalPeeringGateways(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ConnectLocalPeeringGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ConnectLocalPeeringGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ConnectLocalPeeringGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ConnectLocalPeeringGateways")
	assert.NoError(t, err)

	type ConnectLocalPeeringGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ConnectLocalPeeringGatewaysRequest
	}

	var requests []ConnectLocalPeeringGatewaysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ConnectLocalPeeringGateways(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientConnectRemotePeeringConnections(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ConnectRemotePeeringConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ConnectRemotePeeringConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ConnectRemotePeeringConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ConnectRemotePeeringConnections")
	assert.NoError(t, err)

	type ConnectRemotePeeringConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ConnectRemotePeeringConnectionsRequest
	}

	var requests []ConnectRemotePeeringConnectionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ConnectRemotePeeringConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateCpe(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateCpe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCpe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateCpe", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateCpe")
	assert.NoError(t, err)

	type CreateCpeRequestInfo struct {
		ContainerId string
		Request     core.CreateCpeRequest
	}

	var requests []CreateCpeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateCrossConnect(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateCrossConnect")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCrossConnect is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateCrossConnect", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateCrossConnect")
	assert.NoError(t, err)

	type CreateCrossConnectRequestInfo struct {
		ContainerId string
		Request     core.CreateCrossConnectRequest
	}

	var requests []CreateCrossConnectRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateCrossConnectGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateCrossConnectGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCrossConnectGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateCrossConnectGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateCrossConnectGroup")
	assert.NoError(t, err)

	type CreateCrossConnectGroupRequestInfo struct {
		ContainerId string
		Request     core.CreateCrossConnectGroupRequest
	}

	var requests []CreateCrossConnectGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateDhcpOptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateDhcpOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDhcpOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateDhcpOptions", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateDhcpOptions")
	assert.NoError(t, err)

	type CreateDhcpOptionsRequestInfo struct {
		ContainerId string
		Request     core.CreateDhcpOptionsRequest
	}

	var requests []CreateDhcpOptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateDrg(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateDrg")
	assert.NoError(t, err)

	type CreateDrgRequestInfo struct {
		ContainerId string
		Request     core.CreateDrgRequest
	}

	var requests []CreateDrgRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateDrgAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateDrgAttachment")
	assert.NoError(t, err)

	type CreateDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.CreateDrgAttachmentRequest
	}

	var requests []CreateDrgAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateFlowLogConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateFlowLogConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateFlowLogConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateFlowLogConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateFlowLogConfig")
	assert.NoError(t, err)

	type CreateFlowLogConfigRequestInfo struct {
		ContainerId string
		Request     core.CreateFlowLogConfigRequest
	}

	var requests []CreateFlowLogConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateFlowLogConfigAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateFlowLogConfigAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateFlowLogConfigAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateFlowLogConfigAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateFlowLogConfigAttachment")
	assert.NoError(t, err)

	type CreateFlowLogConfigAttachmentRequestInfo struct {
		ContainerId string
		Request     core.CreateFlowLogConfigAttachmentRequest
	}

	var requests []CreateFlowLogConfigAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateIPSecConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateIPSecConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIPSecConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateIPSecConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateIPSecConnection")
	assert.NoError(t, err)

	type CreateIPSecConnectionRequestInfo struct {
		ContainerId string
		Request     core.CreateIPSecConnectionRequest
	}

	var requests []CreateIPSecConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateInternetGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateInternetGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternetGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternetGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternetGateway")
	assert.NoError(t, err)

	type CreateInternetGatewayRequestInfo struct {
		ContainerId string
		Request     core.CreateInternetGatewayRequest
	}

	var requests []CreateInternetGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateLocalPeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateLocalPeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLocalPeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateLocalPeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateLocalPeeringConnection")
	assert.NoError(t, err)

	type CreateLocalPeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.CreateLocalPeeringConnectionRequest
	}

	var requests []CreateLocalPeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateLocalPeeringGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateLocalPeeringGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLocalPeeringGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateLocalPeeringGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateLocalPeeringGateway")
	assert.NoError(t, err)

	type CreateLocalPeeringGatewayRequestInfo struct {
		ContainerId string
		Request     core.CreateLocalPeeringGatewayRequest
	}

	var requests []CreateLocalPeeringGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateNatGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateNatGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateNatGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateNatGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateNatGateway")
	assert.NoError(t, err)

	type CreateNatGatewayRequestInfo struct {
		ContainerId string
		Request     core.CreateNatGatewayRequest
	}

	var requests []CreateNatGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreatePrivateIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreatePrivateIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePrivateIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreatePrivateIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreatePrivateIp")
	assert.NoError(t, err)

	type CreatePrivateIpRequestInfo struct {
		ContainerId string
		Request     core.CreatePrivateIpRequest
	}

	var requests []CreatePrivateIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreatePrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreatePublicIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreatePublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreatePublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreatePublicIp")
	assert.NoError(t, err)

	type CreatePublicIpRequestInfo struct {
		ContainerId string
		Request     core.CreatePublicIpRequest
	}

	var requests []CreatePublicIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreatePublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateRemotePeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateRemotePeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRemotePeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateRemotePeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateRemotePeeringConnection")
	assert.NoError(t, err)

	type CreateRemotePeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.CreateRemotePeeringConnectionRequest
	}

	var requests []CreateRemotePeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateRouteTable(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateRouteTable")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRouteTable is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateRouteTable", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateRouteTable")
	assert.NoError(t, err)

	type CreateRouteTableRequestInfo struct {
		ContainerId string
		Request     core.CreateRouteTableRequest
	}

	var requests []CreateRouteTableRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateSecurityList(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateSecurityList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSecurityList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateSecurityList", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateSecurityList")
	assert.NoError(t, err)

	type CreateSecurityListRequestInfo struct {
		ContainerId string
		Request     core.CreateSecurityListRequest
	}

	var requests []CreateSecurityListRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateServiceGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateServiceGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateServiceGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateServiceGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateServiceGateway")
	assert.NoError(t, err)

	type CreateServiceGatewayRequestInfo struct {
		ContainerId string
		Request     core.CreateServiceGatewayRequest
	}

	var requests []CreateServiceGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateSubnet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateSubnet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSubnet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateSubnet", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateSubnet")
	assert.NoError(t, err)

	type CreateSubnetRequestInfo struct {
		ContainerId string
		Request     core.CreateSubnetRequest
	}

	var requests []CreateSubnetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateVcn(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateVcn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVcn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateVcn", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateVcn")
	assert.NoError(t, err)

	type CreateVcnRequestInfo struct {
		ContainerId string
		Request     core.CreateVcnRequest
	}

	var requests []CreateVcnRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientCreateVirtualCircuit(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "CreateVirtualCircuit")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVirtualCircuit is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateVirtualCircuit", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateVirtualCircuit")
	assert.NoError(t, err)

	type CreateVirtualCircuitRequestInfo struct {
		ContainerId string
		Request     core.CreateVirtualCircuitRequest
	}

	var requests []CreateVirtualCircuitRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteCpe(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteCpe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCpe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteCpe", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteCpe")
	assert.NoError(t, err)

	type DeleteCpeRequestInfo struct {
		ContainerId string
		Request     core.DeleteCpeRequest
	}

	var requests []DeleteCpeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteCrossConnect(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteCrossConnect")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCrossConnect is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteCrossConnect", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteCrossConnect")
	assert.NoError(t, err)

	type DeleteCrossConnectRequestInfo struct {
		ContainerId string
		Request     core.DeleteCrossConnectRequest
	}

	var requests []DeleteCrossConnectRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteCrossConnectGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteCrossConnectGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCrossConnectGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteCrossConnectGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteCrossConnectGroup")
	assert.NoError(t, err)

	type DeleteCrossConnectGroupRequestInfo struct {
		ContainerId string
		Request     core.DeleteCrossConnectGroupRequest
	}

	var requests []DeleteCrossConnectGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteDhcpOptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteDhcpOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDhcpOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteDhcpOptions", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteDhcpOptions")
	assert.NoError(t, err)

	type DeleteDhcpOptionsRequestInfo struct {
		ContainerId string
		Request     core.DeleteDhcpOptionsRequest
	}

	var requests []DeleteDhcpOptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteDrg(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteDrg")
	assert.NoError(t, err)

	type DeleteDrgRequestInfo struct {
		ContainerId string
		Request     core.DeleteDrgRequest
	}

	var requests []DeleteDrgRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteDrgAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteDrgAttachment")
	assert.NoError(t, err)

	type DeleteDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.DeleteDrgAttachmentRequest
	}

	var requests []DeleteDrgAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteFlowLogConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteFlowLogConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteFlowLogConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteFlowLogConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteFlowLogConfig")
	assert.NoError(t, err)

	type DeleteFlowLogConfigRequestInfo struct {
		ContainerId string
		Request     core.DeleteFlowLogConfigRequest
	}

	var requests []DeleteFlowLogConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteFlowLogConfigAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteFlowLogConfigAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteFlowLogConfigAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteFlowLogConfigAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteFlowLogConfigAttachment")
	assert.NoError(t, err)

	type DeleteFlowLogConfigAttachmentRequestInfo struct {
		ContainerId string
		Request     core.DeleteFlowLogConfigAttachmentRequest
	}

	var requests []DeleteFlowLogConfigAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteIPSecConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteIPSecConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteIPSecConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteIPSecConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteIPSecConnection")
	assert.NoError(t, err)

	type DeleteIPSecConnectionRequestInfo struct {
		ContainerId string
		Request     core.DeleteIPSecConnectionRequest
	}

	var requests []DeleteIPSecConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteInternetGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteInternetGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternetGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternetGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternetGateway")
	assert.NoError(t, err)

	type DeleteInternetGatewayRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternetGatewayRequest
	}

	var requests []DeleteInternetGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteLocalPeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteLocalPeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLocalPeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteLocalPeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteLocalPeeringConnection")
	assert.NoError(t, err)

	type DeleteLocalPeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.DeleteLocalPeeringConnectionRequest
	}

	var requests []DeleteLocalPeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteLocalPeeringGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteLocalPeeringGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLocalPeeringGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteLocalPeeringGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteLocalPeeringGateway")
	assert.NoError(t, err)

	type DeleteLocalPeeringGatewayRequestInfo struct {
		ContainerId string
		Request     core.DeleteLocalPeeringGatewayRequest
	}

	var requests []DeleteLocalPeeringGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteNatGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteNatGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteNatGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteNatGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteNatGateway")
	assert.NoError(t, err)

	type DeleteNatGatewayRequestInfo struct {
		ContainerId string
		Request     core.DeleteNatGatewayRequest
	}

	var requests []DeleteNatGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeletePrivateIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeletePrivateIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePrivateIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeletePrivateIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeletePrivateIp")
	assert.NoError(t, err)

	type DeletePrivateIpRequestInfo struct {
		ContainerId string
		Request     core.DeletePrivateIpRequest
	}

	var requests []DeletePrivateIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeletePrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeletePublicIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeletePublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeletePublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeletePublicIp")
	assert.NoError(t, err)

	type DeletePublicIpRequestInfo struct {
		ContainerId string
		Request     core.DeletePublicIpRequest
	}

	var requests []DeletePublicIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeletePublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteRemotePeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteRemotePeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRemotePeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteRemotePeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteRemotePeeringConnection")
	assert.NoError(t, err)

	type DeleteRemotePeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.DeleteRemotePeeringConnectionRequest
	}

	var requests []DeleteRemotePeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteRouteTable(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteRouteTable")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRouteTable is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteRouteTable", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteRouteTable")
	assert.NoError(t, err)

	type DeleteRouteTableRequestInfo struct {
		ContainerId string
		Request     core.DeleteRouteTableRequest
	}

	var requests []DeleteRouteTableRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteSecurityList(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteSecurityList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSecurityList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteSecurityList", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteSecurityList")
	assert.NoError(t, err)

	type DeleteSecurityListRequestInfo struct {
		ContainerId string
		Request     core.DeleteSecurityListRequest
	}

	var requests []DeleteSecurityListRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteServiceGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteServiceGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteServiceGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteServiceGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteServiceGateway")
	assert.NoError(t, err)

	type DeleteServiceGatewayRequestInfo struct {
		ContainerId string
		Request     core.DeleteServiceGatewayRequest
	}

	var requests []DeleteServiceGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteSubnet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteSubnet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSubnet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteSubnet", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteSubnet")
	assert.NoError(t, err)

	type DeleteSubnetRequestInfo struct {
		ContainerId string
		Request     core.DeleteSubnetRequest
	}

	var requests []DeleteSubnetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteVcn(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteVcn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteVcn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteVcn", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteVcn")
	assert.NoError(t, err)

	type DeleteVcnRequestInfo struct {
		ContainerId string
		Request     core.DeleteVcnRequest
	}

	var requests []DeleteVcnRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteVirtualCircuit(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DeleteVirtualCircuit")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteVirtualCircuit is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteVirtualCircuit", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteVirtualCircuit")
	assert.NoError(t, err)

	type DeleteVirtualCircuitRequestInfo struct {
		ContainerId string
		Request     core.DeleteVirtualCircuitRequest
	}

	var requests []DeleteVirtualCircuitRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDetachServiceId(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DetachServiceId")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetachServiceId is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DetachServiceId", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DetachServiceId")
	assert.NoError(t, err)

	type DetachServiceIdRequestInfo struct {
		ContainerId string
		Request     core.DetachServiceIdRequest
	}

	var requests []DetachServiceIdRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DetachServiceId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGenerateLocalPeeringToken(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GenerateLocalPeeringToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateLocalPeeringToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GenerateLocalPeeringToken", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GenerateLocalPeeringToken")
	assert.NoError(t, err)

	type GenerateLocalPeeringTokenRequestInfo struct {
		ContainerId string
		Request     core.GenerateLocalPeeringTokenRequest
	}

	var requests []GenerateLocalPeeringTokenRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GenerateLocalPeeringToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetCpe(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetCpe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCpe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCpe", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCpe")
	assert.NoError(t, err)

	type GetCpeRequestInfo struct {
		ContainerId string
		Request     core.GetCpeRequest
	}

	var requests []GetCpeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetCrossConnect(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetCrossConnect")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCrossConnect is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCrossConnect", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCrossConnect")
	assert.NoError(t, err)

	type GetCrossConnectRequestInfo struct {
		ContainerId string
		Request     core.GetCrossConnectRequest
	}

	var requests []GetCrossConnectRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetCrossConnectGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetCrossConnectGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCrossConnectGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCrossConnectGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCrossConnectGroup")
	assert.NoError(t, err)

	type GetCrossConnectGroupRequestInfo struct {
		ContainerId string
		Request     core.GetCrossConnectGroupRequest
	}

	var requests []GetCrossConnectGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetCrossConnectLetterOfAuthority(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetCrossConnectLetterOfAuthority")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCrossConnectLetterOfAuthority is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCrossConnectLetterOfAuthority", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCrossConnectLetterOfAuthority")
	assert.NoError(t, err)

	type GetCrossConnectLetterOfAuthorityRequestInfo struct {
		ContainerId string
		Request     core.GetCrossConnectLetterOfAuthorityRequest
	}

	var requests []GetCrossConnectLetterOfAuthorityRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCrossConnectLetterOfAuthority(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetCrossConnectStatus(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetCrossConnectStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCrossConnectStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCrossConnectStatus", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCrossConnectStatus")
	assert.NoError(t, err)

	type GetCrossConnectStatusRequestInfo struct {
		ContainerId string
		Request     core.GetCrossConnectStatusRequest
	}

	var requests []GetCrossConnectStatusRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCrossConnectStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetDhcpOptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetDhcpOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDhcpOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetDhcpOptions", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetDhcpOptions")
	assert.NoError(t, err)

	type GetDhcpOptionsRequestInfo struct {
		ContainerId string
		Request     core.GetDhcpOptionsRequest
	}

	var requests []GetDhcpOptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetDrg(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetDrg")
	assert.NoError(t, err)

	type GetDrgRequestInfo struct {
		ContainerId string
		Request     core.GetDrgRequest
	}

	var requests []GetDrgRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetDrgAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetDrgAttachment")
	assert.NoError(t, err)

	type GetDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetDrgAttachmentRequest
	}

	var requests []GetDrgAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetFastConnectProviderService(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetFastConnectProviderService")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFastConnectProviderService is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetFastConnectProviderService", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetFastConnectProviderService")
	assert.NoError(t, err)

	type GetFastConnectProviderServiceRequestInfo struct {
		ContainerId string
		Request     core.GetFastConnectProviderServiceRequest
	}

	var requests []GetFastConnectProviderServiceRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetFastConnectProviderService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetFlowLogConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetFlowLogConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFlowLogConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetFlowLogConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetFlowLogConfig")
	assert.NoError(t, err)

	type GetFlowLogConfigRequestInfo struct {
		ContainerId string
		Request     core.GetFlowLogConfigRequest
	}

	var requests []GetFlowLogConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetFlowLogConfigAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetFlowLogConfigAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFlowLogConfigAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetFlowLogConfigAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetFlowLogConfigAttachment")
	assert.NoError(t, err)

	type GetFlowLogConfigAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetFlowLogConfigAttachmentRequest
	}

	var requests []GetFlowLogConfigAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetIPSecConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetIPSecConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIPSecConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIPSecConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIPSecConnection")
	assert.NoError(t, err)

	type GetIPSecConnectionRequestInfo struct {
		ContainerId string
		Request     core.GetIPSecConnectionRequest
	}

	var requests []GetIPSecConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetIPSecConnectionDeviceConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetIPSecConnectionDeviceConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIPSecConnectionDeviceConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIPSecConnectionDeviceConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIPSecConnectionDeviceConfig")
	assert.NoError(t, err)

	type GetIPSecConnectionDeviceConfigRequestInfo struct {
		ContainerId string
		Request     core.GetIPSecConnectionDeviceConfigRequest
	}

	var requests []GetIPSecConnectionDeviceConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetIPSecConnectionDeviceConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetIPSecConnectionDeviceStatus(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetIPSecConnectionDeviceStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIPSecConnectionDeviceStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIPSecConnectionDeviceStatus", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIPSecConnectionDeviceStatus")
	assert.NoError(t, err)

	type GetIPSecConnectionDeviceStatusRequestInfo struct {
		ContainerId string
		Request     core.GetIPSecConnectionDeviceStatusRequest
	}

	var requests []GetIPSecConnectionDeviceStatusRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetIPSecConnectionDeviceStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetInternetGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetInternetGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternetGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternetGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternetGateway")
	assert.NoError(t, err)

	type GetInternetGatewayRequestInfo struct {
		ContainerId string
		Request     core.GetInternetGatewayRequest
	}

	var requests []GetInternetGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetLocalPeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetLocalPeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLocalPeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetLocalPeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetLocalPeeringConnection")
	assert.NoError(t, err)

	type GetLocalPeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.GetLocalPeeringConnectionRequest
	}

	var requests []GetLocalPeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetLocalPeeringGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetLocalPeeringGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLocalPeeringGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetLocalPeeringGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetLocalPeeringGateway")
	assert.NoError(t, err)

	type GetLocalPeeringGatewayRequestInfo struct {
		ContainerId string
		Request     core.GetLocalPeeringGatewayRequest
	}

	var requests []GetLocalPeeringGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetNatGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetNatGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNatGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetNatGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetNatGateway")
	assert.NoError(t, err)

	type GetNatGatewayRequestInfo struct {
		ContainerId string
		Request     core.GetNatGatewayRequest
	}

	var requests []GetNatGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetPrivateIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetPrivateIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPrivateIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPrivateIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPrivateIp")
	assert.NoError(t, err)

	type GetPrivateIpRequestInfo struct {
		ContainerId string
		Request     core.GetPrivateIpRequest
	}

	var requests []GetPrivateIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetPublicIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetPublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPublicIp")
	assert.NoError(t, err)

	type GetPublicIpRequestInfo struct {
		ContainerId string
		Request     core.GetPublicIpRequest
	}

	var requests []GetPublicIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetPublicIpByIpAddress(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetPublicIpByIpAddress")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPublicIpByIpAddress is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPublicIpByIpAddress", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPublicIpByIpAddress")
	assert.NoError(t, err)

	type GetPublicIpByIpAddressRequestInfo struct {
		ContainerId string
		Request     core.GetPublicIpByIpAddressRequest
	}

	var requests []GetPublicIpByIpAddressRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPublicIpByIpAddress(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetPublicIpByPrivateIpId(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetPublicIpByPrivateIpId")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPublicIpByPrivateIpId is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPublicIpByPrivateIpId", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPublicIpByPrivateIpId")
	assert.NoError(t, err)

	type GetPublicIpByPrivateIpIdRequestInfo struct {
		ContainerId string
		Request     core.GetPublicIpByPrivateIpIdRequest
	}

	var requests []GetPublicIpByPrivateIpIdRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPublicIpByPrivateIpId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetRemotePeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetRemotePeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRemotePeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetRemotePeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetRemotePeeringConnection")
	assert.NoError(t, err)

	type GetRemotePeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.GetRemotePeeringConnectionRequest
	}

	var requests []GetRemotePeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetRouteTable(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetRouteTable")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRouteTable is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetRouteTable", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetRouteTable")
	assert.NoError(t, err)

	type GetRouteTableRequestInfo struct {
		ContainerId string
		Request     core.GetRouteTableRequest
	}

	var requests []GetRouteTableRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetSecurityList(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetSecurityList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSecurityList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetSecurityList", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetSecurityList")
	assert.NoError(t, err)

	type GetSecurityListRequestInfo struct {
		ContainerId string
		Request     core.GetSecurityListRequest
	}

	var requests []GetSecurityListRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetService(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetService")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetService is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetService", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetService")
	assert.NoError(t, err)

	type GetServiceRequestInfo struct {
		ContainerId string
		Request     core.GetServiceRequest
	}

	var requests []GetServiceRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetServiceGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetServiceGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetServiceGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetServiceGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetServiceGateway")
	assert.NoError(t, err)

	type GetServiceGatewayRequestInfo struct {
		ContainerId string
		Request     core.GetServiceGatewayRequest
	}

	var requests []GetServiceGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetSubnet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetSubnet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSubnet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetSubnet", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetSubnet")
	assert.NoError(t, err)

	type GetSubnetRequestInfo struct {
		ContainerId string
		Request     core.GetSubnetRequest
	}

	var requests []GetSubnetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetVcn(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetVcn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVcn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetVcn", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetVcn")
	assert.NoError(t, err)

	type GetVcnRequestInfo struct {
		ContainerId string
		Request     core.GetVcnRequest
	}

	var requests []GetVcnRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetVirtualCircuit(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetVirtualCircuit")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVirtualCircuit is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetVirtualCircuit", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetVirtualCircuit")
	assert.NoError(t, err)

	type GetVirtualCircuitRequestInfo struct {
		ContainerId string
		Request     core.GetVirtualCircuitRequest
	}

	var requests []GetVirtualCircuitRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetVnic(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "GetVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetVnic", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetVnic")
	assert.NoError(t, err)

	type GetVnicRequestInfo struct {
		ContainerId string
		Request     core.GetVnicRequest
	}

	var requests []GetVnicRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListAllowedPeerRegionsForRemotePeering(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListAllowedPeerRegionsForRemotePeering")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAllowedPeerRegionsForRemotePeering is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListAllowedPeerRegionsForRemotePeering", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListAllowedPeerRegionsForRemotePeering")
	assert.NoError(t, err)

	type ListAllowedPeerRegionsForRemotePeeringRequestInfo struct {
		ContainerId string
		Request     core.ListAllowedPeerRegionsForRemotePeeringRequest
	}

	var requests []ListAllowedPeerRegionsForRemotePeeringRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListAllowedPeerRegionsForRemotePeering(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListCpes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListCpes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCpes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCpes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCpes")
	assert.NoError(t, err)

	type ListCpesRequestInfo struct {
		ContainerId string
		Request     core.ListCpesRequest
	}

	var requests []ListCpesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListCpesRequest)
				return c.ListCpes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListCpesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListCpesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListCrossConnectGroups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListCrossConnectGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCrossConnectGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCrossConnectGroups", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCrossConnectGroups")
	assert.NoError(t, err)

	type ListCrossConnectGroupsRequestInfo struct {
		ContainerId string
		Request     core.ListCrossConnectGroupsRequest
	}

	var requests []ListCrossConnectGroupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListCrossConnectGroupsRequest)
				return c.ListCrossConnectGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListCrossConnectGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListCrossConnectGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListCrossConnectLocations(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListCrossConnectLocations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCrossConnectLocations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCrossConnectLocations", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCrossConnectLocations")
	assert.NoError(t, err)

	type ListCrossConnectLocationsRequestInfo struct {
		ContainerId string
		Request     core.ListCrossConnectLocationsRequest
	}

	var requests []ListCrossConnectLocationsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListCrossConnectLocationsRequest)
				return c.ListCrossConnectLocations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListCrossConnectLocationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListCrossConnectLocationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListCrossConnects(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListCrossConnects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCrossConnects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCrossConnects", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCrossConnects")
	assert.NoError(t, err)

	type ListCrossConnectsRequestInfo struct {
		ContainerId string
		Request     core.ListCrossConnectsRequest
	}

	var requests []ListCrossConnectsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListCrossConnectsRequest)
				return c.ListCrossConnects(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListCrossConnectsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListCrossConnectsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListCrossconnectPortSpeedShapes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListCrossconnectPortSpeedShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCrossconnectPortSpeedShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCrossconnectPortSpeedShapes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCrossconnectPortSpeedShapes")
	assert.NoError(t, err)

	type ListCrossconnectPortSpeedShapesRequestInfo struct {
		ContainerId string
		Request     core.ListCrossconnectPortSpeedShapesRequest
	}

	var requests []ListCrossconnectPortSpeedShapesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListCrossconnectPortSpeedShapesRequest)
				return c.ListCrossconnectPortSpeedShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListCrossconnectPortSpeedShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListCrossconnectPortSpeedShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListDhcpOptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListDhcpOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDhcpOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListDhcpOptions", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListDhcpOptions")
	assert.NoError(t, err)

	type ListDhcpOptionsRequestInfo struct {
		ContainerId string
		Request     core.ListDhcpOptionsRequest
	}

	var requests []ListDhcpOptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListDhcpOptionsRequest)
				return c.ListDhcpOptions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDhcpOptionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDhcpOptionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListDrgAttachments(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListDrgAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDrgAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListDrgAttachments", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListDrgAttachments")
	assert.NoError(t, err)

	type ListDrgAttachmentsRequestInfo struct {
		ContainerId string
		Request     core.ListDrgAttachmentsRequest
	}

	var requests []ListDrgAttachmentsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListDrgAttachmentsRequest)
				return c.ListDrgAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDrgAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDrgAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListDrgs(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListDrgs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDrgs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListDrgs", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListDrgs")
	assert.NoError(t, err)

	type ListDrgsRequestInfo struct {
		ContainerId string
		Request     core.ListDrgsRequest
	}

	var requests []ListDrgsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListDrgsRequest)
				return c.ListDrgs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDrgsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDrgsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListFastConnectProviderServices(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListFastConnectProviderServices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFastConnectProviderServices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListFastConnectProviderServices", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListFastConnectProviderServices")
	assert.NoError(t, err)

	type ListFastConnectProviderServicesRequestInfo struct {
		ContainerId string
		Request     core.ListFastConnectProviderServicesRequest
	}

	var requests []ListFastConnectProviderServicesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListFastConnectProviderServicesRequest)
				return c.ListFastConnectProviderServices(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListFastConnectProviderServicesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListFastConnectProviderServicesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListFastConnectProviderVirtualCircuitBandwidthShapes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListFastConnectProviderVirtualCircuitBandwidthShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFastConnectProviderVirtualCircuitBandwidthShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListFastConnectProviderVirtualCircuitBandwidthShapes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListFastConnectProviderVirtualCircuitBandwidthShapes")
	assert.NoError(t, err)

	type ListFastConnectProviderVirtualCircuitBandwidthShapesRequestInfo struct {
		ContainerId string
		Request     core.ListFastConnectProviderVirtualCircuitBandwidthShapesRequest
	}

	var requests []ListFastConnectProviderVirtualCircuitBandwidthShapesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListFastConnectProviderVirtualCircuitBandwidthShapesRequest)
				return c.ListFastConnectProviderVirtualCircuitBandwidthShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListFastConnectProviderVirtualCircuitBandwidthShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListFastConnectProviderVirtualCircuitBandwidthShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListFlowLogConfigAttachments(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListFlowLogConfigAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFlowLogConfigAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListFlowLogConfigAttachments", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListFlowLogConfigAttachments")
	assert.NoError(t, err)

	type ListFlowLogConfigAttachmentsRequestInfo struct {
		ContainerId string
		Request     core.ListFlowLogConfigAttachmentsRequest
	}

	var requests []ListFlowLogConfigAttachmentsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListFlowLogConfigAttachmentsRequest)
				return c.ListFlowLogConfigAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListFlowLogConfigAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListFlowLogConfigAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListFlowLogConfigs(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListFlowLogConfigs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFlowLogConfigs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListFlowLogConfigs", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListFlowLogConfigs")
	assert.NoError(t, err)

	type ListFlowLogConfigsRequestInfo struct {
		ContainerId string
		Request     core.ListFlowLogConfigsRequest
	}

	var requests []ListFlowLogConfigsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListFlowLogConfigsRequest)
				return c.ListFlowLogConfigs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListFlowLogConfigsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListFlowLogConfigsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListIPSecConnections(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListIPSecConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIPSecConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListIPSecConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListIPSecConnections")
	assert.NoError(t, err)

	type ListIPSecConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ListIPSecConnectionsRequest
	}

	var requests []ListIPSecConnectionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListIPSecConnectionsRequest)
				return c.ListIPSecConnections(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListIPSecConnectionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListIPSecConnectionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListInternetGateways(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListInternetGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternetGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternetGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternetGateways")
	assert.NoError(t, err)

	type ListInternetGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ListInternetGatewaysRequest
	}

	var requests []ListInternetGatewaysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternetGatewaysRequest)
				return c.ListInternetGateways(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternetGatewaysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternetGatewaysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListLocalPeeringConnections(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListLocalPeeringConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLocalPeeringConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListLocalPeeringConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListLocalPeeringConnections")
	assert.NoError(t, err)

	type ListLocalPeeringConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ListLocalPeeringConnectionsRequest
	}

	var requests []ListLocalPeeringConnectionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListLocalPeeringConnectionsRequest)
				return c.ListLocalPeeringConnections(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListLocalPeeringConnectionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListLocalPeeringConnectionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListLocalPeeringGateways(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListLocalPeeringGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLocalPeeringGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListLocalPeeringGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListLocalPeeringGateways")
	assert.NoError(t, err)

	type ListLocalPeeringGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ListLocalPeeringGatewaysRequest
	}

	var requests []ListLocalPeeringGatewaysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListLocalPeeringGatewaysRequest)
				return c.ListLocalPeeringGateways(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListLocalPeeringGatewaysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListLocalPeeringGatewaysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListNatGateways(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListNatGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNatGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListNatGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListNatGateways")
	assert.NoError(t, err)

	type ListNatGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ListNatGatewaysRequest
	}

	var requests []ListNatGatewaysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListNatGatewaysRequest)
				return c.ListNatGateways(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListNatGatewaysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListNatGatewaysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListPrivateIps(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListPrivateIps")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPrivateIps is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListPrivateIps", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListPrivateIps")
	assert.NoError(t, err)

	type ListPrivateIpsRequestInfo struct {
		ContainerId string
		Request     core.ListPrivateIpsRequest
	}

	var requests []ListPrivateIpsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListPrivateIpsRequest)
				return c.ListPrivateIps(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListPrivateIpsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListPrivateIpsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListPublicIps(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListPublicIps")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPublicIps is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListPublicIps", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListPublicIps")
	assert.NoError(t, err)

	type ListPublicIpsRequestInfo struct {
		ContainerId string
		Request     core.ListPublicIpsRequest
	}

	var requests []ListPublicIpsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListPublicIpsRequest)
				return c.ListPublicIps(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListPublicIpsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListPublicIpsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListRemotePeeringConnections(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListRemotePeeringConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRemotePeeringConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListRemotePeeringConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListRemotePeeringConnections")
	assert.NoError(t, err)

	type ListRemotePeeringConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ListRemotePeeringConnectionsRequest
	}

	var requests []ListRemotePeeringConnectionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListRemotePeeringConnectionsRequest)
				return c.ListRemotePeeringConnections(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListRemotePeeringConnectionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListRemotePeeringConnectionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListRouteTables(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListRouteTables")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRouteTables is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListRouteTables", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListRouteTables")
	assert.NoError(t, err)

	type ListRouteTablesRequestInfo struct {
		ContainerId string
		Request     core.ListRouteTablesRequest
	}

	var requests []ListRouteTablesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListRouteTablesRequest)
				return c.ListRouteTables(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListRouteTablesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListRouteTablesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListSecurityLists(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListSecurityLists")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSecurityLists is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListSecurityLists", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListSecurityLists")
	assert.NoError(t, err)

	type ListSecurityListsRequestInfo struct {
		ContainerId string
		Request     core.ListSecurityListsRequest
	}

	var requests []ListSecurityListsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListSecurityListsRequest)
				return c.ListSecurityLists(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListSecurityListsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListSecurityListsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListServiceGateways(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListServiceGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListServiceGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListServiceGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListServiceGateways")
	assert.NoError(t, err)

	type ListServiceGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ListServiceGatewaysRequest
	}

	var requests []ListServiceGatewaysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListServiceGatewaysRequest)
				return c.ListServiceGateways(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListServiceGatewaysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListServiceGatewaysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListServices(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListServices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListServices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListServices", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListServices")
	assert.NoError(t, err)

	type ListServicesRequestInfo struct {
		ContainerId string
		Request     core.ListServicesRequest
	}

	var requests []ListServicesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListServicesRequest)
				return c.ListServices(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListServicesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListServicesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListSubnets(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListSubnets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSubnets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListSubnets", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListSubnets")
	assert.NoError(t, err)

	type ListSubnetsRequestInfo struct {
		ContainerId string
		Request     core.ListSubnetsRequest
	}

	var requests []ListSubnetsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListSubnetsRequest)
				return c.ListSubnets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListSubnetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListSubnetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListVcns(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListVcns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVcns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListVcns", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListVcns")
	assert.NoError(t, err)

	type ListVcnsRequestInfo struct {
		ContainerId string
		Request     core.ListVcnsRequest
	}

	var requests []ListVcnsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListVcnsRequest)
				return c.ListVcns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVcnsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVcnsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListVirtualCircuitBandwidthShapes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListVirtualCircuitBandwidthShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVirtualCircuitBandwidthShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListVirtualCircuitBandwidthShapes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListVirtualCircuitBandwidthShapes")
	assert.NoError(t, err)

	type ListVirtualCircuitBandwidthShapesRequestInfo struct {
		ContainerId string
		Request     core.ListVirtualCircuitBandwidthShapesRequest
	}

	var requests []ListVirtualCircuitBandwidthShapesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListVirtualCircuitBandwidthShapesRequest)
				return c.ListVirtualCircuitBandwidthShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVirtualCircuitBandwidthShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVirtualCircuitBandwidthShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListVirtualCircuitPublicPrefixes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListVirtualCircuitPublicPrefixes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVirtualCircuitPublicPrefixes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListVirtualCircuitPublicPrefixes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListVirtualCircuitPublicPrefixes")
	assert.NoError(t, err)

	type ListVirtualCircuitPublicPrefixesRequestInfo struct {
		ContainerId string
		Request     core.ListVirtualCircuitPublicPrefixesRequest
	}

	var requests []ListVirtualCircuitPublicPrefixesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListVirtualCircuitPublicPrefixes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientListVirtualCircuits(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "ListVirtualCircuits")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVirtualCircuits is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListVirtualCircuits", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListVirtualCircuits")
	assert.NoError(t, err)

	type ListVirtualCircuitsRequestInfo struct {
		ContainerId string
		Request     core.ListVirtualCircuitsRequest
	}

	var requests []ListVirtualCircuitsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListVirtualCircuitsRequest)
				return c.ListVirtualCircuits(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVirtualCircuitsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVirtualCircuitsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateCpe(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateCpe")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCpe is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateCpe", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateCpe")
	assert.NoError(t, err)

	type UpdateCpeRequestInfo struct {
		ContainerId string
		Request     core.UpdateCpeRequest
	}

	var requests []UpdateCpeRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateCrossConnect(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateCrossConnect")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCrossConnect is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateCrossConnect", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateCrossConnect")
	assert.NoError(t, err)

	type UpdateCrossConnectRequestInfo struct {
		ContainerId string
		Request     core.UpdateCrossConnectRequest
	}

	var requests []UpdateCrossConnectRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateCrossConnectGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateCrossConnectGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCrossConnectGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateCrossConnectGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateCrossConnectGroup")
	assert.NoError(t, err)

	type UpdateCrossConnectGroupRequestInfo struct {
		ContainerId string
		Request     core.UpdateCrossConnectGroupRequest
	}

	var requests []UpdateCrossConnectGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateDhcpOptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateDhcpOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDhcpOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateDhcpOptions", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateDhcpOptions")
	assert.NoError(t, err)

	type UpdateDhcpOptionsRequestInfo struct {
		ContainerId string
		Request     core.UpdateDhcpOptionsRequest
	}

	var requests []UpdateDhcpOptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateDrg(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateDrg")
	assert.NoError(t, err)

	type UpdateDrgRequestInfo struct {
		ContainerId string
		Request     core.UpdateDrgRequest
	}

	var requests []UpdateDrgRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateDrgAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateDrgAttachment")
	assert.NoError(t, err)

	type UpdateDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.UpdateDrgAttachmentRequest
	}

	var requests []UpdateDrgAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateFlowLogConfig(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateFlowLogConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateFlowLogConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateFlowLogConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateFlowLogConfig")
	assert.NoError(t, err)

	type UpdateFlowLogConfigRequestInfo struct {
		ContainerId string
		Request     core.UpdateFlowLogConfigRequest
	}

	var requests []UpdateFlowLogConfigRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateFlowLogConfigAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateFlowLogConfigAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateFlowLogConfigAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateFlowLogConfigAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateFlowLogConfigAttachment")
	assert.NoError(t, err)

	type UpdateFlowLogConfigAttachmentRequestInfo struct {
		ContainerId string
		Request     core.UpdateFlowLogConfigAttachmentRequest
	}

	var requests []UpdateFlowLogConfigAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateIPSecConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateIPSecConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIPSecConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateIPSecConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateIPSecConnection")
	assert.NoError(t, err)

	type UpdateIPSecConnectionRequestInfo struct {
		ContainerId string
		Request     core.UpdateIPSecConnectionRequest
	}

	var requests []UpdateIPSecConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateInternetGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateInternetGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternetGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternetGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternetGateway")
	assert.NoError(t, err)

	type UpdateInternetGatewayRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternetGatewayRequest
	}

	var requests []UpdateInternetGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateLocalPeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateLocalPeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLocalPeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateLocalPeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateLocalPeeringConnection")
	assert.NoError(t, err)

	type UpdateLocalPeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.UpdateLocalPeeringConnectionRequest
	}

	var requests []UpdateLocalPeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateLocalPeeringGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateLocalPeeringGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLocalPeeringGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateLocalPeeringGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateLocalPeeringGateway")
	assert.NoError(t, err)

	type UpdateLocalPeeringGatewayRequestInfo struct {
		ContainerId string
		Request     core.UpdateLocalPeeringGatewayRequest
	}

	var requests []UpdateLocalPeeringGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateNatGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateNatGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNatGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateNatGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateNatGateway")
	assert.NoError(t, err)

	type UpdateNatGatewayRequestInfo struct {
		ContainerId string
		Request     core.UpdateNatGatewayRequest
	}

	var requests []UpdateNatGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdatePrivateIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdatePrivateIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePrivateIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdatePrivateIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdatePrivateIp")
	assert.NoError(t, err)

	type UpdatePrivateIpRequestInfo struct {
		ContainerId string
		Request     core.UpdatePrivateIpRequest
	}

	var requests []UpdatePrivateIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdatePrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdatePublicIp(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdatePublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdatePublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdatePublicIp")
	assert.NoError(t, err)

	type UpdatePublicIpRequestInfo struct {
		ContainerId string
		Request     core.UpdatePublicIpRequest
	}

	var requests []UpdatePublicIpRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdatePublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateRemotePeeringConnection(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateRemotePeeringConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRemotePeeringConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateRemotePeeringConnection", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateRemotePeeringConnection")
	assert.NoError(t, err)

	type UpdateRemotePeeringConnectionRequestInfo struct {
		ContainerId string
		Request     core.UpdateRemotePeeringConnectionRequest
	}

	var requests []UpdateRemotePeeringConnectionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateRouteTable(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateRouteTable")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRouteTable is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateRouteTable", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateRouteTable")
	assert.NoError(t, err)

	type UpdateRouteTableRequestInfo struct {
		ContainerId string
		Request     core.UpdateRouteTableRequest
	}

	var requests []UpdateRouteTableRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateSecurityList(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateSecurityList")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSecurityList is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateSecurityList", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateSecurityList")
	assert.NoError(t, err)

	type UpdateSecurityListRequestInfo struct {
		ContainerId string
		Request     core.UpdateSecurityListRequest
	}

	var requests []UpdateSecurityListRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateServiceGateway(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateServiceGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateServiceGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateServiceGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateServiceGateway")
	assert.NoError(t, err)

	type UpdateServiceGatewayRequestInfo struct {
		ContainerId string
		Request     core.UpdateServiceGatewayRequest
	}

	var requests []UpdateServiceGatewayRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateSubnet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateSubnet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSubnet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateSubnet", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateSubnet")
	assert.NoError(t, err)

	type UpdateSubnetRequestInfo struct {
		ContainerId string
		Request     core.UpdateSubnetRequest
	}

	var requests []UpdateSubnetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateVcn(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateVcn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVcn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateVcn", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateVcn")
	assert.NoError(t, err)

	type UpdateVcnRequestInfo struct {
		ContainerId string
		Request     core.UpdateVcnRequest
	}

	var requests []UpdateVcnRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateVirtualCircuit(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateVirtualCircuit")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVirtualCircuit is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateVirtualCircuit", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateVirtualCircuit")
	assert.NoError(t, err)

	type UpdateVirtualCircuitRequestInfo struct {
		ContainerId string
		Request     core.UpdateVirtualCircuitRequest
	}

	var requests []UpdateVirtualCircuitRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientUpdateVnic(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "UpdateVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateVnic", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateVnic")
	assert.NoError(t, err)

	type UpdateVnicRequestInfo struct {
		ContainerId string
		Request     core.UpdateVnicRequest
	}

	var requests []UpdateVnicRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
