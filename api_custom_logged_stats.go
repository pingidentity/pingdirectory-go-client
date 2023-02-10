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

// CustomLoggedStatsApiService CustomLoggedStatsApi service
type CustomLoggedStatsApiService service

type ApiAddCustomLoggedStatsRequest struct {
	ctx                         context.Context
	ApiService                  *CustomLoggedStatsApiService
	addCustomLoggedStatsRequest *AddCustomLoggedStatsRequest
}

// Create a new Custom Logged Stats in the config
func (r ApiAddCustomLoggedStatsRequest) AddCustomLoggedStatsRequest(addCustomLoggedStatsRequest AddCustomLoggedStatsRequest) ApiAddCustomLoggedStatsRequest {
	r.addCustomLoggedStatsRequest = &addCustomLoggedStatsRequest
	return r
}

func (r ApiAddCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsResponse, *http.Response, error) {
	return r.ApiService.AddCustomLoggedStatsExecute(r)
}

/*
AddCustomLoggedStats Add a new Custom Logged Stats to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsApiService) AddCustomLoggedStats(ctx context.Context) ApiAddCustomLoggedStatsRequest {
	return ApiAddCustomLoggedStatsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsResponse
func (a *CustomLoggedStatsApiService) AddCustomLoggedStatsExecute(r ApiAddCustomLoggedStatsRequest) (*CustomLoggedStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsApiService.AddCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-logged-stats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addCustomLoggedStatsRequest == nil {
		return localVarReturnValue, nil, reportError("addCustomLoggedStatsRequest is required and must be specified")
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
	localVarPostBody = r.addCustomLoggedStatsRequest
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

type ApiDeleteCustomLoggedStatsRequest struct {
	ctx                   context.Context
	ApiService            *CustomLoggedStatsApiService
	customLoggedStatsName string
}

func (r ApiDeleteCustomLoggedStatsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomLoggedStatsExecute(r)
}

/*
DeleteCustomLoggedStats Delete a Custom Logged Stats

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customLoggedStatsName Name of the Custom Logged Stats to be deleted
	@return ApiDeleteCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsApiService) DeleteCustomLoggedStats(ctx context.Context, customLoggedStatsName string) ApiDeleteCustomLoggedStatsRequest {
	return ApiDeleteCustomLoggedStatsRequest{
		ApiService:            a,
		ctx:                   ctx,
		customLoggedStatsName: customLoggedStatsName,
	}
}

// Execute executes the request
func (a *CustomLoggedStatsApiService) DeleteCustomLoggedStatsExecute(r ApiDeleteCustomLoggedStatsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsApiService.DeleteCustomLoggedStats")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-logged-stats/{custom-logged-stats-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom-logged-stats-name"+"}", url.PathEscape(parameterToString(r.customLoggedStatsName, "")), -1)

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

type ApiGetCustomLoggedStatsRequest struct {
	ctx                   context.Context
	ApiService            *CustomLoggedStatsApiService
	customLoggedStatsName string
}

func (r ApiGetCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsResponse, *http.Response, error) {
	return r.ApiService.GetCustomLoggedStatsExecute(r)
}

/*
GetCustomLoggedStats Returns a single Custom Logged Stats

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customLoggedStatsName Name of the Custom Logged Stats to be read
	@return ApiGetCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsApiService) GetCustomLoggedStats(ctx context.Context, customLoggedStatsName string) ApiGetCustomLoggedStatsRequest {
	return ApiGetCustomLoggedStatsRequest{
		ApiService:            a,
		ctx:                   ctx,
		customLoggedStatsName: customLoggedStatsName,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsResponse
func (a *CustomLoggedStatsApiService) GetCustomLoggedStatsExecute(r ApiGetCustomLoggedStatsRequest) (*CustomLoggedStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsApiService.GetCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-logged-stats/{custom-logged-stats-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom-logged-stats-name"+"}", url.PathEscape(parameterToString(r.customLoggedStatsName, "")), -1)

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

type ApiUpdateCustomLoggedStatsRequest struct {
	ctx                   context.Context
	ApiService            *CustomLoggedStatsApiService
	customLoggedStatsName string
	updateRequest         *UpdateRequest
}

// Update an existing Custom Logged Stats
func (r ApiUpdateCustomLoggedStatsRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateCustomLoggedStatsRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsResponse, *http.Response, error) {
	return r.ApiService.UpdateCustomLoggedStatsExecute(r)
}

/*
UpdateCustomLoggedStats Update an existing Custom Logged Stats by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customLoggedStatsName Name of the Custom Logged Stats to be updated
	@return ApiUpdateCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsApiService) UpdateCustomLoggedStats(ctx context.Context, customLoggedStatsName string) ApiUpdateCustomLoggedStatsRequest {
	return ApiUpdateCustomLoggedStatsRequest{
		ApiService:            a,
		ctx:                   ctx,
		customLoggedStatsName: customLoggedStatsName,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsResponse
func (a *CustomLoggedStatsApiService) UpdateCustomLoggedStatsExecute(r ApiUpdateCustomLoggedStatsRequest) (*CustomLoggedStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsApiService.UpdateCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom-logged-stats/{custom-logged-stats-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom-logged-stats-name"+"}", url.PathEscape(parameterToString(r.customLoggedStatsName, "")), -1)

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
