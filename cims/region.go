// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // RegionEnum Enum with underlying type: string
type RegionEnum string

// Set of constants representing the allowable values for RegionEnum
const (
    RegionDev RegionEnum = "DEV"
    RegionSea RegionEnum = "SEA"
    RegionIntegNext RegionEnum = "INTEG_NEXT"
    RegionIntegStable RegionEnum = "INTEG_STABLE"
    RegionPhx RegionEnum = "PHX"
    RegionIad RegionEnum = "IAD"
    RegionFra RegionEnum = "FRA"
    RegionEuFrankfurt1 RegionEnum = "EU_FRANKFURT_1"
    RegionLhr RegionEnum = "LHR"
    RegionYyz RegionEnum = "YYZ"
    RegionNrt RegionEnum = "NRT"
    RegionUsLangley1 RegionEnum = "US_LANGLEY_1"
    RegionUsLuke1 RegionEnum = "US_LUKE_1"
    RegionIcn RegionEnum = "ICN"
    RegionBom RegionEnum = "BOM"
    RegionGru RegionEnum = "GRU"
    RegionUsGovAshburn1 RegionEnum = "US_GOV_ASHBURN_1"
    RegionUsGovPhoenix1 RegionEnum = "US_GOV_PHOENIX_1"
    RegionUsGovChicago1 RegionEnum = "US_GOV_CHICAGO_1"
    RegionSyd RegionEnum = "SYD"
    RegionZrh RegionEnum = "ZRH"
    RegionJed RegionEnum = "JED"
    RegionAms RegionEnum = "AMS"
    RegionKix RegionEnum = "KIX"
    RegionMel RegionEnum = "MEL"
    RegionYul RegionEnum = "YUL"
    RegionHyd RegionEnum = "HYD"
    RegionYny RegionEnum = "YNY"
    RegionTiw RegionEnum = "TIW"
)

var mappingRegion = map[string]RegionEnum { 
    "DEV": RegionDev,
    "SEA": RegionSea,
    "INTEG_NEXT": RegionIntegNext,
    "INTEG_STABLE": RegionIntegStable,
    "PHX": RegionPhx,
    "IAD": RegionIad,
    "FRA": RegionFra,
    "EU_FRANKFURT_1": RegionEuFrankfurt1,
    "LHR": RegionLhr,
    "YYZ": RegionYyz,
    "NRT": RegionNrt,
    "US_LANGLEY_1": RegionUsLangley1,
    "US_LUKE_1": RegionUsLuke1,
    "ICN": RegionIcn,
    "BOM": RegionBom,
    "GRU": RegionGru,
    "US_GOV_ASHBURN_1": RegionUsGovAshburn1,
    "US_GOV_PHOENIX_1": RegionUsGovPhoenix1,
    "US_GOV_CHICAGO_1": RegionUsGovChicago1,
    "SYD": RegionSyd,
    "ZRH": RegionZrh,
    "JED": RegionJed,
    "AMS": RegionAms,
    "KIX": RegionKix,
    "MEL": RegionMel,
    "YUL": RegionYul,
    "HYD": RegionHyd,
    "YNY": RegionYny,
    "TIW": RegionTiw,
}

// GetRegionEnumValues Enumerates the set of values for RegionEnum
func GetRegionEnumValues() []RegionEnum {
   values := make([]RegionEnum, 0)
   for _, v := range mappingRegion {
       values = append(values, v)
   }
   return values
}


