// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in Government Cloud tenancies. For more information, see
// Information for Oracle Cloud Infrastructure Government Cloud Customers (https://docs.cloud.oracle.com/Content/General/Concepts/govinfo.htm).
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CronExecutionSchedule Specifies the execution schedule of CRON type.
type CronExecutionSchedule struct {

	// The value representing the execution schedule, as defined by cron format.
	Expression *string `mandatory:"true" json:"expression"`

	// Specifies the time zone the schedule is in.
	Timezone ExecutionScheduleTimezoneEnum `mandatory:"true" json:"timezone"`
}

//GetTimezone returns Timezone
func (m CronExecutionSchedule) GetTimezone() ExecutionScheduleTimezoneEnum {
	return m.Timezone
}

func (m CronExecutionSchedule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CronExecutionSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCronExecutionSchedule CronExecutionSchedule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCronExecutionSchedule
	}{
		"cron",
		(MarshalTypeCronExecutionSchedule)(m),
	}

	return json.Marshal(&s)
}