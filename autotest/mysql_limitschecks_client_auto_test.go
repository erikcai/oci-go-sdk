package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/mysql"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createLimitsChecksClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := mysql.NewLimitsChecksClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestLimitsChecksClientGetMysqlaasInstanceLimitsCheck(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetMysqlaasInstanceLimitsCheck")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMysqlaasInstanceLimitsCheck is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "LimitsChecks", "GetMysqlaasInstanceLimitsCheck", createLimitsChecksClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.LimitsChecksClient)

	body, err := testClient.getRequests("mysql", "GetMysqlaasInstanceLimitsCheck")
	assert.NoError(t, err)

	type GetMysqlaasInstanceLimitsCheckRequestInfo struct {
		ContainerId string
		Request     mysql.GetMysqlaasInstanceLimitsCheckRequest
	}

	var requests []GetMysqlaasInstanceLimitsCheckRequestInfo
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
			response, err := c.GetMysqlaasInstanceLimitsCheck(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
