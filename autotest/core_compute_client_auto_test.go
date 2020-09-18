package autotest

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"github.com/oracle/oci-go-sdk/v25/core"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createComputeClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := core.NewComputeClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientAddImageShapeCompatibilityEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AddImageShapeCompatibilityEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddImageShapeCompatibilityEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "AddImageShapeCompatibilityEntry", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "AddImageShapeCompatibilityEntry")
	assert.NoError(t, err)

	type AddImageShapeCompatibilityEntryRequestInfo struct {
		ContainerId string
		Request     core.AddImageShapeCompatibilityEntryRequest
	}

	var requests []AddImageShapeCompatibilityEntryRequestInfo
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
			response, err := c.AddImageShapeCompatibilityEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientAttachBootVolume(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AttachBootVolume")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AttachBootVolume is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "AttachBootVolume", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "AttachBootVolume")
	assert.NoError(t, err)

	type AttachBootVolumeRequestInfo struct {
		ContainerId string
		Request     core.AttachBootVolumeRequest
	}

	var requests []AttachBootVolumeRequestInfo
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
			response, err := c.AttachBootVolume(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientAttachVnic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AttachVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AttachVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "AttachVnic", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "AttachVnic")
	assert.NoError(t, err)

	type AttachVnicRequestInfo struct {
		ContainerId string
		Request     core.AttachVnicRequest
	}

	var requests []AttachVnicRequestInfo
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
			response, err := c.AttachVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientAttachVolume(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "AttachVolume")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AttachVolume is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "AttachVolume", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "AttachVolume")
	assert.NoError(t, err)

	type AttachVolumeRequestInfo struct {
		ContainerId string
		Request     core.AttachVolumeRequest
	}

	var requests []AttachVolumeRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]AttachVolumeRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["AttachVolumeDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "type",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"service_determined": &core.AttachServiceDeterminedVolumeDetails{},
				"emulated":           &core.AttachEmulatedVolumeDetails{},
				"iscsi":              &core.AttachIScsiVolumeDetails{},
				"paravirtualized":    &core.AttachParavirtualizedVolumeDetails{},
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
			response, err := c.AttachVolume(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientCaptureConsoleHistory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CaptureConsoleHistory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CaptureConsoleHistory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CaptureConsoleHistory", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CaptureConsoleHistory")
	assert.NoError(t, err)

	type CaptureConsoleHistoryRequestInfo struct {
		ContainerId string
		Request     core.CaptureConsoleHistoryRequest
	}

	var requests []CaptureConsoleHistoryRequestInfo
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
			response, err := c.CaptureConsoleHistory(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientChangeComputeCapacityReservationCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeComputeCapacityReservationCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeComputeCapacityReservationCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ChangeComputeCapacityReservationCompartment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ChangeComputeCapacityReservationCompartment")
	assert.NoError(t, err)

	type ChangeComputeCapacityReservationCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeComputeCapacityReservationCompartmentRequest
	}

	var requests []ChangeComputeCapacityReservationCompartmentRequestInfo
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
			response, err := c.ChangeComputeCapacityReservationCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientChangeComputeImageCapabilitySchemaCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeComputeImageCapabilitySchemaCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeComputeImageCapabilitySchemaCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ChangeComputeImageCapabilitySchemaCompartment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ChangeComputeImageCapabilitySchemaCompartment")
	assert.NoError(t, err)

	type ChangeComputeImageCapabilitySchemaCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeComputeImageCapabilitySchemaCompartmentRequest
	}

	var requests []ChangeComputeImageCapabilitySchemaCompartmentRequestInfo
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
			response, err := c.ChangeComputeImageCapabilitySchemaCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientChangeDedicatedVmHostCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeDedicatedVmHostCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeDedicatedVmHostCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ChangeDedicatedVmHostCompartment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ChangeDedicatedVmHostCompartment")
	assert.NoError(t, err)

	type ChangeDedicatedVmHostCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeDedicatedVmHostCompartmentRequest
	}

	var requests []ChangeDedicatedVmHostCompartmentRequestInfo
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
			response, err := c.ChangeDedicatedVmHostCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientChangeImageCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeImageCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeImageCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ChangeImageCompartment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ChangeImageCompartment")
	assert.NoError(t, err)

	type ChangeImageCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeImageCompartmentRequest
	}

	var requests []ChangeImageCompartmentRequestInfo
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
			response, err := c.ChangeImageCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientChangeInstanceCompartment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ChangeInstanceCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ChangeInstanceCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ChangeInstanceCompartment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ChangeInstanceCompartment")
	assert.NoError(t, err)

	type ChangeInstanceCompartmentRequestInfo struct {
		ContainerId string
		Request     core.ChangeInstanceCompartmentRequest
	}

	var requests []ChangeInstanceCompartmentRequestInfo
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
			response, err := c.ChangeInstanceCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientCreateAppCatalogSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateAppCatalogSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAppCatalogSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CreateAppCatalogSubscription", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CreateAppCatalogSubscription")
	assert.NoError(t, err)

	type CreateAppCatalogSubscriptionRequestInfo struct {
		ContainerId string
		Request     core.CreateAppCatalogSubscriptionRequest
	}

	var requests []CreateAppCatalogSubscriptionRequestInfo
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
			response, err := c.CreateAppCatalogSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientCreateComputeCapacityReservation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateComputeCapacityReservation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateComputeCapacityReservation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CreateComputeCapacityReservation", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CreateComputeCapacityReservation")
	assert.NoError(t, err)

	type CreateComputeCapacityReservationRequestInfo struct {
		ContainerId string
		Request     core.CreateComputeCapacityReservationRequest
	}

	var requests []CreateComputeCapacityReservationRequestInfo
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
			response, err := c.CreateComputeCapacityReservation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientCreateComputeImageCapabilitySchema(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateComputeImageCapabilitySchema")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateComputeImageCapabilitySchema is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CreateComputeImageCapabilitySchema", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CreateComputeImageCapabilitySchema")
	assert.NoError(t, err)

	type CreateComputeImageCapabilitySchemaRequestInfo struct {
		ContainerId string
		Request     core.CreateComputeImageCapabilitySchemaRequest
	}

	var requests []CreateComputeImageCapabilitySchemaRequestInfo
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
			response, err := c.CreateComputeImageCapabilitySchema(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientCreateDedicatedVmHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateDedicatedVmHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDedicatedVmHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CreateDedicatedVmHost", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CreateDedicatedVmHost")
	assert.NoError(t, err)

	type CreateDedicatedVmHostRequestInfo struct {
		ContainerId string
		Request     core.CreateDedicatedVmHostRequest
	}

	var requests []CreateDedicatedVmHostRequestInfo
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
			response, err := c.CreateDedicatedVmHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientCreateImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CreateImage", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CreateImage")
	assert.NoError(t, err)

	type CreateImageRequestInfo struct {
		ContainerId string
		Request     core.CreateImageRequest
	}

	var requests []CreateImageRequestInfo
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
			response, err := c.CreateImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeServices" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="COM"
func TestComputeClientCreateInstanceConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "CreateInstanceConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateInstanceConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "CreateInstanceConsoleConnection", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "CreateInstanceConsoleConnection")
	assert.NoError(t, err)

	type CreateInstanceConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     core.CreateInstanceConsoleConnectionRequest
	}

	var requests []CreateInstanceConsoleConnectionRequestInfo
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
			response, err := c.CreateInstanceConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientDeleteAppCatalogSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteAppCatalogSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAppCatalogSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteAppCatalogSubscription", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteAppCatalogSubscription")
	assert.NoError(t, err)

	type DeleteAppCatalogSubscriptionRequestInfo struct {
		ContainerId string
		Request     core.DeleteAppCatalogSubscriptionRequest
	}

	var requests []DeleteAppCatalogSubscriptionRequestInfo
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
			response, err := c.DeleteAppCatalogSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientDeleteComputeCapacityReservation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteComputeCapacityReservation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteComputeCapacityReservation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteComputeCapacityReservation", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteComputeCapacityReservation")
	assert.NoError(t, err)

	type DeleteComputeCapacityReservationRequestInfo struct {
		ContainerId string
		Request     core.DeleteComputeCapacityReservationRequest
	}

	var requests []DeleteComputeCapacityReservationRequestInfo
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
			response, err := c.DeleteComputeCapacityReservation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientDeleteComputeImageCapabilitySchema(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteComputeImageCapabilitySchema")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteComputeImageCapabilitySchema is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteComputeImageCapabilitySchema", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteComputeImageCapabilitySchema")
	assert.NoError(t, err)

	type DeleteComputeImageCapabilitySchemaRequestInfo struct {
		ContainerId string
		Request     core.DeleteComputeImageCapabilitySchemaRequest
	}

	var requests []DeleteComputeImageCapabilitySchemaRequestInfo
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
			response, err := c.DeleteComputeImageCapabilitySchema(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientDeleteConsoleHistory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteConsoleHistory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteConsoleHistory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteConsoleHistory", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteConsoleHistory")
	assert.NoError(t, err)

	type DeleteConsoleHistoryRequestInfo struct {
		ContainerId string
		Request     core.DeleteConsoleHistoryRequest
	}

	var requests []DeleteConsoleHistoryRequestInfo
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
			response, err := c.DeleteConsoleHistory(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDeleteDedicatedVmHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteDedicatedVmHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDedicatedVmHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteDedicatedVmHost", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteDedicatedVmHost")
	assert.NoError(t, err)

	type DeleteDedicatedVmHostRequestInfo struct {
		ContainerId string
		Request     core.DeleteDedicatedVmHostRequest
	}

	var requests []DeleteDedicatedVmHostRequestInfo
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
			response, err := c.DeleteDedicatedVmHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientDeleteImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteImage", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteImage")
	assert.NoError(t, err)

	type DeleteImageRequestInfo struct {
		ContainerId string
		Request     core.DeleteImageRequest
	}

	var requests []DeleteImageRequestInfo
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
			response, err := c.DeleteImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeServices" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="COM"
func TestComputeClientDeleteInstanceConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DeleteInstanceConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteInstanceConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DeleteInstanceConsoleConnection", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DeleteInstanceConsoleConnection")
	assert.NoError(t, err)

	type DeleteInstanceConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     core.DeleteInstanceConsoleConnectionRequest
	}

	var requests []DeleteInstanceConsoleConnectionRequestInfo
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
			response, err := c.DeleteInstanceConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientDetachBootVolume(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DetachBootVolume")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetachBootVolume is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DetachBootVolume", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DetachBootVolume")
	assert.NoError(t, err)

	type DetachBootVolumeRequestInfo struct {
		ContainerId string
		Request     core.DetachBootVolumeRequest
	}

	var requests []DetachBootVolumeRequestInfo
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
			response, err := c.DetachBootVolume(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientDetachVnic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DetachVnic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetachVnic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DetachVnic", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DetachVnic")
	assert.NoError(t, err)

	type DetachVnicRequestInfo struct {
		ContainerId string
		Request     core.DetachVnicRequest
	}

	var requests []DetachVnicRequestInfo
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
			response, err := c.DetachVnic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientDetachVolume(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "DetachVolume")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetachVolume is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "DetachVolume", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "DetachVolume")
	assert.NoError(t, err)

	type DetachVolumeRequestInfo struct {
		ContainerId string
		Request     core.DetachVolumeRequest
	}

	var requests []DetachVolumeRequestInfo
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
			response, err := c.DetachVolume(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientExportImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ExportImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ExportImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ExportImage", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ExportImage")
	assert.NoError(t, err)

	type ExportImageRequestInfo struct {
		ContainerId string
		Request     core.ExportImageRequest
	}

	var requests []ExportImageRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]ExportImageRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["ExportImageDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "destinationType",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"objectStorageUri":   &core.ExportImageViaObjectStorageUriDetails{},
				"objectStorageTuple": &core.ExportImageViaObjectStorageTupleDetails{},
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
			response, err := c.ExportImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetAppCatalogListing(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetAppCatalogListing")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAppCatalogListing is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetAppCatalogListing", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetAppCatalogListing")
	assert.NoError(t, err)

	type GetAppCatalogListingRequestInfo struct {
		ContainerId string
		Request     core.GetAppCatalogListingRequest
	}

	var requests []GetAppCatalogListingRequestInfo
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
			response, err := c.GetAppCatalogListing(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetAppCatalogListingAgreements(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetAppCatalogListingAgreements")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAppCatalogListingAgreements is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetAppCatalogListingAgreements", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetAppCatalogListingAgreements")
	assert.NoError(t, err)

	type GetAppCatalogListingAgreementsRequestInfo struct {
		ContainerId string
		Request     core.GetAppCatalogListingAgreementsRequest
	}

	var requests []GetAppCatalogListingAgreementsRequestInfo
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
			response, err := c.GetAppCatalogListingAgreements(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetAppCatalogListingResourceVersion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetAppCatalogListingResourceVersion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAppCatalogListingResourceVersion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetAppCatalogListingResourceVersion", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetAppCatalogListingResourceVersion")
	assert.NoError(t, err)

	type GetAppCatalogListingResourceVersionRequestInfo struct {
		ContainerId string
		Request     core.GetAppCatalogListingResourceVersionRequest
	}

	var requests []GetAppCatalogListingResourceVersionRequestInfo
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
			response, err := c.GetAppCatalogListingResourceVersion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetBootVolumeAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetBootVolumeAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBootVolumeAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetBootVolumeAttachment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetBootVolumeAttachment")
	assert.NoError(t, err)

	type GetBootVolumeAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetBootVolumeAttachmentRequest
	}

	var requests []GetBootVolumeAttachmentRequestInfo
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
			response, err := c.GetBootVolumeAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetComputeCapacityReservation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetComputeCapacityReservation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetComputeCapacityReservation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetComputeCapacityReservation", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetComputeCapacityReservation")
	assert.NoError(t, err)

	type GetComputeCapacityReservationRequestInfo struct {
		ContainerId string
		Request     core.GetComputeCapacityReservationRequest
	}

	var requests []GetComputeCapacityReservationRequestInfo
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
			response, err := c.GetComputeCapacityReservation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetComputeGlobalImageCapabilitySchema(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetComputeGlobalImageCapabilitySchema")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetComputeGlobalImageCapabilitySchema is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetComputeGlobalImageCapabilitySchema", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetComputeGlobalImageCapabilitySchema")
	assert.NoError(t, err)

	type GetComputeGlobalImageCapabilitySchemaRequestInfo struct {
		ContainerId string
		Request     core.GetComputeGlobalImageCapabilitySchemaRequest
	}

	var requests []GetComputeGlobalImageCapabilitySchemaRequestInfo
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
			response, err := c.GetComputeGlobalImageCapabilitySchema(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetComputeGlobalImageCapabilitySchemaVersion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetComputeGlobalImageCapabilitySchemaVersion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetComputeGlobalImageCapabilitySchemaVersion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetComputeGlobalImageCapabilitySchemaVersion", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetComputeGlobalImageCapabilitySchemaVersion")
	assert.NoError(t, err)

	type GetComputeGlobalImageCapabilitySchemaVersionRequestInfo struct {
		ContainerId string
		Request     core.GetComputeGlobalImageCapabilitySchemaVersionRequest
	}

	var requests []GetComputeGlobalImageCapabilitySchemaVersionRequestInfo
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
			response, err := c.GetComputeGlobalImageCapabilitySchemaVersion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetComputeImageCapabilitySchema(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetComputeImageCapabilitySchema")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetComputeImageCapabilitySchema is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetComputeImageCapabilitySchema", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetComputeImageCapabilitySchema")
	assert.NoError(t, err)

	type GetComputeImageCapabilitySchemaRequestInfo struct {
		ContainerId string
		Request     core.GetComputeImageCapabilitySchemaRequest
	}

	var requests []GetComputeImageCapabilitySchemaRequestInfo
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
			response, err := c.GetComputeImageCapabilitySchema(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetConsoleHistory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetConsoleHistory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConsoleHistory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetConsoleHistory", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetConsoleHistory")
	assert.NoError(t, err)

	type GetConsoleHistoryRequestInfo struct {
		ContainerId string
		Request     core.GetConsoleHistoryRequest
	}

	var requests []GetConsoleHistoryRequestInfo
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
			response, err := c.GetConsoleHistory(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetConsoleHistoryContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetConsoleHistoryContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConsoleHistoryContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetConsoleHistoryContent", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetConsoleHistoryContent")
	assert.NoError(t, err)

	type GetConsoleHistoryContentRequestInfo struct {
		ContainerId string
		Request     core.GetConsoleHistoryContentRequest
	}

	var requests []GetConsoleHistoryContentRequestInfo
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
			response, err := c.GetConsoleHistoryContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetDedicatedVmHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetDedicatedVmHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDedicatedVmHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetDedicatedVmHost", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetDedicatedVmHost")
	assert.NoError(t, err)

	type GetDedicatedVmHostRequestInfo struct {
		ContainerId string
		Request     core.GetDedicatedVmHostRequest
	}

	var requests []GetDedicatedVmHostRequestInfo
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
			response, err := c.GetDedicatedVmHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetImage", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetImage")
	assert.NoError(t, err)

	type GetImageRequestInfo struct {
		ContainerId string
		Request     core.GetImageRequest
	}

	var requests []GetImageRequestInfo
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
			response, err := c.GetImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientGetImageShapeCompatibilityEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetImageShapeCompatibilityEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetImageShapeCompatibilityEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetImageShapeCompatibilityEntry", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetImageShapeCompatibilityEntry")
	assert.NoError(t, err)

	type GetImageShapeCompatibilityEntryRequestInfo struct {
		ContainerId string
		Request     core.GetImageShapeCompatibilityEntryRequest
	}

	var requests []GetImageShapeCompatibilityEntryRequestInfo
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
			response, err := c.GetImageShapeCompatibilityEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetInstance", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetInstance")
	assert.NoError(t, err)

	type GetInstanceRequestInfo struct {
		ContainerId string
		Request     core.GetInstanceRequest
	}

	var requests []GetInstanceRequestInfo
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
			response, err := c.GetInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeServices" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="COM"
func TestComputeClientGetInstanceConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetInstanceConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetInstanceConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetInstanceConsoleConnection", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetInstanceConsoleConnection")
	assert.NoError(t, err)

	type GetInstanceConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     core.GetInstanceConsoleConnectionRequest
	}

	var requests []GetInstanceConsoleConnectionRequestInfo
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
			response, err := c.GetInstanceConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetVnicAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetVnicAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVnicAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetVnicAttachment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetVnicAttachment")
	assert.NoError(t, err)

	type GetVnicAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetVnicAttachmentRequest
	}

	var requests []GetVnicAttachmentRequestInfo
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
			response, err := c.GetVnicAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetVolumeAttachment(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetVolumeAttachment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVolumeAttachment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetVolumeAttachment", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetVolumeAttachment")
	assert.NoError(t, err)

	type GetVolumeAttachmentRequestInfo struct {
		ContainerId string
		Request     core.GetVolumeAttachmentRequest
	}

	var requests []GetVolumeAttachmentRequestInfo
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
			response, err := c.GetVolumeAttachment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientGetWindowsInstanceInitialCredentials(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "GetWindowsInstanceInitialCredentials")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWindowsInstanceInitialCredentials is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "GetWindowsInstanceInitialCredentials", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "GetWindowsInstanceInitialCredentials")
	assert.NoError(t, err)

	type GetWindowsInstanceInitialCredentialsRequestInfo struct {
		ContainerId string
		Request     core.GetWindowsInstanceInitialCredentialsRequest
	}

	var requests []GetWindowsInstanceInitialCredentialsRequestInfo
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
			response, err := c.GetWindowsInstanceInitialCredentials(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientInstanceAction(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "InstanceAction")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("InstanceAction is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "InstanceAction", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "InstanceAction")
	assert.NoError(t, err)

	type InstanceActionRequestInfo struct {
		ContainerId string
		Request     core.InstanceActionRequest
	}

	var requests []InstanceActionRequestInfo
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
			response, err := c.InstanceAction(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientLaunchInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "LaunchInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("LaunchInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "LaunchInstance", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "LaunchInstance")
	assert.NoError(t, err)

	type LaunchInstanceRequestInfo struct {
		ContainerId string
		Request     core.LaunchInstanceRequest
	}

	var requests []LaunchInstanceRequestInfo
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
			response, err := c.LaunchInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListAppCatalogListingResourceVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListAppCatalogListingResourceVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAppCatalogListingResourceVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListAppCatalogListingResourceVersions", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListAppCatalogListingResourceVersions")
	assert.NoError(t, err)

	type ListAppCatalogListingResourceVersionsRequestInfo struct {
		ContainerId string
		Request     core.ListAppCatalogListingResourceVersionsRequest
	}

	var requests []ListAppCatalogListingResourceVersionsRequestInfo
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
				r := req.(*core.ListAppCatalogListingResourceVersionsRequest)
				return c.ListAppCatalogListingResourceVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListAppCatalogListingResourceVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListAppCatalogListingResourceVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListAppCatalogListings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListAppCatalogListings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAppCatalogListings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListAppCatalogListings", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListAppCatalogListings")
	assert.NoError(t, err)

	type ListAppCatalogListingsRequestInfo struct {
		ContainerId string
		Request     core.ListAppCatalogListingsRequest
	}

	var requests []ListAppCatalogListingsRequestInfo
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
				r := req.(*core.ListAppCatalogListingsRequest)
				return c.ListAppCatalogListings(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListAppCatalogListingsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListAppCatalogListingsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListAppCatalogSubscriptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListAppCatalogSubscriptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAppCatalogSubscriptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListAppCatalogSubscriptions", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListAppCatalogSubscriptions")
	assert.NoError(t, err)

	type ListAppCatalogSubscriptionsRequestInfo struct {
		ContainerId string
		Request     core.ListAppCatalogSubscriptionsRequest
	}

	var requests []ListAppCatalogSubscriptionsRequestInfo
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
				r := req.(*core.ListAppCatalogSubscriptionsRequest)
				return c.ListAppCatalogSubscriptions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListAppCatalogSubscriptionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListAppCatalogSubscriptionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListBootVolumeAttachments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListBootVolumeAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBootVolumeAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListBootVolumeAttachments", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListBootVolumeAttachments")
	assert.NoError(t, err)

	type ListBootVolumeAttachmentsRequestInfo struct {
		ContainerId string
		Request     core.ListBootVolumeAttachmentsRequest
	}

	var requests []ListBootVolumeAttachmentsRequestInfo
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
				r := req.(*core.ListBootVolumeAttachmentsRequest)
				return c.ListBootVolumeAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListBootVolumeAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListBootVolumeAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListComputeCapacityReservationInstanceShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListComputeCapacityReservationInstanceShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComputeCapacityReservationInstanceShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListComputeCapacityReservationInstanceShapes", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListComputeCapacityReservationInstanceShapes")
	assert.NoError(t, err)

	type ListComputeCapacityReservationInstanceShapesRequestInfo struct {
		ContainerId string
		Request     core.ListComputeCapacityReservationInstanceShapesRequest
	}

	var requests []ListComputeCapacityReservationInstanceShapesRequestInfo
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
				r := req.(*core.ListComputeCapacityReservationInstanceShapesRequest)
				return c.ListComputeCapacityReservationInstanceShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListComputeCapacityReservationInstanceShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListComputeCapacityReservationInstanceShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListComputeCapacityReservationInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListComputeCapacityReservationInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComputeCapacityReservationInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListComputeCapacityReservationInstances", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListComputeCapacityReservationInstances")
	assert.NoError(t, err)

	type ListComputeCapacityReservationInstancesRequestInfo struct {
		ContainerId string
		Request     core.ListComputeCapacityReservationInstancesRequest
	}

	var requests []ListComputeCapacityReservationInstancesRequestInfo
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
				r := req.(*core.ListComputeCapacityReservationInstancesRequest)
				return c.ListComputeCapacityReservationInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListComputeCapacityReservationInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListComputeCapacityReservationInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListComputeCapacityReservations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListComputeCapacityReservations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComputeCapacityReservations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListComputeCapacityReservations", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListComputeCapacityReservations")
	assert.NoError(t, err)

	type ListComputeCapacityReservationsRequestInfo struct {
		ContainerId string
		Request     core.ListComputeCapacityReservationsRequest
	}

	var requests []ListComputeCapacityReservationsRequestInfo
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
				r := req.(*core.ListComputeCapacityReservationsRequest)
				return c.ListComputeCapacityReservations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListComputeCapacityReservationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListComputeCapacityReservationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListComputeGlobalImageCapabilitySchemaVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListComputeGlobalImageCapabilitySchemaVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComputeGlobalImageCapabilitySchemaVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListComputeGlobalImageCapabilitySchemaVersions", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListComputeGlobalImageCapabilitySchemaVersions")
	assert.NoError(t, err)

	type ListComputeGlobalImageCapabilitySchemaVersionsRequestInfo struct {
		ContainerId string
		Request     core.ListComputeGlobalImageCapabilitySchemaVersionsRequest
	}

	var requests []ListComputeGlobalImageCapabilitySchemaVersionsRequestInfo
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
				r := req.(*core.ListComputeGlobalImageCapabilitySchemaVersionsRequest)
				return c.ListComputeGlobalImageCapabilitySchemaVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListComputeGlobalImageCapabilitySchemaVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListComputeGlobalImageCapabilitySchemaVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListComputeGlobalImageCapabilitySchemas(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListComputeGlobalImageCapabilitySchemas")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComputeGlobalImageCapabilitySchemas is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListComputeGlobalImageCapabilitySchemas", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListComputeGlobalImageCapabilitySchemas")
	assert.NoError(t, err)

	type ListComputeGlobalImageCapabilitySchemasRequestInfo struct {
		ContainerId string
		Request     core.ListComputeGlobalImageCapabilitySchemasRequest
	}

	var requests []ListComputeGlobalImageCapabilitySchemasRequestInfo
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
				r := req.(*core.ListComputeGlobalImageCapabilitySchemasRequest)
				return c.ListComputeGlobalImageCapabilitySchemas(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListComputeGlobalImageCapabilitySchemasResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListComputeGlobalImageCapabilitySchemasResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListComputeImageCapabilitySchemas(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListComputeImageCapabilitySchemas")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListComputeImageCapabilitySchemas is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListComputeImageCapabilitySchemas", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListComputeImageCapabilitySchemas")
	assert.NoError(t, err)

	type ListComputeImageCapabilitySchemasRequestInfo struct {
		ContainerId string
		Request     core.ListComputeImageCapabilitySchemasRequest
	}

	var requests []ListComputeImageCapabilitySchemasRequestInfo
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
				r := req.(*core.ListComputeImageCapabilitySchemasRequest)
				return c.ListComputeImageCapabilitySchemas(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListComputeImageCapabilitySchemasResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListComputeImageCapabilitySchemasResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListConsoleHistories(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListConsoleHistories")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListConsoleHistories is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListConsoleHistories", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListConsoleHistories")
	assert.NoError(t, err)

	type ListConsoleHistoriesRequestInfo struct {
		ContainerId string
		Request     core.ListConsoleHistoriesRequest
	}

	var requests []ListConsoleHistoriesRequestInfo
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
				r := req.(*core.ListConsoleHistoriesRequest)
				return c.ListConsoleHistories(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListConsoleHistoriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListConsoleHistoriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListDedicatedVmHostInstanceShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListDedicatedVmHostInstanceShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDedicatedVmHostInstanceShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListDedicatedVmHostInstanceShapes", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListDedicatedVmHostInstanceShapes")
	assert.NoError(t, err)

	type ListDedicatedVmHostInstanceShapesRequestInfo struct {
		ContainerId string
		Request     core.ListDedicatedVmHostInstanceShapesRequest
	}

	var requests []ListDedicatedVmHostInstanceShapesRequestInfo
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
				r := req.(*core.ListDedicatedVmHostInstanceShapesRequest)
				return c.ListDedicatedVmHostInstanceShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDedicatedVmHostInstanceShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDedicatedVmHostInstanceShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListDedicatedVmHostInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListDedicatedVmHostInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDedicatedVmHostInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListDedicatedVmHostInstances", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListDedicatedVmHostInstances")
	assert.NoError(t, err)

	type ListDedicatedVmHostInstancesRequestInfo struct {
		ContainerId string
		Request     core.ListDedicatedVmHostInstancesRequest
	}

	var requests []ListDedicatedVmHostInstancesRequestInfo
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
				r := req.(*core.ListDedicatedVmHostInstancesRequest)
				return c.ListDedicatedVmHostInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDedicatedVmHostInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDedicatedVmHostInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListDedicatedVmHostShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListDedicatedVmHostShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDedicatedVmHostShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListDedicatedVmHostShapes", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListDedicatedVmHostShapes")
	assert.NoError(t, err)

	type ListDedicatedVmHostShapesRequestInfo struct {
		ContainerId string
		Request     core.ListDedicatedVmHostShapesRequest
	}

	var requests []ListDedicatedVmHostShapesRequestInfo
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
				r := req.(*core.ListDedicatedVmHostShapesRequest)
				return c.ListDedicatedVmHostShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDedicatedVmHostShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDedicatedVmHostShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListDedicatedVmHosts(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListDedicatedVmHosts")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDedicatedVmHosts is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListDedicatedVmHosts", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListDedicatedVmHosts")
	assert.NoError(t, err)

	type ListDedicatedVmHostsRequestInfo struct {
		ContainerId string
		Request     core.ListDedicatedVmHostsRequest
	}

	var requests []ListDedicatedVmHostsRequestInfo
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
				r := req.(*core.ListDedicatedVmHostsRequest)
				return c.ListDedicatedVmHosts(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListDedicatedVmHostsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListDedicatedVmHostsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListImageShapeCompatibilityEntries(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListImageShapeCompatibilityEntries")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListImageShapeCompatibilityEntries is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListImageShapeCompatibilityEntries", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListImageShapeCompatibilityEntries")
	assert.NoError(t, err)

	type ListImageShapeCompatibilityEntriesRequestInfo struct {
		ContainerId string
		Request     core.ListImageShapeCompatibilityEntriesRequest
	}

	var requests []ListImageShapeCompatibilityEntriesRequestInfo
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
				r := req.(*core.ListImageShapeCompatibilityEntriesRequest)
				return c.ListImageShapeCompatibilityEntries(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListImageShapeCompatibilityEntriesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListImageShapeCompatibilityEntriesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientListImages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListImages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListImages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListImages", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListImages")
	assert.NoError(t, err)

	type ListImagesRequestInfo struct {
		ContainerId string
		Request     core.ListImagesRequest
	}

	var requests []ListImagesRequestInfo
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
				r := req.(*core.ListImagesRequest)
				return c.ListImages(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListImagesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListImagesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeServices" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="COM"
func TestComputeClientListInstanceConsoleConnections(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInstanceConsoleConnections")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstanceConsoleConnections is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListInstanceConsoleConnections", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListInstanceConsoleConnections")
	assert.NoError(t, err)

	type ListInstanceConsoleConnectionsRequestInfo struct {
		ContainerId string
		Request     core.ListInstanceConsoleConnectionsRequest
	}

	var requests []ListInstanceConsoleConnectionsRequestInfo
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
				r := req.(*core.ListInstanceConsoleConnectionsRequest)
				return c.ListInstanceConsoleConnections(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInstanceConsoleConnectionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInstanceConsoleConnectionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListInstanceDevices(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInstanceDevices")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstanceDevices is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListInstanceDevices", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListInstanceDevices")
	assert.NoError(t, err)

	type ListInstanceDevicesRequestInfo struct {
		ContainerId string
		Request     core.ListInstanceDevicesRequest
	}

	var requests []ListInstanceDevicesRequestInfo
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
				r := req.(*core.ListInstanceDevicesRequest)
				return c.ListInstanceDevices(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInstanceDevicesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInstanceDevicesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListInstances(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListInstances")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListInstances is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListInstances", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListInstances")
	assert.NoError(t, err)

	type ListInstancesRequestInfo struct {
		ContainerId string
		Request     core.ListInstancesRequest
	}

	var requests []ListInstancesRequestInfo
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
				r := req.(*core.ListInstancesRequest)
				return c.ListInstances(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListInstancesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListInstancesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListShapes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListShapes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListShapes", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListShapes")
	assert.NoError(t, err)

	type ListShapesRequestInfo struct {
		ContainerId string
		Request     core.ListShapesRequest
	}

	var requests []ListShapesRequestInfo
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
				r := req.(*core.ListShapesRequest)
				return c.ListShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListVnicAttachments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListVnicAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVnicAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListVnicAttachments", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListVnicAttachments")
	assert.NoError(t, err)

	type ListVnicAttachmentsRequestInfo struct {
		ContainerId string
		Request     core.ListVnicAttachmentsRequest
	}

	var requests []ListVnicAttachmentsRequestInfo
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
				r := req.(*core.ListVnicAttachmentsRequest)
				return c.ListVnicAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVnicAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVnicAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientListVolumeAttachments(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "ListVolumeAttachments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVolumeAttachments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "ListVolumeAttachments", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "ListVolumeAttachments")
	assert.NoError(t, err)

	type ListVolumeAttachmentsRequestInfo struct {
		ContainerId string
		Request     core.ListVolumeAttachmentsRequest
	}

	var requests []ListVolumeAttachmentsRequestInfo
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
				r := req.(*core.ListVolumeAttachmentsRequest)
				return c.ListVolumeAttachments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]core.ListVolumeAttachmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(core.ListVolumeAttachmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientRemoveImageShapeCompatibilityEntry(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "RemoveImageShapeCompatibilityEntry")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveImageShapeCompatibilityEntry is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "RemoveImageShapeCompatibilityEntry", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "RemoveImageShapeCompatibilityEntry")
	assert.NoError(t, err)

	type RemoveImageShapeCompatibilityEntryRequestInfo struct {
		ContainerId string
		Request     core.RemoveImageShapeCompatibilityEntryRequest
	}

	var requests []RemoveImageShapeCompatibilityEntryRequestInfo
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
			response, err := c.RemoveImageShapeCompatibilityEntry(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientTerminateInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "TerminateInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("TerminateInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "TerminateInstance", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "TerminateInstance")
	assert.NoError(t, err)

	type TerminateInstanceRequestInfo struct {
		ContainerId string
		Request     core.TerminateInstanceRequest
	}

	var requests []TerminateInstanceRequestInfo
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
			response, err := c.TerminateInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientUpdateComputeCapacityReservation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateComputeCapacityReservation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateComputeCapacityReservation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateComputeCapacityReservation", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateComputeCapacityReservation")
	assert.NoError(t, err)

	type UpdateComputeCapacityReservationRequestInfo struct {
		ContainerId string
		Request     core.UpdateComputeCapacityReservationRequest
	}

	var requests []UpdateComputeCapacityReservationRequestInfo
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
			response, err := c.UpdateComputeCapacityReservation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientUpdateComputeImageCapabilitySchema(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateComputeImageCapabilitySchema")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateComputeImageCapabilitySchema is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateComputeImageCapabilitySchema", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateComputeImageCapabilitySchema")
	assert.NoError(t, err)

	type UpdateComputeImageCapabilitySchemaRequestInfo struct {
		ContainerId string
		Request     core.UpdateComputeImageCapabilitySchemaRequest
	}

	var requests []UpdateComputeImageCapabilitySchemaRequestInfo
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
			response, err := c.UpdateComputeImageCapabilitySchema(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientUpdateConsoleHistory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateConsoleHistory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateConsoleHistory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateConsoleHistory", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateConsoleHistory")
	assert.NoError(t, err)

	type UpdateConsoleHistoryRequestInfo struct {
		ContainerId string
		Request     core.UpdateConsoleHistoryRequest
	}

	var requests []UpdateConsoleHistoryRequestInfo
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
			response, err := c.UpdateConsoleHistory(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientUpdateDedicatedVmHost(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateDedicatedVmHost")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDedicatedVmHost is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateDedicatedVmHost", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateDedicatedVmHost")
	assert.NoError(t, err)

	type UpdateDedicatedVmHostRequestInfo struct {
		ContainerId string
		Request     core.UpdateDedicatedVmHostRequest
	}

	var requests []UpdateDedicatedVmHostRequestInfo
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
			response, err := c.UpdateDedicatedVmHost(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeImaging" email="imaging_dev_us_grp@oracle.com" jiraProject="COM" opsJiraProject="COM"
func TestComputeClientUpdateImage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateImage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateImage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateImage", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateImage")
	assert.NoError(t, err)

	type UpdateImageRequestInfo struct {
		ContainerId string
		Request     core.UpdateImageRequest
	}

	var requests []UpdateImageRequestInfo
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
			response, err := c.UpdateImage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeSharedOwnershipVmAndBm" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="NONE"
func TestComputeClientUpdateInstance(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInstance")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInstance is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateInstance", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateInstance")
	assert.NoError(t, err)

	type UpdateInstanceRequestInfo struct {
		ContainerId string
		Request     core.UpdateInstanceRequest
	}

	var requests []UpdateInstanceRequestInfo
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
			response, err := c.UpdateInstance(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="computeServices" email="compute_dev_us_grp@oracle.com" jiraProject="BMI" opsJiraProject="COM"
func TestComputeClientUpdateInstanceConsoleConnection(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("core", "UpdateInstanceConsoleConnection")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateInstanceConsoleConnection is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "Compute", "UpdateInstanceConsoleConnection", createComputeClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeClient)

	body, err := testClient.getRequests("core", "UpdateInstanceConsoleConnection")
	assert.NoError(t, err)

	type UpdateInstanceConsoleConnectionRequestInfo struct {
		ContainerId string
		Request     core.UpdateInstanceConsoleConnectionRequest
	}

	var requests []UpdateInstanceConsoleConnectionRequestInfo
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
			response, err := c.UpdateInstanceConsoleConnection(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
