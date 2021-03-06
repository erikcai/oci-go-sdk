package autotest

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/keymanagement"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createKmsCryptoClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
	client, err := keymanagement.NewKmsCryptoClientWithConfigurationProvider(p, testConfig.Endpoint)
	return client, err
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsCryptoClientDecrypt(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "Decrypt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Decrypt is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "Decrypt", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "Decrypt")
	assert.NoError(t, err)

	type DecryptRequestInfo struct {
		ContainerId string
		Request     keymanagement.DecryptRequest
	}

	var requests []DecryptRequestInfo
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
			response, err := c.Decrypt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsCryptoClientEncrypt(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "Encrypt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Encrypt is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "Encrypt", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "Encrypt")
	assert.NoError(t, err)

	type EncryptRequestInfo struct {
		ContainerId string
		Request     keymanagement.EncryptRequest
	}

	var requests []EncryptRequestInfo
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
			response, err := c.Encrypt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsCryptoClientExportKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "ExportKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExportKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "ExportKey", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "ExportKey")
	assert.NoError(t, err)

	type ExportKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.ExportKeyRequest
	}

	var requests []ExportKeyRequestInfo
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
			response, err := c.ExportKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsCryptoClientGenerateDataEncryptionKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "GenerateDataEncryptionKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateDataEncryptionKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "GenerateDataEncryptionKey", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "GenerateDataEncryptionKey")
	assert.NoError(t, err)

	type GenerateDataEncryptionKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.GenerateDataEncryptionKeyRequest
	}

	var requests []GenerateDataEncryptionKeyRequestInfo
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
			response, err := c.GenerateDataEncryptionKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsCryptoClientSign(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "Sign")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Sign is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "Sign", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "Sign")
	assert.NoError(t, err)

	type SignRequestInfo struct {
		ContainerId string
		Request     keymanagement.SignRequest
	}

	var requests []SignRequestInfo
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
			response, err := c.Sign(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsCryptoClientVerify(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "Verify")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Verify is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "Verify", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "Verify")
	assert.NoError(t, err)

	type VerifyRequestInfo struct {
		ContainerId string
		Request     keymanagement.VerifyRequest
	}

	var requests []VerifyRequestInfo
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
			response, err := c.Verify(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
