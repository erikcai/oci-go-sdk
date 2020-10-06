package autotest

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"github.com/oracle/oci-go-sdk/v26/core"

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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientAcceptLocalPeeringToken(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AcceptLocalPeeringToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientAddNetworkSecurityGroupSecurityRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AddNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddNetworkSecurityGroupSecurityRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "AddNetworkSecurityGroupSecurityRules", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "AddNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)

	type AddNetworkSecurityGroupSecurityRulesRequestInfo struct {
		ContainerId string
		Request     core.AddNetworkSecurityGroupSecurityRulesRequest
	}

	var requests []AddNetworkSecurityGroupSecurityRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AddNetworkSecurityGroupSecurityRules(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientAddPublicIpPoolCapacity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AddPublicIpPoolCapacity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddPublicIpPoolCapacity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "AddPublicIpPoolCapacity", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "AddPublicIpPoolCapacity")
	assert.NoError(t, err)

	type AddPublicIpPoolCapacityRequestInfo struct {
		ContainerId string
		Request     core.AddPublicIpPoolCapacityRequest
	}

	var requests []AddPublicIpPoolCapacityRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AddPublicIpPoolCapacity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientAddVcnCidr(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AddVcnCidr")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddVcnCidr is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "AddVcnCidr", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "AddVcnCidr")
	assert.NoError(t, err)

	type AddVcnCidrRequestInfo struct {
		ContainerId string
		Request     core.AddVcnCidrRequest
	}

	var requests []AddVcnCidrRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AddVcnCidr(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientAdvertiseByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AdvertiseByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AdvertiseByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "AdvertiseByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "AdvertiseByoipRange")
	assert.NoError(t, err)

	type AdvertiseByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.AdvertiseByoipRangeRequest
	}

	var requests []AdvertiseByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AdvertiseByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientAttachServiceId(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.AttachServiceId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientBulkAddVirtualCircuitPublicPrefixes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.BulkAddVirtualCircuitPublicPrefixes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientBulkDeleteVirtualCircuitPublicPrefixes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.BulkDeleteVirtualCircuitPublicPrefixes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientChangeByoipRangeCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeByoipRangeCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeByoipRangeCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeByoipRangeCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeByoipRangeCompartment")
	assert.NoError(t, err)

	type ChangeByoipRangeCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeByoipRangeCompartmentRequest
	}

	var requests []ChangeByoipRangeCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeByoipRangeCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientChangeCpeCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeCpeCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCpeCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeCpeCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeCpeCompartment")
	assert.NoError(t, err)

	type ChangeCpeCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeCpeCompartmentRequest
	}

	var requests []ChangeCpeCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeCpeCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientChangeCrossConnectCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeCrossConnectCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCrossConnectCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeCrossConnectCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeCrossConnectCompartment")
	assert.NoError(t, err)

	type ChangeCrossConnectCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeCrossConnectCompartmentRequest
	}

	var requests []ChangeCrossConnectCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeCrossConnectCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientChangeCrossConnectGroupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeCrossConnectGroupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeCrossConnectGroupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeCrossConnectGroupCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeCrossConnectGroupCompartment")
	assert.NoError(t, err)

	type ChangeCrossConnectGroupCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeCrossConnectGroupCompartmentRequest
	}

	var requests []ChangeCrossConnectGroupCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeCrossConnectGroupCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeDhcpOptionsCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeDhcpOptionsCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDhcpOptionsCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeDhcpOptionsCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeDhcpOptionsCompartment")
	assert.NoError(t, err)

	type ChangeDhcpOptionsCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeDhcpOptionsCompartmentRequest
	}

	var requests []ChangeDhcpOptionsCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeDhcpOptionsCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeDrgCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeDrgCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDrgCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeDrgCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeDrgCompartment")
	assert.NoError(t, err)

	type ChangeDrgCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeDrgCompartmentRequest
	}

	var requests []ChangeDrgCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeDrgCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientChangeIPSecConnectionCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeIPSecConnectionCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeIPSecConnectionCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeIPSecConnectionCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeIPSecConnectionCompartment")
	assert.NoError(t, err)

	type ChangeIPSecConnectionCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeIPSecConnectionCompartmentRequest
	}

	var requests []ChangeIPSecConnectionCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeIPSecConnectionCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeInternalDrgCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeInternalDrgCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeInternalDrgCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeInternalDrgCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeInternalDrgCompartment")
	assert.NoError(t, err)

	type ChangeInternalDrgCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeInternalDrgCompartmentRequest
	}

	var requests []ChangeInternalDrgCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeInternalDrgCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeInternalGenericGatewayCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeInternalGenericGatewayCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeInternalGenericGatewayCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeInternalGenericGatewayCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeInternalGenericGatewayCompartment")
	assert.NoError(t, err)

	type ChangeInternalGenericGatewayCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeInternalGenericGatewayCompartmentRequest
	}

	var requests []ChangeInternalGenericGatewayCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeInternalGenericGatewayCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeInternetGatewayCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeInternetGatewayCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeInternetGatewayCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeInternetGatewayCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeInternetGatewayCompartment")
	assert.NoError(t, err)

	type ChangeInternetGatewayCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeInternetGatewayCompartmentRequest
	}

	var requests []ChangeInternetGatewayCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeInternetGatewayCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeLocalPeeringGatewayCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeLocalPeeringGatewayCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeLocalPeeringGatewayCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeLocalPeeringGatewayCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeLocalPeeringGatewayCompartment")
	assert.NoError(t, err)

	type ChangeLocalPeeringGatewayCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeLocalPeeringGatewayCompartmentRequest
	}

	var requests []ChangeLocalPeeringGatewayCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeLocalPeeringGatewayCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="pnp" email="elpaso_ops_us_grp@oracle.com" jiraProject="NAT" opsJiraProject="PNP"
func TestVirtualNetworkClientChangeNatGatewayCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeNatGatewayCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeNatGatewayCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeNatGatewayCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeNatGatewayCompartment")
	assert.NoError(t, err)

	type ChangeNatGatewayCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeNatGatewayCompartmentRequest
	}

	var requests []ChangeNatGatewayCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeNatGatewayCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeNetworkSecurityGroupCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeNetworkSecurityGroupCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeNetworkSecurityGroupCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeNetworkSecurityGroupCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeNetworkSecurityGroupCompartment")
	assert.NoError(t, err)

	type ChangeNetworkSecurityGroupCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeNetworkSecurityGroupCompartmentRequest
	}

	var requests []ChangeNetworkSecurityGroupCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeNetworkSecurityGroupCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientChangePrivateEndpointCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangePrivateEndpointCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangePrivateEndpointCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangePrivateEndpointCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangePrivateEndpointCompartment")
	assert.NoError(t, err)

	type ChangePrivateEndpointCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangePrivateEndpointCompartmentRequest
	}

	var requests []ChangePrivateEndpointCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangePrivateEndpointCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangePublicIpCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangePublicIpCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangePublicIpCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangePublicIpCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangePublicIpCompartment")
	assert.NoError(t, err)

	type ChangePublicIpCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangePublicIpCompartmentRequest
	}

	var requests []ChangePublicIpCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangePublicIpCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientChangePublicIpPoolCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangePublicIpPoolCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangePublicIpPoolCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangePublicIpPoolCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangePublicIpPoolCompartment")
	assert.NoError(t, err)

	type ChangePublicIpPoolCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangePublicIpPoolCompartmentRequest
	}

	var requests []ChangePublicIpPoolCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangePublicIpPoolCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientChangeRemotePeeringConnectionCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeRemotePeeringConnectionCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRemotePeeringConnectionCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeRemotePeeringConnectionCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeRemotePeeringConnectionCompartment")
	assert.NoError(t, err)

	type ChangeRemotePeeringConnectionCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeRemotePeeringConnectionCompartmentRequest
	}

	var requests []ChangeRemotePeeringConnectionCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeRemotePeeringConnectionCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeRouteTableCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeRouteTableCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeRouteTableCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeRouteTableCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeRouteTableCompartment")
	assert.NoError(t, err)

	type ChangeRouteTableCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeRouteTableCompartmentRequest
	}

	var requests []ChangeRouteTableCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeRouteTableCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeSecurityListCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeSecurityListCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSecurityListCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeSecurityListCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeSecurityListCompartment")
	assert.NoError(t, err)

	type ChangeSecurityListCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeSecurityListCompartmentRequest
	}

	var requests []ChangeSecurityListCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeSecurityListCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientChangeServiceGatewayCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeServiceGatewayCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeServiceGatewayCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeServiceGatewayCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeServiceGatewayCompartment")
	assert.NoError(t, err)

	type ChangeServiceGatewayCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeServiceGatewayCompartmentRequest
	}

	var requests []ChangeServiceGatewayCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeServiceGatewayCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeSubnetCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeSubnetCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeSubnetCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeSubnetCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeSubnetCompartment")
	assert.NoError(t, err)

	type ChangeSubnetCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeSubnetCompartmentRequest
	}

	var requests []ChangeSubnetCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeSubnetCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeVcnCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeVcnCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeVcnCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeVcnCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeVcnCompartment")
	assert.NoError(t, err)

	type ChangeVcnCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeVcnCompartmentRequest
	}

	var requests []ChangeVcnCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeVcnCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientChangeVirtualCircuitCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeVirtualCircuitCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeVirtualCircuitCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeVirtualCircuitCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeVirtualCircuitCompartment")
	assert.NoError(t, err)

	type ChangeVirtualCircuitCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeVirtualCircuitCompartmentRequest
	}

	var requests []ChangeVirtualCircuitCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeVirtualCircuitCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientChangeVlanCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeVlanCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeVlanCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ChangeVlanCompartment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ChangeVlanCompartment")
	assert.NoError(t, err)

	type ChangeVlanCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeVlanCompartmentRequest
	}

	var requests []ChangeVlanCompartmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ChangeVlanCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientConnectLocalPeeringConnections(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ConnectLocalPeeringConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientConnectLocalPeeringGateways(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ConnectLocalPeeringGateways(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientConnectRemotePeeringConnections(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ConnectRemotePeeringConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientCreateByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateByoipRange")
	assert.NoError(t, err)

	type CreateByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.CreateByoipRangeRequest
	}

	var requests []CreateByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientCreateCpe(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientCreateCrossConnect(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientCreateCrossConnectGroup(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateDhcpOptions(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateDrg(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientCreateEndpointService(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateEndpointService")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateEndpointService is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateEndpointService", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateEndpointService")
	assert.NoError(t, err)

	type CreateEndpointServiceRequestInfo struct {
		ContainerId string
		Request     core.CreateEndpointServiceRequest
	}

	var requests []CreateEndpointServiceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateEndpointService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateFlowLogConfig(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateFlowLogConfigAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientCreateIPSecConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalDnsRecord(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalDnsRecord")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalDnsRecord is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalDnsRecord", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalDnsRecord")
	assert.NoError(t, err)

	type CreateInternalDnsRecordRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalDnsRecordRequest
	}

	var requests []CreateInternalDnsRecordRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalDnsRecord(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalDrg(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalDrg")
	assert.NoError(t, err)

	type CreateInternalDrgRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalDrgRequest
	}

	var requests []CreateInternalDrgRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalDrgAttachment")
	assert.NoError(t, err)

	type CreateInternalDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalDrgAttachmentRequest
	}

	var requests []CreateInternalDrgAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalGenericGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalGenericGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalGenericGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalGenericGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalGenericGateway")
	assert.NoError(t, err)

	type CreateInternalGenericGatewayRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalGenericGatewayRequest
	}

	var requests []CreateInternalGenericGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalGenericGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalPublicIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalPublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalPublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalPublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalPublicIp")
	assert.NoError(t, err)

	type CreateInternalPublicIpRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalPublicIpRequest
	}

	var requests []CreateInternalPublicIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalPublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalVnic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalVnic", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalVnic")
	assert.NoError(t, err)

	type CreateInternalVnicRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalVnicRequest
	}

	var requests []CreateInternalVnicRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternalVnicAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInternalVnicAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInternalVnicAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateInternalVnicAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateInternalVnicAttachment")
	assert.NoError(t, err)

	type CreateInternalVnicAttachmentRequestInfo struct {
		ContainerId string
		Request     core.CreateInternalVnicAttachmentRequest
	}

	var requests []CreateInternalVnicAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternalVnicAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateInternetGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateIpv6(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateIpv6")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIpv6 is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateIpv6", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateIpv6")
	assert.NoError(t, err)

	type CreateIpv6RequestInfo struct {
		ContainerId string
		Request     core.CreateIpv6Request
	}

	var requests []CreateIpv6RequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateIpv6(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateLocalPeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateLocalPeeringGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="pnp" email="elpaso_ops_us_grp@oracle.com" jiraProject="NAT" opsJiraProject="PNP"
func TestVirtualNetworkClientCreateNatGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateNetworkSecurityGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateNetworkSecurityGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateNetworkSecurityGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateNetworkSecurityGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateNetworkSecurityGroup")
	assert.NoError(t, err)

	type CreateNetworkSecurityGroupRequestInfo struct {
		ContainerId string
		Request     core.CreateNetworkSecurityGroupRequest
	}

	var requests []CreateNetworkSecurityGroupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateNetworkSecurityGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientCreatePrivateAccessGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreatePrivateAccessGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePrivateAccessGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreatePrivateAccessGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreatePrivateAccessGateway")
	assert.NoError(t, err)

	type CreatePrivateAccessGatewayRequestInfo struct {
		ContainerId string
		Request     core.CreatePrivateAccessGatewayRequest
	}

	var requests []CreatePrivateAccessGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePrivateAccessGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientCreatePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreatePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreatePrivateEndpoint", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreatePrivateEndpoint")
	assert.NoError(t, err)

	type CreatePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     core.CreatePrivateEndpointRequest
	}

	var requests []CreatePrivateEndpointRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreatePrivateIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreatePrivateIpNextHop(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreatePrivateIpNextHop")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePrivateIpNextHop is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreatePrivateIpNextHop", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreatePrivateIpNextHop")
	assert.NoError(t, err)

	type CreatePrivateIpNextHopRequestInfo struct {
		ContainerId string
		Request     core.CreatePrivateIpNextHopRequest
	}

	var requests []CreatePrivateIpNextHopRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePrivateIpNextHop(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreatePublicIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientCreatePublicIpPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreatePublicIpPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePublicIpPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreatePublicIpPool", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreatePublicIpPool")
	assert.NoError(t, err)

	type CreatePublicIpPoolRequestInfo struct {
		ContainerId string
		Request     core.CreatePublicIpPoolRequest
	}

	var requests []CreatePublicIpPoolRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreatePublicIpPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientCreateRemotePeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateRouteTable(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateSecurityList(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientCreateServiceGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateSubnet(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateVcn(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientCreateVirtualCircuit(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateVlan(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateVlan")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVlan is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateVlan", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateVlan")
	assert.NoError(t, err)

	type CreateVlanRequestInfo struct {
		ContainerId string
		Request     core.CreateVlanRequest
	}

	var requests []CreateVlanRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateVlan(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientCreateVnicWorker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateVnicWorker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVnicWorker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "CreateVnicWorker", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "CreateVnicWorker")
	assert.NoError(t, err)

	type CreateVnicWorkerRequestInfo struct {
		ContainerId string
		Request     core.CreateVnicWorkerRequest
	}

	var requests []CreateVnicWorkerRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateVnicWorker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientDeleteByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteByoipRange")
	assert.NoError(t, err)

	type DeleteByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.DeleteByoipRangeRequest
	}

	var requests []DeleteByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientDeleteCpe(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientDeleteCrossConnect(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientDeleteCrossConnectGroup(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteDhcpOptions(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteDrg(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientDeleteEndpointService(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteEndpointService")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteEndpointService is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteEndpointService", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteEndpointService")
	assert.NoError(t, err)

	type DeleteEndpointServiceRequestInfo struct {
		ContainerId string
		Request     core.DeleteEndpointServiceRequest
	}

	var requests []DeleteEndpointServiceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteEndpointService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteFlowLogConfig(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteFlowLogConfigAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientDeleteIPSecConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalDnsRecord(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalDnsRecord")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalDnsRecord is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalDnsRecord", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalDnsRecord")
	assert.NoError(t, err)

	type DeleteInternalDnsRecordRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalDnsRecordRequest
	}

	var requests []DeleteInternalDnsRecordRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalDnsRecord(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalDrg(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalDrg")
	assert.NoError(t, err)

	type DeleteInternalDrgRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalDrgRequest
	}

	var requests []DeleteInternalDrgRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalDrgAttachment")
	assert.NoError(t, err)

	type DeleteInternalDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalDrgAttachmentRequest
	}

	var requests []DeleteInternalDrgAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalGenericGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalGenericGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalGenericGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalGenericGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalGenericGateway")
	assert.NoError(t, err)

	type DeleteInternalGenericGatewayRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalGenericGatewayRequest
	}

	var requests []DeleteInternalGenericGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalGenericGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalPublicIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalPublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalPublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalPublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalPublicIp")
	assert.NoError(t, err)

	type DeleteInternalPublicIpRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalPublicIpRequest
	}

	var requests []DeleteInternalPublicIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalPublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalVnic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalVnic", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalVnic")
	assert.NoError(t, err)

	type DeleteInternalVnicRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalVnicRequest
	}

	var requests []DeleteInternalVnicRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternalVnicAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInternalVnicAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInternalVnicAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteInternalVnicAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteInternalVnicAttachment")
	assert.NoError(t, err)

	type DeleteInternalVnicAttachmentRequestInfo struct {
		ContainerId string
		Request     core.DeleteInternalVnicAttachmentRequest
	}

	var requests []DeleteInternalVnicAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternalVnicAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteInternetGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteIpv6(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteIpv6")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteIpv6 is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteIpv6", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteIpv6")
	assert.NoError(t, err)

	type DeleteIpv6RequestInfo struct {
		ContainerId string
		Request     core.DeleteIpv6Request
	}

	var requests []DeleteIpv6RequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteIpv6(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteLocalPeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteLocalPeeringGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="pnp" email="elpaso_ops_us_grp@oracle.com" jiraProject="NAT" opsJiraProject="PNP"
func TestVirtualNetworkClientDeleteNatGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteNetworkSecurityGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteNetworkSecurityGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteNetworkSecurityGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteNetworkSecurityGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteNetworkSecurityGroup")
	assert.NoError(t, err)

	type DeleteNetworkSecurityGroupRequestInfo struct {
		ContainerId string
		Request     core.DeleteNetworkSecurityGroupRequest
	}

	var requests []DeleteNetworkSecurityGroupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteNetworkSecurityGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientDeletePrivateAccessGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeletePrivateAccessGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePrivateAccessGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeletePrivateAccessGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeletePrivateAccessGateway")
	assert.NoError(t, err)

	type DeletePrivateAccessGatewayRequestInfo struct {
		ContainerId string
		Request     core.DeletePrivateAccessGatewayRequest
	}

	var requests []DeletePrivateAccessGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePrivateAccessGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientDeletePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeletePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeletePrivateEndpoint", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeletePrivateEndpoint")
	assert.NoError(t, err)

	type DeletePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     core.DeletePrivateEndpointRequest
	}

	var requests []DeletePrivateEndpointRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeletePrivateIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeletePrivateIpNextHop(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeletePrivateIpNextHop")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePrivateIpNextHop is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeletePrivateIpNextHop", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeletePrivateIpNextHop")
	assert.NoError(t, err)

	type DeletePrivateIpNextHopRequestInfo struct {
		ContainerId string
		Request     core.DeletePrivateIpNextHopRequest
	}

	var requests []DeletePrivateIpNextHopRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePrivateIpNextHop(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeletePublicIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientDeletePublicIpPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeletePublicIpPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePublicIpPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeletePublicIpPool", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeletePublicIpPool")
	assert.NoError(t, err)

	type DeletePublicIpPoolRequestInfo struct {
		ContainerId string
		Request     core.DeletePublicIpPoolRequest
	}

	var requests []DeletePublicIpPoolRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeletePublicIpPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientDeleteRemotePeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientDeleteReverseConnectionNatIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteReverseConnectionNatIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteReverseConnectionNatIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteReverseConnectionNatIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteReverseConnectionNatIp")
	assert.NoError(t, err)

	type DeleteReverseConnectionNatIpRequestInfo struct {
		ContainerId string
		Request     core.DeleteReverseConnectionNatIpRequest
	}

	var requests []DeleteReverseConnectionNatIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteReverseConnectionNatIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteRouteTable(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteSecurityList(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientDeleteServiceGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteSubnet(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteVcn(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientDeleteVirtualCircuit(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteVlan(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteVlan")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteVlan is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteVlan", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteVlan")
	assert.NoError(t, err)

	type DeleteVlanRequestInfo struct {
		ContainerId string
		Request     core.DeleteVlanRequest
	}

	var requests []DeleteVlanRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteVlan(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDeleteVnicWorker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteVnicWorker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteVnicWorker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DeleteVnicWorker", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DeleteVnicWorker")
	assert.NoError(t, err)

	type DeleteVnicWorkerRequestInfo struct {
		ContainerId string
		Request     core.DeleteVnicWorkerRequest
	}

	var requests []DeleteVnicWorkerRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DeleteVnicWorker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientDetachServiceId(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DetachServiceId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientDisableReverseConnections(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DisableReverseConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DisableReverseConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DisableReverseConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DisableReverseConnections")
	assert.NoError(t, err)

	type DisableReverseConnectionsRequestInfo struct {
		ContainerId string
		Request     core.DisableReverseConnectionsRequest
	}

	var requests []DisableReverseConnectionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DisableReverseConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientDisableVnicWorker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DisableVnicWorker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DisableVnicWorker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "DisableVnicWorker", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "DisableVnicWorker")
	assert.NoError(t, err)

	type DisableVnicWorkerRequestInfo struct {
		ContainerId string
		Request     core.DisableVnicWorkerRequest
	}

	var requests []DisableVnicWorkerRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.DisableVnicWorker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientEnableReverseConnections(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "EnableReverseConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EnableReverseConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "EnableReverseConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "EnableReverseConnections")
	assert.NoError(t, err)

	type EnableReverseConnectionsRequestInfo struct {
		ContainerId string
		Request     core.EnableReverseConnectionsRequest
	}

	var requests []EnableReverseConnectionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.EnableReverseConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientEnableVnicWorker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "EnableVnicWorker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EnableVnicWorker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "EnableVnicWorker", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "EnableVnicWorker")
	assert.NoError(t, err)

	type EnableVnicWorkerRequestInfo struct {
		ContainerId string
		Request     core.EnableVnicWorkerRequest
	}

	var requests []EnableVnicWorkerRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.EnableVnicWorker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGenerateLocalPeeringToken(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GenerateLocalPeeringToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientGetByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetByoipRange")
	assert.NoError(t, err)

	type GetByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.GetByoipRangeRequest
	}

	var requests []GetByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCpe(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCpeDeviceConfigContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetCpeDeviceConfigContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCpeDeviceConfigContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCpeDeviceConfigContent", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCpeDeviceConfigContent")
	assert.NoError(t, err)

	type GetCpeDeviceConfigContentRequestInfo struct {
		ContainerId string
		Request     core.GetCpeDeviceConfigContentRequest
	}

	var requests []GetCpeDeviceConfigContentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCpeDeviceConfigContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCpeDeviceShape(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetCpeDeviceShape")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCpeDeviceShape is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCpeDeviceShape", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCpeDeviceShape")
	assert.NoError(t, err)

	type GetCpeDeviceShapeRequestInfo struct {
		ContainerId string
		Request     core.GetCpeDeviceShapeRequest
	}

	var requests []GetCpeDeviceShapeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCpeDeviceShape(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetCreateReverseConnectionNatIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetCreateReverseConnectionNatIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCreateReverseConnectionNatIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetCreateReverseConnectionNatIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetCreateReverseConnectionNatIp")
	assert.NoError(t, err)

	type GetCreateReverseConnectionNatIpRequestInfo struct {
		ContainerId string
		Request     core.GetCreateReverseConnectionNatIpRequest
	}

	var requests []GetCreateReverseConnectionNatIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCreateReverseConnectionNatIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCrossConnect(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCrossConnectGroup(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCrossConnectLetterOfAuthority(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCrossConnectLetterOfAuthority(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetCrossConnectStatus(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetCrossConnectStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetDhcpOptions(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetDrg(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetDrgRedundancyStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetDrgRedundancyStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDrgRedundancyStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetDrgRedundancyStatus", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetDrgRedundancyStatus")
	assert.NoError(t, err)

	type GetDrgRedundancyStatusRequestInfo struct {
		ContainerId string
		Request     core.GetDrgRedundancyStatusRequest
	}

	var requests []GetDrgRedundancyStatusRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetDrgRedundancyStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetEndpointService(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetEndpointService")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetEndpointService is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetEndpointService", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetEndpointService")
	assert.NoError(t, err)

	type GetEndpointServiceRequestInfo struct {
		ContainerId string
		Request     core.GetEndpointServiceRequest
	}

	var requests []GetEndpointServiceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetEndpointService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetEndpointServiceNextHop(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetEndpointServiceNextHop")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetEndpointServiceNextHop is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetEndpointServiceNextHop", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetEndpointServiceNextHop")
	assert.NoError(t, err)

	type GetEndpointServiceNextHopRequestInfo struct {
		ContainerId string
		Request     core.GetEndpointServiceNextHopRequest
	}

	var requests []GetEndpointServiceNextHopRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetEndpointServiceNextHop(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetFastConnectProviderService(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetFastConnectProviderService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetFastConnectProviderServiceKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetFastConnectProviderServiceKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFastConnectProviderServiceKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetFastConnectProviderServiceKey", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetFastConnectProviderServiceKey")
	assert.NoError(t, err)

	type GetFastConnectProviderServiceKeyRequestInfo struct {
		ContainerId string
		Request     core.GetFastConnectProviderServiceKeyRequest
	}

	var requests []GetFastConnectProviderServiceKeyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetFastConnectProviderServiceKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetFlowLogConfig(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetFlowLogConfigAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetIPSecConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetIPSecConnectionDeviceConfig(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIPSecConnectionDeviceConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetIPSecConnectionDeviceStatus(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIPSecConnectionDeviceStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetIPSecConnectionTunnel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetIPSecConnectionTunnel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIPSecConnectionTunnel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIPSecConnectionTunnel", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIPSecConnectionTunnel")
	assert.NoError(t, err)

	type GetIPSecConnectionTunnelRequestInfo struct {
		ContainerId string
		Request     core.GetIPSecConnectionTunnelRequest
	}

	var requests []GetIPSecConnectionTunnelRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIPSecConnectionTunnel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetIPSecConnectionTunnelSharedSecret(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetIPSecConnectionTunnelSharedSecret")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIPSecConnectionTunnelSharedSecret is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIPSecConnectionTunnelSharedSecret", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIPSecConnectionTunnelSharedSecret")
	assert.NoError(t, err)

	type GetIPSecConnectionTunnelSharedSecretRequestInfo struct {
		ContainerId string
		Request     core.GetIPSecConnectionTunnelSharedSecretRequest
	}

	var requests []GetIPSecConnectionTunnelSharedSecretRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIPSecConnectionTunnelSharedSecret(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalDnsRecord(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalDnsRecord")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalDnsRecord is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalDnsRecord", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalDnsRecord")
	assert.NoError(t, err)

	type GetInternalDnsRecordRequestInfo struct {
		ContainerId string
		Request     core.GetInternalDnsRecordRequest
	}

	var requests []GetInternalDnsRecordRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalDnsRecord(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalDrg(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalDrg")
	assert.NoError(t, err)

	type GetInternalDrgRequestInfo struct {
		ContainerId string
		Request     core.GetInternalDrgRequest
	}

	var requests []GetInternalDrgRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalDrgAttachment")
	assert.NoError(t, err)

	type GetInternalDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetInternalDrgAttachmentRequest
	}

	var requests []GetInternalDrgAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalGenericGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalGenericGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalGenericGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalGenericGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalGenericGateway")
	assert.NoError(t, err)

	type GetInternalGenericGatewayRequestInfo struct {
		ContainerId string
		Request     core.GetInternalGenericGatewayRequest
	}

	var requests []GetInternalGenericGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalGenericGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalGenericGatewayByGatewayId(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalGenericGatewayByGatewayId")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalGenericGatewayByGatewayId is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalGenericGatewayByGatewayId", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalGenericGatewayByGatewayId")
	assert.NoError(t, err)

	type GetInternalGenericGatewayByGatewayIdRequestInfo struct {
		ContainerId string
		Request     core.GetInternalGenericGatewayByGatewayIdRequest
	}

	var requests []GetInternalGenericGatewayByGatewayIdRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalGenericGatewayByGatewayId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalPublicIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalPublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalPublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalPublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalPublicIp")
	assert.NoError(t, err)

	type GetInternalPublicIpRequestInfo struct {
		ContainerId string
		Request     core.GetInternalPublicIpRequest
	}

	var requests []GetInternalPublicIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalPublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalSubnet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalSubnet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalSubnet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalSubnet", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalSubnet")
	assert.NoError(t, err)

	type GetInternalSubnetRequestInfo struct {
		ContainerId string
		Request     core.GetInternalSubnetRequest
	}

	var requests []GetInternalSubnetRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalVnic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalVnic", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalVnic")
	assert.NoError(t, err)

	type GetInternalVnicRequestInfo struct {
		ContainerId string
		Request     core.GetInternalVnicRequest
	}

	var requests []GetInternalVnicRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternalVnicAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInternalVnicAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInternalVnicAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetInternalVnicAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetInternalVnicAttachment")
	assert.NoError(t, err)

	type GetInternalVnicAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetInternalVnicAttachmentRequest
	}

	var requests []GetInternalVnicAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternalVnicAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetInternetGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetIpsecCpeDeviceConfigContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetIpsecCpeDeviceConfigContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIpsecCpeDeviceConfigContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIpsecCpeDeviceConfigContent", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIpsecCpeDeviceConfigContent")
	assert.NoError(t, err)

	type GetIpsecCpeDeviceConfigContentRequestInfo struct {
		ContainerId string
		Request     core.GetIpsecCpeDeviceConfigContentRequest
	}

	var requests []GetIpsecCpeDeviceConfigContentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIpsecCpeDeviceConfigContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetIpv6(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetIpv6")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIpv6 is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetIpv6", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetIpv6")
	assert.NoError(t, err)

	type GetIpv6RequestInfo struct {
		ContainerId string
		Request     core.GetIpv6Request
	}

	var requests []GetIpv6RequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetIpv6(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetLocalPeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetLocalPeeringGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="pnp" email="elpaso_ops_us_grp@oracle.com" jiraProject="NAT" opsJiraProject="PNP"
func TestVirtualNetworkClientGetNatGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetNetworkSecurityGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetNetworkSecurityGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNetworkSecurityGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetNetworkSecurityGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetNetworkSecurityGroup")
	assert.NoError(t, err)

	type GetNetworkSecurityGroupRequestInfo struct {
		ContainerId string
		Request     core.GetNetworkSecurityGroupRequest
	}

	var requests []GetNetworkSecurityGroupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetNetworkSecurityGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetPrivateAccessGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetPrivateAccessGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPrivateAccessGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPrivateAccessGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPrivateAccessGateway")
	assert.NoError(t, err)

	type GetPrivateAccessGatewayRequestInfo struct {
		ContainerId string
		Request     core.GetPrivateAccessGatewayRequest
	}

	var requests []GetPrivateAccessGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPrivateAccessGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetPrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetPrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPrivateEndpoint", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPrivateEndpoint")
	assert.NoError(t, err)

	type GetPrivateEndpointRequestInfo struct {
		ContainerId string
		Request     core.GetPrivateEndpointRequest
	}

	var requests []GetPrivateEndpointRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetPrivateEndpointAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetPrivateEndpointAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPrivateEndpointAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPrivateEndpointAssociation", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPrivateEndpointAssociation")
	assert.NoError(t, err)

	type GetPrivateEndpointAssociationRequestInfo struct {
		ContainerId string
		Request     core.GetPrivateEndpointAssociationRequest
	}

	var requests []GetPrivateEndpointAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPrivateEndpointAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetPrivateIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetPrivateIpNextHop(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetPrivateIpNextHop")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPrivateIpNextHop is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPrivateIpNextHop", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPrivateIpNextHop")
	assert.NoError(t, err)

	type GetPrivateIpNextHopRequestInfo struct {
		ContainerId string
		Request     core.GetPrivateIpNextHopRequest
	}

	var requests []GetPrivateIpNextHopRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPrivateIpNextHop(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetPublicIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetPublicIpByIpAddress(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPublicIpByIpAddress(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetPublicIpByPrivateIpId(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPublicIpByPrivateIpId(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientGetPublicIpPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetPublicIpPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPublicIpPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetPublicIpPool", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetPublicIpPool")
	assert.NoError(t, err)

	type GetPublicIpPoolRequestInfo struct {
		ContainerId string
		Request     core.GetPublicIpPoolRequest
	}

	var requests []GetPublicIpPoolRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetPublicIpPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetRemotePeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetReverseConnectionNatIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetReverseConnectionNatIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetReverseConnectionNatIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetReverseConnectionNatIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetReverseConnectionNatIp")
	assert.NoError(t, err)

	type GetReverseConnectionNatIpRequestInfo struct {
		ContainerId string
		Request     core.GetReverseConnectionNatIpRequest
	}

	var requests []GetReverseConnectionNatIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetReverseConnectionNatIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetRouteTable(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetSecurityList(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetService(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientGetServiceGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetSubnet(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetTunnelCpeDeviceConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetTunnelCpeDeviceConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTunnelCpeDeviceConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetTunnelCpeDeviceConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetTunnelCpeDeviceConfig")
	assert.NoError(t, err)

	type GetTunnelCpeDeviceConfigRequestInfo struct {
		ContainerId string
		Request     core.GetTunnelCpeDeviceConfigRequest
	}

	var requests []GetTunnelCpeDeviceConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetTunnelCpeDeviceConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetTunnelCpeDeviceConfigContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetTunnelCpeDeviceConfigContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTunnelCpeDeviceConfigContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetTunnelCpeDeviceConfigContent", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetTunnelCpeDeviceConfigContent")
	assert.NoError(t, err)

	type GetTunnelCpeDeviceConfigContentRequestInfo struct {
		ContainerId string
		Request     core.GetTunnelCpeDeviceConfigContentRequest
	}

	var requests []GetTunnelCpeDeviceConfigContentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetTunnelCpeDeviceConfigContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetVcn(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestVirtualNetworkClientGetVcnDnsResolverAssociation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetVcnDnsResolverAssociation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVcnDnsResolverAssociation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetVcnDnsResolverAssociation", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetVcnDnsResolverAssociation")
	assert.NoError(t, err)

	type GetVcnDnsResolverAssociationRequestInfo struct {
		ContainerId string
		Request     core.GetVcnDnsResolverAssociationRequest
	}

	var requests []GetVcnDnsResolverAssociationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVcnDnsResolverAssociation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientGetVirtualCircuit(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetVlan(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetVlan")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVlan is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetVlan", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetVlan")
	assert.NoError(t, err)

	type GetVlanRequestInfo struct {
		ContainerId string
		Request     core.GetVlanRequest
	}

	var requests []GetVlanRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVlan(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetVnic(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientGetVnicWorker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetVnicWorker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVnicWorker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "GetVnicWorker", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "GetVnicWorker")
	assert.NoError(t, err)

	type GetVnicWorkerRequestInfo struct {
		ContainerId string
		Request     core.GetVnicWorkerRequest
	}

	var requests []GetVnicWorkerRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.GetVnicWorker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListAllowedPeerRegionsForRemotePeering(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ListAllowedPeerRegionsForRemotePeering(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientListByoipAllocatedRanges(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListByoipAllocatedRanges")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListByoipAllocatedRanges is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListByoipAllocatedRanges", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListByoipAllocatedRanges")
	assert.NoError(t, err)

	type ListByoipAllocatedRangesRequestInfo struct {
		ContainerId string
		Request     core.ListByoipAllocatedRangesRequest
	}

	var requests []ListByoipAllocatedRangesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListByoipAllocatedRangesRequest)
				return c.ListByoipAllocatedRanges(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListByoipAllocatedRangesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListByoipAllocatedRangesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientListByoipRanges(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListByoipRanges")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListByoipRanges is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListByoipRanges", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListByoipRanges")
	assert.NoError(t, err)

	type ListByoipRangesRequestInfo struct {
		ContainerId string
		Request     core.ListByoipRangesRequest
	}

	var requests []ListByoipRangesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListByoipRangesRequest)
				return c.ListByoipRanges(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListByoipRangesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListByoipRangesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCpeDeviceShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListCpeDeviceShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCpeDeviceShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCpeDeviceShapes", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCpeDeviceShapes")
	assert.NoError(t, err)

	type ListCpeDeviceShapesRequestInfo struct {
		ContainerId string
		Request     core.ListCpeDeviceShapesRequest
	}

	var requests []ListCpeDeviceShapesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListCpeDeviceShapesRequest)
				return c.ListCpeDeviceShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListCpeDeviceShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListCpeDeviceShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCpes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCrossConnectGroups(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCrossConnectLocations(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCrossConnectMappings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListCrossConnectMappings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCrossConnectMappings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListCrossConnectMappings", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListCrossConnectMappings")
	assert.NoError(t, err)

	type ListCrossConnectMappingsRequestInfo struct {
		ContainerId string
		Request     core.ListCrossConnectMappingsRequest
	}

	var requests []ListCrossConnectMappingsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ListCrossConnectMappings(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCrossConnects(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListCrossconnectPortSpeedShapes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListDhcpOptions(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListDrgAttachments(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListDrgs(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListEndpointServices(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListEndpointServices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEndpointServices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListEndpointServices", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListEndpointServices")
	assert.NoError(t, err)

	type ListEndpointServicesRequestInfo struct {
		ContainerId string
		Request     core.ListEndpointServicesRequest
	}

	var requests []ListEndpointServicesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListEndpointServicesRequest)
				return c.ListEndpointServices(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListEndpointServicesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListEndpointServicesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListFastConnectProviderServices(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListFastConnectProviderVirtualCircuitBandwidthShapes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListFlowLogConfigAttachments(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListFlowLogConfigs(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListIPSecConnectionTunnels(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListIPSecConnectionTunnels")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIPSecConnectionTunnels is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListIPSecConnectionTunnels", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListIPSecConnectionTunnels")
	assert.NoError(t, err)

	type ListIPSecConnectionTunnelsRequestInfo struct {
		ContainerId string
		Request     core.ListIPSecConnectionTunnelsRequest
	}

	var requests []ListIPSecConnectionTunnelsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListIPSecConnectionTunnelsRequest)
				return c.ListIPSecConnectionTunnels(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListIPSecConnectionTunnelsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListIPSecConnectionTunnelsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListIPSecConnections(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternalDnsRecords(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInternalDnsRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternalDnsRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternalDnsRecords", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternalDnsRecords")
	assert.NoError(t, err)

	type ListInternalDnsRecordsRequestInfo struct {
		ContainerId string
		Request     core.ListInternalDnsRecordsRequest
	}

	var requests []ListInternalDnsRecordsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternalDnsRecordsRequest)
				return c.ListInternalDnsRecords(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternalDnsRecordsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternalDnsRecordsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternalDrgAttachments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInternalDrgAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternalDrgAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternalDrgAttachments", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternalDrgAttachments")
	assert.NoError(t, err)

	type ListInternalDrgAttachmentsRequestInfo struct {
		ContainerId string
		Request     core.ListInternalDrgAttachmentsRequest
	}

	var requests []ListInternalDrgAttachmentsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternalDrgAttachmentsRequest)
				return c.ListInternalDrgAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternalDrgAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternalDrgAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternalDrgs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInternalDrgs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternalDrgs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternalDrgs", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternalDrgs")
	assert.NoError(t, err)

	type ListInternalDrgsRequestInfo struct {
		ContainerId string
		Request     core.ListInternalDrgsRequest
	}

	var requests []ListInternalDrgsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternalDrgsRequest)
				return c.ListInternalDrgs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternalDrgsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternalDrgsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternalGenericGateways(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInternalGenericGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternalGenericGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternalGenericGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternalGenericGateways")
	assert.NoError(t, err)

	type ListInternalGenericGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ListInternalGenericGatewaysRequest
	}

	var requests []ListInternalGenericGatewaysRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternalGenericGatewaysRequest)
				return c.ListInternalGenericGateways(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternalGenericGatewaysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternalGenericGatewaysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternalPublicIps(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInternalPublicIps")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternalPublicIps is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternalPublicIps", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternalPublicIps")
	assert.NoError(t, err)

	type ListInternalPublicIpsRequestInfo struct {
		ContainerId string
		Request     core.ListInternalPublicIpsRequest
	}

	var requests []ListInternalPublicIpsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternalPublicIpsRequest)
				return c.ListInternalPublicIps(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternalPublicIpsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternalPublicIpsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternalVnics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInternalVnics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInternalVnics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListInternalVnics", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListInternalVnics")
	assert.NoError(t, err)

	type ListInternalVnicsRequestInfo struct {
		ContainerId string
		Request     core.ListInternalVnicsRequest
	}

	var requests []ListInternalVnicsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListInternalVnicsRequest)
				return c.ListInternalVnics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInternalVnicsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInternalVnicsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListInternetGateways(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListIpv6s(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListIpv6s")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIpv6s is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListIpv6s", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListIpv6s")
	assert.NoError(t, err)

	type ListIpv6sRequestInfo struct {
		ContainerId string
		Request     core.ListIpv6sRequest
	}

	var requests []ListIpv6sRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListIpv6sRequest)
				return c.ListIpv6s(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListIpv6sResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListIpv6sResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListLocalPeeringConnections(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListLocalPeeringGateways(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="pnp" email="elpaso_ops_us_grp@oracle.com" jiraProject="NAT" opsJiraProject="PNP"
func TestVirtualNetworkClientListNatGateways(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListNetworkSecurityGroupSecurityRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNetworkSecurityGroupSecurityRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListNetworkSecurityGroupSecurityRules", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)

	type ListNetworkSecurityGroupSecurityRulesRequestInfo struct {
		ContainerId string
		Request     core.ListNetworkSecurityGroupSecurityRulesRequest
	}

	var requests []ListNetworkSecurityGroupSecurityRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListNetworkSecurityGroupSecurityRulesRequest)
				return c.ListNetworkSecurityGroupSecurityRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListNetworkSecurityGroupSecurityRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListNetworkSecurityGroupSecurityRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListNetworkSecurityGroupVnics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListNetworkSecurityGroupVnics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNetworkSecurityGroupVnics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListNetworkSecurityGroupVnics", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListNetworkSecurityGroupVnics")
	assert.NoError(t, err)

	type ListNetworkSecurityGroupVnicsRequestInfo struct {
		ContainerId string
		Request     core.ListNetworkSecurityGroupVnicsRequest
	}

	var requests []ListNetworkSecurityGroupVnicsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListNetworkSecurityGroupVnicsRequest)
				return c.ListNetworkSecurityGroupVnics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListNetworkSecurityGroupVnicsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListNetworkSecurityGroupVnicsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListNetworkSecurityGroups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListNetworkSecurityGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNetworkSecurityGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListNetworkSecurityGroups", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListNetworkSecurityGroups")
	assert.NoError(t, err)

	type ListNetworkSecurityGroupsRequestInfo struct {
		ContainerId string
		Request     core.ListNetworkSecurityGroupsRequest
	}

	var requests []ListNetworkSecurityGroupsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListNetworkSecurityGroupsRequest)
				return c.ListNetworkSecurityGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListNetworkSecurityGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListNetworkSecurityGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListNextHops(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListNextHops")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListNextHops is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListNextHops", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListNextHops")
	assert.NoError(t, err)

	type ListNextHopsRequestInfo struct {
		ContainerId string
		Request     core.ListNextHopsRequest
	}

	var requests []ListNextHopsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ListNextHops(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListPrivateAccessGateways(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListPrivateAccessGateways")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPrivateAccessGateways is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListPrivateAccessGateways", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListPrivateAccessGateways")
	assert.NoError(t, err)

	type ListPrivateAccessGatewaysRequestInfo struct {
		ContainerId string
		Request     core.ListPrivateAccessGatewaysRequest
	}

	var requests []ListPrivateAccessGatewaysRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListPrivateAccessGatewaysRequest)
				return c.ListPrivateAccessGateways(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListPrivateAccessGatewaysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListPrivateAccessGatewaysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListPrivateEndpointAssociations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListPrivateEndpointAssociations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPrivateEndpointAssociations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListPrivateEndpointAssociations", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListPrivateEndpointAssociations")
	assert.NoError(t, err)

	type ListPrivateEndpointAssociationsRequestInfo struct {
		ContainerId string
		Request     core.ListPrivateEndpointAssociationsRequest
	}

	var requests []ListPrivateEndpointAssociationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListPrivateEndpointAssociationsRequest)
				return c.ListPrivateEndpointAssociations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListPrivateEndpointAssociationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListPrivateEndpointAssociationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListPrivateEndpoints(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListPrivateEndpoints")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPrivateEndpoints is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListPrivateEndpoints", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListPrivateEndpoints")
	assert.NoError(t, err)

	type ListPrivateEndpointsRequestInfo struct {
		ContainerId string
		Request     core.ListPrivateEndpointsRequest
	}

	var requests []ListPrivateEndpointsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListPrivateEndpointsRequest)
				return c.ListPrivateEndpoints(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListPrivateEndpointsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListPrivateEndpointsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListPrivateIps(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientListPublicIpPools(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListPublicIpPools")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPublicIpPools is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListPublicIpPools", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListPublicIpPools")
	assert.NoError(t, err)

	type ListPublicIpPoolsRequestInfo struct {
		ContainerId string
		Request     core.ListPublicIpPoolsRequest
	}

	var requests []ListPublicIpPoolsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListPublicIpPoolsRequest)
				return c.ListPublicIpPools(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListPublicIpPoolsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListPublicIpPoolsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListPublicIps(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListRemotePeeringConnections(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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
func TestVirtualNetworkClientListReverseConnectionNatIpCidrs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListReverseConnectionNatIpCidrs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListReverseConnectionNatIpCidrs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListReverseConnectionNatIpCidrs", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListReverseConnectionNatIpCidrs")
	assert.NoError(t, err)

	type ListReverseConnectionNatIpCidrsRequestInfo struct {
		ContainerId string
		Request     core.ListReverseConnectionNatIpCidrsRequest
	}

	var requests []ListReverseConnectionNatIpCidrsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ListReverseConnectionNatIpCidrs(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListReverseConnectionNatIps(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListReverseConnectionNatIps")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListReverseConnectionNatIps is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListReverseConnectionNatIps", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListReverseConnectionNatIps")
	assert.NoError(t, err)

	type ListReverseConnectionNatIpsRequestInfo struct {
		ContainerId string
		Request     core.ListReverseConnectionNatIpsRequest
	}

	var requests []ListReverseConnectionNatIpsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListReverseConnectionNatIpsRequest)
				return c.ListReverseConnectionNatIps(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListReverseConnectionNatIpsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListReverseConnectionNatIpsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListRouteTables(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListSecurityLists(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListServiceGateways(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientListServices(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListSubnets(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListVcns(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListVirtualCircuitBandwidthShapes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListVirtualCircuitPublicPrefixes(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ListVirtualCircuitPublicPrefixes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientListVirtualCircuits(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
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

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListVlans(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListVlans")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVlans is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListVlans", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListVlans")
	assert.NoError(t, err)

	type ListVlansRequestInfo struct {
		ContainerId string
		Request     core.ListVlansRequest
	}

	var requests []ListVlansRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListVlansRequest)
				return c.ListVlans(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVlansResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVlansResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientListVnicWorkers(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListVnicWorkers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVnicWorkers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ListVnicWorkers", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ListVnicWorkers")
	assert.NoError(t, err)

	type ListVnicWorkersRequestInfo struct {
		ContainerId string
		Request     core.ListVnicWorkersRequest
	}

	var requests []ListVnicWorkersRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*core.ListVnicWorkersRequest)
				return c.ListVnicWorkers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVnicWorkersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVnicWorkersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientMigrateDrg(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "MigrateDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MigrateDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "MigrateDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "MigrateDrg")
	assert.NoError(t, err)

	type MigrateDrgRequestInfo struct {
		ContainerId string
		Request     core.MigrateDrgRequest
	}

	var requests []MigrateDrgRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.MigrateDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientModifyReverseConnections(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ModifyReverseConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ModifyReverseConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ModifyReverseConnections", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ModifyReverseConnections")
	assert.NoError(t, err)

	type ModifyReverseConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ModifyReverseConnectionsRequest
	}

	var requests []ModifyReverseConnectionsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ModifyReverseConnections(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientModifyVcnCidr(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ModifyVcnCidr")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ModifyVcnCidr is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ModifyVcnCidr", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ModifyVcnCidr")
	assert.NoError(t, err)

	type ModifyVcnCidrRequestInfo struct {
		ContainerId string
		Request     core.ModifyVcnCidrRequest
	}

	var requests []ModifyVcnCidrRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ModifyVcnCidr(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientRemoveNetworkSecurityGroupSecurityRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "RemoveNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveNetworkSecurityGroupSecurityRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "RemoveNetworkSecurityGroupSecurityRules", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "RemoveNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)

	type RemoveNetworkSecurityGroupSecurityRulesRequestInfo struct {
		ContainerId string
		Request     core.RemoveNetworkSecurityGroupSecurityRulesRequest
	}

	var requests []RemoveNetworkSecurityGroupSecurityRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RemoveNetworkSecurityGroupSecurityRules(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientRemovePublicIpPoolCapacity(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "RemovePublicIpPoolCapacity")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemovePublicIpPoolCapacity is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "RemovePublicIpPoolCapacity", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "RemovePublicIpPoolCapacity")
	assert.NoError(t, err)

	type RemovePublicIpPoolCapacityRequestInfo struct {
		ContainerId string
		Request     core.RemovePublicIpPoolCapacityRequest
	}

	var requests []RemovePublicIpPoolCapacityRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RemovePublicIpPoolCapacity(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientRemoveVcnCidr(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "RemoveVcnCidr")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveVcnCidr is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "RemoveVcnCidr", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "RemoveVcnCidr")
	assert.NoError(t, err)

	type RemoveVcnCidrRequestInfo struct {
		ContainerId string
		Request     core.RemoveVcnCidrRequest
	}

	var requests []RemoveVcnCidrRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.RemoveVcnCidr(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientUpdateByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateByoipRange")
	assert.NoError(t, err)

	type UpdateByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.UpdateByoipRangeRequest
	}

	var requests []UpdateByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateCpe(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateCpe(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateCrossConnect(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateCrossConnect(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateCrossConnectGroup(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateCrossConnectGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateDhcpOptions(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDhcpOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateDrg(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientUpdateEndpointService(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateEndpointService")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateEndpointService is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateEndpointService", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateEndpointService")
	assert.NoError(t, err)

	type UpdateEndpointServiceRequestInfo struct {
		ContainerId string
		Request     core.UpdateEndpointServiceRequest
	}

	var requests []UpdateEndpointServiceRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateEndpointService(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientUpdateEndpointServiceNextHop(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateEndpointServiceNextHop")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateEndpointServiceNextHop is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateEndpointServiceNextHop", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateEndpointServiceNextHop")
	assert.NoError(t, err)

	type UpdateEndpointServiceNextHopRequestInfo struct {
		ContainerId string
		Request     core.UpdateEndpointServiceNextHopRequest
	}

	var requests []UpdateEndpointServiceNextHopRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateEndpointServiceNextHop(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateFlowLogConfig(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateFlowLogConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateFlowLogConfigAttachment(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateFlowLogConfigAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateIPSecConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateIPSecConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateIPSecConnectionTunnel(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateIPSecConnectionTunnel")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIPSecConnectionTunnel is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateIPSecConnectionTunnel", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateIPSecConnectionTunnel")
	assert.NoError(t, err)

	type UpdateIPSecConnectionTunnelRequestInfo struct {
		ContainerId string
		Request     core.UpdateIPSecConnectionTunnelRequest
	}

	var requests []UpdateIPSecConnectionTunnelRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateIPSecConnectionTunnel(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateIPSecConnectionTunnelSharedSecret(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateIPSecConnectionTunnelSharedSecret")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIPSecConnectionTunnelSharedSecret is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateIPSecConnectionTunnelSharedSecret", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateIPSecConnectionTunnelSharedSecret")
	assert.NoError(t, err)

	type UpdateIPSecConnectionTunnelSharedSecretRequestInfo struct {
		ContainerId string
		Request     core.UpdateIPSecConnectionTunnelSharedSecretRequest
	}

	var requests []UpdateIPSecConnectionTunnelSharedSecretRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateIPSecConnectionTunnelSharedSecret(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternalDnsRecord(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInternalDnsRecord")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternalDnsRecord is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternalDnsRecord", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternalDnsRecord")
	assert.NoError(t, err)

	type UpdateInternalDnsRecordRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternalDnsRecordRequest
	}

	var requests []UpdateInternalDnsRecordRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternalDnsRecord(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternalDrg(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInternalDrg")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternalDrg is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternalDrg", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternalDrg")
	assert.NoError(t, err)

	type UpdateInternalDrgRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternalDrgRequest
	}

	var requests []UpdateInternalDrgRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternalDrg(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternalDrgAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInternalDrgAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternalDrgAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternalDrgAttachment", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternalDrgAttachment")
	assert.NoError(t, err)

	type UpdateInternalDrgAttachmentRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternalDrgAttachmentRequest
	}

	var requests []UpdateInternalDrgAttachmentRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternalDrgAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternalGenericGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInternalGenericGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternalGenericGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternalGenericGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternalGenericGateway")
	assert.NoError(t, err)

	type UpdateInternalGenericGatewayRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternalGenericGatewayRequest
	}

	var requests []UpdateInternalGenericGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternalGenericGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternalPublicIp(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInternalPublicIp")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternalPublicIp is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternalPublicIp", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternalPublicIp")
	assert.NoError(t, err)

	type UpdateInternalPublicIpRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternalPublicIpRequest
	}

	var requests []UpdateInternalPublicIpRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternalPublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternalVnic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInternalVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInternalVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateInternalVnic", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateInternalVnic")
	assert.NoError(t, err)

	type UpdateInternalVnicRequestInfo struct {
		ContainerId string
		Request     core.UpdateInternalVnicRequest
	}

	var requests []UpdateInternalVnicRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternalVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateInternetGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateInternetGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateIpv6(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateIpv6")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIpv6 is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateIpv6", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateIpv6")
	assert.NoError(t, err)

	type UpdateIpv6RequestInfo struct {
		ContainerId string
		Request     core.UpdateIpv6Request
	}

	var requests []UpdateIpv6RequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateIpv6(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateLocalPeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateLocalPeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateLocalPeeringGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateLocalPeeringGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="pnp" email="elpaso_ops_us_grp@oracle.com" jiraProject="NAT" opsJiraProject="PNP"
func TestVirtualNetworkClientUpdateNatGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateNatGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateNetworkSecurityGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateNetworkSecurityGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNetworkSecurityGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateNetworkSecurityGroup", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateNetworkSecurityGroup")
	assert.NoError(t, err)

	type UpdateNetworkSecurityGroupRequestInfo struct {
		ContainerId string
		Request     core.UpdateNetworkSecurityGroupRequest
	}

	var requests []UpdateNetworkSecurityGroupRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateNetworkSecurityGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateNetworkSecurityGroupSecurityRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNetworkSecurityGroupSecurityRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateNetworkSecurityGroupSecurityRules", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateNetworkSecurityGroupSecurityRules")
	assert.NoError(t, err)

	type UpdateNetworkSecurityGroupSecurityRulesRequestInfo struct {
		ContainerId string
		Request     core.UpdateNetworkSecurityGroupSecurityRulesRequest
	}

	var requests []UpdateNetworkSecurityGroupSecurityRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateNetworkSecurityGroupSecurityRules(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientUpdatePrivateAccessGateway(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdatePrivateAccessGateway")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePrivateAccessGateway is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdatePrivateAccessGateway", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdatePrivateAccessGateway")
	assert.NoError(t, err)

	type UpdatePrivateAccessGatewayRequestInfo struct {
		ContainerId string
		Request     core.UpdatePrivateAccessGatewayRequest
	}

	var requests []UpdatePrivateAccessGatewayRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdatePrivateAccessGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="privateEndpoint" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientUpdatePrivateEndpoint(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdatePrivateEndpoint")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePrivateEndpoint is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdatePrivateEndpoint", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdatePrivateEndpoint")
	assert.NoError(t, err)

	type UpdatePrivateEndpointRequestInfo struct {
		ContainerId string
		Request     core.UpdatePrivateEndpointRequest
	}

	var requests []UpdatePrivateEndpointRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdatePrivateEndpoint(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdatePrivateIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdatePrivateIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdatePrivateIpNextHop(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdatePrivateIpNextHop")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePrivateIpNextHop is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdatePrivateIpNextHop", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdatePrivateIpNextHop")
	assert.NoError(t, err)

	type UpdatePrivateIpNextHopRequestInfo struct {
		ContainerId string
		Request     core.UpdatePrivateIpNextHopRequest
	}

	var requests []UpdatePrivateIpNextHopRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdatePrivateIpNextHop(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdatePublicIp(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdatePublicIp(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientUpdatePublicIpPool(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdatePublicIpPool")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePublicIpPool is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdatePublicIpPool", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdatePublicIpPool")
	assert.NoError(t, err)

	type UpdatePublicIpPoolRequestInfo struct {
		ContainerId string
		Request     core.UpdatePublicIpPoolRequest
	}

	var requests []UpdatePublicIpPoolRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdatePublicIpPool(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateRemotePeeringConnection(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateRemotePeeringConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateRouteTable(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateRouteTable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateSecurityList(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateSecurityList(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="serviceGateway" email="oci_sgw_ops_us_grp@oracle.com" jiraProject="SG" opsJiraProject="SGW"
func TestVirtualNetworkClientUpdateServiceGateway(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateServiceGateway(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateSubnet(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateSubnet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateTunnelCpeDeviceConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateTunnelCpeDeviceConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTunnelCpeDeviceConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateTunnelCpeDeviceConfig", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateTunnelCpeDeviceConfig")
	assert.NoError(t, err)

	type UpdateTunnelCpeDeviceConfigRequestInfo struct {
		ContainerId string
		Request     core.UpdateTunnelCpeDeviceConfigRequest
	}

	var requests []UpdateTunnelCpeDeviceConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateTunnelCpeDeviceConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateVcn(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVcn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="c3" email="c3_scrum_team_us_grp@oracle.com" jiraProject="RSC" opsJiraProject="RSC"
func TestVirtualNetworkClientUpdateVirtualCircuit(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVirtualCircuit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateVlan(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateVlan")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVlan is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateVlan", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateVlan")
	assert.NoError(t, err)

	type UpdateVlanRequestInfo struct {
		ContainerId string
		Request     core.UpdateVlanRequest
	}

	var requests []UpdateVlanRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVlan(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateVnic(t *testing.T) {
	defer failTestOnPanic(t)

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
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="virtualNetwork" email="bmc_vcn_cp_us_grp@oracle.com" jiraProject="VCN" opsJiraProject="VN"
func TestVirtualNetworkClientUpdateVnicWorker(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateVnicWorker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVnicWorker is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "UpdateVnicWorker", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "UpdateVnicWorker")
	assert.NoError(t, err)

	type UpdateVnicWorkerRequestInfo struct {
		ContainerId string
		Request     core.UpdateVnicWorkerRequest
	}

	var requests []UpdateVnicWorkerRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.UpdateVnicWorker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientValidateByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ValidateByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ValidateByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "ValidateByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "ValidateByoipRange")
	assert.NoError(t, err)

	type ValidateByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.ValidateByoipRangeRequest
	}

	var requests []ValidateByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.ValidateByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="vcnip" email="vcn_ip_mgmt_grp@oracle.com" jiraProject="VCNIP" opsJiraProject="VCNIP"
func TestVirtualNetworkClientWithdrawByoipRange(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "WithdrawByoipRange")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("WithdrawByoipRange is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "VirtualNetwork", "WithdrawByoipRange", createVirtualNetworkClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.VirtualNetworkClient)

	body, err := testClient.getRequests("core", "WithdrawByoipRange")
	assert.NoError(t, err)

	type WithdrawByoipRangeRequestInfo struct {
		ContainerId string
		Request     core.WithdrawByoipRangeRequest
	}

	var requests []WithdrawByoipRangeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.WithdrawByoipRange(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
