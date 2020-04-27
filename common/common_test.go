// Copyright (c) 2016 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	// OC1
	region := StringToRegion("us-phoenix-1")
	endpoint := region.Endpoint("foo")
	assert.Equal(t, "foo.us-phoenix-1.oraclecloud.com", endpoint)

	region = StringToRegion("us-ashburn-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-ashburn-1.oraclecloud.com", endpoint)

	// OC2
	region = StringToRegion("us-langley-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-langley-1.oraclegovcloud.com", endpoint)

	// OC3
	region = StringToRegion("us-gov-ashburn-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-gov-ashburn-1.oraclegovcloud.com", endpoint)

	// OC4
	region = StringToRegion("uk-gov-cardiff-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.uk-gov-cardiff-1.oraclegovcloud.uk", endpoint)

	//OC5
	region = StringToRegion("us-tacoma-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.us-tacoma-1.oracleonsrcloud.com", endpoint)

	//OC8
	region = StringToRegion("ap-chiyoda-1")
	endpoint = region.Endpoint("bar")
	assert.Equal(t, "bar.ap-chiyoda-1.oraclecloud8.com", endpoint)
}

func TestEndpointForTemplate(t *testing.T) {
	type testData struct {
		region           Region
		service          string
		endpointTemplate string
		expected         string
	}
	testDataSet := []testData{
		{
			// template with region
			region:           StringToRegion("us-phoenix-1"),
			service:          "test",
			endpointTemplate: "https://foo.{region}.bar.com",
			expected:         "https://foo.us-phoenix-1.bar.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("us-phoenix-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("ap-sydney-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("ap-hyderabad-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("ap-chuncheon-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("uk-cardiff-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("us-sanjose-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("ap-singapore-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with second level domain
			region:           StringToRegion("me-dubai-1"),
			service:          "test",
			endpointTemplate: "https://foo.region.{secondLevelDomain}",
			expected:         "https://foo.region.oraclecloud.com",
		},
		{
			// template with everything for OC2
			region:           StringToRegion("us-langley-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.us-langley-1.oraclegovcloud.com",
		},
		{
			// template with everything for OC3
			region:           StringToRegion("us-gov-ashburn-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.us-gov-ashburn-1.oraclegovcloud.com",
		},
		{
			// template with everything for OC4
			region:           StringToRegion("uk-gov-cardiff-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.uk-gov-cardiff-1.oraclegovcloud.uk",
		},
		{
			// template with everything for OC5
			region:           StringToRegion("us-tacoma-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.us-tacoma-1.oracleonsrcloud.com",
		},
		{
			// template with everything for OC8
			region:           StringToRegion("ap-chiyoda-1"),
			service:          "test",
			endpointTemplate: "https://test.{region}.{secondLevelDomain}",
			expected:         "https://test.ap-chiyoda-1.oraclecloud8.com",
		},
	}

	for _, testData := range testDataSet {
		endpoint := testData.region.EndpointForTemplate(testData.service, testData.endpointTemplate)
		assert.Equal(t, testData.expected, endpoint)
	}
}

func TestStringToRegion(t *testing.T) {
	region := StringToRegion("yyz")
	assert.Equal(t, RegionCAToronto1, region)

	region = StringToRegion("nrt")
	assert.Equal(t, RegionAPTokyo1, region)

	region = StringToRegion("gru")
	assert.Equal(t, RegionSASaopaulo1, region)

	region = StringToRegion("yny")
	assert.Equal(t, RegionAPChuncheon1, region)

	region = StringToRegion("cwl")
	assert.Equal(t, RegionUKCardiff1, region)

	region = StringToRegion("hyd")
	assert.Equal(t, RegionAPHyderabad1, region)

	region = StringToRegion("syd")
	assert.Equal(t, RegionAPSydney1, region)

	region = StringToRegion("sjc")
	assert.Equal(t, RegionSJC1, region)

	region = StringToRegion("dxb")
	assert.Equal(t, RegionMEDubai1, region)

	region = StringToRegion("sin")
	assert.Equal(t, RegionAPSingapore1, region)

	region = StringToRegion("tiw")
	assert.Equal(t, RegionUSTacoma1, region)

	region = StringToRegion("nja")
	assert.Equal(t, RegionAPChiyoda1, region)
}
