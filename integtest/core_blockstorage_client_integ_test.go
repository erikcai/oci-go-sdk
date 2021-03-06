// Copyright (c) 2016, 2017, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package integtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/erikcai/oci-go-sdk/v33/common"
	"github.com/erikcai/oci-go-sdk/v33/core"
	"github.com/stretchr/testify/assert"
)

var (
	tick    = time.Tick(500 * time.Millisecond)
	timeout = time.After(30 * time.Second)
)

//Volumes CRUDL
func TestBlockstorageClient_CreateVolume(t *testing.T) {

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.CreateVolumeRequest{}

	request.AvailabilityDomain = common.String(validAD())
	request.CompartmentId = common.String(getTenancyID())
	request.DisplayName = common.String("GoSDK2_TestBlockStorageVolume_" + getUuid())

	resCreate, err := c.CreateVolume(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, resCreate.Id)

	requestGet := core.GetVolumeRequest{}
	requestGet.VolumeId = resCreate.Id
	requestGet.RequestMetadata = common.RequestMetadata{
		RetryPolicy: &common.RetryPolicy{
			MaximumNumberAttempts: 10,
			ShouldRetryOperation: func(response common.OCIOperationResponse) bool {
				if response.Error != nil {
					return false
				}
				volResponse, ok := response.Response.(core.GetVolumeResponse)
				if !ok {
					return false
				}
				return !(volResponse.LifecycleState == core.VolumeLifecycleStateAvailable)
			},
			NextDuration: func(response common.OCIOperationResponse) time.Duration {
				return 10 * time.Second
			},
		},
	}

	c.GetVolume(context.Background(), requestGet)

	//update
	updateTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.UpdateVolumeRequest{}
		request.VolumeId = resCreate.Id
		request.DisplayName = common.String("GoSDK2_TestBlockStorageVolumeUpdate_" + getUuid())
		r, err := c.UpdateVolume(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		assert.NoError(t, err)
		return
	}
	updateTest(t)

	//list
	listTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.ListVolumesRequest{}
		request.CompartmentId = common.String(getTenancyID())
		r, err := c.ListVolumes(context.Background(), request)
		assert.NoError(t, err)
		assert.NotEmpty(t, r.Items)
		return
	}
	listTest(t)

	//delete
	deleteTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)

		volumeState := resCreate.LifecycleState
		for volumeState != core.VolumeLifecycleStateAvailable {
			time.Sleep(15 * time.Second)
			getVolumeRequest, _ := c.GetVolume(context.Background(), core.GetVolumeRequest{
				VolumeId: resCreate.Id,
			})
			volumeState = getVolumeRequest.LifecycleState
		}
		getRequest := core.GetVolumeRequest{}
		getRequest.VolumeId = resCreate.Id
		c.GetVolume(context.Background(), getRequest)

		request := core.DeleteVolumeRequest{}
		request.VolumeId = resCreate.Id
		_, err := c.DeleteVolume(context.Background(), request)
		assert.NoError(t, err)
		return
	}
	deleteTest(t)

	return
}

func TestBlockstorageClient_CreateVolumeBackup(t *testing.T) {

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	var volumeId *string
	//First create a volume
	createVol := func(t *testing.T) {
		rCreateVol := core.CreateVolumeRequest{}
		rCreateVol.AvailabilityDomain = common.String(validAD())
		rCreateVol.CompartmentId = common.String(getTenancyID())
		rCreateVol.DisplayName = common.String("GoSDK2_TestBlockStorageVolume_" + getUuid())
		resCreate, err := c.CreateVolume(context.Background(), rCreateVol)
		volumeId = resCreate.Id
		failIfError(t, err)
	}

	deleteVol := func() {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.DeleteVolumeRequest{}
		request.VolumeId = volumeId
		c.DeleteVolume(context.Background(), request)
		return
	}

	readVol := func() (interface{}, error) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.GetVolumeRequest{}
		request.VolumeId = volumeId
		resRead, err := c.GetVolume(context.Background(), request)
		return resRead, err
	}

	createVol(t)
	defer deleteVol()
	failIfError(t,
		retryUntilTrueOrError(readVol,
			checkLifecycleState(string(core.VolumeLifecycleStateAvailable)), tick, time.After(120*time.Second)))

	//Create
	request := core.CreateVolumeBackupRequest{}
	request.VolumeId = volumeId
	request.DisplayName = common.String("GoSDK2_TestBlockStorageVolumeBackup_" + getUuid())
	r, err := c.CreateVolumeBackup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	failIfError(t, err)

	//Read
	readTest := func() (interface{}, error) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.GetVolumeBackupRequest{}
		request.VolumeBackupId = r.Id
		rRead, err := c.GetVolumeBackup(context.Background(), request)
		return rRead, err
	}

	failIfError(t,
		retryUntilTrueOrError(readTest,
			checkLifecycleState(string(core.VolumeBackupLifecycleStateAvailable)), tick, time.After(120*time.Second)))

	//List
	listTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.ListVolumeBackupsRequest{}
		request.CompartmentId = common.String(getTenancyID())
		request.VolumeId = volumeId
		r, err := c.ListVolumeBackups(context.Background(), request)
		failIfError(t, err)
		assert.True(t, len(r.Items) > 0)
		return
	}
	listTest(t)

	//Update
	updateTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.UpdateVolumeBackupRequest{}
		request.VolumeBackupId = r.Id
		request.DisplayName = common.String("GoSDK2_TestBlockStorageVolumeBackupUpdate_" + getUuid())
		r, err := c.UpdateVolumeBackup(context.Background(), request)
		assert.NotEmpty(t, r, fmt.Sprint(r))
		failIfError(t, err)
		return
	}
	updateTest(t)

	//Delete
	deleteTest := func(t *testing.T) {
		c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
		failIfError(t, clerr)
		request := core.DeleteVolumeBackupRequest{}
		request.VolumeBackupId = r.Id
		_, err := c.DeleteVolumeBackup(context.Background(), request)
		failIfError(t, err)
		return
	}

	deleteTest(t)
	return
}

func TestBlockstorageClient_ListBootVolumes(t *testing.T) {
	bootVolumes := listBootVolumes(t)
	assert.NotEmpty(t, bootVolumes)
	return
}

func TestBlockstorageClient_GetBootVolume(t *testing.T) {
	// get list of boot volumes and make sure it's not empty
	bootVolumes := listBootVolumes(t)
	assert.NotEmpty(t, bootVolumes)

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := core.GetBootVolumeRequest{
		BootVolumeId: bootVolumes[0].Id,
	}

	r, err := c.GetBootVolume(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r.Id)
	return
}
