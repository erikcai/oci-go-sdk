package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/cims"
	"github.com/oracle/oci-go-sdk/v25/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createCimsUserClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := cims.NewUserClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_ops_cims_dev_us_grp@oracle.com" jiraProject="CIMS" opsJiraProject="CIMS"
func TestCimsUserClientCreateUser(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("cims", "CreateUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("cims", "User", "CreateUser", createCimsUserClientWithProvider)
	assert.NoError(t, err)
	c := cc.(cims.UserClient)

	body, err := testClient.getRequests("cims", "CreateUser")
	assert.NoError(t, err)

	type CreateUserRequestInfo struct {
		ContainerId string
		Request     cims.CreateUserRequest
	}

	var requests []CreateUserRequestInfo
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
			response, err := c.CreateUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
