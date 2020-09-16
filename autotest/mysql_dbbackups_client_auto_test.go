package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"github.com/oracle/oci-go-sdk/v25/mysql"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDbBackupsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := mysql.NewDbBackupsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbBackupsClientCreateBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "CreateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbBackups", "CreateBackup", createDbBackupsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbBackupsClient)

	body, err := testClient.getRequests("mysql", "CreateBackup")
	assert.NoError(t, err)

	type CreateBackupRequestInfo struct {
		ContainerId string
		Request     mysql.CreateBackupRequest
	}

	var requests []CreateBackupRequestInfo
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
			response, err := c.CreateBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbBackupsClientDeleteBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "DeleteBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbBackups", "DeleteBackup", createDbBackupsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbBackupsClient)

	body, err := testClient.getRequests("mysql", "DeleteBackup")
	assert.NoError(t, err)

	type DeleteBackupRequestInfo struct {
		ContainerId string
		Request     mysql.DeleteBackupRequest
	}

	var requests []DeleteBackupRequestInfo
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
			response, err := c.DeleteBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbBackupsClientGetBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "GetBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbBackups", "GetBackup", createDbBackupsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbBackupsClient)

	body, err := testClient.getRequests("mysql", "GetBackup")
	assert.NoError(t, err)

	type GetBackupRequestInfo struct {
		ContainerId string
		Request     mysql.GetBackupRequest
	}

	var requests []GetBackupRequestInfo
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
			response, err := c.GetBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbBackupsClientListBackups(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "ListBackups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbBackups", "ListBackups", createDbBackupsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbBackupsClient)

	body, err := testClient.getRequests("mysql", "ListBackups")
	assert.NoError(t, err)

	type ListBackupsRequestInfo struct {
		ContainerId string
		Request     mysql.ListBackupsRequest
	}

	var requests []ListBackupsRequestInfo
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
				r := req.(*mysql.ListBackupsRequest)
				return c.ListBackups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]mysql.ListBackupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(mysql.ListBackupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="mysql-cloud-dev_ww_grp@oracle.com" jiraProject="MY" opsJiraProject="MYOPS"
func TestDbBackupsClientUpdateBackup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("mysql", "UpdateBackup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("mysql", "DbBackups", "UpdateBackup", createDbBackupsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(mysql.DbBackupsClient)

	body, err := testClient.getRequests("mysql", "UpdateBackup")
	assert.NoError(t, err)

	type UpdateBackupRequestInfo struct {
		ContainerId string
		Request     mysql.UpdateBackupRequest
	}

	var requests []UpdateBackupRequestInfo
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
			response, err := c.UpdateBackup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
