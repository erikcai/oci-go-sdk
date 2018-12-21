package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/dns"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDnsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := dns.NewDnsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientCreateZone(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "CreateZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateZone is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "CreateZone")
	assert.NoError(t, err)

	type CreateZoneRequestInfo struct {
		ContainerId string
		Request     dns.CreateZoneRequest
	}

	var requests []CreateZoneRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientDeleteDomainRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "DeleteDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDomainRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "DeleteDomainRecords")
	assert.NoError(t, err)

	type DeleteDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.DeleteDomainRecordsRequest
	}

	var requests []DeleteDomainRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDomainRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientDeleteRRSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "DeleteRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRRSet is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "DeleteRRSet")
	assert.NoError(t, err)

	type DeleteRRSetRequestInfo struct {
		ContainerId string
		Request     dns.DeleteRRSetRequest
	}

	var requests []DeleteRRSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteRRSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientDeleteZone(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "DeleteZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteZone is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "DeleteZone")
	assert.NoError(t, err)

	type DeleteZoneRequestInfo struct {
		ContainerId string
		Request     dns.DeleteZoneRequest
	}

	var requests []DeleteZoneRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientGetDomainRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDomainRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetDomainRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetDomainRecords")
	assert.NoError(t, err)

	type GetDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.GetDomainRecordsRequest
	}

	var requests []GetDomainRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dns.GetDomainRecordsRequest)
				return c.GetDomainRecords(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.GetDomainRecordsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.GetDomainRecordsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientGetRRSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRRSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetRRSet", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetRRSet")
	assert.NoError(t, err)

	type GetRRSetRequestInfo struct {
		ContainerId string
		Request     dns.GetRRSetRequest
	}

	var requests []GetRRSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dns.GetRRSetRequest)
				return c.GetRRSet(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.GetRRSetResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.GetRRSetResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientGetZone(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetZone is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "GetZone")
	assert.NoError(t, err)

	type GetZoneRequestInfo struct {
		ContainerId string
		Request     dns.GetZoneRequest
	}

	var requests []GetZoneRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientGetZoneRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetZoneRecords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "GetZoneRecords", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "GetZoneRecords")
	assert.NoError(t, err)

	type GetZoneRecordsRequestInfo struct {
		ContainerId string
		Request     dns.GetZoneRecordsRequest
	}

	var requests []GetZoneRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dns.GetZoneRecordsRequest)
				return c.GetZoneRecords(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.GetZoneRecordsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.GetZoneRecordsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientListZones(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "ListZones")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListZones is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("dns", "Dns", "ListZones", createDnsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(dns.DnsClient)

	body, err := testClient.getRequests("dns", "ListZones")
	assert.NoError(t, err)

	type ListZonesRequestInfo struct {
		ContainerId string
		Request     dns.ListZonesRequest
	}

	var requests []ListZonesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dns.ListZonesRequest)
				return c.ListZones(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListZonesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListZonesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientPatchDomainRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "PatchDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PatchDomainRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "PatchDomainRecords")
	assert.NoError(t, err)

	type PatchDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.PatchDomainRecordsRequest
	}

	var requests []PatchDomainRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.PatchDomainRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientPatchRRSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "PatchRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PatchRRSet is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "PatchRRSet")
	assert.NoError(t, err)

	type PatchRRSetRequestInfo struct {
		ContainerId string
		Request     dns.PatchRRSetRequest
	}

	var requests []PatchRRSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.PatchRRSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientPatchZoneRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "PatchZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PatchZoneRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "PatchZoneRecords")
	assert.NoError(t, err)

	type PatchZoneRecordsRequestInfo struct {
		ContainerId string
		Request     dns.PatchZoneRecordsRequest
	}

	var requests []PatchZoneRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.PatchZoneRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientUpdateDomainRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "UpdateDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDomainRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "UpdateDomainRecords")
	assert.NoError(t, err)

	type UpdateDomainRecordsRequestInfo struct {
		ContainerId string
		Request     dns.UpdateDomainRecordsRequest
	}

	var requests []UpdateDomainRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDomainRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientUpdateRRSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "UpdateRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRRSet is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "UpdateRRSet")
	assert.NoError(t, err)

	type UpdateRRSetRequestInfo struct {
		ContainerId string
		Request     dns.UpdateRRSetRequest
	}

	var requests []UpdateRRSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateRRSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientUpdateZone(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "UpdateZone")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateZone is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "UpdateZone")
	assert.NoError(t, err)

	type UpdateZoneRequestInfo struct {
		ContainerId string
		Request     dns.UpdateZoneRequest
	}

	var requests []UpdateZoneRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateZone(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestDnsClientUpdateZoneRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "UpdateZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateZoneRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "UpdateZoneRecords")
	assert.NoError(t, err)

	type UpdateZoneRecordsRequestInfo struct {
		ContainerId string
		Request     dns.UpdateZoneRecordsRequest
	}

	var requests []UpdateZoneRecordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateZoneRecords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
