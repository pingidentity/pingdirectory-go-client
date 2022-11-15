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


// UncachedEntryCriteriaApiService UncachedEntryCriteriaApi service
type UncachedEntryCriteriaApiService service

type ApiAddUncachedEntryCriteriaRequest struct {
	ctx context.Context
	ApiService *UncachedEntryCriteriaApiService
	addUncachedEntryCriteriaRequest *AddUncachedEntryCriteriaRequest
}

// Create a new Uncached Entry Criteria in the config
func (r ApiAddUncachedEntryCriteriaRequest) AddUncachedEntryCriteriaRequest(addUncachedEntryCriteriaRequest AddUncachedEntryCriteriaRequest) ApiAddUncachedEntryCriteriaRequest {
	r.addUncachedEntryCriteriaRequest = &addUncachedEntryCriteriaRequest
	return r
}

func (r ApiAddUncachedEntryCriteriaRequest) Execute() (*AddUncachedEntryCriteria200Response, *http.Response, error) {
	return r.ApiService.AddUncachedEntryCriteriaExecute(r)
}

/*
AddUncachedEntryCriteria Add a new Uncached Entry Criteria to the config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddUncachedEntryCriteriaRequest
*/
func (a *UncachedEntryCriteriaApiService) AddUncachedEntryCriteria(ctx context.Context) ApiAddUncachedEntryCriteriaRequest {
	return ApiAddUncachedEntryCriteriaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddUncachedEntryCriteria200Response
func (a *UncachedEntryCriteriaApiService) AddUncachedEntryCriteriaExecute(r ApiAddUncachedEntryCriteriaRequest) (*AddUncachedEntryCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddUncachedEntryCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UncachedEntryCriteriaApiService.AddUncachedEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/uncached-entry-criteria"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addUncachedEntryCriteriaRequest == nil {
		return localVarReturnValue, nil, reportError("addUncachedEntryCriteriaRequest is required and must be specified")
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
	localVarPostBody = r.addUncachedEntryCriteriaRequest
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

type ApiDeleteUncachedEntryCriteriaRequest struct {
	ctx context.Context
	ApiService *UncachedEntryCriteriaApiService
	uncachedEntryCriteriaName string
}

func (r ApiDeleteUncachedEntryCriteriaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteUncachedEntryCriteriaExecute(r)
}

/*
DeleteUncachedEntryCriteria Delete a Uncached Entry Criteria

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uncachedEntryCriteriaName Name of the Uncached Entry Criteria to be deleted
 @return ApiDeleteUncachedEntryCriteriaRequest
*/
func (a *UncachedEntryCriteriaApiService) DeleteUncachedEntryCriteria(ctx context.Context, uncachedEntryCriteriaName string) ApiDeleteUncachedEntryCriteriaRequest {
	return ApiDeleteUncachedEntryCriteriaRequest{
		ApiService: a,
		ctx: ctx,
		uncachedEntryCriteriaName: uncachedEntryCriteriaName,
	}
}

// Execute executes the request
func (a *UncachedEntryCriteriaApiService) DeleteUncachedEntryCriteriaExecute(r ApiDeleteUncachedEntryCriteriaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UncachedEntryCriteriaApiService.DeleteUncachedEntryCriteria")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/uncached-entry-criteria/{uncached-entry-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"uncached-entry-criteria-name"+"}", url.PathEscape(parameterToString(r.uncachedEntryCriteriaName, "")), -1)

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

type ApiGetUncachedEntryCriteriaRequest struct {
	ctx context.Context
	ApiService *UncachedEntryCriteriaApiService
	uncachedEntryCriteriaName string
}

func (r ApiGetUncachedEntryCriteriaRequest) Execute() (*AddUncachedEntryCriteria200Response, *http.Response, error) {
	return r.ApiService.GetUncachedEntryCriteriaExecute(r)
}

/*
GetUncachedEntryCriteria Returns a single Uncached Entry Criteria

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uncachedEntryCriteriaName Name of the Uncached Entry Criteria to be read
 @return ApiGetUncachedEntryCriteriaRequest
*/
func (a *UncachedEntryCriteriaApiService) GetUncachedEntryCriteria(ctx context.Context, uncachedEntryCriteriaName string) ApiGetUncachedEntryCriteriaRequest {
	return ApiGetUncachedEntryCriteriaRequest{
		ApiService: a,
		ctx: ctx,
		uncachedEntryCriteriaName: uncachedEntryCriteriaName,
	}
}

// Execute executes the request
//  @return AddUncachedEntryCriteria200Response
func (a *UncachedEntryCriteriaApiService) GetUncachedEntryCriteriaExecute(r ApiGetUncachedEntryCriteriaRequest) (*AddUncachedEntryCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddUncachedEntryCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UncachedEntryCriteriaApiService.GetUncachedEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/uncached-entry-criteria/{uncached-entry-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"uncached-entry-criteria-name"+"}", url.PathEscape(parameterToString(r.uncachedEntryCriteriaName, "")), -1)

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

type ApiUpdateUncachedEntryCriteriaRequest struct {
	ctx context.Context
	ApiService *UncachedEntryCriteriaApiService
	uncachedEntryCriteriaName string
	updateRequest *UpdateRequest
}

// Update an existing Uncached Entry Criteria
func (r ApiUpdateUncachedEntryCriteriaRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateUncachedEntryCriteriaRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateUncachedEntryCriteriaRequest) Execute() (*AddUncachedEntryCriteria200Response, *http.Response, error) {
	return r.ApiService.UpdateUncachedEntryCriteriaExecute(r)
}

/*
UpdateUncachedEntryCriteria Update an existing Uncached Entry Criteria by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uncachedEntryCriteriaName Name of the Uncached Entry Criteria to be updated
 @return ApiUpdateUncachedEntryCriteriaRequest
*/
func (a *UncachedEntryCriteriaApiService) UpdateUncachedEntryCriteria(ctx context.Context, uncachedEntryCriteriaName string) ApiUpdateUncachedEntryCriteriaRequest {
	return ApiUpdateUncachedEntryCriteriaRequest{
		ApiService: a,
		ctx: ctx,
		uncachedEntryCriteriaName: uncachedEntryCriteriaName,
	}
}

// Execute executes the request
//  @return AddUncachedEntryCriteria200Response
func (a *UncachedEntryCriteriaApiService) UpdateUncachedEntryCriteriaExecute(r ApiUpdateUncachedEntryCriteriaRequest) (*AddUncachedEntryCriteria200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddUncachedEntryCriteria200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UncachedEntryCriteriaApiService.UpdateUncachedEntryCriteria")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/uncached-entry-criteria/{uncached-entry-criteria-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"uncached-entry-criteria-name"+"}", url.PathEscape(parameterToString(r.uncachedEntryCriteriaName, "")), -1)

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
