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
func TestApplicationCatalogResourceClientGetAppCatalogApplicationMapping(t *testing.T) {
	enabled, err := testClient.isApiEnabled("marketplace", "GetAppCatalogApplicationMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAppCatalogApplicationMapping is not enabled by the testing service")
	}
	c, err := marketplace.NewApplicationCatalogResourceClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("marketplace", "GetAppCatalogApplicationMapping")
	assert.NoError(t, err)

	type GetAppCatalogApplicationMappingRequestInfo struct {
		ContainerId string
		Request     marketplace.GetAppCatalogApplicationMappingRequest
	}

	var requests []GetAppCatalogApplicationMappingRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
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
