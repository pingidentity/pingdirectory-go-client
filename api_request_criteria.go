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

// RequestCriteriaApiService RequestCriteriaApi service
type RequestCriteriaApiService service

type ApiAddRequestCriteriaRequest struct {
	ctx                       context.Context
	ApiService                *RequestCriteriaApiService
	addRequestCriteriaRequest *AddRequestCriteriaRequest
}

// Create a new Request Criteria in the config
func (r ApiAddRequestCriteriaRequest) AddRequestCriteriaRequest(addRequestCriteriaRequest AddRequestCriteriaRequest) ApiAddRequestCriteriaRequest {
	r.addRequestCriteriaRequest = &addRequestCriteriaRequest
	return r
}

func (r ApiAddRequestCriteriaRequest) Execute() (*AddRequestCriteria200Response, *http.Response, error) {
	return r.ApiService.AddRequestCriteriaExecute(r)
}

/*
AddRequestCriteria Add a new Request Criteria to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddRequestCriteriaRequest
*/
func (a *RequestCriteriaApiService) AddRequestCriteria(ctx context.Context) ApiAddRequestCriteriaRequest {
	return ApiAddRequestCriteriaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddRequestCriteria200Response
func (a *RequestCriteriaApiService) AddRequestCriteriaExecute(r ApiAddRequestCriteriaRequest) (*AddRequestCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddRequestCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestCriteriaApiService.AddRequestCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/request-criteria"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addRequestCriteriaRequest == nil {
		return localVarReturnValue, nil, reportError("addRequestCriteriaRequest is required and must be specified")
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
	localVarPostBody = r.addRequestCriteriaRequest
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

type ApiDeleteRequestCriteriaRequest struct {
	ctx                 context.Context
	ApiService          *RequestCriteriaApiService
	requestCriteriaName string
}

func (r ApiDeleteRequestCriteriaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRequestCriteriaExecute(r)
}

/*
DeleteRequestCriteria Delete a Request Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestCriteriaName Name of the Request Criteria
	@return ApiDeleteRequestCriteriaRequest
*/
func (a *RequestCriteriaApiService) DeleteRequestCriteria(ctx context.Context, requestCriteriaName string) ApiDeleteRequestCriteriaRequest {
	return ApiDeleteRequestCriteriaRequest{
		ApiService:          a,
		ctx:                 ctx,
		requestCriteriaName: requestCriteriaName,
	}
}

// Execute executes the request
func (a *RequestCriteriaApiService) DeleteRequestCriteriaExecute(r ApiDeleteRequestCriteriaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestCriteriaApiService.DeleteRequestCriteria")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/request-criteria/{request-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"request-criteria-name"+"}", url.PathEscape(parameterToString(r.requestCriteriaName, "")), -1)

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

type ApiGetRequestCriteriaRequest struct {
	ctx                 context.Context
	ApiService          *RequestCriteriaApiService
	requestCriteriaName string
}

func (r ApiGetRequestCriteriaRequest) Execute() (*AddRequestCriteria200Response, *http.Response, error) {
	return r.ApiService.GetRequestCriteriaExecute(r)
}

/*
GetRequestCriteria Returns a single Request Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestCriteriaName Name of the Request Criteria
	@return ApiGetRequestCriteriaRequest
*/
func (a *RequestCriteriaApiService) GetRequestCriteria(ctx context.Context, requestCriteriaName string) ApiGetRequestCriteriaRequest {
	return ApiGetRequestCriteriaRequest{
		ApiService:          a,
		ctx:                 ctx,
		requestCriteriaName: requestCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddRequestCriteria200Response
func (a *RequestCriteriaApiService) GetRequestCriteriaExecute(r ApiGetRequestCriteriaRequest) (*AddRequestCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddRequestCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestCriteriaApiService.GetRequestCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/request-criteria/{request-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"request-criteria-name"+"}", url.PathEscape(parameterToString(r.requestCriteriaName, "")), -1)

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

type ApiUpdateRequestCriteriaRequest struct {
	ctx                 context.Context
	ApiService          *RequestCriteriaApiService
	requestCriteriaName string
	updateRequest       *UpdateRequest
}

// Update an existing Request Criteria
func (r ApiUpdateRequestCriteriaRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateRequestCriteriaRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateRequestCriteriaRequest) Execute() (*AddRequestCriteria200Response, *http.Response, error) {
	return r.ApiService.UpdateRequestCriteriaExecute(r)
}

/*
UpdateRequestCriteria Update an existing Request Criteria by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestCriteriaName Name of the Request Criteria
	@return ApiUpdateRequestCriteriaRequest
*/
func (a *RequestCriteriaApiService) UpdateRequestCriteria(ctx context.Context, requestCriteriaName string) ApiUpdateRequestCriteriaRequest {
	return ApiUpdateRequestCriteriaRequest{
		ApiService:          a,
		ctx:                 ctx,
		requestCriteriaName: requestCriteriaName,
	}
}

// Execute executes the request
//
//	@return AddRequestCriteria200Response
func (a *RequestCriteriaApiService) UpdateRequestCriteriaExecute(r ApiUpdateRequestCriteriaRequest) (*AddRequestCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddRequestCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestCriteriaApiService.UpdateRequestCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/request-criteria/{request-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"request-criteria-name"+"}", url.PathEscape(parameterToString(r.requestCriteriaName, "")), -1)

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
