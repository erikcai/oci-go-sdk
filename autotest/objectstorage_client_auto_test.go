package autotest

import (
    "github.com/oracle/oci-go-sdk/objectstorage"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientAbortMultipartUpload(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "AbortMultipartUpload")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("AbortMultipartUpload is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "AbortMultipartUpload")
    assert.NoError(t, err)

    type AbortMultipartUploadRequestInfo struct {
        ContainerId string
        Request objectstorage.AbortMultipartUploadRequest
    }

    var requests []AbortMultipartUploadRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.AbortMultipartUpload(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCancelWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "CancelWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CancelWorkRequest is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "CancelWorkRequest")
    assert.NoError(t, err)

    type CancelWorkRequestRequestInfo struct {
        ContainerId string
        Request objectstorage.CancelWorkRequestRequest
    }

    var requests []CancelWorkRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CancelWorkRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCommitMultipartUpload(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "CommitMultipartUpload")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CommitMultipartUpload is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "CommitMultipartUpload")
    assert.NoError(t, err)

    type CommitMultipartUploadRequestInfo struct {
        ContainerId string
        Request objectstorage.CommitMultipartUploadRequest
    }

    var requests []CommitMultipartUploadRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CommitMultipartUpload(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCopyObject(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "CopyObject")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CopyObject is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "CopyObject")
    assert.NoError(t, err)

    type CopyObjectRequestInfo struct {
        ContainerId string
        Request objectstorage.CopyObjectRequest
    }

    var requests []CopyObjectRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CopyObject(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreateBucket(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "CreateBucket")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateBucket is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "CreateBucket")
    assert.NoError(t, err)

    type CreateBucketRequestInfo struct {
        ContainerId string
        Request objectstorage.CreateBucketRequest
    }

    var requests []CreateBucketRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateBucket(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreateMultipartUpload(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "CreateMultipartUpload")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateMultipartUpload is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "CreateMultipartUpload")
    assert.NoError(t, err)

    type CreateMultipartUploadRequestInfo struct {
        ContainerId string
        Request objectstorage.CreateMultipartUploadRequest
    }

    var requests []CreateMultipartUploadRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateMultipartUpload(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientCreatePreauthenticatedRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "CreatePreauthenticatedRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreatePreauthenticatedRequest is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "CreatePreauthenticatedRequest")
    assert.NoError(t, err)

    type CreatePreauthenticatedRequestRequestInfo struct {
        ContainerId string
        Request objectstorage.CreatePreauthenticatedRequestRequest
    }

    var requests []CreatePreauthenticatedRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreatePreauthenticatedRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteBucket(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "DeleteBucket")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteBucket is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "DeleteBucket")
    assert.NoError(t, err)

    type DeleteBucketRequestInfo struct {
        ContainerId string
        Request objectstorage.DeleteBucketRequest
    }

    var requests []DeleteBucketRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteBucket(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteObject(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "DeleteObject")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteObject is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "DeleteObject")
    assert.NoError(t, err)

    type DeleteObjectRequestInfo struct {
        ContainerId string
        Request objectstorage.DeleteObjectRequest
    }

    var requests []DeleteObjectRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteObject(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeleteObjectLifecyclePolicy(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "DeleteObjectLifecyclePolicy")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteObjectLifecyclePolicy is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "DeleteObjectLifecyclePolicy")
    assert.NoError(t, err)

    type DeleteObjectLifecyclePolicyRequestInfo struct {
        ContainerId string
        Request objectstorage.DeleteObjectLifecyclePolicyRequest
    }

    var requests []DeleteObjectLifecyclePolicyRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteObjectLifecyclePolicy(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientDeletePreauthenticatedRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "DeletePreauthenticatedRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeletePreauthenticatedRequest is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "DeletePreauthenticatedRequest")
    assert.NoError(t, err)

    type DeletePreauthenticatedRequestRequestInfo struct {
        ContainerId string
        Request objectstorage.DeletePreauthenticatedRequestRequest
    }

    var requests []DeletePreauthenticatedRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeletePreauthenticatedRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetBucket(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetBucket")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetBucket is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetBucket")
    assert.NoError(t, err)

    type GetBucketRequestInfo struct {
        ContainerId string
        Request objectstorage.GetBucketRequest
    }

    var requests []GetBucketRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetBucket(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetNamespace(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetNamespace")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetNamespace is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetNamespace")
    assert.NoError(t, err)

    type GetNamespaceRequestInfo struct {
        ContainerId string
        Request objectstorage.GetNamespaceRequest
    }

    var requests []GetNamespaceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetNamespace(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetNamespaceMetadata(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetNamespaceMetadata")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetNamespaceMetadata is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetNamespaceMetadata")
    assert.NoError(t, err)

    type GetNamespaceMetadataRequestInfo struct {
        ContainerId string
        Request objectstorage.GetNamespaceMetadataRequest
    }

    var requests []GetNamespaceMetadataRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetNamespaceMetadata(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetObject(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetObject")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetObject is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetObject")
    assert.NoError(t, err)

    type GetObjectRequestInfo struct {
        ContainerId string
        Request objectstorage.GetObjectRequest
    }

    var requests []GetObjectRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetObject(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetObjectLifecyclePolicy(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetObjectLifecyclePolicy")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetObjectLifecyclePolicy is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetObjectLifecyclePolicy")
    assert.NoError(t, err)

    type GetObjectLifecyclePolicyRequestInfo struct {
        ContainerId string
        Request objectstorage.GetObjectLifecyclePolicyRequest
    }

    var requests []GetObjectLifecyclePolicyRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetObjectLifecyclePolicy(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetPreauthenticatedRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetPreauthenticatedRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetPreauthenticatedRequest is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetPreauthenticatedRequest")
    assert.NoError(t, err)

    type GetPreauthenticatedRequestRequestInfo struct {
        ContainerId string
        Request objectstorage.GetPreauthenticatedRequestRequest
    }

    var requests []GetPreauthenticatedRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetPreauthenticatedRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientGetWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "GetWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetWorkRequest is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "GetWorkRequest")
    assert.NoError(t, err)

    type GetWorkRequestRequestInfo struct {
        ContainerId string
        Request objectstorage.GetWorkRequestRequest
    }

    var requests []GetWorkRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetWorkRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientHeadBucket(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "HeadBucket")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("HeadBucket is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "HeadBucket")
    assert.NoError(t, err)

    type HeadBucketRequestInfo struct {
        ContainerId string
        Request objectstorage.HeadBucketRequest
    }

    var requests []HeadBucketRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.HeadBucket(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientHeadObject(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "HeadObject")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("HeadObject is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "HeadObject")
    assert.NoError(t, err)

    type HeadObjectRequestInfo struct {
        ContainerId string
        Request objectstorage.HeadObjectRequest
    }

    var requests []HeadObjectRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.HeadObject(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListBuckets(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListBuckets")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListBuckets is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListBuckets")
    assert.NoError(t, err)

    type ListBucketsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListBucketsRequest
    }

    var requests []ListBucketsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListMultipartUploadParts(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListMultipartUploadParts")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListMultipartUploadParts is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListMultipartUploadParts")
    assert.NoError(t, err)

    type ListMultipartUploadPartsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListMultipartUploadPartsRequest
    }

    var requests []ListMultipartUploadPartsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListMultipartUploads(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListMultipartUploads")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListMultipartUploads is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListMultipartUploads")
    assert.NoError(t, err)

    type ListMultipartUploadsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListMultipartUploadsRequest
    }

    var requests []ListMultipartUploadsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListObjects(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListObjects")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListObjects is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListObjects")
    assert.NoError(t, err)

    type ListObjectsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListObjectsRequest
    }

    var requests []ListObjectsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ListObjects(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListPreauthenticatedRequests(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListPreauthenticatedRequests")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListPreauthenticatedRequests is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListPreauthenticatedRequests")
    assert.NoError(t, err)

    type ListPreauthenticatedRequestsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListPreauthenticatedRequestsRequest
    }

    var requests []ListPreauthenticatedRequestsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListWorkRequestErrors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListWorkRequestErrors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestErrors is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListWorkRequestErrors")
    assert.NoError(t, err)

    type ListWorkRequestErrorsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListWorkRequestErrorsRequest
    }

    var requests []ListWorkRequestErrorsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListWorkRequestLogs(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListWorkRequestLogs")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestLogs is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListWorkRequestLogs")
    assert.NoError(t, err)

    type ListWorkRequestLogsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListWorkRequestLogsRequest
    }

    var requests []ListWorkRequestLogsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientListWorkRequests(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ListWorkRequests")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequests is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ListWorkRequests")
    assert.NoError(t, err)

    type ListWorkRequestsRequestInfo struct {
        ContainerId string
        Request objectstorage.ListWorkRequestsRequest
    }

    var requests []ListWorkRequestsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
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

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientMergeObjectMetadata(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "MergeObjectMetadata")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("MergeObjectMetadata is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "MergeObjectMetadata")
    assert.NoError(t, err)

    type MergeObjectMetadataRequestInfo struct {
        ContainerId string
        Request objectstorage.MergeObjectMetadataRequest
    }

    var requests []MergeObjectMetadataRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.MergeObjectMetadata(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientPutObject(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "PutObject")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("PutObject is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "PutObject")
    assert.NoError(t, err)

    type PutObjectRequestInfo struct {
        ContainerId string
        Request objectstorage.PutObjectRequest
    }

    var requests []PutObjectRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.PutObject(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientPutObjectLifecyclePolicy(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "PutObjectLifecyclePolicy")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("PutObjectLifecyclePolicy is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "PutObjectLifecyclePolicy")
    assert.NoError(t, err)

    type PutObjectLifecyclePolicyRequestInfo struct {
        ContainerId string
        Request objectstorage.PutObjectLifecyclePolicyRequest
    }

    var requests []PutObjectLifecyclePolicyRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.PutObjectLifecyclePolicy(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientRenameObject(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "RenameObject")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("RenameObject is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "RenameObject")
    assert.NoError(t, err)

    type RenameObjectRequestInfo struct {
        ContainerId string
        Request objectstorage.RenameObjectRequest
    }

    var requests []RenameObjectRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.RenameObject(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientReplaceObjectMetadata(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "ReplaceObjectMetadata")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ReplaceObjectMetadata is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "ReplaceObjectMetadata")
    assert.NoError(t, err)

    type ReplaceObjectMetadataRequestInfo struct {
        ContainerId string
        Request objectstorage.ReplaceObjectMetadataRequest
    }

    var requests []ReplaceObjectMetadataRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ReplaceObjectMetadata(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientRestoreObjects(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "RestoreObjects")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("RestoreObjects is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "RestoreObjects")
    assert.NoError(t, err)

    type RestoreObjectsRequestInfo struct {
        ContainerId string
        Request objectstorage.RestoreObjectsRequest
    }

    var requests []RestoreObjectsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.RestoreObjects(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateBucket(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "UpdateBucket")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateBucket is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "UpdateBucket")
    assert.NoError(t, err)

    type UpdateBucketRequestInfo struct {
        ContainerId string
        Request objectstorage.UpdateBucketRequest
    }

    var requests []UpdateBucketRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateBucket(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUpdateNamespaceMetadata(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "UpdateNamespaceMetadata")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateNamespaceMetadata is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "UpdateNamespaceMetadata")
    assert.NoError(t, err)

    type UpdateNamespaceMetadataRequestInfo struct {
        ContainerId string
        Request objectstorage.UpdateNamespaceMetadataRequest
    }

    var requests []UpdateNamespaceMetadataRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateNamespaceMetadata(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo email="opc_casper_us_grp@oracle.com" jiraProject="CASPER" opsJiraProject="IOS"
func TestObjectStorageClientUploadPart(t *testing.T) {
    enabled, err := testClient.isApiEnabled("objectstorage", "UploadPart")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UploadPart is not enabled by the testing service")
    }
    c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("objectstorage", "UploadPart")
    assert.NoError(t, err)

    type UploadPartRequestInfo struct {
        ContainerId string
        Request objectstorage.UploadPartRequest
    }

    var requests []UploadPartRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UploadPart(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
