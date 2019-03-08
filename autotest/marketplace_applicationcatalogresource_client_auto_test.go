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

func createApplicationCatalogResourceClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := marketplace.NewApplicationCatalogResourceClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestApplicationCatalogResourceClientGetAppCatalogApplicationMapping(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("marketplace", "GetAppCatalogApplicationMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAppCatalogApplicationMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("marketplace", "ApplicationCatalogResource", "GetAppCatalogApplicationMapping", createApplicationCatalogResourceClientWithProvider)
	assert.NoError(t, err)
	c := cc.(marketplace.ApplicationCatalogResourceClient)

	body, err := testClient.getRequests("marketplace", "GetAppCatalogApplicationMapping")
	assert.NoError(t, err)

	type GetAppCatalogApplicationMappingRequestInfo struct {
		ContainerId string
		Request     marketplace.GetAppCatalogApplicationMappingRequest
	}

	var requests []GetAppCatalogApplicationMappingRequestInfo
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

			response, err := c.GetAppCatalogApplicationMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
