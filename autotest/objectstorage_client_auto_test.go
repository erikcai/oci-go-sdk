package autotest

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/objectstorage"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createObjectStorageClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientAbortMultipartUpload(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "AbortMultipartUpload")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AbortMultipartUpload is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "AbortMultipartUpload", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "AbortMultipartUpload")
	assert.NoError(t, err)

	type AbortMultipartUploadRequestInfo struct {
		ContainerId string
		Request     objectstorage.AbortMultipartUploadRequest
	}

	var requests []AbortMultipartUploadRequestInfo
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
			response, err := c.AbortMultipartUpload(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientBulkCopyObjects(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "BulkCopyObjects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("BulkCopyObjects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "BulkCopyObjects", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "BulkCopyObjects")
	assert.NoError(t, err)

	type BulkCopyObjectsRequestInfo struct {
		ContainerId string
		Request     objectstorage.BulkCopyObjectsRequest
	}

	var requests []BulkCopyObjectsRequestInfo
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
			response, err := c.BulkCopyObjects(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCancelWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CancelWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CancelWorkRequest", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CancelWorkRequest")
	assert.NoError(t, err)

	type CancelWorkRequestRequestInfo struct {
		ContainerId string
		Request     objectstorage.CancelWorkRequestRequest
	}

	var requests []CancelWorkRequestRequestInfo
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
			response, err := c.CancelWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCheckObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CheckObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CheckObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CheckObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CheckObject")
	assert.NoError(t, err)

	type CheckObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.CheckObjectRequest
	}

	var requests []CheckObjectRequestInfo
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
			response, err := c.CheckObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCommitMultipartUpload(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CommitMultipartUpload")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CommitMultipartUpload is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CommitMultipartUpload", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CommitMultipartUpload")
	assert.NoError(t, err)

	type CommitMultipartUploadRequestInfo struct {
		ContainerId string
		Request     objectstorage.CommitMultipartUploadRequest
	}

	var requests []CommitMultipartUploadRequestInfo
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
			response, err := c.CommitMultipartUpload(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCopyObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CopyObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CopyObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CopyObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CopyObject")
	assert.NoError(t, err)

	type CopyObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.CopyObjectRequest
	}

	var requests []CopyObjectRequestInfo
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
			response, err := c.CopyObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCopyPart(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CopyPart")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CopyPart is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CopyPart", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CopyPart")
	assert.NoError(t, err)

	type CopyPartRequestInfo struct {
		ContainerId string
		Request     objectstorage.CopyPartRequest
	}

	var requests []CopyPartRequestInfo
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
			response, err := c.CopyPart(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreateBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CreateBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CreateBucket", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CreateBucket")
	assert.NoError(t, err)

	type CreateBucketRequestInfo struct {
		ContainerId string
		Request     objectstorage.CreateBucketRequest
	}

	var requests []CreateBucketRequestInfo
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
			response, err := c.CreateBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreateMultipartUpload(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CreateMultipartUpload")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateMultipartUpload is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CreateMultipartUpload", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CreateMultipartUpload")
	assert.NoError(t, err)

	type CreateMultipartUploadRequestInfo struct {
		ContainerId string
		Request     objectstorage.CreateMultipartUploadRequest
	}

	var requests []CreateMultipartUploadRequestInfo
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
			response, err := c.CreateMultipartUpload(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreatePreauthenticatedRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CreatePreauthenticatedRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePreauthenticatedRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CreatePreauthenticatedRequest", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CreatePreauthenticatedRequest")
	assert.NoError(t, err)

	type CreatePreauthenticatedRequestRequestInfo struct {
		ContainerId string
		Request     objectstorage.CreatePreauthenticatedRequestRequest
	}

	var requests []CreatePreauthenticatedRequestRequestInfo
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
			response, err := c.CreatePreauthenticatedRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreateReplicationPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CreateReplicationPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateReplicationPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CreateReplicationPolicy", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CreateReplicationPolicy")
	assert.NoError(t, err)

	type CreateReplicationPolicyRequestInfo struct {
		ContainerId string
		Request     objectstorage.CreateReplicationPolicyRequest
	}

	var requests []CreateReplicationPolicyRequestInfo
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
			response, err := c.CreateReplicationPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreateRetentionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "CreateRetentionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRetentionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "CreateRetentionRule", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "CreateRetentionRule")
	assert.NoError(t, err)

	type CreateRetentionRuleRequestInfo struct {
		ContainerId string
		Request     objectstorage.CreateRetentionRuleRequest
	}

	var requests []CreateRetentionRuleRequestInfo
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
			response, err := c.CreateRetentionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "DeleteBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "DeleteBucket", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "DeleteBucket")
	assert.NoError(t, err)

	type DeleteBucketRequestInfo struct {
		ContainerId string
		Request     objectstorage.DeleteBucketRequest
	}

	var requests []DeleteBucketRequestInfo
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
			response, err := c.DeleteBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "DeleteObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "DeleteObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "DeleteObject")
	assert.NoError(t, err)

	type DeleteObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.DeleteObjectRequest
	}

	var requests []DeleteObjectRequestInfo
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
			response, err := c.DeleteObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteObjectLifecyclePolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "DeleteObjectLifecyclePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteObjectLifecyclePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "DeleteObjectLifecyclePolicy", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "DeleteObjectLifecyclePolicy")
	assert.NoError(t, err)

	type DeleteObjectLifecyclePolicyRequestInfo struct {
		ContainerId string
		Request     objectstorage.DeleteObjectLifecyclePolicyRequest
	}

	var requests []DeleteObjectLifecyclePolicyRequestInfo
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
			response, err := c.DeleteObjectLifecyclePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeletePreauthenticatedRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "DeletePreauthenticatedRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePreauthenticatedRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "DeletePreauthenticatedRequest", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "DeletePreauthenticatedRequest")
	assert.NoError(t, err)

	type DeletePreauthenticatedRequestRequestInfo struct {
		ContainerId string
		Request     objectstorage.DeletePreauthenticatedRequestRequest
	}

	var requests []DeletePreauthenticatedRequestRequestInfo
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
			response, err := c.DeletePreauthenticatedRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteReplicationPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "DeleteReplicationPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteReplicationPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "DeleteReplicationPolicy", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "DeleteReplicationPolicy")
	assert.NoError(t, err)

	type DeleteReplicationPolicyRequestInfo struct {
		ContainerId string
		Request     objectstorage.DeleteReplicationPolicyRequest
	}

	var requests []DeleteReplicationPolicyRequestInfo
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
			response, err := c.DeleteReplicationPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteRetentionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "DeleteRetentionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteRetentionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "DeleteRetentionRule", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "DeleteRetentionRule")
	assert.NoError(t, err)

	type DeleteRetentionRuleRequestInfo struct {
		ContainerId string
		Request     objectstorage.DeleteRetentionRuleRequest
	}

	var requests []DeleteRetentionRuleRequestInfo
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
			response, err := c.DeleteRetentionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetBucket", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetBucket")
	assert.NoError(t, err)

	type GetBucketRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetBucketRequest
	}

	var requests []GetBucketRequestInfo
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
			response, err := c.GetBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetBucketOptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetBucketOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBucketOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetBucketOptions", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetBucketOptions")
	assert.NoError(t, err)

	type GetBucketOptionsRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetBucketOptionsRequest
	}

	var requests []GetBucketOptionsRequestInfo
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
			response, err := c.GetBucketOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetNamespace(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetNamespace", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetNamespace")
	assert.NoError(t, err)

	type GetNamespaceRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetNamespaceRequest
	}

	var requests []GetNamespaceRequestInfo
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
			response, err := c.GetNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetNamespaceMetadata(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetNamespaceMetadata")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetNamespaceMetadata is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetNamespaceMetadata", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetNamespaceMetadata")
	assert.NoError(t, err)

	type GetNamespaceMetadataRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetNamespaceMetadataRequest
	}

	var requests []GetNamespaceMetadataRequestInfo
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
			response, err := c.GetNamespaceMetadata(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetObject")
	assert.NoError(t, err)

	type GetObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetObjectRequest
	}

	var requests []GetObjectRequestInfo
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
			response, err := c.GetObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetObjectLifecyclePolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetObjectLifecyclePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetObjectLifecyclePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetObjectLifecyclePolicy", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetObjectLifecyclePolicy")
	assert.NoError(t, err)

	type GetObjectLifecyclePolicyRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetObjectLifecyclePolicyRequest
	}

	var requests []GetObjectLifecyclePolicyRequestInfo
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
			response, err := c.GetObjectLifecyclePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetPreauthenticatedRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetPreauthenticatedRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPreauthenticatedRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetPreauthenticatedRequest", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetPreauthenticatedRequest")
	assert.NoError(t, err)

	type GetPreauthenticatedRequestRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetPreauthenticatedRequestRequest
	}

	var requests []GetPreauthenticatedRequestRequestInfo
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
			response, err := c.GetPreauthenticatedRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetReplicationPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetReplicationPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetReplicationPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetReplicationPolicy", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetReplicationPolicy")
	assert.NoError(t, err)

	type GetReplicationPolicyRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetReplicationPolicyRequest
	}

	var requests []GetReplicationPolicyRequestInfo
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
			response, err := c.GetReplicationPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetRetentionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetRetentionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetRetentionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetRetentionRule", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetRetentionRule")
	assert.NoError(t, err)

	type GetRetentionRuleRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetRetentionRuleRequest
	}

	var requests []GetRetentionRuleRequestInfo
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
			response, err := c.GetRetentionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "GetWorkRequest", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     objectstorage.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
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
			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientHeadBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "HeadBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("HeadBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "HeadBucket", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "HeadBucket")
	assert.NoError(t, err)

	type HeadBucketRequestInfo struct {
		ContainerId string
		Request     objectstorage.HeadBucketRequest
	}

	var requests []HeadBucketRequestInfo
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
			response, err := c.HeadBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientHeadObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "HeadObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("HeadObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "HeadObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "HeadObject")
	assert.NoError(t, err)

	type HeadObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.HeadObjectRequest
	}

	var requests []HeadObjectRequestInfo
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
			response, err := c.HeadObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListBuckets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListBuckets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBuckets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListBuckets", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListBuckets")
	assert.NoError(t, err)

	type ListBucketsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListBucketsRequest
	}

	var requests []ListBucketsRequestInfo
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
				r := req.(*objectstorage.ListBucketsRequest)
				return c.ListBuckets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListBucketsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListBucketsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListMultipartUploadParts(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListMultipartUploadParts")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMultipartUploadParts is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListMultipartUploadParts", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListMultipartUploadParts")
	assert.NoError(t, err)

	type ListMultipartUploadPartsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListMultipartUploadPartsRequest
	}

	var requests []ListMultipartUploadPartsRequestInfo
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
				r := req.(*objectstorage.ListMultipartUploadPartsRequest)
				return c.ListMultipartUploadParts(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListMultipartUploadPartsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListMultipartUploadPartsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListMultipartUploads(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListMultipartUploads")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMultipartUploads is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListMultipartUploads", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListMultipartUploads")
	assert.NoError(t, err)

	type ListMultipartUploadsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListMultipartUploadsRequest
	}

	var requests []ListMultipartUploadsRequestInfo
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
				r := req.(*objectstorage.ListMultipartUploadsRequest)
				return c.ListMultipartUploads(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListMultipartUploadsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListMultipartUploadsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListObjectVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListObjectVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListObjectVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListObjectVersions", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListObjectVersions")
	assert.NoError(t, err)

	type ListObjectVersionsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListObjectVersionsRequest
	}

	var requests []ListObjectVersionsRequestInfo
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
				r := req.(*objectstorage.ListObjectVersionsRequest)
				return c.ListObjectVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListObjectVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListObjectVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListObjects(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListObjects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListObjects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListObjects", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListObjects")
	assert.NoError(t, err)

	type ListObjectsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListObjectsRequest
	}

	var requests []ListObjectsRequestInfo
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
			response, err := c.ListObjects(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListPreauthenticatedRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListPreauthenticatedRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPreauthenticatedRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListPreauthenticatedRequests", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListPreauthenticatedRequests")
	assert.NoError(t, err)

	type ListPreauthenticatedRequestsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListPreauthenticatedRequestsRequest
	}

	var requests []ListPreauthenticatedRequestsRequestInfo
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
				r := req.(*objectstorage.ListPreauthenticatedRequestsRequest)
				return c.ListPreauthenticatedRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListPreauthenticatedRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListPreauthenticatedRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListReplicationPolicies(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListReplicationPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListReplicationPolicies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListReplicationPolicies", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListReplicationPolicies")
	assert.NoError(t, err)

	type ListReplicationPoliciesRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListReplicationPoliciesRequest
	}

	var requests []ListReplicationPoliciesRequestInfo
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
				r := req.(*objectstorage.ListReplicationPoliciesRequest)
				return c.ListReplicationPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListReplicationPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListReplicationPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListReplicationSources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListReplicationSources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListReplicationSources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListReplicationSources", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListReplicationSources")
	assert.NoError(t, err)

	type ListReplicationSourcesRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListReplicationSourcesRequest
	}

	var requests []ListReplicationSourcesRequestInfo
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
				r := req.(*objectstorage.ListReplicationSourcesRequest)
				return c.ListReplicationSources(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListReplicationSourcesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListReplicationSourcesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListRetentionRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListRetentionRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRetentionRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListRetentionRules", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListRetentionRules")
	assert.NoError(t, err)

	type ListRetentionRulesRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListRetentionRulesRequest
	}

	var requests []ListRetentionRulesRequestInfo
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
				r := req.(*objectstorage.ListRetentionRulesRequest)
				return c.ListRetentionRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListRetentionRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListRetentionRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListWorkRequestErrors", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListWorkRequestErrorsRequest
	}

	var requests []ListWorkRequestErrorsRequestInfo
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
				r := req.(*objectstorage.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListWorkRequestLogs", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListWorkRequestLogsRequest
	}

	var requests []ListWorkRequestLogsRequestInfo
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
				r := req.(*objectstorage.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ListWorkRequests", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     objectstorage.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
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
				r := req.(*objectstorage.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]objectstorage.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(objectstorage.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientMakeBucketWritable(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "MakeBucketWritable")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MakeBucketWritable is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "MakeBucketWritable", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "MakeBucketWritable")
	assert.NoError(t, err)

	type MakeBucketWritableRequestInfo struct {
		ContainerId string
		Request     objectstorage.MakeBucketWritableRequest
	}

	var requests []MakeBucketWritableRequestInfo
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
			response, err := c.MakeBucketWritable(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientMergeObjectMetadata(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "MergeObjectMetadata")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("MergeObjectMetadata is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "MergeObjectMetadata", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "MergeObjectMetadata")
	assert.NoError(t, err)

	type MergeObjectMetadataRequestInfo struct {
		ContainerId string
		Request     objectstorage.MergeObjectMetadataRequest
	}

	var requests []MergeObjectMetadataRequestInfo
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
			response, err := c.MergeObjectMetadata(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientPutObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "PutObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PutObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "PutObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "PutObject")
	assert.NoError(t, err)

	type PutObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.PutObjectRequest
	}

	var requests []PutObjectRequestInfo
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
			response, err := c.PutObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientPutObjectLifecyclePolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "PutObjectLifecyclePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PutObjectLifecyclePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "PutObjectLifecyclePolicy", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "PutObjectLifecyclePolicy")
	assert.NoError(t, err)

	type PutObjectLifecyclePolicyRequestInfo struct {
		ContainerId string
		Request     objectstorage.PutObjectLifecyclePolicyRequest
	}

	var requests []PutObjectLifecyclePolicyRequestInfo
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
			response, err := c.PutObjectLifecyclePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientReencryptBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ReencryptBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReencryptBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ReencryptBucket", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ReencryptBucket")
	assert.NoError(t, err)

	type ReencryptBucketRequestInfo struct {
		ContainerId string
		Request     objectstorage.ReencryptBucketRequest
	}

	var requests []ReencryptBucketRequestInfo
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
			response, err := c.ReencryptBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientReencryptObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ReencryptObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReencryptObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ReencryptObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ReencryptObject")
	assert.NoError(t, err)

	type ReencryptObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.ReencryptObjectRequest
	}

	var requests []ReencryptObjectRequestInfo
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
			response, err := c.ReencryptObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientRenameObject(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "RenameObject")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RenameObject is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "RenameObject", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "RenameObject")
	assert.NoError(t, err)

	type RenameObjectRequestInfo struct {
		ContainerId string
		Request     objectstorage.RenameObjectRequest
	}

	var requests []RenameObjectRequestInfo
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
			response, err := c.RenameObject(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientReplaceObjectMetadata(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "ReplaceObjectMetadata")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ReplaceObjectMetadata is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "ReplaceObjectMetadata", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "ReplaceObjectMetadata")
	assert.NoError(t, err)

	type ReplaceObjectMetadataRequestInfo struct {
		ContainerId string
		Request     objectstorage.ReplaceObjectMetadataRequest
	}

	var requests []ReplaceObjectMetadataRequestInfo
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
			response, err := c.ReplaceObjectMetadata(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientRestoreObjects(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "RestoreObjects")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RestoreObjects is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "RestoreObjects", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "RestoreObjects")
	assert.NoError(t, err)

	type RestoreObjectsRequestInfo struct {
		ContainerId string
		Request     objectstorage.RestoreObjectsRequest
	}

	var requests []RestoreObjectsRequestInfo
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
			response, err := c.RestoreObjects(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateBucket(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "UpdateBucket")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBucket is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "UpdateBucket", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "UpdateBucket")
	assert.NoError(t, err)

	type UpdateBucketRequestInfo struct {
		ContainerId string
		Request     objectstorage.UpdateBucketRequest
	}

	var requests []UpdateBucketRequestInfo
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
			response, err := c.UpdateBucket(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateBucketOptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "UpdateBucketOptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBucketOptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "UpdateBucketOptions", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "UpdateBucketOptions")
	assert.NoError(t, err)

	type UpdateBucketOptionsRequestInfo struct {
		ContainerId string
		Request     objectstorage.UpdateBucketOptionsRequest
	}

	var requests []UpdateBucketOptionsRequestInfo
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
			response, err := c.UpdateBucketOptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateNamespaceMetadata(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "UpdateNamespaceMetadata")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateNamespaceMetadata is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "UpdateNamespaceMetadata", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "UpdateNamespaceMetadata")
	assert.NoError(t, err)

	type UpdateNamespaceMetadataRequestInfo struct {
		ContainerId string
		Request     objectstorage.UpdateNamespaceMetadataRequest
	}

	var requests []UpdateNamespaceMetadataRequestInfo
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
			response, err := c.UpdateNamespaceMetadata(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateObjectStorageTier(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "UpdateObjectStorageTier")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateObjectStorageTier is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "UpdateObjectStorageTier", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "UpdateObjectStorageTier")
	assert.NoError(t, err)

	type UpdateObjectStorageTierRequestInfo struct {
		ContainerId string
		Request     objectstorage.UpdateObjectStorageTierRequest
	}

	var requests []UpdateObjectStorageTierRequestInfo
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
			response, err := c.UpdateObjectStorageTier(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateRetentionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "UpdateRetentionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateRetentionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "UpdateRetentionRule", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "UpdateRetentionRule")
	assert.NoError(t, err)

	type UpdateRetentionRuleRequestInfo struct {
		ContainerId string
		Request     objectstorage.UpdateRetentionRuleRequest
	}

	var requests []UpdateRetentionRuleRequestInfo
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
			response, err := c.UpdateRetentionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUploadPart(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("objectstorage", "UploadPart")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UploadPart is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("objectstorage", "ObjectStorage", "UploadPart", createObjectStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(objectstorage.ObjectStorageClient)

	body, err := testClient.getRequests("objectstorage", "UploadPart")
	assert.NoError(t, err)

	type UploadPartRequestInfo struct {
		ContainerId string
		Request     objectstorage.UploadPartRequest
	}

	var requests []UploadPartRequestInfo
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
			response, err := c.UploadPart(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
