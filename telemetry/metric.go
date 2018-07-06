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


    
 // Metric The properties that define a metric.
 // For more information on monitoring metrics, see Monitoring Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Monitoring/Concepts/overview.htm).
type Metric struct {
    
 // The name of the metric.
 // Example: `CpuUtilization`
    Name *string `mandatory:"true" json:"name"`
    
 // The source service or application of the metric.
 // Example: `oci/compute`
    Namespace *string `mandatory:"true" json:"namespace"`
    
 // The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing
 // the resources monitored by the metric.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
 // Each dimension takes the form of a key-value pair.
 // Example: "resourceId": "<instance_OCID>"
    Dimensions map[string]string `mandatory:"false" json:"dimensions"`
    
 // The references provided in a metric definition to indicate extra information about the metric.
 // Example: `"unit": "milliseconds"`
    Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m Metric) String() string {
    return common.PointerString(m)
}





