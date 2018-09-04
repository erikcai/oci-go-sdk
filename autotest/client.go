// Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.

// Package autotest contains all auto generated integration test cases
package autotest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
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

func failIfError(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
		t.FailNow()
	}
}

//OCITestClient a client for the oci-testing-service
type OCITestClient struct {
	HTTPClient      http.Client
	ServiceEndpoint string
	SessionID       string
	Log             *log.Logger
}

//NewOCITestClient creates a client for the oci-testing-service
func NewOCITestClient() *OCITestClient {
	var testLog = log.New(os.Stderr, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	if _, isDebugEnabled := os.LookupEnv("OCI_GO_AUTOTEST_DEBUG"); !isDebugEnabled {
		testLog.SetOutput(ioutil.Discard)
	}
	return &OCITestClient{
		HTTPClient:      http.Client{},
		ServiceEndpoint: defaultOCITestingServiceEndpoint,
		Log:             testLog,
	}
}

func (client OCITestClient) getEndpointForService(serviceName, clientName string) (string, error) {
	endPoint := pathJoin(client.ServiceEndpoint, "request", "enable")
	request, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return "", err
	}
	request.URL.Query().Add("sessionId", client.SessionID)
	request.URL.Query().Add("serviceName", serviceName)
	request.URL.Query().Add("clientName", clientName)

	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}

func (client OCITestClient) isApiEnabled(serviceName, operationName string) (bool, error) {
	endPoint := pathJoin(client.ServiceEndpoint, "request", "enable")
	request, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return false, err
	}
	request.URL.Query().Add("serviceName", serviceName)
	request.URL.Query().Add("apiName", operationName)
	request.URL.Query().Add("lang", "Go")

	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
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

// validateError the result of SDK API call for failure case
func (client OCITestClient) validateError(containerId string, req interface{}, responseError error) (string, error) {
	data := ErrorToValidate{ContainerID: containerId}

	var err error
	reqType := reflect.TypeOf(req)
	data.RequestClass = fmt.Sprintf(requestClassTemplate, path.Base(reqType.PkgPath()), reqType.Name())
	data.RequestJSON, err = marshal(reflect.ValueOf(req))
	if err != nil {
		return "", err
	}
	data.ErrorJSON, err = marshal(reflect.ValueOf(responseError))
	if err != nil {
		return "", err
	}

	client.Log.Println("Error:" + data.ErrorJSON)

	body, _ := json.Marshal(data)

	client.Log.Println("ErrorToValidate:" + string(body))

	endPoint := pathJoin(client.ServiceEndpoint, fmt.Sprintf("error?sessionId=%s", client.SessionID))
	request, err := http.NewRequest("POST", endPoint, bytes.NewReader(body))

	if err != nil {
		return "", err
	}

	response, err := client.HTTPClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)

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
		data.RequestJSON, err = marshalStruct(reflect.ValueOf(req))
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
	responseJSON, err = marshal(reflect.ValueOf(res))
	return
}

func (client *OCITestClient) callValidate(body []byte) (string, error) {
	client.Log.Println("DataToValidate:", string(body))

	endPoint := pathJoin(client.ServiceEndpoint, fmt.Sprintf("response?sessionId=%s", client.SessionID))
	request, err := http.NewRequest("POST", endPoint, bytes.NewReader(body))

	if err != nil {
		return "", err
	}

	response, err := client.HTTPClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (client *OCITestClient) startSession() error {
	clientInfo := ClientInfo{Language: "Go"} // Java and Go share the same serializer
	body, _ := json.Marshal(clientInfo)

	client.Log.Println("ClientInfo:" + string(body))

	endPoint := pathJoin(client.ServiceEndpoint, "sessions")
	request, err := http.NewRequest("POST", endPoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = checkHttpResponse(response)
	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	sessionId := string(body)
	client.SessionID = sessionId
	client.Log.Println("SessionId:", client.SessionID)
	return nil
}

func (client OCITestClient) endSession() error {
	if client.SessionID == "" {
		return fmt.Errorf("can not end session if none has been started")
	}

	endPoint := pathJoin(client.ServiceEndpoint, "sessions", client.SessionID)
	request, err := http.NewRequest("DELETE", endPoint, nil)
	if err != nil {
		return err
	}

	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = checkHttpResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func (client OCITestClient) getRequests(serviceName string, apiName string) (string, error) {
	endPoint := pathJoin(client.ServiceEndpoint, "request")
	request, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		return "", err
	}

	query := request.URL.Query()
	query.Add("sessionId", client.SessionID)
	query.Add("serviceName", serviceName)
	query.Add("apiName", apiName)
	query.Add("lang", "Go")
	request.URL.RawQuery = query.Encode()

	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	err = checkHttpResponse(response)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	client.Log.Println("Request:" + string(body))
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
		//Report the error since all page-able responses need to have a nextPage
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

func checkHttpResponse(response *http.Response) error {
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf("response is not successful %d", response.StatusCode)
	}
	return nil
}

func marshal(val reflect.Value) (string, error) {
	if !val.IsValid() {
		return "", fmt.Errorf("marshaling value %#v is not supported", val)
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		return marshalSlice(val)
	case reflect.Struct:
		return marshalStruct(val)
	default:
		return "", fmt.Errorf("marshaling of value %#v is not supported", val)
	}
}
func marshalSlice(val reflect.Value) (string, error) {
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		return "", fmt.Errorf("can not marshal a not slice as a slice")
	}

	allFieldsMarshaled := make([]string, 0)
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i)
		m, e := marshal(v)
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
func marshalStruct(val reflect.Value) (string, error) {
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

		if contributesToOk || presentInOk || "servicefailure" == valueType.Name() {
			if !isFieldMandatory(sf) && isFieldEmpty(sv) {
				//	fmt.Println("ignoring: ", sf.Name, " is empty")
				continue
			}
			bs, e := json.Marshal(sv.Interface())
			if e != nil {
				return "", e
			}
			allFieldsMarshaled = append(allFieldsMarshaled,
				fmt.Sprintf(`"%s":%s`, uncapitalize(sf.Name), string(bs)))
		}
	}

	allJson := strings.Join(allFieldsMarshaled, ",")
	rs := "{" + allJson + "}"
	return rs, nil
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
