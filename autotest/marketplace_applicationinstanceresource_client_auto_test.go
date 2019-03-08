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

func createApplicationInstanceResourceClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := marketplace.NewApplicationInstanceResourceClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestApplicationInstanceResourceClientGetInstallConfiguration(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("marketplace", "GetInstallConfiguration")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstallConfiguration is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("marketplace", "ApplicationInstanceResource", "GetInstallConfiguration", createApplicationInstanceResourceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(marketplace.ApplicationInstanceResourceClient)

	body, err := testClient.getRequests("marketplace", "GetInstallConfiguration")
	assert.NoError(t, err)

	type GetInstallConfigurationRequestInfo struct {
		ContainerId string
		Request     marketplace.GetInstallConfigurationRequest
	}

	var requests []GetInstallConfigurationRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestApplicationInstanceResourceClientGetTermsOfUses(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("marketplace", "GetTermsOfUses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTermsOfUses is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("marketplace", "ApplicationInstanceResource", "GetTermsOfUses", createApplicationInstanceResourceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(marketplace.ApplicationInstanceResourceClient)

	body, err := testClient.getRequests("marketplace", "GetTermsOfUses")
	assert.NoError(t, err)

	type GetTermsOfUsesRequestInfo struct {
		ContainerId string
		Request     marketplace.GetTermsOfUsesRequest
	}

	var requests []GetTermsOfUsesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
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
