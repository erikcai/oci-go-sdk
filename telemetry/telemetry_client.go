// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Telemetry API
//
// Use the Telemetry API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// For information about metrics, see Telemetry Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Telemetry/Concepts/telemetryoverview.htm).
// For information about alarms, see Alarms Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Alarms/Concepts/alarmsoverview.htm).
//

package telemetry

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//TelemetryClient a client for Telemetry
type TelemetryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTelemetryClientWithConfigurationProvider Creates a new default Telemetry client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTelemetryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TelemetryClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = TelemetryClient{BaseClient: baseClient}
	client.BasePath = "20180401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TelemetryClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "telemetry", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TelemetryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *TelemetryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAlarm Creates a new alarm in the specified compartment.
func (client TelemetryClient) CreateAlarm(ctx context.Context, request CreateAlarmRequest) (response CreateAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateAlarmResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAlarmResponse")
	}
	return
}

// createAlarm implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) createAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms")
	if err != nil {
		return nil, err
	}

	var response CreateAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAlarm Deletes the specified alarm.
func (client TelemetryClient) DeleteAlarm(ctx context.Context, request DeleteAlarmRequest) (response DeleteAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteAlarmResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAlarmResponse")
	}
	return
}

// deleteAlarm implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) deleteAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/alarms/{alarmId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlarm Gets the specified alarm.
func (client TelemetryClient) GetAlarm(ctx context.Context, request GetAlarmRequest) (response GetAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAlarmResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlarmResponse")
	}
	return
}

// getAlarm implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) getAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/{alarmId}")
	if err != nil {
		return nil, err
	}

	var response GetAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlarmHistory Get the history of the specified alarm.
func (client TelemetryClient) GetAlarmHistory(ctx context.Context, request GetAlarmHistoryRequest) (response GetAlarmHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlarmHistory, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetAlarmHistoryResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlarmHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlarmHistoryResponse")
	}
	return
}

// getAlarmHistory implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) getAlarmHistory(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/{alarmId}/history")
	if err != nil {
		return nil, err
	}

	var response GetAlarmHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlarms Lists the alarms for the specified compartment.
func (client TelemetryClient) ListAlarms(ctx context.Context, request ListAlarmsRequest) (response ListAlarmsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlarms, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListAlarmsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlarmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlarmsResponse")
	}
	return
}

// listAlarms implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) listAlarms(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms")
	if err != nil {
		return nil, err
	}

	var response ListAlarmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMetrics Returns metric definitions that match the criteria specified in the request. Compartment OCID required.
// For more information on monitoring metrics, see Telemetry Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Telemetry/Concepts/telemetryoverview.htm).
func (client TelemetryClient) ListMetrics(ctx context.Context, request ListMetricsRequest) (response ListMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListMetricsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetricsResponse")
	}
	return
}

// listMetrics implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) listMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics/actions/listMetrics")
	if err != nil {
		return nil, err
	}

	var response ListMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAlarmSuppression Removes any existing suppression for the specified alarm.
func (client TelemetryClient) RemoveAlarmSuppression(ctx context.Context, request RemoveAlarmSuppressionRequest) (response RemoveAlarmSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeAlarmSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			response = RemoveAlarmSuppressionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveAlarmSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveAlarmSuppressionResponse")
	}
	return
}

// removeAlarmSuppression implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) removeAlarmSuppression(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/removeSuppression")
	if err != nil {
		return nil, err
	}

	var response RemoveAlarmSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeMetricsData Returns aggregated data that match the criteria specified in the request. Compartment OCID required.
// For more information on monitoring metrics, see Telemetry Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Telemetry/Concepts/telemetryoverview.htm).
func (client TelemetryClient) SummarizeMetricsData(ctx context.Context, request SummarizeMetricsDataRequest) (response SummarizeMetricsDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeMetricsData, policy)
	if err != nil {
		if ociResponse != nil {
			response = SummarizeMetricsDataResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeMetricsDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeMetricsDataResponse")
	}
	return
}

// summarizeMetricsData implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) summarizeMetricsData(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics/actions/summarizeMetricsData")
	if err != nil {
		return nil, err
	}

	var response SummarizeMetricsDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAlarm Updates the specified alarm.
func (client TelemetryClient) UpdateAlarm(ctx context.Context, request UpdateAlarmRequest) (response UpdateAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateAlarmResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAlarmResponse")
	}
	return
}

// updateAlarm implements the OCIOperation interface (enables retrying operations)
func (client TelemetryClient) updateAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/alarms/{alarmId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
