// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Telemetry Service API
// 
 // Use the Telemetry Service API to access metric data about the health, capacity, and performance of your cloud resources.
 // For more information on monitoring metrics, see Monitoring Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Monitoring/Concepts/overview.htm).
//

package telemetry

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // MetricData The set of aggregated data returned for a metric.
 // For more information on monitoring metrics, see Monitoring Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Monitoring/Concepts/overview.htm).
type MetricData struct {
    
 // The reference provided in a metric definition to indicate the source service or
 // application of the posted data point.
 // Example: `oci/compute`
    Namespace *string `mandatory:"true" json:"namespace"`
    
 // The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the
 // resources from which the aggregated data was returned.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // The name of the metric.
 // Example: `CpuUtilization`
    Name *string `mandatory:"true" json:"name"`
    
 // Qualifiers provided in the definition of the returned metric.
 // Available dimensions vary by metric namespace. Each dimension takes the form of a key-value pair.
 // Example: `resourceId`
    Dimensions map[string]string `mandatory:"true" json:"dimensions"`
    
 // The list of timestamp-value pairs returned for the specified request. Metric values are rolled up to the start time specified in the request.
    AggregatedDatapoints []AggregatedDatapoint `mandatory:"true" json:"aggregatedDatapoints"`
    
 // The references provided in a metric definition to indicate extra information about the metric.
 // Example: `unit`
    Metadata map[string]string `mandatory:"false" json:"metadata"`
    
 // The resolution value specified in the request. Resolution controls the time between
 // calculated aggregation windows.
 // Example: `1m`
    Resolution *string `mandatory:"false" json:"resolution"`
}

func (m MetricData) String() string {
    return common.PointerString(m)
}





