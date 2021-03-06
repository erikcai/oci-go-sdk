// Copyright (c) 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Package autotest contains all auto generated integration test cases
package autotest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/erikcai/oci-go-sdk/v33/common"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// DataToValidate defines the data that needs to be sent back to oci testing service for validation in a successful scenario.
type DataToValidate struct {
	// The id of container that generates the request.
	ContainerID string `json:"containerId"`

	// The type of request object. For example: "com.oracle.bmc.core.requests.CreatePublicIpRequest"
	RequestClass string `json:"requestClass"`

	// The request object in json string.
	RequestJSON string `json:"requestJson"`

	// The type of response object.
	ResponseClass string `json:"responseClass"`

	// The response object in json string.
	ResponseJSON string `json:"responseJson"`

	// IsListResponse if the response is a list of responses
	IsListResponse bool `json:"listResponse"`
}

func (d DataToValidate) String() string {
	buffer := bytes.Buffer{}
	json.Indent(&buffer, []byte(d.RequestJSON), "", "  ")
	prettyReq := buffer.String()
	json.Indent(&buffer, []byte(d.ResponseJSON), "", "  ")
	prettyRes := buffer.String()

	return fmt.Sprintf("<ContainerID: %s\nRequestClass: %s\nRequestJSON: %s\nResponseClass: %s\nResponseJSON:%s>",
		d.ContainerID, d.RequestClass, prettyReq, d.ResponseClass, prettyRes)
}

// ErrorToValidate  struct defines the data that needs to be sent back to oci testing service for validation in  failure scenario.
type ErrorToValidate struct {
	// The id of container that generates the request.
	ContainerID string `json:"containerId"`

	// The type of request object. For example: "com.oracle.bmc.core.requests.CreatePublicIpRequest"
	RequestClass string `json:"requestClass"`

	// The request object in json string.
	RequestJSON string `json:"requestJson"`

	// The error object in json string. For example:
	// "{\"statusCode\":400,\"code\":\"InvalidParameter\",\"message\":\"compartmentId size must be between 1 and 255\"}"
	ErrorJSON string `json:"errorJson"`
}

func (d ErrorToValidate) String() string {
	buffer := bytes.Buffer{}
	json.Indent(&buffer, []byte(d.RequestJSON), "", "  ")
	prettyReq := buffer.String()
	json.Indent(&buffer, []byte(d.ErrorJSON), "", "  ")
	prettyErr := buffer.String()

	return fmt.Sprintf("<ContainerID: %s\nRequestClass: %s\nRequestJSON: %s\nErrorJSON:%s>",
		d.ContainerID, d.RequestClass, prettyReq, prettyErr)
}

// The struct defines the data for client information.
type ClientInfo struct {
	Language string `json:"language"`
}

const defaultOCITestingServiceEndpoint = "http://localhost:8090/SDKTestingService/"
const requestClassTemplate = "com.oracle.bmc.%s.requests.%s"
const responseClassTemplate = "com.oracle.bmc.%s.responses.%s"
const mockMode = "mock"

var firstPageOnly = false
var withRetry = true

func failIfError(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
		t.FailNow()
	}
}

//OCITestClient a client for the oci-testing-service
type OCITestClient struct {
	HTTPClient       http.Client
	ServiceEndpoint  string
	SessionID        string
	Log              *log.Logger
	IgnoreHttpBodies bool
}

//NewOCITestClient creates a client for the oci-testing-service
func NewOCITestClient() *OCITestClient {
	var testLog = log.New(os.Stderr, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	if _, isDebugEnabled := os.LookupEnv("OCI_GO_AUTOTEST_DEBUG"); !isDebugEnabled {
		testLog.SetOutput(ioutil.Discard)
	}

	testingServiceEndpoint := defaultOCITestingServiceEndpoint
	if endpoint, ok := os.LookupEnv("OCI_GO_TESTING_SERVICE_ENDPOINT"); ok {
		testingServiceEndpoint = endpoint
	}

	return &OCITestClient{
		HTTPClient:      http.Client{},
		ServiceEndpoint: testingServiceEndpoint,
		Log:             testLog,
	}
}

func (client OCITestClient) createClientForOperation(serviceName, clientName, operation string,
	constructor func(common.ConfigurationProvider, TestingConfig) (interface{}, error)) (interface{}, error) {
	confProvider, testConfig, err := client.configurationProvider(serviceName, clientName, operation)
	if err != nil {
		return nil, err
	}

	return constructor(confProvider, testConfig)
}

func (client OCITestClient) getEndpointForService(serviceName, clientName, operationName string) (string, error) {
	endpoint := pathJoin(client.ServiceEndpoint, "endpoint")
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	q := url.Values{}
	q.Add("sessionId", client.SessionID)
	q.Add("serviceName", serviceName)
	q.Add("clientName", clientName+"Client")
	q.Add("apiName", operationName)
	q.Add("lang", "Go")
	request.URL.RawQuery = q.Encode()

	body, err := client.callService(request)
	if err != nil {
		return "", err
	}

	return string(body), err
}

type TestingConfig struct {
	Region         string `json:"region"`
	TenantID       string `json:"tenantId"`
	CompartmentID  string `json:"compartmentId"`
	UserID         string `json:"userId"`
	Endpoint       string `json:"endpoint"`
	Fingerprint    string `json:"fingerprint"`
	PassPhrase     string `json:"passPhrase"`
	KeyFile        string `json:"keyFile"`
	KeyFileContent string `json:"keyFileContent"`
	TestMode       string `json:"testMode"`
}

func (client OCITestClient) getConfiguration(serviceName, clientName, operationName string) (config TestingConfig, err error) {
	endpoint := pathJoin(client.ServiceEndpoint, "config")
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}
	q := url.Values{}
	q.Add("serviceName", serviceName)
	q.Add("clientName", clientName)
	q.Add("apiName", operationName)
	q.Add("lang", "Go")
	request.URL.RawQuery = q.Encode()

	body, err := client.callService(request)
	if err != nil {
		return
	}

	config = TestingConfig{}
	err = json.Unmarshal(body, &config)

	testEndpoint, err := client.getEndpointForService(serviceName, clientName, operationName)
	if err != nil {
		return
	}
	config.Endpoint = testEndpoint

	// Mock mode
	if config.TestMode == mockMode {
		firstPageOnly = true
		withRetry = false
	}

	client.Log.Printf("Server configuration acquired: %#v\n", config)
	return
}

func (client OCITestClient) configurationProvider(serviceName, clientName, operationName string) (provider common.ConfigurationProvider, testConfig TestingConfig, err error) {
	testConfig, err = client.getConfiguration(serviceName, clientName, operationName)

	if err != nil {
		return
	}

	provider = common.NewRawConfigurationProvider(
		testConfig.TenantID,
		testConfig.UserID,
		testConfig.Region,
		testConfig.Fingerprint,
		testConfig.KeyFileContent,
		&testConfig.PassPhrase)

	var ok bool
	if ok, err = common.IsConfigurationProviderValid(provider); !ok {
		err = fmt.Errorf("server configuration is not valid %s", err.Error())
		return
	}

	return
}

func (client OCITestClient) isApiEnabled(serviceName, operationName string) (bool, error) {
	endpoint := pathJoin(client.ServiceEndpoint, "request", "enable")
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return false, err
	}
	q := url.Values{}
	q.Add("serviceName", serviceName)
	q.Add("apiName", operationName)
	q.Add("lang", "Go")
	request.URL.RawQuery = q.Encode()

	body, err := client.callService(request)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(body))
}

// validateResult the result of SDK API call.
func (client OCITestClient) validateResult(containerId string, request interface{}, response interface{}, error error) (string, error) {
	if error != nil {
		return client.validateError(containerId, request, error)
	}
	return client.validateResponse(containerId, request, response)
}

//clientSideError represents an error in the client side of the sdk
type clientSideError struct {
	Message string `json:"message"`
}

func (ce clientSideError) Error() string {
	return ce.Message
}

// validateError the result of SDK API call for failure case
func (client OCITestClient) validateError(containerId string, req interface{}, responseError error) (string, error) {
	client.Log.Println("Validating error")
	if _, ok := common.IsServiceError(responseError); !ok {
		responseError = clientSideError{"client side error: " + responseError.Error()}
	}

	data := ErrorToValidate{ContainerID: containerId}

	var err error
	reqType := reflect.TypeOf(req)
	data.RequestClass = fmt.Sprintf(requestClassTemplate, path.Base(reqType.PkgPath()), reqType.Name())
	data.RequestJSON, err = marshal(reflect.ValueOf(req), client.Log)
	if err != nil {
		return "", err
	}

	data.ErrorJSON, err = marshal(reflect.ValueOf(responseError), client.Log)
	if err != nil {
		return "", err
	}

	body, _ := json.Marshal(data)

	client.Log.Println("ErrorToValidate:" + prettyJson(body))

	endpoint := pathJoin(client.ServiceEndpoint, fmt.Sprintf("error?sessionId=%s&lang=Go", client.SessionID))
	request, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))

	if err != nil {
		return "", err
	}

	body, err = client.callService(request)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Validate the result of SDK API call for happy case
func (client OCITestClient) validateResponse(containerId string, req interface{}, res interface{}) (string, error) {
	client.Log.Println("Validating response")
	data := DataToValidate{ContainerID: containerId}
	var err error
	var resType reflect.Type

	data.IsListResponse, data.ResponseClass, data.ResponseJSON, resType, err = client.marshalResponse(res)
	if err != nil {
		return "", err
	}

	if req != nil {
		reqType := reflect.TypeOf(req)
		data.RequestClass = fmt.Sprintf(requestClassTemplate, path.Base(reqType.PkgPath()), reqType.Name())
		data.RequestJSON, err = marshalStruct(reflect.ValueOf(req), client.Log)
		if err != nil {
			return "", err
		}
	} else {
		pkgName := path.Base(resType.PkgPath())
		reqTypeName := strings.Replace(resType.Name(), "Response", "Request", 1)
		data.RequestClass = fmt.Sprintf(requestClassTemplate, pkgName, reqTypeName)
		data.RequestJSON = "{}"
	}

	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return client.callValidate(body)
}

func (client *OCITestClient) marshalResponse(res interface{}) (isListResponse bool, responseClass, responseJSON string,
	resType reflect.Type, err error) {
	resType = reflect.TypeOf(res)
	isListResponse = false
	responseClass = fmt.Sprintf(responseClassTemplate, path.Base(resType.PkgPath()), resType.Name())
	if resType.Kind() == reflect.Array || resType.Kind() == reflect.Slice {
		isListResponse = true
		resType = reflect.TypeOf(res).Elem()
		responseClass = fmt.Sprintf(responseClassTemplate, path.Base(resType.PkgPath()), resType.Name())
	}
	responseJSON, err = marshal(reflect.ValueOf(res), client.Log)
	return
}

func (client *OCITestClient) callValidate(body []byte) (string, error) {
	client.Log.Println("DataToValidate:", prettyJson(body))

	endpoint := pathJoin(client.ServiceEndpoint, fmt.Sprintf("response?sessionId=%s&lang=Go", client.SessionID))
	request, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))

	if err != nil {
		return "", err
	}

	responseBody, err := client.callService(request)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func (client *OCITestClient) startSession() error {
	clientInfo := ClientInfo{Language: "Go"} // Java and Go share the same serializer
	body, _ := json.Marshal(clientInfo)

	client.Log.Println("ClientInfo:" + string(body))

	endpoint := pathJoin(client.ServiceEndpoint, "sessions")
	request, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	responseBody, err := client.callService(request)
	if err != nil {
		return err
	}

	sessionId := string(responseBody)
	client.SessionID = sessionId
	client.Log.Println("SessionId:", client.SessionID)
	return nil
}

func (client OCITestClient) endSession() error {
	if client.SessionID == "" {
		return fmt.Errorf("can not end session if none has been started")
	}

	client.Log.Println("Terminating session: ", client.SessionID)
	endpoint := pathJoin(client.ServiceEndpoint, "sessions", client.SessionID)
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	_, err = client.callService(request)
	if err != nil {
		return err
	}

	return nil
}

func (client OCITestClient) getRequests(serviceName string, apiName string) (string, error) {
	endpoint := pathJoin(client.ServiceEndpoint, "request")
	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return "", err
	}

	query := request.URL.Query()
	query.Add("sessionId", client.SessionID)
	query.Add("serviceName", serviceName)
	query.Add("apiName", apiName)
	query.Add("lang", "Go")
	request.URL.RawQuery = query.Encode()

	body, err := client.callService(request)
	if err != nil {
		return "", err
	}

	client.Log.Println("Server sent requests:" + prettyJson(body))
	return string(body), nil
}

func (client OCITestClient) generateListResponses(request common.OCIRequest,
	listFn func(common.OCIRequest) (common.OCIResponse, error)) (listResponses []common.OCIResponse, err error) {

	listResponses = make([]common.OCIResponse, 0)
	firstListResponse, err := listFn(request)
	if err != nil {
		return
	}
	listResponses = append(listResponses, firstListResponse)

	nextPageToken, err := getFieldValue(firstListResponse, "OpcNextPage")
	if err != nil {
		typeName := reflect.TypeOf(firstListResponse).Name()
		client.Log.Printf("%s does not have OpcNextPage, eventhoug is a List API ", typeName)
		// do not fail there are some apis that do not have OpcNextPage
		err = nil
		return
	}

	//No more pages return
	if isFieldEmpty(reflect.ValueOf(nextPageToken)) {
		return
	}

	err = setFieldValue(request, "Page", nextPageToken)
	if err != nil {
		return
	}

	if firstPageOnly {
		return
	}

	nextResponse, err := listFn(request)
	if err != nil {
		return
	}
	listResponses = append(listResponses, nextResponse)

	prevToken, err := getFieldValue(nextResponse, "OpcPrevPage")
	if err != nil {
		//clear the error since some request might not have a PrevPage
		err = nil
		return
	}

	err = setFieldValue(request, "Page", prevToken)
	if err != nil {
		return
	}

	prevResponse, err := listFn(request)
	if err != nil {
		return
	}

	listResponses = append(listResponses, prevResponse)
	return
}

func (client OCITestClient) callService(request *http.Request) (body []byte, err error) {

	requestLog, _ := httputil.DumpRequestOut(request, !client.IgnoreHttpBodies)
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		client.Log.Printf("===HttpWiredata\nRequest %s\n===EndWireData", requestLog)
		return nil, err
	}

	defer response.Body.Close()
	responseLog, _ := httputil.DumpResponse(response, !client.IgnoreHttpBodies)

	httpWireData := fmt.Sprintf("====HttpWiredata\nRequest: %s\nResponse: %s\nEndWireData=====", requestLog, responseLog)
	client.Log.Println(httpWireData)

	err = checkHttpResponse(response)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(response.Body)
	return

}

func checkHttpResponse(response *http.Response) error {
	if response.StatusCode < 200 || response.StatusCode > 299 {
		bodyContent := new(bytes.Buffer)
		io.Copy(bodyContent, response.Body)
		return fmt.Errorf("response is not successful status code: %d\n%v", response.StatusCode, bodyContent.String())
	}
	return nil
}

func marshal(val reflect.Value, logger *log.Logger) (string, error) {
	if !val.IsValid() {
		return "", fmt.Errorf("marshaling value %#v is not supported", val)
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		return marshalSlice(val, logger)
	case reflect.Struct:
		return marshalStruct(val, logger)
	case reflect.String:
		strval, e := json.Marshal(val.Interface())
		return string(strval), e
	default:
		return "", fmt.Errorf("marshaling of value %#v is not supported", val)
	}
}
func marshalSlice(val reflect.Value, logger *log.Logger) (string, error) {
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		return "", fmt.Errorf("can not marshal a not slice as a slice")
	}

	allFieldsMarshaled := make([]string, 0)
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i)
		m, e := marshal(v, logger)
		if e != nil {
			return "", e
		}
		allFieldsMarshaled = append(allFieldsMarshaled, m)
	}

	allJson := strings.Join(allFieldsMarshaled, ",")
	return `[` + allJson + `]`, nil

}

// marshalStruct A customized marshalStruct method for request and response objects. It
// marshals each of the fields in the struct
func marshalStruct(val reflect.Value, logger *log.Logger) (string, error) {
	valueType := val.Type()
	allFieldsMarshaled := make([]string, 0)

	for i := 0; i < valueType.NumField(); i++ {
		sf := valueType.Field(i)

		// unexported
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		sv := val.Field(i)
		_, contributesToOk := sf.Tag.Lookup("contributesTo")
		_, presentInOk := sf.Tag.Lookup("presentIn")
		errorInterface := reflect.TypeOf((*error)(nil)).Elem()
		isError := valueType.Implements(errorInterface)

		if contributesToOk || presentInOk || isError {
			if !isFieldMandatory(sf) && isFieldEmpty(sv) {
				//	fmt.Println("ignoring: ", sf.Name, " is empty")
				continue
			}

			var bs []byte
			var e error
			valToMarshal := sv.Interface()
			fieldName := sf.Name
			if rc, ok := valToMarshal.(io.ReadCloser); ok {
				valToMarshal, e = ioutil.ReadAll(rc)
				if e != nil {
					return "", e
				}

				//https://jira.oci.oraclecorp.com/browse/DEX-5466
				if fieldName == "Content" {
					fieldName = "inputstream"
				}
			}

			bs, e = json.Marshal(valToMarshal)
			if e != nil {
				return "", e
			}

			//Cleans up json fields that the testing service might not be able to process
			adjustedJSON, e := removeNilFieldsInJSONWithTaggedStruct(bs, sv, logger)
			if e != nil {
				return "", e
			}
			allFieldsMarshaled = append(allFieldsMarshaled,
				fmt.Sprintf(`"%s":%s`, uncapitalize(fieldName), string(adjustedJSON)))
		}
	}

	allJson := strings.Join(allFieldsMarshaled, ",")
	rs := "{" + allJson + "}"
	return rs, nil
}

// removeNilFieldsInJSONWithTaggedStruct remove struct fields tagged with json and mandatory false
// that are nil
func removeNilFieldsInJSONWithTaggedStruct(rawJSON []byte, value reflect.Value, logger *log.Logger) ([]byte, error) {
	var rawInterface interface{}
	decoder := json.NewDecoder(bytes.NewBuffer(rawJSON))
	decoder.UseNumber()
	var err error
	if err = decoder.Decode(&rawInterface); err != nil {
		return nil, err
	}

	fixedMap, err := omitNilFieldsInJSON(rawInterface, value, logger)
	if err != nil {
		return nil, err
	}
	return json.Marshal(fixedMap)
}

var timeType = reflect.TypeOf(common.SDKTime{})
var timeTypePtr = reflect.TypeOf(&common.SDKTime{})

var sdkDateType = reflect.TypeOf(common.SDKDate{})
var sdkDateTypePtr = reflect.TypeOf(&common.SDKDate{})

// omitNilFieldsInJSON, removes json keys whose struct value is nil, and the field is tagged with the json and
// mandatory:false tags
func omitNilFieldsInJSON(data interface{}, value reflect.Value, logger *log.Logger) (interface{}, error) {
	switch value.Kind() {
	case reflect.Struct:
		if value.Type() == timeType || value.Type() == timeTypePtr ||
			value.Type() == sdkDateType || value.Type() == sdkDateTypePtr {
			return data, nil
		}
		jsonMap, ok := data.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("can not read raw map from data for type: %s", value.Type().String())
		}

		fieldType := value.Type()
		for i := 0; i < fieldType.NumField(); i++ {
			currentField := fieldType.Field(i)
			//unexported skip
			if currentField.PkgPath != "" {
				continue
			}

			//Does not have json tag, no-op
			if _, ok := currentField.Tag.Lookup("json"); !ok {
				continue
			}

			currentFieldValue := value.Field(i)
			ok, jsonFieldName, err := getTaggedNilFieldNameOrError(currentField, currentFieldValue, logger)
			if err != nil {
				return nil, fmt.Errorf("can not omit nil fields for field: %s, due to: %s",
					currentField.Name, err.Error())
			}

			//Delete the struct field from the json representation
			if ok {
				delete(jsonMap, jsonFieldName)
				continue
			}

			// Check to make sure the field is part of the json representation of the value
			if _, contains := jsonMap[jsonFieldName]; !contains {
				logger.Printf("Field %s is not present in json, omitting", jsonFieldName)
				continue
			}

			if currentFieldValue.Type() == timeType || currentFieldValue.Type() == timeTypePtr ||
				currentField.Type == sdkDateType || currentField.Type == sdkDateTypePtr {
				continue
			}
			// does it need to be adjusted?
			var adjustedValue interface{}
			adjustedValue, err = omitNilFieldsInJSON(jsonMap[jsonFieldName], currentFieldValue, logger)
			if err != nil {
				return nil, fmt.Errorf("can not omit nil fields for field: %s, due to: %s",
					currentField.Name, err.Error())
			}

			jsonMap[jsonFieldName] = adjustedValue
		}
		return jsonMap, nil
	case reflect.Slice, reflect.Array:
		// Special case: a []byte may have been marshalled as a string
		if data != nil && reflect.TypeOf(data).Kind() == reflect.String && value.Type().Elem().Kind() == reflect.Uint8 {
			return data, nil
		}
		jsonList, ok := data.([]interface{})
		if !ok {
			logger.Printf("can not omit nil fields, data was expected to be a not-nil list")
			return data, nil
		}
		newList := make([]interface{}, len(jsonList))
		var err error
		for i, val := range jsonList {
			newList[i], err = omitNilFieldsInJSON(val, value.Index(i), logger)
			if err != nil {
				return nil, err
			}
		}
		return newList, nil
	case reflect.Map:
		jsonMap, ok := data.(map[string]interface{})
		if !ok {
			logger.Printf("can not omit nil fields, data was expected to be a not-nil map")
			return data, nil
		}
		newMap := make(map[string]interface{}, len(jsonMap))
		var err error
		for key, val := range jsonMap {
			newMap[key], err = omitNilFieldsInJSON(val, value.MapIndex(reflect.ValueOf(key)), logger)
			if err != nil {
				return nil, err
			}
		}
		return newMap, nil
	case reflect.Ptr:
		valPtr := value.Elem()
		return omitNilFieldsInJSON(data, valPtr, logger)
	case reflect.Interface:
		if _, ok := value.Interface().(io.ReadCloser); ok {
			return data, nil
		}
		valPtr := value.Elem()
		return omitNilFieldsInJSON(data, valPtr, logger)
	default:
		//Otherwise no-op
		return data, nil
	}
}

// getTaggedNilFieldNameOrError, evaluates if a field with json and  non mandatory tags is nil
// returns the json tag name, or an error if the tags are incorrectly present
func getTaggedNilFieldNameOrError(field reflect.StructField, fieldValue reflect.Value, logger *log.Logger) (bool, string, error) {
	currentTag := field.Tag
	jsonTag := currentTag.Get("json")

	if jsonTag == "" {
		return false, "", fmt.Errorf("json tag is not valid for field %s", field.Name)
	}

	partsJSONTag := strings.Split(jsonTag, ",")
	nameJSONField := partsJSONTag[0]

	if _, ok := currentTag.Lookup("mandatory"); !ok {
		//No mandatory field set, no-op
		return false, nameJSONField, nil
	}
	isMandatory, err := strconv.ParseBool(currentTag.Get("mandatory"))
	if err != nil {
		return false, "", fmt.Errorf("mandatory tag is not valid for field %s", field.Name)
	}

	// If the field is marked as mandatory, no-op
	if isMandatory {
		if strings.HasSuffix(field.Type.Name(), "Enum") && fieldValue.String() == "" {
			logger.Printf("Enum Field: %s of type %s will be omitted because its value is empty despite being mandatory", field.Name, field.Type.String())
			return true, nameJSONField, nil
		}
		return false, nameJSONField, nil
	}

	logger.Printf("Adjusting tag: mandatory is false and json tag is valid on field: %s", field.Name)

	// If the field can not be nil, then no-op
	if !isNillableType(&fieldValue) {
		logger.Printf("WARNING json field is tagged with mandatory flags, but the type can not be nil, field name: %s", field.Name)
		return false, nameJSONField, nil
	}

	// If field value is nil, tag it as omitEmpty
	return fieldValue.IsNil(), nameJSONField, nil

}

// isNillableType returns true if the filed can be nil
func isNillableType(value *reflect.Value) bool {
	k := value.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return true
	}
	return false
}

func isFieldMandatory(field reflect.StructField) bool {
	tagVal, ok := field.Tag.Lookup("mandatory")
	if !ok {
		return true
	}
	mandatory, err := strconv.ParseBool(tagVal)
	if err != nil {
		return false
	}
	return mandatory
}

func isFieldEmpty(val reflect.Value) bool {
	realVal := val
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		realVal = val.Elem()
	}
	switch realVal.Kind() {
	case reflect.String:
		return realVal.String() == ""
	case reflect.Slice, reflect.Array:
		return realVal.Len() == 0
	default:
		return false
	}
}

func getField(value reflect.Value, fieldName string) (reflect.Value, error) {
	valType := value.Type()
	if valType.Kind() == reflect.Ptr {
		return getField(value.Elem(), fieldName)
	}
	field := value.FieldByName(fieldName)
	if !field.IsValid() {
		return reflect.Value{}, fmt.Errorf("%s is not part of %s", fieldName, valType.Name())
	}
	return field, nil
}

func getFieldValue(s interface{}, fieldName string) (interface{}, error) {
	value := reflect.ValueOf(s)
	field, err := getField(value, fieldName)
	if err != nil {
		return nil, err
	}
	return field.Interface(), nil
}

func setFieldValue(s interface{}, fieldName string, fieldValue interface{}) error {
	valType := reflect.TypeOf(s)
	value := reflect.ValueOf(s)
	field, err := getField(value, fieldName)
	if err != nil {
		return err
	}

	if !field.CanSet() {
		return fmt.Errorf("%s can not bet set in %s", fieldName, valType.Name())
	}
	field.Set(reflect.ValueOf(fieldValue))
	return nil
}

func pathJoin(parts ...string) string {
	all := make([]string, 0)
	for _, part := range parts {
		p := part
		if strings.HasSuffix(part, "/") {
			p = part[0 : len(part)-1]
		}
		all = append(all, p)

	}
	urlString := strings.Join(all, "/")
	return urlString
}

// uncapitalize the given string
func uncapitalize(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}

	return ""
}

//PolymorphicRequestUnmarshallingInfo holds all the information to marshal a polymorphic request into a concrete type
type PolymorphicRequestUnmarshallingInfo struct {
	//The name of polymorphic discriminator
	DiscriminatorName string
	//A map of values of the discriminator field to non-zero interfaces that can be used to unmarshall json for the value of
	//the discriminator
	DiscriminatorValuesAndTypes map[string]interface{}
}

//shouldReplace returns true when a field has been replaced, false otherwise. If field has been replaced it returns the new value of the field as a second paramater
// otherwise nil. The value is obtained by marshalling the current requestFieldValue into a type specified by the polymorphicUnmarshallingInfo struct.
// The PolymorphicRequestUnmarshallingInfo struct contains a values of the discriminator fields to a concrete type describing such data.
// fieldNames is an array due to testing service json objects not always match the field type, usually they either match the field name or the field type name
func shouldReplace(polymorhpicUnmarshallingInfo map[string]PolymorphicRequestUnmarshallingInfo, fieldNames []string, requestFieldValue interface{}, logger *log.Logger) (bool, interface{}, error) {
	var polymorphicInfo PolymorphicRequestUnmarshallingInfo
	var ok bool

	//Sigh transform polymorphicUnmarshallingInfo to map[string]interface{}
	genericUnmarshalInfo := make(map[string]interface{})
	for k, v := range polymorhpicUnmarshallingInfo {
		genericUnmarshalInfo[k] = v
	}

	//Is the current field a polymorhpic field
	if polInfo := findValWithKeys(genericUnmarshalInfo, fieldNames...); polInfo == nil {
		logger.Println("polymorphic field not found for fields name:", fieldNames)
		return false, nil, nil
	} else {
		polymorphicInfo = polInfo.(PolymorphicRequestUnmarshallingInfo)
	}

	//the field value present in the request has been instially unmarshalled to a raw map
	rawMap, ok := requestFieldValue.(map[string]interface{})
	if !ok {
		err := fmt.Errorf("something went wrong, map as expected")
		return false, nil, err
	}

	discriminatorName := polymorphicInfo.DiscriminatorName
	discriminatorValue, ok := rawMap[discriminatorName].(string)
	if !ok {
		err := fmt.Errorf("found unspecified polymorphic type %s", requestFieldValue)
		return false, nil, err
	}

	//find the concrete type for the discriminator value
	concreteType, ok := polymorphicInfo.DiscriminatorValuesAndTypes[discriminatorValue]
	logger.Println("polymorphic field of type ", reflect.TypeOf(concreteType).String())
	if !ok {
		err := fmt.Errorf("did not find discrimantor value :%s in the set of possible values for field %s", discriminatorValue, fieldNames)
		return false, nil, err
	}

	//marshal the map back to json, so we can later unmarshal to the proper data structure
	data, err := json.Marshal(rawMap)
	if err != nil {
		return false, nil, err
	}

	//create a new pointer of the type stored in our polymorphic information
	//the objects in PolymorphicRequestUnmarshallInfo are pointers, thus we need
	//a new pointer to the underlying type
	newConcreteValue := reflect.New(reflect.TypeOf(concreteType).Elem())
	//finally unmarshal the json value to the right type
	err = json.Unmarshal(data, newConcreteValue.Interface())
	if err != nil {
		return false, nil, err
	}
	return true, newConcreteValue.Interface(), nil

}

func conditionalStructCopy(srcData interface{}, dst interface{}, polymorphicRequestInfo map[string]PolymorphicRequestUnmarshallingInfo, logger *log.Logger) (err error) {
	return recursiveStructCopy(srcData, reflect.ValueOf(dst), polymorphicRequestInfo, logger)
}

// recursiveStructCopy copies src to destination. It copies keys,values in src  that match fieldNames in dst. This is done for all fields
// expect for the fields names that match keys in polymorphicRequestInfo map
// for the fields that match a key in polymorphicRequestInfo map, the value of the key is used to extract a type. This type is then used to
// to json marshall the matching contents of src. This is done to all the fields in the dstValue struct.
// The dstValue parameter needs to be settable
func recursiveStructCopy(srcData interface{}, dstValue reflect.Value, polymorphicRequestInfo map[string]PolymorphicRequestUnmarshallingInfo, logger *log.Logger) (err error) {
	if dstValue.Kind() != reflect.Ptr {
		err = fmt.Errorf("dstValue need to be a pointer to a struct")
		return
	}

	val := dstValue.Elem()
	dstType := val.Type()
	if dstType.Kind() != reflect.Struct {
		err = fmt.Errorf("dstValue need to be a pointer to struct. is it value?")
		return
	}

	mapData, ok := srcData.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("src needs to be a map")
		return
	}

	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		dstFieldVal := val.Field(i)
		fieldName := dstField.Name

		srcFieldVal := findValWithKeys(mapData, fieldName, uncapitalize(fieldName))
		if srcFieldVal == nil {
			logger.Printf("can not find value for field: %s. Skipping", fieldName)
			continue
		}

		if dstField.Type.Kind() == reflect.Struct {
			recursiveStructCopy(srcFieldVal, dstFieldVal.Addr(), polymorphicRequestInfo, logger)
		} else {
			replacementVal := reflect.ValueOf(srcFieldVal)
			action := "setting"
			if ok, replacement, err := shouldReplace(polymorphicRequestInfo, []string{dstField.Name, dstField.Type.Name()}, srcFieldVal, logger); ok {
				action = "replacing"
				replacementVal = reflect.ValueOf(replacement)
			} else {
				if err != nil {
					return err
				}
			}

			logger.Printf("%s field name: %s, of %s --> val: %s", action, fieldName, dstFieldVal.String(), replacementVal.String())
			switch dstFieldVal.Type().Kind() {
			case reflect.Ptr:
				newVal := getNumberFieldValueOrSame(dstFieldVal.Type().Elem(), replacementVal, logger)
				//create a new pointer
				dstFieldVal.Set(reflect.New(dstFieldVal.Type().Elem()))
				//set the value of the pointer
				dstFieldVal.Elem().Set(newVal)
				break
			case reflect.Interface:
				logger.Printf("Setting interface of type: %s to type: %s", dstFieldVal.Type(), replacementVal.Elem().Type())
				dstFieldVal.Set(reflect.New(replacementVal.Elem().Type()))
				dstFieldVal.Set(replacementVal.Elem())
				break
			default:
				dstFieldVal.Set(replacementVal)
				break
			}
		}
	}
	return nil
}

// gets the json value casted to the appropiate number size if the the field type is a number
// otherwise returns the value as is
func getNumberFieldValueOrSame(fieldType reflect.Type, jsonValue reflect.Value, logger *log.Logger) reflect.Value {
	// type of json number in raw marshalling is float64
	if jsonValue.Kind() != reflect.Float64 {
		return jsonValue
	}

	logger.Printf("Unmarshaling json number from float to integer type: %s", fieldType.Kind())
	jsonNumber := jsonValue.Float()
	// Convert float to the kind of the field
	switch fieldType.Kind() {
	case reflect.Int32:
		return reflect.ValueOf(int32(jsonNumber))
	case reflect.Int:
		return reflect.ValueOf(int(jsonNumber))
	case reflect.Int64:
		return reflect.ValueOf(int64(jsonNumber))
	default:
		return reflect.ValueOf(jsonNumber)
	}
}

// unmarshalRequestInfo unmarshals request information type data structures, returns a list of requestInfo data structures unmarshalled or an error
// a requestInfo type data structures are structures containing metadata information at the top level and a json object
// representing a request. The json object representing the request usually contains the json values flattened in the body, see unmarshalRequestEmbeddedJson
// Usually modeled as
// type requestInfo  struct {
// 		ContainerId string 				<=== First field Named "ContainerId"
//		Request  *some request type* 	<==== Second field named "Request"
//}
func unmarshalRequestInfo(holder []map[string]interface{}, requestInfo interface{}, logger *log.Logger) error {
	var err error
	if holder == nil {
		return fmt.Errorf("data to be unmarshaled can not be nil")
	}

	pointerSliceType := reflect.ValueOf(requestInfo).Type()
	if pointerSliceType.Kind() != reflect.Ptr {
		return fmt.Errorf("Type of requestInfo needs to be a pointer")
	}

	listRequest := reflect.ValueOf(requestInfo).Elem()
	sliceType := pointerSliceType.Elem()
	elemType := sliceType.Elem()

	for _, val := range holder {
		newRequestInfo := reflect.New(elemType)
		containerId := findValWithKeys(val, "containerId", "ContainerId").(string)
		newRequestInfo.Elem().FieldByName("ContainerId").SetString(containerId)

		requestField := newRequestInfo.Elem().FieldByName("Request")
		bts, _ := json.Marshal(findValWithKeys(val, "request", "Request"))
		err = unmarshalRequestEmbeddedJson(requestField.Addr().Interface(), bts, true, logger)
		if err != nil {
			return err
		}
		listRequest.Set(reflect.Append(listRequest, newRequestInfo.Elem()))
	}
	return nil
}

//unmarshalRequestEmbeddedJson unmarshals a json payload that contain embedded values to an struct with matching embedded fields
// Iterates over fields of the request,
// if there is an anonymous body field collect all embedded json values in a map
// marshal that map back to  to json
// unmarshal the map to the proper field marked as body
// for all other fields umarshal them from values whose keys match the field name
func unmarshalRequestEmbeddedJson(request interface{}, data []byte, embeddedValues bool, logger *log.Logger) (err error) {
	if embeddedValues == false {
		return fmt.Errorf("only embedded or flatten json values are supported at this point")
	}

	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}

	//umarshal to map
	var holder map[string]interface{}
	err = json.Unmarshal(data, &holder)
	if err != nil {
		return
	}

	if len(holder) == 0 {
		logger.Println("skipping unmarshaling, data is empty")
		return
	}

	val := reflect.ValueOf(request).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		structField := typ.Field(i)

		//unexported
		if structField.PkgPath != "" {
			continue
		}

		tag := structField.Tag.Get("contributesTo")
		fieldName := structField.Name
		fieldValue := val.Field(i)

		//If field is anonymous and contains body data
		if structField.Anonymous && tag == "body" {
			var exportedFieldNames []string
			// Get all exported subfields names of this struct field
			exportedFieldNames, err = getAllExportedFieldNamesOfStruct(fieldValue)
			if err != nil {
				return
			}

			//Find keys in the map that match the above field names. And put them in a map
			structOnlyMap := filterKeyValues(holder, exportedFieldNames...)
			if len(structOnlyMap) == 0 {
				logger.Printf("did not find any values in json for fields: %v. Skipping", exportedFieldNames)
				continue
			}

			//Marshal that map to json
			var jsonStructOnly []byte
			jsonStructOnly, err = json.Marshal(structOnlyMap)
			logger.Printf("for field: %v will unmarshal data: %v", fieldName, string(jsonStructOnly))
			if err != nil {
				return fmt.Errorf("could not marshal: %v to json while unmarshaling field: %v, due to: %v", structOnlyMap, fieldName, err)
			}

			//Finally unmarshal json into proper field
			err = json.Unmarshal(jsonStructOnly, fieldValue.Addr().Interface())
			if err != nil {
				return fmt.Errorf("could not unmarshal into field: %s with json: %v. Due to: %v", fieldName, string(jsonStructOnly), err)
			}
			continue
		}

		//Otherwise try to unmarshal into the field
		if fieldVal := findValWithKeys(holder, fieldName, uncapitalize(fieldName)); fieldVal != nil {
			fieldByteData, _ := json.Marshal(fieldVal)
			logger.Printf("unmarshaling field: %s, with type: %s with data: %v\n", fieldName, fieldValue.Type().Name(), string(fieldByteData))
			readCloserType := reflect.TypeOf((*io.ReadCloser)(nil)).Elem()
			if fieldValue.Type() == readCloserType {
				logger.Printf("unmarshaling binary content to a ReadCloser")
				buffer := bytes.NewBuffer(fieldByteData)
				fieldValue.Set(reflect.ValueOf(ioutil.NopCloser(buffer)))

				//TODO the testing service is not passing content len here
				if lenVal := val.FieldByName("ContentLength"); lenVal.IsValid() {
					logger.Println("WARN-- forcing the Content Length to be set")
					contentLen := int64(buffer.Len())
					lenVal.Set(reflect.ValueOf(&contentLen))
				}
			} else {
				err = json.Unmarshal(fieldByteData, fieldValue.Addr().Interface())
			}
		} else {
			logger.Printf("not found data for field: %s.Skipping\n", fieldName)
		}
	}
	return nil
}

func getAllExportedFieldNamesOfStruct(structValue reflect.Value) (allFieldNames []string, err error) {
	structType := structValue.Type()
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		//unexported
		if structField.PkgPath != "" {
			continue
		}
		allFieldNames = append(allFieldNames, structField.Name)
	}
	return
}

func filterKeyValues(src map[string]interface{}, fieldNames ...string) (dst map[string]interface{}) {
	dst = make(map[string]interface{})
	for _, fieldName := range fieldNames {
		if val := findValWithKeys(src, fieldName, uncapitalize(fieldName)); val != nil {
			dst[uncapitalize(fieldName)] = val
		}
	}
	return
}

func findValWithKeys(data map[string]interface{}, keys ...string) interface{} {
	for _, k := range keys {
		if val, ok := data[k]; ok {
			return val
		}
	}
	return nil
}

func prettyJson(jsonstring []byte) string {
	var buffer bytes.Buffer
	err := json.Indent(&buffer, jsonstring, "", "\t")
	if err != nil {
		return string(jsonstring)
	}
	return string(buffer.Bytes())
}

func failTestOnPanic(t *testing.T) {
	if r := recover(); r != nil {
		e := fmt.Errorf("failed due to panic: %s %s", r, debug.Stack())
		t.Error(e)
	}
}
