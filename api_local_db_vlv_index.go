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

// LocalDbVlvIndexApiService LocalDbVlvIndexApi service
type LocalDbVlvIndexApiService service

type ApiAddLocalDbVlvIndexRequest struct {
	ctx                       context.Context
	ApiService                *LocalDbVlvIndexApiService
	backendName               string
	addLocalDbVlvIndexRequest *AddLocalDbVlvIndexRequest
}

// Create a new Local DB VLV Index in the config
func (r ApiAddLocalDbVlvIndexRequest) AddLocalDbVlvIndexRequest(addLocalDbVlvIndexRequest AddLocalDbVlvIndexRequest) ApiAddLocalDbVlvIndexRequest {
	r.addLocalDbVlvIndexRequest = &addLocalDbVlvIndexRequest
	return r
}

func (r ApiAddLocalDbVlvIndexRequest) Execute() (*LocalDbVlvIndexResponse, *http.Response, error) {
	return r.ApiService.AddLocalDbVlvIndexExecute(r)
}

/*
AddLocalDbVlvIndex Add a new Local DB VLV Index to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param backendName Name of the Backend
	@return ApiAddLocalDbVlvIndexRequest
*/
func (a *LocalDbVlvIndexApiService) AddLocalDbVlvIndex(ctx context.Context, backendName string) ApiAddLocalDbVlvIndexRequest {
	return ApiAddLocalDbVlvIndexRequest{
		ApiService:  a,
		ctx:         ctx,
		backendName: backendName,
	}
}

// Execute executes the request
//
//	@return LocalDbVlvIndexResponse
func (a *LocalDbVlvIndexApiService) AddLocalDbVlvIndexExecute(r ApiAddLocalDbVlvIndexRequest) (*LocalDbVlvIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LocalDbVlvIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocalDbVlvIndexApiService.AddLocalDbVlvIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/backends/{backend-name}/local-db-vlv-indexes"
	localVarPath = strings.Replace(localVarPath, "{"+"backend-name"+"}", url.PathEscape(parameterToString(r.backendName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addLocalDbVlvIndexRequest == nil {
		return localVarReturnValue, nil, reportError("addLocalDbVlvIndexRequest is required and must be specified")
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
	localVarPostBody = r.addLocalDbVlvIndexRequest
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

type ApiDeleteLocalDbVlvIndexRequest struct {
	ctx                 context.Context
	ApiService          *LocalDbVlvIndexApiService
	localDbVlvIndexName string
	backendName         string
}

func (r ApiDeleteLocalDbVlvIndexRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLocalDbVlvIndexExecute(r)
}

/*
DeleteLocalDbVlvIndex Delete a Local DB VLV Index

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param localDbVlvIndexName Name of the Local DB VLV Index
	@param backendName Name of the Backend
	@return ApiDeleteLocalDbVlvIndexRequest
*/
func (a *LocalDbVlvIndexApiService) DeleteLocalDbVlvIndex(ctx context.Context, localDbVlvIndexName string, backendName string) ApiDeleteLocalDbVlvIndexRequest {
	return ApiDeleteLocalDbVlvIndexRequest{
		ApiService:          a,
		ctx:                 ctx,
		localDbVlvIndexName: localDbVlvIndexName,
		backendName:         backendName,
	}
}

// Execute executes the request
func (a *LocalDbVlvIndexApiService) DeleteLocalDbVlvIndexExecute(r ApiDeleteLocalDbVlvIndexRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocalDbVlvIndexApiService.DeleteLocalDbVlvIndex")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/backends/{backend-name}/local-db-vlv-indexes/{local-db-vlv-index-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"local-db-vlv-index-name"+"}", url.PathEscape(parameterToString(r.localDbVlvIndexName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backend-name"+"}", url.PathEscape(parameterToString(r.backendName, "")), -1)

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

type ApiGetLocalDbVlvIndexRequest struct {
	ctx                 context.Context
	ApiService          *LocalDbVlvIndexApiService
	localDbVlvIndexName string
	backendName         string
}

func (r ApiGetLocalDbVlvIndexRequest) Execute() (*LocalDbVlvIndexResponse, *http.Response, error) {
	return r.ApiService.GetLocalDbVlvIndexExecute(r)
}

/*
GetLocalDbVlvIndex Returns a single Local DB VLV Index

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param localDbVlvIndexName Name of the Local DB VLV Index
	@param backendName Name of the Backend
	@return ApiGetLocalDbVlvIndexRequest
*/
func (a *LocalDbVlvIndexApiService) GetLocalDbVlvIndex(ctx context.Context, localDbVlvIndexName string, backendName string) ApiGetLocalDbVlvIndexRequest {
	return ApiGetLocalDbVlvIndexRequest{
		ApiService:          a,
		ctx:                 ctx,
		localDbVlvIndexName: localDbVlvIndexName,
		backendName:         backendName,
	}
}

// Execute executes the request
//
//	@return LocalDbVlvIndexResponse
func (a *LocalDbVlvIndexApiService) GetLocalDbVlvIndexExecute(r ApiGetLocalDbVlvIndexRequest) (*LocalDbVlvIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LocalDbVlvIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocalDbVlvIndexApiService.GetLocalDbVlvIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/backends/{backend-name}/local-db-vlv-indexes/{local-db-vlv-index-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"local-db-vlv-index-name"+"}", url.PathEscape(parameterToString(r.localDbVlvIndexName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backend-name"+"}", url.PathEscape(parameterToString(r.backendName, "")), -1)

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

type ApiUpdateLocalDbVlvIndexRequest struct {
	ctx                 context.Context
	ApiService          *LocalDbVlvIndexApiService
	localDbVlvIndexName string
	backendName         string
	updateRequest       *UpdateRequest
}

// Update an existing Local DB VLV Index
func (r ApiUpdateLocalDbVlvIndexRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateLocalDbVlvIndexRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateLocalDbVlvIndexRequest) Execute() (*LocalDbVlvIndexResponse, *http.Response, error) {
	return r.ApiService.UpdateLocalDbVlvIndexExecute(r)
}

/*
UpdateLocalDbVlvIndex Update an existing Local DB VLV Index by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param localDbVlvIndexName Name of the Local DB VLV Index
	@param backendName Name of the Backend
	@return ApiUpdateLocalDbVlvIndexRequest
*/
func (a *LocalDbVlvIndexApiService) UpdateLocalDbVlvIndex(ctx context.Context, localDbVlvIndexName string, backendName string) ApiUpdateLocalDbVlvIndexRequest {
	return ApiUpdateLocalDbVlvIndexRequest{
		ApiService:          a,
		ctx:                 ctx,
		localDbVlvIndexName: localDbVlvIndexName,
		backendName:         backendName,
	}
}

// Execute executes the request
//
//	@return LocalDbVlvIndexResponse
func (a *LocalDbVlvIndexApiService) UpdateLocalDbVlvIndexExecute(r ApiUpdateLocalDbVlvIndexRequest) (*LocalDbVlvIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LocalDbVlvIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocalDbVlvIndexApiService.UpdateLocalDbVlvIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/backends/{backend-name}/local-db-vlv-indexes/{local-db-vlv-index-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"local-db-vlv-index-name"+"}", url.PathEscape(parameterToString(r.localDbVlvIndexName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backend-name"+"}", url.PathEscape(parameterToString(r.backendName, "")), -1)

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
