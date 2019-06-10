// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for Load Balancing Service API
//

package example

import (
	"context"
	"fmt"
	"time"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/oracle/oci-go-sdk/loadbalancer"
)

const (
	loadbalancerDisplayName = "OCI-GO-Sample-LB"
)

func ExampleCreateLoadbalancer() {

	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	ctx := context.Background()
	helpers.FatalIfError(clerr)
	request := loadbalancer.CreateLoadBalancerRequest{}
	request.CompartmentId = helpers.CompartmentID()
	request.DisplayName = common.String(loadbalancerDisplayName)

	subnet1 := CreateOrGetSubnet()
	fmt.Println("create subnet1 complete")

	// create a subnet in different availability domain
	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	req := identity.ListAvailabilityDomainsRequest{}
	req.CompartmentId = helpers.CompartmentID()
	response, err := identityClient.ListAvailabilityDomains(ctx, req)
	helpers.FatalIfError(err)
	availableDomain := response.Items[1].Name

	subnet2 := CreateOrGetSubnetWithDetails(common.String(subnetDisplayName2), common.String("10.0.1.0/24"), common.String("subnetdns2"), availableDomain)
	fmt.Println("create subnet2 complete")

	request.SubnetIds = []string{*subnet1.Id, *subnet2.Id}

	shapes := listLoadBalancerShapes(ctx, c)
	fmt.Println("list load balancer shapes complete")

	request.ShapeName = shapes[0].Name

	ruleSets := map[string]loadbalancer.RuleSetDetails{
		"ruleset1": loadbalancer.RuleSetDetails{
			Items: []loadbalancer.Rule{
				loadbalancer.AddHttpRequestHeaderRule{
					Header: common.String("some-header-name-to-add"),
					Value:  common.String("some-value-for-header"),
				},
				loadbalancer.RemoveHttpResponseHeaderRule{
					Header: common.String("some-header-name-to-remove"),
				},
				loadbalancer.ExtendHttpRequestHeaderValueRule{
					Header: common.String("some-other-header-name-to-alter"),
					Prefix: common.String("some-prefix-value-for-header"),
					Suffix: common.String("some-suffix-value-for-header"),
				},
				loadbalancer.ControlAccessUsingHttpMethodsRule{
					AllowedMethods: []string{
						"PUT",
						"POST",
					},
					StatusCode: common.Int(403),
				},
			}},
	}
	request.RuleSets = ruleSets

	_, err = c.CreateLoadBalancer(ctx, request)
	helpers.FatalIfError(err)

	fmt.Println("create load balancer complete")

	// get created loadbalancer
	getLoadBalancer := func() *loadbalancer.LoadBalancer {
		loadbalancers := listLoadBalancers(ctx, c, loadbalancer.LoadBalancerLifecycleStateActive)
		for _, element := range loadbalancers {
			if *element.DisplayName == loadbalancerDisplayName {
				// found it, return
				return &element
			}
		}

		return nil
	}

	// use to check the lifecycle states of new load balancer
	loadBalancerLifecycleStateCheck := func() (interface{}, error) {
		loadBalancer := getLoadBalancer()
		if loadBalancer != nil {
			return loadBalancer, nil
		}

		return loadbalancer.LoadBalancer{}, nil
	}

	// wait for instance lifecyle state become running
	helpers.FatalIfError(
		helpers.RetryUntilTrueOrError(
			loadBalancerLifecycleStateCheck,
			helpers.CheckLifecycleState(string(loadbalancer.LoadBalancerLifecycleStateActive)),
			time.Tick(10*time.Second),
			time.After((5 * time.Minute))))

	newCreatedLoadBalancer := getLoadBalancer()
	fmt.Printf("new loadbalancer LifecycleState is: %s\n", newCreatedLoadBalancer.LifecycleState)

	loadBalancerRuleSets := listRuleSets(ctx, c, newCreatedLoadBalancer.Id)
	fmt.Printf("Rule Sets from GET: %+v", loadBalancerRuleSets)

	newRuleSetResponse, err := addRuleSet(ctx, c, newCreatedLoadBalancer.Id)
	fmt.Printf("New rule set response: %+v", newRuleSetResponse)

	// clean up resources
	defer func() {
		deleteLoadbalancer(ctx, c, newCreatedLoadBalancer.Id)

		client, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
		helpers.FatalIfError(clerr)

		vcnID := subnet1.VcnId
		deleteSubnet(ctx, client, subnet1.Id)
		deleteSubnet(ctx, client, subnet2.Id)
		deleteVcn(ctx, client, vcnID)
	}()

	// Output:
	// create subnet1 complete
	// create subnet2 complete
	// list load balancer shapes complete
	// create load balancer complete
	// new loadbalancer LifecycleState is: ACTIVE
	// Rule Sets from GET: {}
	// New rule set response: {}
	// deleting load balancer
	// load balancer deleted
	// deleteing subnet
	// subnet deleted
	// deleteing subnet
	// subnet deleted
	// deleteing VCN
	// VCN deleted
}

func listLoadBalancerShapes(ctx context.Context, client loadbalancer.LoadBalancerClient) []loadbalancer.LoadBalancerShape {
	request := loadbalancer.ListShapesRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	r, err := client.ListShapes(ctx, request)
	helpers.FatalIfError(err)
	return r.Items
}

func listLoadBalancers(ctx context.Context, client loadbalancer.LoadBalancerClient, lifecycleState loadbalancer.LoadBalancerLifecycleStateEnum) []loadbalancer.LoadBalancer {
	request := loadbalancer.ListLoadBalancersRequest{
		CompartmentId:  helpers.CompartmentID(),
		DisplayName:    common.String(loadbalancerDisplayName),
		LifecycleState: lifecycleState,
	}

	r, err := client.ListLoadBalancers(ctx, request)
	helpers.FatalIfError(err)
	return r.Items
}

func deleteLoadbalancer(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string) {
	request := loadbalancer.DeleteLoadBalancerRequest{
		LoadBalancerId: id,
	}

	_, err := client.DeleteLoadBalancer(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("deleting load balancer")

	// get loadbalancer
	getLoadBalancer := func() *loadbalancer.LoadBalancer {
		loadbalancers := listLoadBalancers(ctx, client, loadbalancer.LoadBalancerLifecycleStateDeleting)
		for _, element := range loadbalancers {
			if *element.DisplayName == loadbalancerDisplayName {
				// found it, return
				return &element
			}
		}

		return nil
	}

	// use to check the lifecycle state of load balancer
	loadBalancerLifecycleStateCheck := func() (interface{}, error) {
		loadBalancer := getLoadBalancer()
		if loadBalancer != nil {
			return loadBalancer, nil
		}

		// cannot find load balancer which means it's been deleted
		return loadbalancer.LoadBalancer{LifecycleState: loadbalancer.LoadBalancerLifecycleStateDeleted}, nil
	}

	// wait for load balancer been deleted
	helpers.FatalIfError(
		helpers.RetryUntilTrueOrError(
			loadBalancerLifecycleStateCheck,
			helpers.CheckLifecycleState(string(loadbalancer.LoadBalancerLifecycleStateDeleted)),
			time.Tick(10*time.Second),
			time.After((10 * time.Minute))))

	fmt.Println("load balancer deleted")
}

func addRuleSet(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string) (loadbalancer.CreateRuleSetResponse, error) {
	request := loadbalancer.CreateRuleSetRequest{}
	request.LoadBalancerId = id
	ruleSetDetails := loadbalancer.CreateRuleSetDetails{
		Name: common.String("ruleset2"),
		Items: []loadbalancer.Rule{
			loadbalancer.AddHttpResponseHeaderRule{
				Header: common.String("some-second-header-name-to-add"),
				Value:  common.String("some-second-value-for-header"),
			},
			loadbalancer.RemoveHttpRequestHeaderRule{
				Header: common.String("some-second-header-name-to-remove"),
			},
		},
	}
	request.CreateRuleSetDetails = ruleSetDetails

	response, err := client.CreateRuleSet(ctx, request)
	helpers.FatalIfError(err)
	println("ruleset added")
	return response, err
}

func listRuleSets(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string) []loadbalancer.RuleSet {
	request := loadbalancer.ListRuleSetsRequest{
		LoadBalancerId: id,
	}

	r, err := client.ListRuleSets(ctx, request)
	helpers.FatalIfError(err)
	return r.Items
}
