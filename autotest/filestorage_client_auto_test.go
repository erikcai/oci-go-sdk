package autotest

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"github.com/oracle/oci-go-sdk/v29/filestorage"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createFileStorageClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := filestorage.NewFileStorageClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientChangeFileSystemCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ChangeFileSystemCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeFileSystemCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ChangeFileSystemCompartment", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ChangeFileSystemCompartment")
	assert.NoError(t, err)

	type ChangeFileSystemCompartmentRequestInfo struct {
		ContainerId string
		Request     filestorage.ChangeFileSystemCompartmentRequest
	}

	var requests []ChangeFileSystemCompartmentRequestInfo
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
			response, err := c.ChangeFileSystemCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientChangeMountTargetCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ChangeMountTargetCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeMountTargetCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ChangeMountTargetCompartment", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ChangeMountTargetCompartment")
	assert.NoError(t, err)

	type ChangeMountTargetCompartmentRequestInfo struct {
		ContainerId string
		Request     filestorage.ChangeMountTargetCompartmentRequest
	}

	var requests []ChangeMountTargetCompartmentRequestInfo
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
			response, err := c.ChangeMountTargetCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientChangeOutboundConnectorCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ChangeOutboundConnectorCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeOutboundConnectorCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ChangeOutboundConnectorCompartment", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ChangeOutboundConnectorCompartment")
	assert.NoError(t, err)

	type ChangeOutboundConnectorCompartmentRequestInfo struct {
		ContainerId string
		Request     filestorage.ChangeOutboundConnectorCompartmentRequest
	}

	var requests []ChangeOutboundConnectorCompartmentRequestInfo
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
			response, err := c.ChangeOutboundConnectorCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientCreateExport(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "CreateExport")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateExport is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "CreateExport", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "CreateExport")
	assert.NoError(t, err)

	type CreateExportRequestInfo struct {
		ContainerId string
		Request     filestorage.CreateExportRequest
	}

	var requests []CreateExportRequestInfo
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
			response, err := c.CreateExport(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientCreateFileSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "CreateFileSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateFileSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "CreateFileSystem", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "CreateFileSystem")
	assert.NoError(t, err)

	type CreateFileSystemRequestInfo struct {
		ContainerId string
		Request     filestorage.CreateFileSystemRequest
	}

	var requests []CreateFileSystemRequestInfo
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
			response, err := c.CreateFileSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientCreateMountTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "CreateMountTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateMountTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "CreateMountTarget", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "CreateMountTarget")
	assert.NoError(t, err)

	type CreateMountTargetRequestInfo struct {
		ContainerId string
		Request     filestorage.CreateMountTargetRequest
	}

	var requests []CreateMountTargetRequestInfo
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
			response, err := c.CreateMountTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientCreateOutboundConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "CreateOutboundConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOutboundConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "CreateOutboundConnector", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "CreateOutboundConnector")
	assert.NoError(t, err)

	type CreateOutboundConnectorRequestInfo struct {
		ContainerId string
		Request     filestorage.CreateOutboundConnectorRequest
	}

	var requests []CreateOutboundConnectorRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateOutboundConnectorRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateOutboundConnectorDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "connectorType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"LDAPBIND": &filestorage.CreateLdapBindAccountDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			if withRetry == true {
				retryPolicy = retryPolicyForTests()
			}
			req.Request.RequestMetadata.RetryPolicy = retryPolicy
			response, err := c.CreateOutboundConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientCreateSnapshot(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "CreateSnapshot")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSnapshot is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "CreateSnapshot", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "CreateSnapshot")
	assert.NoError(t, err)

	type CreateSnapshotRequestInfo struct {
		ContainerId string
		Request     filestorage.CreateSnapshotRequest
	}

	var requests []CreateSnapshotRequestInfo
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
			response, err := c.CreateSnapshot(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientDeleteExport(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "DeleteExport")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteExport is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "DeleteExport", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "DeleteExport")
	assert.NoError(t, err)

	type DeleteExportRequestInfo struct {
		ContainerId string
		Request     filestorage.DeleteExportRequest
	}

	var requests []DeleteExportRequestInfo
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
			response, err := c.DeleteExport(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientDeleteFileSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "DeleteFileSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteFileSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "DeleteFileSystem", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "DeleteFileSystem")
	assert.NoError(t, err)

	type DeleteFileSystemRequestInfo struct {
		ContainerId string
		Request     filestorage.DeleteFileSystemRequest
	}

	var requests []DeleteFileSystemRequestInfo
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
			response, err := c.DeleteFileSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientDeleteMountTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "DeleteMountTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteMountTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "DeleteMountTarget", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "DeleteMountTarget")
	assert.NoError(t, err)

	type DeleteMountTargetRequestInfo struct {
		ContainerId string
		Request     filestorage.DeleteMountTargetRequest
	}

	var requests []DeleteMountTargetRequestInfo
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
			response, err := c.DeleteMountTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientDeleteOutboundConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "DeleteOutboundConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteOutboundConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "DeleteOutboundConnector", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "DeleteOutboundConnector")
	assert.NoError(t, err)

	type DeleteOutboundConnectorRequestInfo struct {
		ContainerId string
		Request     filestorage.DeleteOutboundConnectorRequest
	}

	var requests []DeleteOutboundConnectorRequestInfo
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
			response, err := c.DeleteOutboundConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientDeleteSnapshot(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "DeleteSnapshot")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSnapshot is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "DeleteSnapshot", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "DeleteSnapshot")
	assert.NoError(t, err)

	type DeleteSnapshotRequestInfo struct {
		ContainerId string
		Request     filestorage.DeleteSnapshotRequest
	}

	var requests []DeleteSnapshotRequestInfo
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
			response, err := c.DeleteSnapshot(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientDiscardKerberosKeytab(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "DiscardKerberosKeytab")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DiscardKerberosKeytab is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "DiscardKerberosKeytab", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "DiscardKerberosKeytab")
	assert.NoError(t, err)

	type DiscardKerberosKeytabRequestInfo struct {
		ContainerId string
		Request     filestorage.DiscardKerberosKeytabRequest
	}

	var requests []DiscardKerberosKeytabRequestInfo
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
			response, err := c.DiscardKerberosKeytab(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientGetExport(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "GetExport")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExport is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "GetExport", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "GetExport")
	assert.NoError(t, err)

	type GetExportRequestInfo struct {
		ContainerId string
		Request     filestorage.GetExportRequest
	}

	var requests []GetExportRequestInfo
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
			response, err := c.GetExport(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientGetExportSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "GetExportSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetExportSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "GetExportSet", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "GetExportSet")
	assert.NoError(t, err)

	type GetExportSetRequestInfo struct {
		ContainerId string
		Request     filestorage.GetExportSetRequest
	}

	var requests []GetExportSetRequestInfo
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
			response, err := c.GetExportSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientGetFileSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "GetFileSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetFileSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "GetFileSystem", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "GetFileSystem")
	assert.NoError(t, err)

	type GetFileSystemRequestInfo struct {
		ContainerId string
		Request     filestorage.GetFileSystemRequest
	}

	var requests []GetFileSystemRequestInfo
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
			response, err := c.GetFileSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientGetMountTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "GetMountTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMountTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "GetMountTarget", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "GetMountTarget")
	assert.NoError(t, err)

	type GetMountTargetRequestInfo struct {
		ContainerId string
		Request     filestorage.GetMountTargetRequest
	}

	var requests []GetMountTargetRequestInfo
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
			response, err := c.GetMountTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientGetOutboundConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "GetOutboundConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetOutboundConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "GetOutboundConnector", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "GetOutboundConnector")
	assert.NoError(t, err)

	type GetOutboundConnectorRequestInfo struct {
		ContainerId string
		Request     filestorage.GetOutboundConnectorRequest
	}

	var requests []GetOutboundConnectorRequestInfo
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
			response, err := c.GetOutboundConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientGetSnapshot(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "GetSnapshot")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSnapshot is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "GetSnapshot", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "GetSnapshot")
	assert.NoError(t, err)

	type GetSnapshotRequestInfo struct {
		ContainerId string
		Request     filestorage.GetSnapshotRequest
	}

	var requests []GetSnapshotRequestInfo
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
			response, err := c.GetSnapshot(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientListExportSets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ListExportSets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListExportSets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ListExportSets", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ListExportSets")
	assert.NoError(t, err)

	type ListExportSetsRequestInfo struct {
		ContainerId string
		Request     filestorage.ListExportSetsRequest
	}

	var requests []ListExportSetsRequestInfo
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
				r := req.(*filestorage.ListExportSetsRequest)
				return c.ListExportSets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]filestorage.ListExportSetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(filestorage.ListExportSetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientListExports(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ListExports")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListExports is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ListExports", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ListExports")
	assert.NoError(t, err)

	type ListExportsRequestInfo struct {
		ContainerId string
		Request     filestorage.ListExportsRequest
	}

	var requests []ListExportsRequestInfo
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
				r := req.(*filestorage.ListExportsRequest)
				return c.ListExports(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]filestorage.ListExportsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(filestorage.ListExportsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientListFileSystems(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ListFileSystems")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFileSystems is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ListFileSystems", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ListFileSystems")
	assert.NoError(t, err)

	type ListFileSystemsRequestInfo struct {
		ContainerId string
		Request     filestorage.ListFileSystemsRequest
	}

	var requests []ListFileSystemsRequestInfo
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
				r := req.(*filestorage.ListFileSystemsRequest)
				return c.ListFileSystems(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]filestorage.ListFileSystemsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(filestorage.ListFileSystemsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientListMountTargets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ListMountTargets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMountTargets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ListMountTargets", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ListMountTargets")
	assert.NoError(t, err)

	type ListMountTargetsRequestInfo struct {
		ContainerId string
		Request     filestorage.ListMountTargetsRequest
	}

	var requests []ListMountTargetsRequestInfo
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
				r := req.(*filestorage.ListMountTargetsRequest)
				return c.ListMountTargets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]filestorage.ListMountTargetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(filestorage.ListMountTargetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientListOutboundConnectors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ListOutboundConnectors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOutboundConnectors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ListOutboundConnectors", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ListOutboundConnectors")
	assert.NoError(t, err)

	type ListOutboundConnectorsRequestInfo struct {
		ContainerId string
		Request     filestorage.ListOutboundConnectorsRequest
	}

	var requests []ListOutboundConnectorsRequestInfo
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
				r := req.(*filestorage.ListOutboundConnectorsRequest)
				return c.ListOutboundConnectors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]filestorage.ListOutboundConnectorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(filestorage.ListOutboundConnectorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientListSnapshots(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ListSnapshots")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSnapshots is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ListSnapshots", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ListSnapshots")
	assert.NoError(t, err)

	type ListSnapshotsRequestInfo struct {
		ContainerId string
		Request     filestorage.ListSnapshotsRequest
	}

	var requests []ListSnapshotsRequestInfo
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
				r := req.(*filestorage.ListSnapshotsRequest)
				return c.ListSnapshots(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]filestorage.ListSnapshotsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(filestorage.ListSnapshotsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientModifyRootdirAttributes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "ModifyRootdirAttributes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ModifyRootdirAttributes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "ModifyRootdirAttributes", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "ModifyRootdirAttributes")
	assert.NoError(t, err)

	type ModifyRootdirAttributesRequestInfo struct {
		ContainerId string
		Request     filestorage.ModifyRootdirAttributesRequest
	}

	var requests []ModifyRootdirAttributesRequestInfo
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
			response, err := c.ModifyRootdirAttributes(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientTestOutboundConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "TestOutboundConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TestOutboundConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "TestOutboundConnector", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "TestOutboundConnector")
	assert.NoError(t, err)

	type TestOutboundConnectorRequestInfo struct {
		ContainerId string
		Request     filestorage.TestOutboundConnectorRequest
	}

	var requests []TestOutboundConnectorRequestInfo
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
			response, err := c.TestOutboundConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUpdateExport(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UpdateExport")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExport is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UpdateExport", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UpdateExport")
	assert.NoError(t, err)

	type UpdateExportRequestInfo struct {
		ContainerId string
		Request     filestorage.UpdateExportRequest
	}

	var requests []UpdateExportRequestInfo
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
			response, err := c.UpdateExport(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUpdateExportSet(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UpdateExportSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateExportSet is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UpdateExportSet", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UpdateExportSet")
	assert.NoError(t, err)

	type UpdateExportSetRequestInfo struct {
		ContainerId string
		Request     filestorage.UpdateExportSetRequest
	}

	var requests []UpdateExportSetRequestInfo
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
			response, err := c.UpdateExportSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUpdateFileSystem(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UpdateFileSystem")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateFileSystem is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UpdateFileSystem", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UpdateFileSystem")
	assert.NoError(t, err)

	type UpdateFileSystemRequestInfo struct {
		ContainerId string
		Request     filestorage.UpdateFileSystemRequest
	}

	var requests []UpdateFileSystemRequestInfo
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
			response, err := c.UpdateFileSystem(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUpdateMountTarget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UpdateMountTarget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateMountTarget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UpdateMountTarget", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UpdateMountTarget")
	assert.NoError(t, err)

	type UpdateMountTargetRequestInfo struct {
		ContainerId string
		Request     filestorage.UpdateMountTargetRequest
	}

	var requests []UpdateMountTargetRequestInfo
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
			response, err := c.UpdateMountTarget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUpdateOutboundConnector(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UpdateOutboundConnector")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOutboundConnector is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UpdateOutboundConnector", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UpdateOutboundConnector")
	assert.NoError(t, err)

	type UpdateOutboundConnectorRequestInfo struct {
		ContainerId string
		Request     filestorage.UpdateOutboundConnectorRequest
	}

	var requests []UpdateOutboundConnectorRequestInfo
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
			response, err := c.UpdateOutboundConnector(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUpdateSnapshot(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UpdateSnapshot")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSnapshot is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UpdateSnapshot", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UpdateSnapshot")
	assert.NoError(t, err)

	type UpdateSnapshotRequestInfo struct {
		ContainerId string
		Request     filestorage.UpdateSnapshotRequest
	}

	var requests []UpdateSnapshotRequestInfo
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
			response, err := c.UpdateSnapshot(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_ffsw_us_grp@oracle.com" jiraProject="FFSW" opsJiraProject="FSS"
func TestFileStorageClientUploadKerberosKeytab(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("filestorage", "UploadKerberosKeytab")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UploadKerberosKeytab is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("filestorage", "FileStorage", "UploadKerberosKeytab", createFileStorageClientWithProvider)
	assert.NoError(t, err)
	c := cc.(filestorage.FileStorageClient)

	body, err := testClient.getRequests("filestorage", "UploadKerberosKeytab")
	assert.NoError(t, err)

	type UploadKerberosKeytabRequestInfo struct {
		ContainerId string
		Request     filestorage.UploadKerberosKeytabRequest
	}

	var requests []UploadKerberosKeytabRequestInfo
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
			response, err := c.UploadKerberosKeytab(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
