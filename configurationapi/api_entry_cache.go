/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// EntryCacheApiService EntryCacheApi service
type EntryCacheApiService service

type ApiAddEntryCacheRequest struct {
	ctx                      context.Context
	ApiService               *EntryCacheApiService
	addFifoEntryCacheRequest *AddFifoEntryCacheRequest
}

// Create a new Entry Cache in the config
func (r ApiAddEntryCacheRequest) AddFifoEntryCacheRequest(addFifoEntryCacheRequest AddFifoEntryCacheRequest) ApiAddEntryCacheRequest {
	r.addFifoEntryCacheRequest = &addFifoEntryCacheRequest
	return r
}

func (r ApiAddEntryCacheRequest) Execute() (*FifoEntryCacheResponse, *http.Response, error) {
	return r.ApiService.AddEntryCacheExecute(r)
}

/*
AddEntryCache Add a new Entry Cache to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddEntryCacheRequest
*/
func (a *EntryCacheApiService) AddEntryCache(ctx context.Context) ApiAddEntryCacheRequest {
	return ApiAddEntryCacheRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FifoEntryCacheResponse
func (a *EntryCacheApiService) AddEntryCacheExecute(r ApiAddEntryCacheRequest) (*FifoEntryCacheResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FifoEntryCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryCacheApiService.AddEntryCache")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entry-caches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addFifoEntryCacheRequest == nil {
		return localVarReturnValue, nil, reportError("addFifoEntryCacheRequest is required and must be specified")
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
	localVarPostBody = r.addFifoEntryCacheRequest
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

type ApiDeleteEntryCacheRequest struct {
	ctx            context.Context
	ApiService     *EntryCacheApiService
	entryCacheName string
}

func (r ApiDeleteEntryCacheRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEntryCacheExecute(r)
}

/*
DeleteEntryCache Delete a Entry Cache

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryCacheName Name of the Entry Cache
	@return ApiDeleteEntryCacheRequest
*/
func (a *EntryCacheApiService) DeleteEntryCache(ctx context.Context, entryCacheName string) ApiDeleteEntryCacheRequest {
	return ApiDeleteEntryCacheRequest{
		ApiService:     a,
		ctx:            ctx,
		entryCacheName: entryCacheName,
	}
}

// Execute executes the request
func (a *EntryCacheApiService) DeleteEntryCacheExecute(r ApiDeleteEntryCacheRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryCacheApiService.DeleteEntryCache")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entry-caches/{entry-cache-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"entry-cache-name"+"}", url.PathEscape(parameterToString(r.entryCacheName, "")), -1)

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

type ApiGetEntryCacheRequest struct {
	ctx            context.Context
	ApiService     *EntryCacheApiService
	entryCacheName string
}

func (r ApiGetEntryCacheRequest) Execute() (*FifoEntryCacheResponse, *http.Response, error) {
	return r.ApiService.GetEntryCacheExecute(r)
}

/*
GetEntryCache Returns a single Entry Cache

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryCacheName Name of the Entry Cache
	@return ApiGetEntryCacheRequest
*/
func (a *EntryCacheApiService) GetEntryCache(ctx context.Context, entryCacheName string) ApiGetEntryCacheRequest {
	return ApiGetEntryCacheRequest{
		ApiService:     a,
		ctx:            ctx,
		entryCacheName: entryCacheName,
	}
}

// Execute executes the request
//
//	@return FifoEntryCacheResponse
func (a *EntryCacheApiService) GetEntryCacheExecute(r ApiGetEntryCacheRequest) (*FifoEntryCacheResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FifoEntryCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryCacheApiService.GetEntryCache")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entry-caches/{entry-cache-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"entry-cache-name"+"}", url.PathEscape(parameterToString(r.entryCacheName, "")), -1)

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

type ApiUpdateEntryCacheRequest struct {
	ctx            context.Context
	ApiService     *EntryCacheApiService
	entryCacheName string
	updateRequest  *UpdateRequest
}

// Update an existing Entry Cache
func (r ApiUpdateEntryCacheRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateEntryCacheRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateEntryCacheRequest) Execute() (*FifoEntryCacheResponse, *http.Response, error) {
	return r.ApiService.UpdateEntryCacheExecute(r)
}

/*
UpdateEntryCache Update an existing Entry Cache by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryCacheName Name of the Entry Cache
	@return ApiUpdateEntryCacheRequest
*/
func (a *EntryCacheApiService) UpdateEntryCache(ctx context.Context, entryCacheName string) ApiUpdateEntryCacheRequest {
	return ApiUpdateEntryCacheRequest{
		ApiService:     a,
		ctx:            ctx,
		entryCacheName: entryCacheName,
	}
}

// Execute executes the request
//
//	@return FifoEntryCacheResponse
func (a *EntryCacheApiService) UpdateEntryCacheExecute(r ApiUpdateEntryCacheRequest) (*FifoEntryCacheResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FifoEntryCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryCacheApiService.UpdateEntryCache")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entry-caches/{entry-cache-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"entry-cache-name"+"}", url.PathEscape(parameterToString(r.entryCacheName, "")), -1)

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