// Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.
//
// Example code for Core Services API
//
//

/**
 * This class provides an example of how you can create an AutoScalingConfiguration and use that with a InstancePools. It will:
 * <ul>
 * <li>Create the InstanceConfiguration</li>
 * <li>Create a pool based off that configuration.</li>
 * <li>Create an AutoScalingConfiguration for that pool.</li>
 * <li>Clean everything up.</li>
 * </ul>
 */

package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/oracle/oci-go-sdk/autoscaling"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

var (
	imageId, ad, subnetId string
)

// Example to showcase autoscaling configuration creation and eventual teardown
func ExampleCreateAndDeleteAutoscalingConfiguration() {
	AutoscalingParseEnvironmentVariables()

	ctx := context.Background()

	computeMgmtClient, err := core.NewComputeManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	autoscalingClient, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	createInstanceConfigurationResponse, _ := createInstanceConfiguration(ctx, computeMgmtClient, imageId, compartmentId)
	fmt.Println("Instance configuration created")

	instanceConfiguration := createInstanceConfigurationResponse.InstanceConfiguration

	instancePool, _ := createInstancePool(ctx, computeMgmtClient, *instanceConfiguration.Id, subnetId, ad, compartmentId)
	fmt.Println("Instance pool created")

	autoscalingConfiguration, _ := createAutoscalingConfiguration(ctx, autoscalingClient, *instancePool.Id, compartmentId)

	fmt.Println("Autoscaling configuration created")

	// clean up resources
	defer func() {
		deleteAutoscalingConfiguration(ctx, autoscalingClient, *autoscalingConfiguration.Id)
		fmt.Println("Deleted Autoscaling Configuration")

		terminateInstancePool(ctx, computeMgmtClient, *instancePool.Id)
		fmt.Println("Terminated Instance Pool")

		deleteInstanceConfiguration(ctx, computeMgmtClient, *instanceConfiguration.Id)
		fmt.Println("Deleted Instance Configuration")
	}()

	// Output:
	// Instance configuration created
	// Instance pool created
	// Autoscaling configuration created
	// Deleted Autoscaling Configuration
	// Terminated Instance Pool
	// Deleted Instance Configuration
}

// Usage printing
func AutoscalingUsage() {
	log.Printf("Please set the following environment variables to run ChangeAutoscalingCompartment sample")
	log.Printf(" ")
	log.Printf("   IMAGE_ID       # Required: Image Id to use")
	log.Printf("   COMPARTMENT_ID    # Required: Compartment Id to use")
	log.Printf("   AD          # Required: AD to use")
	log.Printf("   SUBNET_ID   # Required: Subnet to use")
	log.Printf(" ")
	os.Exit(1)
}

// Args parser
func AutoscalingParseEnvironmentVariables() {

	imageId = os.Getenv("IMAGE_ID")
	compartmentId = os.Getenv("COMPARTMENT_ID")
	ad = os.Getenv("AD")
	subnetId = os.Getenv("SUBNET_ID")

	if imageId == "" ||
		compartmentId == "" ||
		ad == "" ||
		subnetId == "" {
		AutoscalingUsage()
	}

	log.Printf("IMAGE_ID     : %s", imageId)
	log.Printf("COMPARTMENT_ID  : %s", compartmentId)
	log.Printf("AD     : %s", ad)
	log.Printf("SUBNET_ID  : %s", subnetId)
}

// helper method to create an instance configuration
func createInstanceConfiguration(ctx context.Context, client core.ComputeManagementClient, imageId string, compartmentId string) (response core.CreateInstanceConfigurationResponse, err error) {
	vnicDetails := core.InstanceConfigurationCreateVnicDetails{}

	sourceDetails := core.InstanceConfigurationInstanceSourceViaImageDetails{
		ImageId: &imageId,
	}

	displayName := "Instance Configuration Example"
	shape := "VM.Standard2.1"

	launchDetails := core.InstanceConfigurationLaunchInstanceDetails{
		CompartmentId:     &compartmentId,
		DisplayName:       &displayName,
		CreateVnicDetails: &vnicDetails,
		Shape:             &shape,
		SourceDetails:     &sourceDetails,
	}

	instanceDetails := core.ComputeInstanceDetails{
		LaunchDetails: &launchDetails,
	}

	configurationDetails := core.CreateInstanceConfigurationDetails{
		DisplayName:     &displayName,
		CompartmentId:   &compartmentId,
		InstanceDetails: &instanceDetails,
	}

	req := core.CreateInstanceConfigurationRequest{
		CreateInstanceConfiguration: &configurationDetails,
	}

	response, err = client.CreateInstanceConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}

// helper method to create an instance pool
func createInstancePool(ctx context.Context, client core.ComputeManagementClient, instanceConfigurationId string,
	subnetId string, availabilityDomain string, compartmentId string) (response core.CreateInstancePoolResponse, err error) {

	displayName := "Instance Pool Example"
	size := 1

	req := core.CreateInstancePoolRequest{
		CreateInstancePoolDetails: core.CreateInstancePoolDetails{
			CompartmentId:           &compartmentId,
			InstanceConfigurationId: &instanceConfigurationId,
			PlacementConfigurations: []core.CreateInstancePoolPlacementConfigurationDetails{
				{
					PrimarySubnetId:    &subnetId,
					AvailabilityDomain: &availabilityDomain,
				},
			},
			Size:        &size,
			DisplayName: &displayName,
		},
	}

	response, err = client.CreateInstancePool(ctx, req)
	return
}

// helper method to create an autoscaling configuration
func createAutoscalingConfiguration(ctx context.Context, client autoscaling.AutoScalingClient,
	instancePoolId string, compartmentId string) (response autoscaling.CreateAutoScalingConfigurationResponse, err error) {

	displayName := "Autoscaling Example"

	scaleInThreshold := 30
	scaleInChange := -1
	scaleOutThreshold := 70
	scaleOutChange := 1
	capInit := 2
	capMax := 3
	capMin := 1

	resource := autoscaling.InstancePoolResource{
		Id: &instancePoolId,
	}

	capacity := autoscaling.Capacity{
		Initial: &capInit,
		Max:     &capMax,
		Min:     &capMin,
	}

	// scale in params
	lowerBound := autoscaling.Threshold{
		Operator: autoscaling.ThresholdOperatorLt,
		Value:    &scaleInThreshold,
	}

	scaleInAction := autoscaling.Action{
		Type:  autoscaling.ActionTypeBy,
		Value: &scaleInChange,
	}

	scaleInMetric := autoscaling.Metric{
		Threshold:  &lowerBound,
		MetricType: autoscaling.MetricMetricTypeCpuUtilization,
	}

	scaleInRule := autoscaling.CreateConditionDetails{
		Action: &scaleInAction,
		Metric: &scaleInMetric,
	}

	// scale out params
	upperBound := autoscaling.Threshold{
		Operator: autoscaling.ThresholdOperatorGt,
		Value:    &scaleOutThreshold,
	}

	scaleOutAction := autoscaling.Action{
		Type:  autoscaling.ActionTypeBy,
		Value: &scaleOutChange,
	}

	scaleOutMetric := autoscaling.Metric{
		Threshold:  &upperBound,
		MetricType: autoscaling.MetricMetricTypeCpuUtilization,
	}

	scaleOutRule := autoscaling.CreateConditionDetails{
		Action: &scaleOutAction,
		Metric: &scaleOutMetric,
	}

	// defining the threshold policy
	policy := autoscaling.CreateThresholdPolicyDetails{
		Capacity: &capacity,
		Rules: []autoscaling.CreateConditionDetails{
			scaleInRule,
			scaleOutRule,
		},
	}

	// defining the autoscaling configuration
	createAutoscalingConfigurationDetails := autoscaling.CreateAutoScalingConfigurationDetails{
		DisplayName:   &displayName,
		CompartmentId: &compartmentId,
		Resource:      &resource,
		Policies:      []autoscaling.CreateAutoScalingPolicyDetails{&policy},
	}

	req := autoscaling.CreateAutoScalingConfigurationRequest{
		CreateAutoScalingConfigurationDetails: createAutoscalingConfigurationDetails,
	}

	response, err = client.CreateAutoScalingConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}

// helper method to delete an instance configuration
func deleteAutoscalingConfiguration(ctx context.Context, client autoscaling.AutoScalingClient,
	autoscalingConfigurationId string) (response autoscaling.DeleteAutoScalingConfigurationResponse, err error) {

	req := autoscaling.DeleteAutoScalingConfigurationRequest{
		AutoScalingConfigurationId: &autoscalingConfigurationId,
	}

	response, err = client.DeleteAutoScalingConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}

// helper method to termiante an instance configuration
func terminateInstancePool(ctx context.Context, client core.ComputeManagementClient,
	poolId string) (response core.TerminateInstancePoolResponse, err error) {

	req := core.TerminateInstancePoolRequest{
		InstancePoolId: &poolId,
	}

	response, err = client.TerminateInstancePool(ctx, req)
	helpers.FatalIfError(err)

	return
}

// helper method to delete an instance configuration
func deleteInstanceConfiguration(ctx context.Context, client core.ComputeManagementClient,
	instanceConfigurationId string) (response core.DeleteInstanceConfigurationResponse, err error) {

	req := core.DeleteInstanceConfigurationRequest{
		InstanceConfigurationId: &instanceConfigurationId,
	}

	response, err = client.DeleteInstanceConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}
