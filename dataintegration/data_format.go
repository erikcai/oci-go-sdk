// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DataFormat The data format object.
type DataFormat struct {
	FormatAttribute AbstractFormatAttribute `mandatory:"false" json:"formatAttribute"`

	// type
	Type DataFormatTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m DataFormat) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DataFormat) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FormatAttribute abstractformatattribute `json:"formatAttribute"`
		Type            DataFormatTypeEnum      `json:"type"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.FormatAttribute.UnmarshalPolymorphicJSON(model.FormatAttribute.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FormatAttribute = nn.(AbstractFormatAttribute)
	} else {
		m.FormatAttribute = nil
	}

	m.Type = model.Type
	return
}

// DataFormatTypeEnum Enum with underlying type: string
type DataFormatTypeEnum string

// Set of constants representing the allowable values for DataFormatTypeEnum
const (
	DataFormatTypeXml     DataFormatTypeEnum = "XML"
	DataFormatTypeJson    DataFormatTypeEnum = "JSON"
	DataFormatTypeCsv     DataFormatTypeEnum = "CSV"
	DataFormatTypeAvro    DataFormatTypeEnum = "AVRO"
	DataFormatTypeOrc     DataFormatTypeEnum = "ORC"
	DataFormatTypeParquet DataFormatTypeEnum = "PARQUET"
)

var mappingDataFormatType = map[string]DataFormatTypeEnum{
	"XML":     DataFormatTypeXml,
	"JSON":    DataFormatTypeJson,
	"CSV":     DataFormatTypeCsv,
	"AVRO":    DataFormatTypeAvro,
	"ORC":     DataFormatTypeOrc,
	"PARQUET": DataFormatTypeParquet,
}

// GetDataFormatTypeEnumValues Enumerates the set of values for DataFormatTypeEnum
func GetDataFormatTypeEnumValues() []DataFormatTypeEnum {
	values := make([]DataFormatTypeEnum, 0)
	for _, v := range mappingDataFormatType {
		values = append(values, v)
	}
	return values
}
