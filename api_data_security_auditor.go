/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DataSecurityAuditorApiService DataSecurityAuditorApi service
type DataSecurityAuditorApiService service

type ApiAddDataSecurityAuditorRequest struct {
	ctx                           context.Context
	ApiService                    *DataSecurityAuditorApiService
	addDataSecurityAuditorRequest *AddDataSecurityAuditorRequest
}

// Create a new Data Security Auditor in the config
func (r ApiAddDataSecurityAuditorRequest) AddDataSecurityAuditorRequest(addDataSecurityAuditorRequest AddDataSecurityAuditorRequest) ApiAddDataSecurityAuditorRequest {
	r.addDataSecurityAuditorRequest = &addDataSecurityAuditorRequest
	return r
}

func (r ApiAddDataSecurityAuditorRequest) Execute() (*AddDataSecurityAuditor200Response, *http.Response, error) {
	return r.ApiService.AddDataSecurityAuditorExecute(r)
}

/*
AddDataSecurityAuditor Add a new Data Security Auditor to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddDataSecurityAuditorRequest
*/
func (a *DataSecurityAuditorApiService) AddDataSecurityAuditor(ctx context.Context) ApiAddDataSecurityAuditorRequest {
	return ApiAddDataSecurityAuditorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddDataSecurityAuditor200Response
func (a *DataSecurityAuditorApiService) AddDataSecurityAuditorExecute(r ApiAddDataSecurityAuditorRequest) (*AddDataSecurityAuditor200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddDataSecurityAuditor200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataSecurityAuditorApiService.AddDataSecurityAuditor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-security-auditors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addDataSecurityAuditorRequest == nil {
		return localVarReturnValue, nil, reportError("addDataSecurityAuditorRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addDataSecurityAuditorRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteDataSecurityAuditorRequest struct {
	ctx                     context.Context
	ApiService              *DataSecurityAuditorApiService
	dataSecurityAuditorName string
}

func (r ApiDeleteDataSecurityAuditorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDataSecurityAuditorExecute(r)
}

/*
DeleteDataSecurityAuditor Delete a Data Security Auditor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dataSecurityAuditorName Name of the Data Security Auditor
	@return ApiDeleteDataSecurityAuditorRequest
*/
func (a *DataSecurityAuditorApiService) DeleteDataSecurityAuditor(ctx context.Context, dataSecurityAuditorName string) ApiDeleteDataSecurityAuditorRequest {
	return ApiDeleteDataSecurityAuditorRequest{
		ApiService:              a,
		ctx:                     ctx,
		dataSecurityAuditorName: dataSecurityAuditorName,
	}
}

// Execute executes the request
func (a *DataSecurityAuditorApiService) DeleteDataSecurityAuditorExecute(r ApiDeleteDataSecurityAuditorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataSecurityAuditorApiService.DeleteDataSecurityAuditor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-security-auditors/{data-security-auditor-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"data-security-auditor-name"+"}", url.PathEscape(parameterToString(r.dataSecurityAuditorName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetDataSecurityAuditorRequest struct {
	ctx                     context.Context
	ApiService              *DataSecurityAuditorApiService
	dataSecurityAuditorName string
}

func (r ApiGetDataSecurityAuditorRequest) Execute() (*AddDataSecurityAuditor200Response, *http.Response, error) {
	return r.ApiService.GetDataSecurityAuditorExecute(r)
}

/*
GetDataSecurityAuditor Returns a single Data Security Auditor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dataSecurityAuditorName Name of the Data Security Auditor
	@return ApiGetDataSecurityAuditorRequest
*/
func (a *DataSecurityAuditorApiService) GetDataSecurityAuditor(ctx context.Context, dataSecurityAuditorName string) ApiGetDataSecurityAuditorRequest {
	return ApiGetDataSecurityAuditorRequest{
		ApiService:              a,
		ctx:                     ctx,
		dataSecurityAuditorName: dataSecurityAuditorName,
	}
}

// Execute executes the request
//
//	@return AddDataSecurityAuditor200Response
func (a *DataSecurityAuditorApiService) GetDataSecurityAuditorExecute(r ApiGetDataSecurityAuditorRequest) (*AddDataSecurityAuditor200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddDataSecurityAuditor200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataSecurityAuditorApiService.GetDataSecurityAuditor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-security-auditors/{data-security-auditor-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"data-security-auditor-name"+"}", url.PathEscape(parameterToString(r.dataSecurityAuditorName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateDataSecurityAuditorRequest struct {
	ctx                     context.Context
	ApiService              *DataSecurityAuditorApiService
	dataSecurityAuditorName string
	updateRequest           *UpdateRequest
}

// Update an existing Data Security Auditor
func (r ApiUpdateDataSecurityAuditorRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateDataSecurityAuditorRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateDataSecurityAuditorRequest) Execute() (*AddDataSecurityAuditor200Response, *http.Response, error) {
	return r.ApiService.UpdateDataSecurityAuditorExecute(r)
}

/*
UpdateDataSecurityAuditor Update an existing Data Security Auditor by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dataSecurityAuditorName Name of the Data Security Auditor
	@return ApiUpdateDataSecurityAuditorRequest
*/
func (a *DataSecurityAuditorApiService) UpdateDataSecurityAuditor(ctx context.Context, dataSecurityAuditorName string) ApiUpdateDataSecurityAuditorRequest {
	return ApiUpdateDataSecurityAuditorRequest{
		ApiService:              a,
		ctx:                     ctx,
		dataSecurityAuditorName: dataSecurityAuditorName,
	}
}

// Execute executes the request
//
//	@return AddDataSecurityAuditor200Response
func (a *DataSecurityAuditorApiService) UpdateDataSecurityAuditorExecute(r ApiUpdateDataSecurityAuditorRequest) (*AddDataSecurityAuditor200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddDataSecurityAuditor200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataSecurityAuditorApiService.UpdateDataSecurityAuditor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-security-auditors/{data-security-auditor-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"data-security-auditor-name"+"}", url.PathEscape(parameterToString(r.dataSecurityAuditorName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
