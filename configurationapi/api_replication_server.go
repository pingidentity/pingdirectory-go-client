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

// ReplicationServerAPIService ReplicationServerAPI service
type ReplicationServerAPIService service

type ApiGetReplicationServerRequest struct {
	ctx                         context.Context
	ApiService                  *ReplicationServerAPIService
	synchronizationProviderName string
}

func (r ApiGetReplicationServerRequest) Execute() (*ReplicationServerResponse, *http.Response, error) {
	return r.ApiService.GetReplicationServerExecute(r)
}

/*
GetReplicationServer Returns a single Replication Server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param synchronizationProviderName Name of the Synchronization Provider
	@return ApiGetReplicationServerRequest
*/
func (a *ReplicationServerAPIService) GetReplicationServer(ctx context.Context, synchronizationProviderName string) ApiGetReplicationServerRequest {
	return ApiGetReplicationServerRequest{
		ApiService:                  a,
		ctx:                         ctx,
		synchronizationProviderName: synchronizationProviderName,
	}
}

// Execute executes the request
//
//	@return ReplicationServerResponse
func (a *ReplicationServerAPIService) GetReplicationServerExecute(r ApiGetReplicationServerRequest) (*ReplicationServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationServerAPIService.GetReplicationServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization-providers/{synchronization-provider-name}/replication-server"
	localVarPath = strings.Replace(localVarPath, "{"+"synchronization-provider-name"+"}", url.PathEscape(parameterValueToString(r.synchronizationProviderName, "synchronizationProviderName")), -1)

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

type ApiUpdateReplicationServerRequest struct {
	ctx                         context.Context
	ApiService                  *ReplicationServerAPIService
	synchronizationProviderName string
	updateRequest               *UpdateRequest
}

// Update an existing Replication Server
func (r ApiUpdateReplicationServerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateReplicationServerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateReplicationServerRequest) Execute() (*ReplicationServerResponse, *http.Response, error) {
	return r.ApiService.UpdateReplicationServerExecute(r)
}

/*
UpdateReplicationServer Update an existing Replication Server by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param synchronizationProviderName Name of the Synchronization Provider
	@return ApiUpdateReplicationServerRequest
*/
func (a *ReplicationServerAPIService) UpdateReplicationServer(ctx context.Context, synchronizationProviderName string) ApiUpdateReplicationServerRequest {
	return ApiUpdateReplicationServerRequest{
		ApiService:                  a,
		ctx:                         ctx,
		synchronizationProviderName: synchronizationProviderName,
	}
}

// Execute executes the request
//
//	@return ReplicationServerResponse
func (a *ReplicationServerAPIService) UpdateReplicationServerExecute(r ApiUpdateReplicationServerRequest) (*ReplicationServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationServerAPIService.UpdateReplicationServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization-providers/{synchronization-provider-name}/replication-server"
	localVarPath = strings.Replace(localVarPath, "{"+"synchronization-provider-name"+"}", url.PathEscape(parameterValueToString(r.synchronizationProviderName, "synchronizationProviderName")), -1)

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
