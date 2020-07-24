package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/optimizer"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createOptimizerEnrollmentStatusClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := optimizer.NewEnrollmentStatusClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerEnrollmentStatusClientGetEnrollmentStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "GetEnrollmentStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetEnrollmentStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "EnrollmentStatus", "GetEnrollmentStatus", createOptimizerEnrollmentStatusClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.EnrollmentStatusClient)

	body, err := testClient.getRequests("optimizer", "GetEnrollmentStatus")
	assert.NoError(t, err)

	type GetEnrollmentStatusRequestInfo struct {
		ContainerId string
		Request     optimizer.GetEnrollmentStatusRequest
	}

	var requests []GetEnrollmentStatusRequestInfo
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
			response, err := c.GetEnrollmentStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerEnrollmentStatusClientListEnrollmentStatuses(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "ListEnrollmentStatuses")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEnrollmentStatuses is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "EnrollmentStatus", "ListEnrollmentStatuses", createOptimizerEnrollmentStatusClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.EnrollmentStatusClient)

	body, err := testClient.getRequests("optimizer", "ListEnrollmentStatuses")
	assert.NoError(t, err)

	type ListEnrollmentStatusesRequestInfo struct {
		ContainerId string
		Request     optimizer.ListEnrollmentStatusesRequest
	}

	var requests []ListEnrollmentStatusesRequestInfo
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
				r := req.(*optimizer.ListEnrollmentStatusesRequest)
				return c.ListEnrollmentStatuses(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]optimizer.ListEnrollmentStatusesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(optimizer.ListEnrollmentStatusesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oracle_cloud_optimizer_us_grp@oracle.com" jiraProject="OPTIMIZER" opsJiraProject="OPTIMIZER"
func TestOptimizerEnrollmentStatusClientUpdateEnrollmentStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("optimizer", "UpdateEnrollmentStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateEnrollmentStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("optimizer", "EnrollmentStatus", "UpdateEnrollmentStatus", createOptimizerEnrollmentStatusClientWithProvider)
	assert.NoError(t, err)
	c := cc.(optimizer.EnrollmentStatusClient)

	body, err := testClient.getRequests("optimizer", "UpdateEnrollmentStatus")
	assert.NoError(t, err)

	type UpdateEnrollmentStatusRequestInfo struct {
		ContainerId string
		Request     optimizer.UpdateEnrollmentStatusRequest
	}

	var requests []UpdateEnrollmentStatusRequestInfo
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
			response, err := c.UpdateEnrollmentStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
