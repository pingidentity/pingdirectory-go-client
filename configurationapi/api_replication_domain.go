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

// ReplicationDomainApiService ReplicationDomainApi service
type ReplicationDomainApiService service

type ApiGetReplicationDomainRequest struct {
	ctx                         context.Context
	ApiService                  *ReplicationDomainApiService
	replicationDomainName       string
	synchronizationProviderName string
}

func (r ApiGetReplicationDomainRequest) Execute() (*ReplicationDomainResponse, *http.Response, error) {
	return r.ApiService.GetReplicationDomainExecute(r)
}

/*
GetReplicationDomain Returns a single Replication Domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param replicationDomainName Name of the Replication Domain
	@param synchronizationProviderName Name of the Synchronization Provider
	@return ApiGetReplicationDomainRequest
*/
func (a *ReplicationDomainApiService) GetReplicationDomain(ctx context.Context, replicationDomainName string, synchronizationProviderName string) ApiGetReplicationDomainRequest {
	return ApiGetReplicationDomainRequest{
		ApiService:                  a,
		ctx:                         ctx,
		replicationDomainName:       replicationDomainName,
		synchronizationProviderName: synchronizationProviderName,
	}
}

// Execute executes the request
//
//	@return ReplicationDomainResponse
func (a *ReplicationDomainApiService) GetReplicationDomainExecute(r ApiGetReplicationDomainRequest) (*ReplicationDomainResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationDomainResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationDomainApiService.GetReplicationDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization-providers/{synchronization-provider-name}/replication-domains/{replication-domain-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-domain-name"+"}", url.PathEscape(parameterValueToString(r.replicationDomainName, "replicationDomainName")), -1)
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

type ApiUpdateReplicationDomainRequest struct {
	ctx                         context.Context
	ApiService                  *ReplicationDomainApiService
	replicationDomainName       string
	synchronizationProviderName string
	updateRequest               *UpdateRequest
}

// Update an existing Replication Domain
func (r ApiUpdateReplicationDomainRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateReplicationDomainRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateReplicationDomainRequest) Execute() (*ReplicationDomainResponse, *http.Response, error) {
	return r.ApiService.UpdateReplicationDomainExecute(r)
}

/*
UpdateReplicationDomain Update an existing Replication Domain by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param replicationDomainName Name of the Replication Domain
	@param synchronizationProviderName Name of the Synchronization Provider
	@return ApiUpdateReplicationDomainRequest
*/
func (a *ReplicationDomainApiService) UpdateReplicationDomain(ctx context.Context, replicationDomainName string, synchronizationProviderName string) ApiUpdateReplicationDomainRequest {
	return ApiUpdateReplicationDomainRequest{
		ApiService:                  a,
		ctx:                         ctx,
		replicationDomainName:       replicationDomainName,
		synchronizationProviderName: synchronizationProviderName,
	}
}

// Execute executes the request
//
//	@return ReplicationDomainResponse
func (a *ReplicationDomainApiService) UpdateReplicationDomainExecute(r ApiUpdateReplicationDomainRequest) (*ReplicationDomainResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationDomainResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationDomainApiService.UpdateReplicationDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synchronization-providers/{synchronization-provider-name}/replication-domains/{replication-domain-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-domain-name"+"}", url.PathEscape(parameterValueToString(r.replicationDomainName, "replicationDomainName")), -1)
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
