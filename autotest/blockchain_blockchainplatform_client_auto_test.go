package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/blockchain"
	"github.com/oracle/oci-go-sdk/v25/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBlockchainPlatformClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := blockchain.NewBlockchainPlatformClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientChangeBlockchainPlatformCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ChangeBlockchainPlatformCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeBlockchainPlatformCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ChangeBlockchainPlatformCompartment", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ChangeBlockchainPlatformCompartment")
	assert.NoError(t, err)

	type ChangeBlockchainPlatformCompartmentRequestInfo struct {
		ContainerId string
		Request     blockchain.ChangeBlockchainPlatformCompartmentRequest
	}

	var requests []ChangeBlockchainPlatformCompartmentRequestInfo
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
			response, err := c.ChangeBlockchainPlatformCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientCreateBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "CreateBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "CreateBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "CreateBlockchainPlatform")
	assert.NoError(t, err)

	type CreateBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.CreateBlockchainPlatformRequest
	}

	var requests []CreateBlockchainPlatformRequestInfo
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
			response, err := c.CreateBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientCreateOsn(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "CreateOsn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOsn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "CreateOsn", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "CreateOsn")
	assert.NoError(t, err)

	type CreateOsnRequestInfo struct {
		ContainerId string
		Request     blockchain.CreateOsnRequest
	}

	var requests []CreateOsnRequestInfo
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
			response, err := c.CreateOsn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientCreatePeer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "CreatePeer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePeer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "CreatePeer", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "CreatePeer")
	assert.NoError(t, err)

	type CreatePeerRequestInfo struct {
		ContainerId string
		Request     blockchain.CreatePeerRequest
	}

	var requests []CreatePeerRequestInfo
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
			response, err := c.CreatePeer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientDeleteBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "DeleteBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "DeleteBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "DeleteBlockchainPlatform")
	assert.NoError(t, err)

	type DeleteBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.DeleteBlockchainPlatformRequest
	}

	var requests []DeleteBlockchainPlatformRequestInfo
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
			response, err := c.DeleteBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientDeleteOsn(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "DeleteOsn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteOsn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "DeleteOsn", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "DeleteOsn")
	assert.NoError(t, err)

	type DeleteOsnRequestInfo struct {
		ContainerId string
		Request     blockchain.DeleteOsnRequest
	}

	var requests []DeleteOsnRequestInfo
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
			response, err := c.DeleteOsn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientDeletePeer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "DeletePeer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePeer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "DeletePeer", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "DeletePeer")
	assert.NoError(t, err)

	type DeletePeerRequestInfo struct {
		ContainerId string
		Request     blockchain.DeletePeerRequest
	}

	var requests []DeletePeerRequestInfo
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
			response, err := c.DeletePeer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientDeleteWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "DeleteWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "DeleteWorkRequest", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "DeleteWorkRequest")
	assert.NoError(t, err)

	type DeleteWorkRequestRequestInfo struct {
		ContainerId string
		Request     blockchain.DeleteWorkRequestRequest
	}

	var requests []DeleteWorkRequestRequestInfo
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
			response, err := c.DeleteWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientGetBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "GetBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "GetBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "GetBlockchainPlatform")
	assert.NoError(t, err)

	type GetBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.GetBlockchainPlatformRequest
	}

	var requests []GetBlockchainPlatformRequestInfo
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
			response, err := c.GetBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientGetOsn(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "GetOsn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetOsn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "GetOsn", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "GetOsn")
	assert.NoError(t, err)

	type GetOsnRequestInfo struct {
		ContainerId string
		Request     blockchain.GetOsnRequest
	}

	var requests []GetOsnRequestInfo
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
			response, err := c.GetOsn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientGetPeer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "GetPeer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPeer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "GetPeer", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "GetPeer")
	assert.NoError(t, err)

	type GetPeerRequestInfo struct {
		ContainerId string
		Request     blockchain.GetPeerRequest
	}

	var requests []GetPeerRequestInfo
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
			response, err := c.GetPeer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "GetWorkRequest", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     blockchain.GetWorkRequestRequest
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

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientListBlockchainPlatforms(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ListBlockchainPlatforms")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBlockchainPlatforms is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ListBlockchainPlatforms", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ListBlockchainPlatforms")
	assert.NoError(t, err)

	type ListBlockchainPlatformsRequestInfo struct {
		ContainerId string
		Request     blockchain.ListBlockchainPlatformsRequest
	}

	var requests []ListBlockchainPlatformsRequestInfo
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
				r := req.(*blockchain.ListBlockchainPlatformsRequest)
				return c.ListBlockchainPlatforms(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]blockchain.ListBlockchainPlatformsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(blockchain.ListBlockchainPlatformsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientListOsns(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ListOsns")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListOsns is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ListOsns", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ListOsns")
	assert.NoError(t, err)

	type ListOsnsRequestInfo struct {
		ContainerId string
		Request     blockchain.ListOsnsRequest
	}

	var requests []ListOsnsRequestInfo
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
				r := req.(*blockchain.ListOsnsRequest)
				return c.ListOsns(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]blockchain.ListOsnsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(blockchain.ListOsnsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientListPeers(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ListPeers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPeers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ListPeers", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ListPeers")
	assert.NoError(t, err)

	type ListPeersRequestInfo struct {
		ContainerId string
		Request     blockchain.ListPeersRequest
	}

	var requests []ListPeersRequestInfo
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
				r := req.(*blockchain.ListPeersRequest)
				return c.ListPeers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]blockchain.ListPeersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(blockchain.ListPeersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientListWorkRequestErrors(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ListWorkRequestErrors")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestErrors is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ListWorkRequestErrors", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ListWorkRequestErrors")
	assert.NoError(t, err)

	type ListWorkRequestErrorsRequestInfo struct {
		ContainerId string
		Request     blockchain.ListWorkRequestErrorsRequest
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
				r := req.(*blockchain.ListWorkRequestErrorsRequest)
				return c.ListWorkRequestErrors(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]blockchain.ListWorkRequestErrorsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(blockchain.ListWorkRequestErrorsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientListWorkRequestLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ListWorkRequestLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequestLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ListWorkRequestLogs", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ListWorkRequestLogs")
	assert.NoError(t, err)

	type ListWorkRequestLogsRequestInfo struct {
		ContainerId string
		Request     blockchain.ListWorkRequestLogsRequest
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
				r := req.(*blockchain.ListWorkRequestLogsRequest)
				return c.ListWorkRequestLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]blockchain.ListWorkRequestLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(blockchain.ListWorkRequestLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ListWorkRequests", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     blockchain.ListWorkRequestsRequest
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
				r := req.(*blockchain.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]blockchain.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(blockchain.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientPreviewScaleBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "PreviewScaleBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PreviewScaleBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "PreviewScaleBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "PreviewScaleBlockchainPlatform")
	assert.NoError(t, err)

	type PreviewScaleBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.PreviewScaleBlockchainPlatformRequest
	}

	var requests []PreviewScaleBlockchainPlatformRequestInfo
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
			response, err := c.PreviewScaleBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientScaleBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "ScaleBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ScaleBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "ScaleBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "ScaleBlockchainPlatform")
	assert.NoError(t, err)

	type ScaleBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.ScaleBlockchainPlatformRequest
	}

	var requests []ScaleBlockchainPlatformRequestInfo
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
			response, err := c.ScaleBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientStartBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "StartBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StartBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "StartBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "StartBlockchainPlatform")
	assert.NoError(t, err)

	type StartBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.StartBlockchainPlatformRequest
	}

	var requests []StartBlockchainPlatformRequestInfo
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
			response, err := c.StartBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientStopBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "StopBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("StopBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "StopBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "StopBlockchainPlatform")
	assert.NoError(t, err)

	type StopBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.StopBlockchainPlatformRequest
	}

	var requests []StopBlockchainPlatformRequestInfo
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
			response, err := c.StopBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientUpdateBlockchainPlatform(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "UpdateBlockchainPlatform")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBlockchainPlatform is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "UpdateBlockchainPlatform", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "UpdateBlockchainPlatform")
	assert.NoError(t, err)

	type UpdateBlockchainPlatformRequestInfo struct {
		ContainerId string
		Request     blockchain.UpdateBlockchainPlatformRequest
	}

	var requests []UpdateBlockchainPlatformRequestInfo
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
			response, err := c.UpdateBlockchainPlatform(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientUpdateOsn(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "UpdateOsn")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateOsn is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "UpdateOsn", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "UpdateOsn")
	assert.NoError(t, err)

	type UpdateOsnRequestInfo struct {
		ContainerId string
		Request     blockchain.UpdateOsnRequest
	}

	var requests []UpdateOsnRequestInfo
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
			response, err := c.UpdateOsn(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="bcs_devops_ww_grp@oracle.com" jiraProject="OBP" opsJiraProject="OBP"
func TestBlockchainPlatformClientUpdatePeer(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("blockchain", "UpdatePeer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePeer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("blockchain", "BlockchainPlatform", "UpdatePeer", createBlockchainPlatformClientWithProvider)
	assert.NoError(t, err)
	c := cc.(blockchain.BlockchainPlatformClient)

	body, err := testClient.getRequests("blockchain", "UpdatePeer")
	assert.NoError(t, err)

	type UpdatePeerRequestInfo struct {
		ContainerId string
		Request     blockchain.UpdatePeerRequest
	}

	var requests []UpdatePeerRequestInfo
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
			response, err := c.UpdatePeer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
