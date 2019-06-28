// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Resource Ticket Item
type Resource struct {
	Item Item `mandatory:"false" json:"item"`

	Region ResourceRegionEnum `mandatory:"false" json:"region,omitempty"`

	AvailabilityDomain ResourceAvailabilityDomainEnum `mandatory:"false" json:"availabilityDomain,omitempty"`
}

func (m Resource) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Resource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Item               item                           `json:"item"`
		Region             ResourceRegionEnum             `json:"region"`
		AvailabilityDomain ResourceAvailabilityDomainEnum `json:"availabilityDomain"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	nn, e := model.Item.UnmarshalPolymorphicJSON(model.Item.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Item = nn.(Item)
	} else {
		m.Item = nil
	}
	m.Region = model.Region
	m.AvailabilityDomain = model.AvailabilityDomain
	return
}

// ResourceRegionEnum Enum with underlying type: string
type ResourceRegionEnum string

// Set of constants representing the allowable values for ResourceRegionEnum
const (
	ResourceRegionDev          ResourceRegionEnum = "DEV"
	ResourceRegionSea          ResourceRegionEnum = "SEA"
	ResourceRegionIntegNext    ResourceRegionEnum = "INTEG_NEXT"
	ResourceRegionIntegStable  ResourceRegionEnum = "INTEG_STABLE"
	ResourceRegionPhx          ResourceRegionEnum = "PHX"
	ResourceRegionIad          ResourceRegionEnum = "IAD"
	ResourceRegionEuFrankfurt1 ResourceRegionEnum = "EU_FRANKFURT_1"
	ResourceRegionUkLondon1    ResourceRegionEnum = "UK_LONDON_1"
	ResourceRegionCaToronto1   ResourceRegionEnum = "CA_TORONTO_1"
	ResourceRegionApTokyo1     ResourceRegionEnum = "AP_TOKYO_1"
	ResourceRegionUsLangley1   ResourceRegionEnum = "US_LANGLEY_1"
	ResourceRegionUsLuke1      ResourceRegionEnum = "US_LUKE_1"
)

var mappingResourceRegion = map[string]ResourceRegionEnum{
	"DEV":            ResourceRegionDev,
	"SEA":            ResourceRegionSea,
	"INTEG_NEXT":     ResourceRegionIntegNext,
	"INTEG_STABLE":   ResourceRegionIntegStable,
	"PHX":            ResourceRegionPhx,
	"IAD":            ResourceRegionIad,
	"EU_FRANKFURT_1": ResourceRegionEuFrankfurt1,
	"UK_LONDON_1":    ResourceRegionUkLondon1,
	"CA_TORONTO_1":   ResourceRegionCaToronto1,
	"AP_TOKYO_1":     ResourceRegionApTokyo1,
	"US_LANGLEY_1":   ResourceRegionUsLangley1,
	"US_LUKE_1":      ResourceRegionUsLuke1,
}

// GetResourceRegionEnumValues Enumerates the set of values for ResourceRegionEnum
func GetResourceRegionEnumValues() []ResourceRegionEnum {
	values := make([]ResourceRegionEnum, 0)
	for _, v := range mappingResourceRegion {
		values = append(values, v)
	}
	return values
}

// ResourceAvailabilityDomainEnum Enum with underlying type: string
type ResourceAvailabilityDomainEnum string

// Set of constants representing the allowable values for ResourceAvailabilityDomainEnum
const (
	ResourceAvailabilityDomainDev1             ResourceAvailabilityDomainEnum = "DEV_1"
	ResourceAvailabilityDomainDev2             ResourceAvailabilityDomainEnum = "DEV_2"
	ResourceAvailabilityDomainDev3             ResourceAvailabilityDomainEnum = "DEV_3"
	ResourceAvailabilityDomainIntegNext1       ResourceAvailabilityDomainEnum = "INTEG_NEXT_1"
	ResourceAvailabilityDomainIntegStable1     ResourceAvailabilityDomainEnum = "INTEG_STABLE_1"
	ResourceAvailabilityDomainSeaAd1           ResourceAvailabilityDomainEnum = "SEA_AD_1"
	ResourceAvailabilityDomainSeaAd2           ResourceAvailabilityDomainEnum = "SEA_AD_2"
	ResourceAvailabilityDomainSeaAd3           ResourceAvailabilityDomainEnum = "SEA_AD_3"
	ResourceAvailabilityDomainPhxAd1           ResourceAvailabilityDomainEnum = "PHX_AD_1"
	ResourceAvailabilityDomainPhxAd2           ResourceAvailabilityDomainEnum = "PHX_AD_2"
	ResourceAvailabilityDomainPhxAd3           ResourceAvailabilityDomainEnum = "PHX_AD_3"
	ResourceAvailabilityDomainIadAd1           ResourceAvailabilityDomainEnum = "IAD_AD_1"
	ResourceAvailabilityDomainIadAd2           ResourceAvailabilityDomainEnum = "IAD_AD_2"
	ResourceAvailabilityDomainIadAd3           ResourceAvailabilityDomainEnum = "IAD_AD_3"
	ResourceAvailabilityDomainIadPop1          ResourceAvailabilityDomainEnum = "IAD_POP_1"
	ResourceAvailabilityDomainIadPop2          ResourceAvailabilityDomainEnum = "IAD_POP_2"
	ResourceAvailabilityDomainEuFrankfurt1Ad1  ResourceAvailabilityDomainEnum = "EU_FRANKFURT_1_AD_1"
	ResourceAvailabilityDomainEuFrankfurt1Ad2  ResourceAvailabilityDomainEnum = "EU_FRANKFURT_1_AD_2"
	ResourceAvailabilityDomainEuFrankfurt1Ad3  ResourceAvailabilityDomainEnum = "EU_FRANKFURT_1_AD_3"
	ResourceAvailabilityDomainEuFrankfurt1Pop1 ResourceAvailabilityDomainEnum = "EU_FRANKFURT_1_POP_1"
	ResourceAvailabilityDomainUkLondon1Ad1     ResourceAvailabilityDomainEnum = "UK_LONDON_1_AD_1"
	ResourceAvailabilityDomainUkLondon1Ad2     ResourceAvailabilityDomainEnum = "UK_LONDON_1_AD_2"
	ResourceAvailabilityDomainUkLondon1Ad3     ResourceAvailabilityDomainEnum = "UK_LONDON_1_AD_3"
	ResourceAvailabilityDomainUkLondon1Pop1    ResourceAvailabilityDomainEnum = "UK_LONDON_1_POP_1"
	ResourceAvailabilityDomainUkLondon1Pop2    ResourceAvailabilityDomainEnum = "UK_LONDON_1_POP_2"
	ResourceAvailabilityDomainCaToronto1Ad1    ResourceAvailabilityDomainEnum = "CA_TORONTO_1_AD_1"
	ResourceAvailabilityDomainCaToronto1Pop1   ResourceAvailabilityDomainEnum = "CA_TORONTO_1_POP_1"
	ResourceAvailabilityDomainApTokyo1Ad1      ResourceAvailabilityDomainEnum = "AP_TOKYO_1_AD_1"
	ResourceAvailabilityDomainApTokyo1Pop1     ResourceAvailabilityDomainEnum = "AP_TOKYO_1_POP_1"
	ResourceAvailabilityDomainUsLuke1Ad1       ResourceAvailabilityDomainEnum = "US_LUKE_1_AD_1"
	ResourceAvailabilityDomainUsLuke1Pop1      ResourceAvailabilityDomainEnum = "US_LUKE_1_POP_1"
	ResourceAvailabilityDomainUsLangley1Ad1    ResourceAvailabilityDomainEnum = "US_LANGLEY_1_AD_1"
	ResourceAvailabilityDomainUsLangley1Pop1   ResourceAvailabilityDomainEnum = "US_LANGLEY_1_POP_1"
	ResourceAvailabilityDomainNoAd             ResourceAvailabilityDomainEnum = "NO_AD"
)

var mappingResourceAvailabilityDomain = map[string]ResourceAvailabilityDomainEnum{
	"DEV_1":                ResourceAvailabilityDomainDev1,
	"DEV_2":                ResourceAvailabilityDomainDev2,
	"DEV_3":                ResourceAvailabilityDomainDev3,
	"INTEG_NEXT_1":         ResourceAvailabilityDomainIntegNext1,
	"INTEG_STABLE_1":       ResourceAvailabilityDomainIntegStable1,
	"SEA_AD_1":             ResourceAvailabilityDomainSeaAd1,
	"SEA_AD_2":             ResourceAvailabilityDomainSeaAd2,
	"SEA_AD_3":             ResourceAvailabilityDomainSeaAd3,
	"PHX_AD_1":             ResourceAvailabilityDomainPhxAd1,
	"PHX_AD_2":             ResourceAvailabilityDomainPhxAd2,
	"PHX_AD_3":             ResourceAvailabilityDomainPhxAd3,
	"IAD_AD_1":             ResourceAvailabilityDomainIadAd1,
	"IAD_AD_2":             ResourceAvailabilityDomainIadAd2,
	"IAD_AD_3":             ResourceAvailabilityDomainIadAd3,
	"IAD_POP_1":            ResourceAvailabilityDomainIadPop1,
	"IAD_POP_2":            ResourceAvailabilityDomainIadPop2,
	"EU_FRANKFURT_1_AD_1":  ResourceAvailabilityDomainEuFrankfurt1Ad1,
	"EU_FRANKFURT_1_AD_2":  ResourceAvailabilityDomainEuFrankfurt1Ad2,
	"EU_FRANKFURT_1_AD_3":  ResourceAvailabilityDomainEuFrankfurt1Ad3,
	"EU_FRANKFURT_1_POP_1": ResourceAvailabilityDomainEuFrankfurt1Pop1,
	"UK_LONDON_1_AD_1":     ResourceAvailabilityDomainUkLondon1Ad1,
	"UK_LONDON_1_AD_2":     ResourceAvailabilityDomainUkLondon1Ad2,
	"UK_LONDON_1_AD_3":     ResourceAvailabilityDomainUkLondon1Ad3,
	"UK_LONDON_1_POP_1":    ResourceAvailabilityDomainUkLondon1Pop1,
	"UK_LONDON_1_POP_2":    ResourceAvailabilityDomainUkLondon1Pop2,
	"CA_TORONTO_1_AD_1":    ResourceAvailabilityDomainCaToronto1Ad1,
	"CA_TORONTO_1_POP_1":   ResourceAvailabilityDomainCaToronto1Pop1,
	"AP_TOKYO_1_AD_1":      ResourceAvailabilityDomainApTokyo1Ad1,
	"AP_TOKYO_1_POP_1":     ResourceAvailabilityDomainApTokyo1Pop1,
	"US_LUKE_1_AD_1":       ResourceAvailabilityDomainUsLuke1Ad1,
	"US_LUKE_1_POP_1":      ResourceAvailabilityDomainUsLuke1Pop1,
	"US_LANGLEY_1_AD_1":    ResourceAvailabilityDomainUsLangley1Ad1,
	"US_LANGLEY_1_POP_1":   ResourceAvailabilityDomainUsLangley1Pop1,
	"NO_AD":                ResourceAvailabilityDomainNoAd,
}

// GetResourceAvailabilityDomainEnumValues Enumerates the set of values for ResourceAvailabilityDomainEnum
func GetResourceAvailabilityDomainEnumValues() []ResourceAvailabilityDomainEnum {
	values := make([]ResourceAvailabilityDomainEnum, 0)
	for _, v := range mappingResourceAvailabilityDomain {
		values = append(values, v)
	}
	return values
}
