package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/marketplace"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestPublisherResourcesClientGetPublisher(t *testing.T) {
	enabled, err := testClient.isApiEnabled("marketplace", "GetPublisher")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPublisher is not enabled by the testing service")
	}
	c, err := marketplace.NewPublisherResourcesClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("marketplace", "GetPublisher")
	assert.NoError(t, err)

	type GetPublisherRequestInfo struct {
		ContainerId string
		Request     marketplace.GetPublisherRequest
	}

	var requests []GetPublisherRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPublisher(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestPublisherResourcesClientListPublisherApplications(t *testing.T) {
	enabled, err := testClient.isApiEnabled("marketplace", "ListPublisherApplications")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPublisherApplications is not enabled by the testing service")
	}
	c, err := marketplace.NewPublisherResourcesClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("marketplace", "ListPublisherApplications")
	assert.NoError(t, err)

	type ListPublisherApplicationsRequestInfo struct {
		ContainerId string
		Request     marketplace.ListPublisherApplicationsRequest
	}

	var requests []ListPublisherApplicationsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*marketplace.ListPublisherApplicationsRequest)
				return c.ListPublisherApplications(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]marketplace.ListPublisherApplicationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(marketplace.ListPublisherApplicationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
