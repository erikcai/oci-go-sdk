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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientCreateSteeringPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "CreateSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSteeringPolicy is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "CreateSteeringPolicy")
	assert.NoError(t, err)

	type CreateSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.CreateSteeringPolicyRequest
	}

	var requests []CreateSteeringPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientCreateSteeringPolicyAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "CreateSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSteeringPolicyAttachment is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "CreateSteeringPolicyAttachment")
	assert.NoError(t, err)

	type CreateSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.CreateSteeringPolicyAttachmentRequest
	}

	var requests []CreateSteeringPolicyAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientDeleteSteeringPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "DeleteSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSteeringPolicy is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "DeleteSteeringPolicy")
	assert.NoError(t, err)

	type DeleteSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.DeleteSteeringPolicyRequest
	}

	var requests []DeleteSteeringPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientDeleteSteeringPolicyAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "DeleteSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSteeringPolicyAttachment is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "DeleteSteeringPolicyAttachment")
	assert.NoError(t, err)

	type DeleteSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.DeleteSteeringPolicyAttachmentRequest
	}

	var requests []DeleteSteeringPolicyAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientGetDomainRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetDomainRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDomainRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientGetRRSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetRRSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRRSet is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientGetSteeringPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSteeringPolicy is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "GetSteeringPolicy")
	assert.NoError(t, err)

	type GetSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.GetSteeringPolicyRequest
	}

	var requests []GetSteeringPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientGetSteeringPolicyAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSteeringPolicyAttachment is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "GetSteeringPolicyAttachment")
	assert.NoError(t, err)

	type GetSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.GetSteeringPolicyAttachmentRequest
	}

	var requests []GetSteeringPolicyAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientGetZoneRecords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "GetZoneRecords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetZoneRecords is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientListSteeringPolicies(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "ListSteeringPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSteeringPolicies is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "ListSteeringPolicies")
	assert.NoError(t, err)

	type ListSteeringPoliciesRequestInfo struct {
		ContainerId string
		Request     dns.ListSteeringPoliciesRequest
	}

	var requests []ListSteeringPoliciesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dns.ListSteeringPoliciesRequest)
				return c.ListSteeringPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListSteeringPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListSteeringPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientListSteeringPolicyAttachments(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "ListSteeringPolicyAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSteeringPolicyAttachments is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "ListSteeringPolicyAttachments")
	assert.NoError(t, err)

	type ListSteeringPolicyAttachmentsRequestInfo struct {
		ContainerId string
		Request     dns.ListSteeringPolicyAttachmentsRequest
	}

	var requests []ListSteeringPolicyAttachmentsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*dns.ListSteeringPolicyAttachmentsRequest)
				return c.ListSteeringPolicyAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]dns.ListSteeringPolicyAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(dns.ListSteeringPolicyAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientListZones(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "ListZones")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListZones is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientUpdateSteeringPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "UpdateSteeringPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSteeringPolicy is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "UpdateSteeringPolicy")
	assert.NoError(t, err)

	type UpdateSteeringPolicyRequestInfo struct {
		ContainerId string
		Request     dns.UpdateSteeringPolicyRequest
	}

	var requests []UpdateSteeringPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSteeringPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
func TestDnsClientUpdateSteeringPolicyAttachment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("dns", "UpdateSteeringPolicyAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSteeringPolicyAttachment is not enabled by the testing service")
	}
	c, err := dns.NewDnsClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("dns", "UpdateSteeringPolicyAttachment")
	assert.NoError(t, err)

	type UpdateSteeringPolicyAttachmentRequestInfo struct {
		ContainerId string
		Request     dns.UpdateSteeringPolicyAttachmentRequest
	}

	var requests []UpdateSteeringPolicyAttachmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSteeringPolicyAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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

// IssueRoutingInfo email="email-dev_us_grp@oracle.com" jiraProject="Email Delivery (ED)" opsJiraProject="Email Delivery"
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
