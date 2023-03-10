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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// TopologyAdminUserApiService TopologyAdminUserApi service
type TopologyAdminUserApiService service

type ApiAddTopologyAdminUserRequest struct {
	ctx                         context.Context
	ApiService                  *TopologyAdminUserApiService
	addTopologyAdminUserRequest *AddTopologyAdminUserRequest
}

// Create a new Topology Admin User in the config
func (r ApiAddTopologyAdminUserRequest) AddTopologyAdminUserRequest(addTopologyAdminUserRequest AddTopologyAdminUserRequest) ApiAddTopologyAdminUserRequest {
	r.addTopologyAdminUserRequest = &addTopologyAdminUserRequest
	return r
}

func (r ApiAddTopologyAdminUserRequest) Execute() (*TopologyAdminUserResponse, *http.Response, error) {
	return r.ApiService.AddTopologyAdminUserExecute(r)
}

/*
AddTopologyAdminUser Add a new Topology Admin User to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddTopologyAdminUserRequest
*/
func (a *TopologyAdminUserApiService) AddTopologyAdminUser(ctx context.Context) ApiAddTopologyAdminUserRequest {
	return ApiAddTopologyAdminUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TopologyAdminUserResponse
func (a *TopologyAdminUserApiService) AddTopologyAdminUserExecute(r ApiAddTopologyAdminUserRequest) (*TopologyAdminUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TopologyAdminUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologyAdminUserApiService.AddTopologyAdminUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/topology-admin-users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addTopologyAdminUserRequest == nil {
		return localVarReturnValue, nil, reportError("addTopologyAdminUserRequest is required and must be specified")
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
	localVarPostBody = r.addTopologyAdminUserRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteTopologyAdminUserRequest struct {
	ctx                   context.Context
	ApiService            *TopologyAdminUserApiService
	topologyAdminUserName string
}

func (r ApiDeleteTopologyAdminUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTopologyAdminUserExecute(r)
}

/*
DeleteTopologyAdminUser Delete a Topology Admin User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param topologyAdminUserName Name of the Topology Admin User
	@return ApiDeleteTopologyAdminUserRequest
*/
func (a *TopologyAdminUserApiService) DeleteTopologyAdminUser(ctx context.Context, topologyAdminUserName string) ApiDeleteTopologyAdminUserRequest {
	return ApiDeleteTopologyAdminUserRequest{
		ApiService:            a,
		ctx:                   ctx,
		topologyAdminUserName: topologyAdminUserName,
	}
}

// Execute executes the request
func (a *TopologyAdminUserApiService) DeleteTopologyAdminUserExecute(r ApiDeleteTopologyAdminUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologyAdminUserApiService.DeleteTopologyAdminUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/topology-admin-users/{topology-admin-user-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"topology-admin-user-name"+"}", url.PathEscape(parameterValueToString(r.topologyAdminUserName, "topologyAdminUserName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetTopologyAdminUserRequest struct {
	ctx                   context.Context
	ApiService            *TopologyAdminUserApiService
	topologyAdminUserName string
}

func (r ApiGetTopologyAdminUserRequest) Execute() (*TopologyAdminUserResponse, *http.Response, error) {
	return r.ApiService.GetTopologyAdminUserExecute(r)
}

/*
GetTopologyAdminUser Returns a single Topology Admin User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param topologyAdminUserName Name of the Topology Admin User
	@return ApiGetTopologyAdminUserRequest
*/
func (a *TopologyAdminUserApiService) GetTopologyAdminUser(ctx context.Context, topologyAdminUserName string) ApiGetTopologyAdminUserRequest {
	return ApiGetTopologyAdminUserRequest{
		ApiService:            a,
		ctx:                   ctx,
		topologyAdminUserName: topologyAdminUserName,
	}
}

// Execute executes the request
//
//	@return TopologyAdminUserResponse
func (a *TopologyAdminUserApiService) GetTopologyAdminUserExecute(r ApiGetTopologyAdminUserRequest) (*TopologyAdminUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TopologyAdminUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologyAdminUserApiService.GetTopologyAdminUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/topology-admin-users/{topology-admin-user-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"topology-admin-user-name"+"}", url.PathEscape(parameterValueToString(r.topologyAdminUserName, "topologyAdminUserName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateTopologyAdminUserRequest struct {
	ctx                   context.Context
	ApiService            *TopologyAdminUserApiService
	topologyAdminUserName string
	updateRequest         *UpdateRequest
}

// Update an existing Topology Admin User
func (r ApiUpdateTopologyAdminUserRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateTopologyAdminUserRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateTopologyAdminUserRequest) Execute() (*TopologyAdminUserResponse, *http.Response, error) {
	return r.ApiService.UpdateTopologyAdminUserExecute(r)
}

/*
UpdateTopologyAdminUser Update an existing Topology Admin User by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param topologyAdminUserName Name of the Topology Admin User
	@return ApiUpdateTopologyAdminUserRequest
*/
func (a *TopologyAdminUserApiService) UpdateTopologyAdminUser(ctx context.Context, topologyAdminUserName string) ApiUpdateTopologyAdminUserRequest {
	return ApiUpdateTopologyAdminUserRequest{
		ApiService:            a,
		ctx:                   ctx,
		topologyAdminUserName: topologyAdminUserName,
	}
}

// Execute executes the request
//
//	@return TopologyAdminUserResponse
func (a *TopologyAdminUserApiService) UpdateTopologyAdminUserExecute(r ApiUpdateTopologyAdminUserRequest) (*TopologyAdminUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TopologyAdminUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologyAdminUserApiService.UpdateTopologyAdminUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/topology-admin-users/{topology-admin-user-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"topology-admin-user-name"+"}", url.PathEscape(parameterValueToString(r.topologyAdminUserName, "topologyAdminUserName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
