// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Load Balancing Service API
// This example creates a new loadbalancer with SSL cipher suites. After that it creates a new listener with the SSL configuration and updates the backend set with it. Finally it updates the loadbalancer cipher suites.
// The loadblancer allow rule and conditional add/drop header rules are now enhanced to restrict traffic based on the original client IP in the HTTP Real Client IP header in addition to the previous hop source IP.
package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/core"
	"github.com/oracle/oci-go-sdk/v32/example/helpers"
	"github.com/oracle/oci-go-sdk/v32/identity"
	"github.com/oracle/oci-go-sdk/v32/loadbalancer"
	"math"
	"time"
)

const (
	loadbalancerDisplayName = "OCI-GO-Sample-LB"
	nsgDisplayNameOne       = "OCI-GOSDK-Sample-NSG-1"
	nsgDisplayNameTwo       = "OCI-GOSDK-Sample-NSG-2"
	listenerDisplayName     = "GO_SDK_Listener"
	rulesetOneName          = "ruleset1"
	backendSetOneName       = "backendset1"
	cipherName              = "test-cipher"
	certificateName         = "example-certificate"
	publicCert              = `-----BEGIN CERTIFICATE-----
    publicKeyGoesHere
    -----END CERTIFICATE-----`

	privateKey = `-----BEGIN RSA PRIVATE KEY-----
    PrivateKeyGoesHere
    -----END RSA PRIVATE KEY-----`
	cidrBlocksOneName  = "cidrBlocks1"
	cidrBlocksIPv4Addr = "129.213.176.0/24"
	cidrBlocksIPv6Addr = "2002::1234:abcd:ffff:c0a8:101/64"
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

	cidrBlocks := map[string]loadbalancer.CidrBlocksDetails{
		cidrBlocksOneName: {
			Name:  common.String("cidrBlocks1"),
			Items: []string{cidrBlocksIPv4Addr, cidrBlocksIPv6Addr},
		},
	}
	request.CidrBlocks = cidrBlocks

	// Create rulesets to modify response / request headers or control access types based on REST request
	ruleSets := map[string]loadbalancer.RuleSetDetails{
		rulesetOneName: {
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
				loadbalancer.AllowRule{
					Description: common.String("Allow traffic from internet clients"),
					Conditions: []loadbalancer.RuleCondition{
						loadbalancer.RealIpAddressCondition{
							AttributeValue: common.String("cidrBlocks1"),
							HeaderName:     common.String("X-Real-IP"),
						},
					},
				},
				loadbalancer.AddHttpRequestHeaderRule{
					Header: common.String("some-header-name-to-add-cidrBlock"),
					Value:  common.String("some-value-for-header-cidrBlock"),
					Conditions: []loadbalancer.RuleCondition{
						loadbalancer.RealIpAddressCondition{
							AttributeValue: common.String("cidrBlocks1"),
							HeaderName:     common.String("X-Real-IP"),
						},
					},
				},
				loadbalancer.RemoveHttpRequestHeaderRule{
					Header: common.String("some-header-name-to-remove-cidrBlock"),
					Conditions: []loadbalancer.RuleCondition{
						loadbalancer.RealIpAddressCondition{
							AttributeValue: common.String("cidrBlocks1"),
							HeaderName:     common.String("X-Real-IP"),
						},
					},
				},
				loadbalancer.RedirectRule{
					ResponseCode: common.Int(302),
					Conditions: []loadbalancer.RuleCondition{
						loadbalancer.PathMatchCondition{
							AttributeValue: common.String("/xyz"),
							Operator:       loadbalancer.PathMatchConditionOperatorEnum("FORCE_LONGEST_PREFIX_MATCH"),
						},
					},
					RedirectUri: &loadbalancer.RedirectUri{
						Protocol: common.String("https"),
						Host:     common.String("def.com"),
						Port:     common.Int(8990),
						Path:     common.String("/xxx"),
					},
				},
				loadbalancer.AllowRule{
					Description: common.String("Allow traffic from internet clients"),
					Conditions: []loadbalancer.RuleCondition{
						loadbalancer.SourceIpAddressCondition{
							AttributeValue: common.String("111.111.111.111/32"),
						},
					},
				},
			}},
	}
	request.RuleSets = ruleSets

	// ssl cipher suites for the load balancer

	sslCiphers := map[string]loadbalancer.SslCipherSuiteDetails{
		cipherName: {
			Name:    common.String(cipherName),
			Ciphers: []string{"AES128-SHA", "AES256-SHA"},
		},
	}

	request.SslCipherSuites = sslCiphers

	cert := map[string]loadbalancer.CertificateDetails{
		certificateName: {
			CertificateName:   common.String(certificateName),
			PrivateKey:        common.String(privateKey),
			PublicCertificate: common.String(publicCert),
			CaCertificate:     common.String(publicCert),
		},
	}

	request.Certificates = cert

	// Backend Sets for our new LB. Includes an LB Cookie session persistence configuration. Note that this is
	//   mutually exclusive with a session persistence configuration.
	backendSets := map[string]loadbalancer.BackendSetDetails{
		backendSetOneName: {
			Policy: common.String("ROUND_ROBIN"),
			HealthChecker: &loadbalancer.HealthCheckerDetails{
				Protocol: common.String("HTTP"),
				UrlPath:  common.String("/health"),
				Port:     common.Int(80),
			},
			Backends: []loadbalancer.BackendDetails{
				{
					IpAddress: common.String("10.11.10.5"),
					Port:      common.Int(80),
				},
				{
					IpAddress: common.String("10.12.20.3"),
					Port:      common.Int(80),
				},
			},
			LbCookieSessionPersistenceConfiguration: &loadbalancer.LbCookieSessionPersistenceConfigurationDetails{
				CookieName:      common.String("X-Oracle-OCI-cookie-1"),
				DisableFallback: common.Bool(true),
				Domain:          common.String("www.example.org"),
				Path:            common.String("/cookiepath1"),
				MaxAgeInSeconds: common.Int(300),
				IsSecure:        common.Bool(false),
				IsHttpOnly:      common.Bool(false),
			},
			SslConfiguration: &loadbalancer.SslConfigurationDetails{
				CertificateName:       common.String("example-certificate"),
				VerifyPeerCertificate: common.Bool(true),
				CipherSuiteName:       common.String(cipherName),
				Protocols:             []string{"TLSv1.1"},
			},
		},
	}
	request.BackendSets = backendSets

	listeners := map[string]loadbalancer.ListenerDetails{
		listenerDisplayName: {
			DefaultBackendSetName: common.String(backendSetOneName),
			Port:                  common.Int(80),
			Protocol:              common.String("HTTP"),
			RuleSetNames:          []string{rulesetOneName},
			SslConfiguration: &loadbalancer.SslConfigurationDetails{
				CertificateName:       common.String("example-certificate"),
				VerifyPeerCertificate: common.Bool(true),
				CipherSuiteName:       common.String(cipherName),
				Protocols:             []string{"TLSv1.1"},
			},
		},
	}

	request.Listeners = listeners

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

	attempts := uint(10)
	retryIfLBNotReady := func(r common.OCIOperationResponse) bool {
		loadBalancer := getLoadBalancer()
		if loadBalancer != nil {
			fieldLifecycle, err := helpers.FindLifecycleFieldValue(loadBalancer)

			if err != nil {
				common.Logf("Error getting lifecycleState. Error is %v", err)
				return true
			}

			lifecycleState := string(loadbalancer.LoadBalancerLifecycleStateActive)
			isEqual := fieldLifecycle == lifecycleState
			if isEqual {
				return false
			}
			common.Logf("Current lifecycle state is %s, waiting for it to become %s", fieldLifecycle, lifecycleState)
			return true
		}
		common.Logf("LB not available, waiting...")
		return true
	}

	nextDuration := func(r common.OCIOperationResponse) time.Duration {
		// this function will return the duration as:
		// 1s, 2s, 4s, 8s, 16s, 32s, 64s etc...
		return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
	}

	defaultRetryPolicy := common.NewRetryPolicy(attempts, retryIfLBNotReady, nextDuration)

	request.RequestMetadata = common.RequestMetadata{
		RetryPolicy: &defaultRetryPolicy,
	}

	_, err = c.CreateLoadBalancer(ctx, request)
	helpers.FatalIfError(err)

	newCreatedLoadBalancer := getLoadBalancer()
	fmt.Printf("New loadbalancer LifecycleState is: %s\n\n", newCreatedLoadBalancer.LifecycleState)

	//Update cipher suites
	updateSSLCiphers(ctx, c, newCreatedLoadBalancer.Id)

	loadBalancerRuleSets := listRuleSets(ctx, c, newCreatedLoadBalancer.Id)
	fmt.Printf("Rule Sets from GET: %+v\n\n", loadBalancerRuleSets)

	newRuleSetResponse, err := addRuleSet(ctx, c, newCreatedLoadBalancer.Id)
	fmt.Printf("New rule set response: %+v\n\n", newRuleSetResponse)

	newBackendSetResponse, err := addBackendSet(ctx, c, newCreatedLoadBalancer.Id)
	fmt.Printf("New backend set: %+v\n\n", newBackendSetResponse)

	getListenerRulesResponse := listListenerRules(ctx, c, newCreatedLoadBalancer.Id, common.String(listenerDisplayName))
	fmt.Printf("Listener Rules: %+v\n\n", getListenerRulesResponse)

	// Change Compartment (Requires second compartment to move the LB into)
	secondCompartmentId := helpers.RootCompartmentID()
	changeCompartmentResponse, err := changeLBCompartment(ctx, c, newCreatedLoadBalancer.Id, secondCompartmentId)
	fmt.Printf("Load balancer compartment changed: %+v", changeCompartmentResponse)

	vnClient, vclerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(vclerr)
	vcnId := subnet1.VcnId

	nsg1 := createNsg(ctx, vnClient, nsgDisplayNameOne, helpers.CompartmentID(), vcnId)
	nsg2 := createNsg(ctx, vnClient, nsgDisplayNameTwo, helpers.CompartmentID(), vcnId)

	var nsgArray = []string{*(nsg1.Id), *(nsg2.Id)}

	//Update nsg call
	updateNsgWithLbr(ctx, c, newCreatedLoadBalancer.Id, nsgArray)

	// clean up resources
	defer func() {
		deleteLoadbalancer(ctx, c, newCreatedLoadBalancer.Id)

		vcnID := subnet1.VcnId
		deleteSubnet(ctx, vnClient, subnet1.Id)
		deleteSubnet(ctx, vnClient, subnet2.Id)
		deleteNsg(ctx, vnClient, nsg1.Id)
		deleteNsg(ctx, vnClient, nsg2.Id)
		deleteVcn(ctx, vnClient, vcnID)
	}()

	// Output:
	// create subnet1 complete
	// create subnet2 complete
	// list load balancer shapes complete
	// create load balancer complete
	// new loadbalancer LifecycleState is: ACTIVE
	// Rule Sets from GET: {}
	// New rule set response: {}
	// New backend set: {}
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

// Add a new ruleset to an existing LB
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

// Get a list of rulesets from a given LB
func listRuleSets(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string) []loadbalancer.RuleSet {
	request := loadbalancer.ListRuleSetsRequest{
		LoadBalancerId: id,
	}

	r, err := client.ListRuleSets(ctx, request)
	helpers.FatalIfError(err)
	return r.Items
}

// Add a new backend set to an existing LB
func addBackendSet(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string) (loadbalancer.CreateBackendSetResponse, error) {
	request := loadbalancer.CreateBackendSetRequest{}
	request.LoadBalancerId = id
	backendSetDetails := loadbalancer.CreateBackendSetDetails{
		Name:   common.String("backendset2"),
		Policy: common.String("ROUND_ROBIN"),
		HealthChecker: &loadbalancer.HealthCheckerDetails{
			Protocol: common.String("HTTP"),
			UrlPath:  common.String("/health"),
			Port:     common.Int(80),
		},
		Backends: []loadbalancer.BackendDetails{
			{
				IpAddress: common.String("10.11.10.5"),
				Port:      common.Int(80),
			},
			{
				IpAddress: common.String("10.12.20.3"),
				Port:      common.Int(80),
			},
		},
		LbCookieSessionPersistenceConfiguration: &loadbalancer.LbCookieSessionPersistenceConfigurationDetails{
			CookieName:      common.String("X-Oracle-OCI-cookie-2"),
			DisableFallback: common.Bool(true),
			Domain:          common.String("www.example.org"),
			Path:            common.String("/cookiepath2"),
			MaxAgeInSeconds: common.Int(300),
			IsSecure:        common.Bool(false),
			IsHttpOnly:      common.Bool(false),
		},
	}
	request.CreateBackendSetDetails = backendSetDetails

	response, err := client.CreateBackendSet(ctx, request)
	helpers.FatalIfError(err)
	println("backendset added")
	return response, err
}

// Move the LB to a new compartment
func changeLBCompartment(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string, compartmentId *string) (loadbalancer.ChangeLoadBalancerCompartmentResponse, error) {
	request := loadbalancer.ChangeLoadBalancerCompartmentRequest{}
	request.LoadBalancerId = id
	changeCompartmentDetails := loadbalancer.ChangeLoadBalancerCompartmentDetails{
		CompartmentId: compartmentId,
	}
	request.ChangeLoadBalancerCompartmentDetails = changeCompartmentDetails
	response, err := client.ChangeLoadBalancerCompartment(ctx, request)
	helpers.FatalIfError(err)

	return response, err
}

//Update ssl ciphers
func updateSSLCiphers(ctx context.Context, c loadbalancer.LoadBalancerClient, loadBalancerId *string) {
	request := loadbalancer.UpdateSSLCipherSuiteRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	details := loadbalancer.UpdateSslCipherSuiteDetails{
		Ciphers: []string{"ECDHE-ECDSA-AES256-GCM-SHA384"},
	}
	request.LoadBalancerId = loadBalancerId
	request.UpdateSslCipherSuiteDetails = details
	_, err := c.UpdateSSLCipherSuite(ctx, request)
	helpers.FatalIfError(err)
}

// Create network security group
func createNsg(ctx context.Context, c core.VirtualNetworkClient, displayName string, compartmentId, vcnId *string) core.NetworkSecurityGroup {
	// create a new nsg
	createNsgRequest := core.CreateNetworkSecurityGroupRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	cnsgDetails := core.CreateNetworkSecurityGroupDetails{}
	cnsgDetails.CompartmentId = compartmentId
	cnsgDetails.DisplayName = common.String(displayName)
	cnsgDetails.VcnId = vcnId

	createNsgRequest.CreateNetworkSecurityGroupDetails = cnsgDetails

	r, err := c.CreateNetworkSecurityGroup(ctx, createNsgRequest)
	helpers.FatalIfError(err)
	return r.NetworkSecurityGroup
}

// Delete network security group
func deleteNsg(ctx context.Context, c core.VirtualNetworkClient, nsgId *string) {
	//delete the nsg
	deleteNsgRequest := core.DeleteNetworkSecurityGroupRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}

	deleteNsgRequest.NetworkSecurityGroupId = nsgId

	_, err := c.DeleteNetworkSecurityGroup(ctx, deleteNsgRequest)
	helpers.FatalIfError(err)
}

// Update nsg list with load balancer
func updateNsgWithLbr(ctx context.Context, c loadbalancer.LoadBalancerClient, loadBalancerId *string, networkSecurityGroupIds []string) {
	request := loadbalancer.UpdateNetworkSecurityGroupsRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	details := loadbalancer.UpdateNetworkSecurityGroupsDetails{
		NetworkSecurityGroupIds: networkSecurityGroupIds,
	}
	request.LoadBalancerId = loadBalancerId
	request.UpdateNetworkSecurityGroupsDetails = details

	_, err := c.UpdateNetworkSecurityGroups(ctx, request)
	helpers.FatalIfError(err)
}

// List Listener Rules for a given listener (by load balancer id and listener name)
func listListenerRules(ctx context.Context, client loadbalancer.LoadBalancerClient, id *string, name *string) []loadbalancer.ListenerRuleSummary {
	request := loadbalancer.ListListenerRulesRequest{
		LoadBalancerId: id,
		ListenerName:   name,
	}
	r, err := client.ListListenerRules(ctx, request)
	helpers.FatalIfError(err)
	return r.Items
}
