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

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestApplicationInstanceResourceClientGetInstallConfiguration(t *testing.T) {
	enabled, err := testClient.isApiEnabled("marketplace", "GetInstallConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstallConfiguration is not enabled by the testing service")
	}
	c, err := marketplace.NewApplicationInstanceResourceClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("marketplace", "GetInstallConfiguration")
	assert.NoError(t, err)

	type GetInstallConfigurationRequestInfo struct {
		ContainerId string
		Request     marketplace.GetInstallConfigurationRequest
	}

	var requests []GetInstallConfigurationRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetInstallConfiguration(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo email="" jiraProject="" opsJiraProject=""
func TestApplicationInstanceResourceClientGetTermsOfUses(t *testing.T) {
	enabled, err := testClient.isApiEnabled("marketplace", "GetTermsOfUses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTermsOfUses is not enabled by the testing service")
	}
	c, err := marketplace.NewApplicationInstanceResourceClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("marketplace", "GetTermsOfUses")
	assert.NoError(t, err)

	type GetTermsOfUsesRequestInfo struct {
		ContainerId string
		Request     marketplace.GetTermsOfUsesRequest
	}

	var requests []GetTermsOfUsesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetTermsOfUses(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
