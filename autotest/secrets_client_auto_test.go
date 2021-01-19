package autotest

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"github.com/oracle/oci-go-sdk/v33/secrets"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSecretsClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := secrets.NewSecretsClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestSecretsClientGetSecretBundle(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("secrets", "GetSecretBundle")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSecretBundle is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("secrets", "Secrets", "GetSecretBundle", createSecretsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(secrets.SecretsClient)

	body, err := testClient.getRequests("secrets", "GetSecretBundle")
	assert.NoError(t, err)

	type GetSecretBundleRequestInfo struct {
		ContainerId string
		Request     secrets.GetSecretBundleRequest
	}

	var requests []GetSecretBundleRequestInfo
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
			response, err := c.GetSecretBundle(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_vault_us_grp@oracle.com" jiraProject="SECSVC" opsJiraProject="SI"
func TestSecretsClientListSecretBundleVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("secrets", "ListSecretBundleVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSecretBundleVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("secrets", "Secrets", "ListSecretBundleVersions", createSecretsClientWithProvider)
	assert.NoError(t, err)
	c := cc.(secrets.SecretsClient)

	body, err := testClient.getRequests("secrets", "ListSecretBundleVersions")
	assert.NoError(t, err)

	type ListSecretBundleVersionsRequestInfo struct {
		ContainerId string
		Request     secrets.ListSecretBundleVersionsRequest
	}

	var requests []ListSecretBundleVersionsRequestInfo
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
				r := req.(*secrets.ListSecretBundleVersionsRequest)
				return c.ListSecretBundleVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]secrets.ListSecretBundleVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(secrets.ListSecretBundleVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
