package integtest

import (
	"context"
	"testing"

	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/storagegateway"
	"github.com/stretchr/testify/assert"
)

func TestStorageGateway_Create(t *testing.T) {
	t.Skip("Not implemented")
	p := configurationProvider()
	cId, _ := p.TenancyOCID()
	c, _ := storagegateway.NewStorageGatewayClientWithConfigurationProvider(p)

	res, err := c.CreateStorageGateway(context.Background(), storagegateway.CreateStorageGatewayRequest{
		CreateStorageGatewayDetails: storagegateway.CreateStorageGatewayDetails{
			CompartmentId: common.String(cId),
			DisplayName:   common.String("TestStorage"),
			Description:   common.String("some sample des"),
		},
	})

	assert.NoError(t, err)
	defer deleteStorageGateway(res.Id, t)
}

func deleteStorageGateway(id *string, t *testing.T) {
	p := configurationProvider()
	c, _ := storagegateway.NewStorageGatewayClientWithConfigurationProvider(p)

	_, err := c.DeleteStorageGateway(context.Background(), storagegateway.DeleteStorageGatewayRequest{
		StorageGatewayId: id,
	})

	assert.NoError(t, err)

}
