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


// ReplicationAssurancePolicyApiService ReplicationAssurancePolicyApi service
type ReplicationAssurancePolicyApiService service

type ApiAddReplicationAssurancePolicyRequest struct {
	ctx context.Context
	ApiService *ReplicationAssurancePolicyApiService
	addReplicationAssurancePolicyRequest *AddReplicationAssurancePolicyRequest
}

// Create a new Replication Assurance Policy in the config
func (r ApiAddReplicationAssurancePolicyRequest) AddReplicationAssurancePolicyRequest(addReplicationAssurancePolicyRequest AddReplicationAssurancePolicyRequest) ApiAddReplicationAssurancePolicyRequest {
	r.addReplicationAssurancePolicyRequest = &addReplicationAssurancePolicyRequest
	return r
}

func (r ApiAddReplicationAssurancePolicyRequest) Execute() (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	return r.ApiService.AddReplicationAssurancePolicyExecute(r)
}

/*
AddReplicationAssurancePolicy Add a new Replication Assurance Policy to the config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyApiService) AddReplicationAssurancePolicy(ctx context.Context) ApiAddReplicationAssurancePolicyRequest {
	return ApiAddReplicationAssurancePolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReplicationAssurancePolicyResponse
func (a *ReplicationAssurancePolicyApiService) AddReplicationAssurancePolicyExecute(r ApiAddReplicationAssurancePolicyRequest) (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReplicationAssurancePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyApiService.AddReplicationAssurancePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addReplicationAssurancePolicyRequest == nil {
		return localVarReturnValue, nil, reportError("addReplicationAssurancePolicyRequest is required and must be specified")
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
	localVarPostBody = r.addReplicationAssurancePolicyRequest
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

type ApiDeleteReplicationAssurancePolicyRequest struct {
	ctx context.Context
	ApiService *ReplicationAssurancePolicyApiService
	replicationAssurancePolicyName string
}

func (r ApiDeleteReplicationAssurancePolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteReplicationAssurancePolicyExecute(r)
}

/*
DeleteReplicationAssurancePolicy Delete a Replication Assurance Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param replicationAssurancePolicyName Name of the Replication Assurance Policy to be deleted
 @return ApiDeleteReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyApiService) DeleteReplicationAssurancePolicy(ctx context.Context, replicationAssurancePolicyName string) ApiDeleteReplicationAssurancePolicyRequest {
	return ApiDeleteReplicationAssurancePolicyRequest{
		ApiService: a,
		ctx: ctx,
		replicationAssurancePolicyName: replicationAssurancePolicyName,
	}
}

// Execute executes the request
func (a *ReplicationAssurancePolicyApiService) DeleteReplicationAssurancePolicyExecute(r ApiDeleteReplicationAssurancePolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyApiService.DeleteReplicationAssurancePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies/{replication-assurance-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-assurance-policy-name"+"}", url.PathEscape(parameterToString(r.replicationAssurancePolicyName, "")), -1)

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

type ApiGetReplicationAssurancePolicyRequest struct {
	ctx context.Context
	ApiService *ReplicationAssurancePolicyApiService
	replicationAssurancePolicyName string
}

func (r ApiGetReplicationAssurancePolicyRequest) Execute() (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	return r.ApiService.GetReplicationAssurancePolicyExecute(r)
}

/*
GetReplicationAssurancePolicy Returns a single Replication Assurance Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param replicationAssurancePolicyName Name of the Replication Assurance Policy to be read
 @return ApiGetReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyApiService) GetReplicationAssurancePolicy(ctx context.Context, replicationAssurancePolicyName string) ApiGetReplicationAssurancePolicyRequest {
	return ApiGetReplicationAssurancePolicyRequest{
		ApiService: a,
		ctx: ctx,
		replicationAssurancePolicyName: replicationAssurancePolicyName,
	}
}

// Execute executes the request
//  @return ReplicationAssurancePolicyResponse
func (a *ReplicationAssurancePolicyApiService) GetReplicationAssurancePolicyExecute(r ApiGetReplicationAssurancePolicyRequest) (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReplicationAssurancePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyApiService.GetReplicationAssurancePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies/{replication-assurance-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-assurance-policy-name"+"}", url.PathEscape(parameterToString(r.replicationAssurancePolicyName, "")), -1)

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

type ApiUpdateReplicationAssurancePolicyRequest struct {
	ctx context.Context
	ApiService *ReplicationAssurancePolicyApiService
	replicationAssurancePolicyName string
	updateRequest *UpdateRequest
}

// Update an existing Replication Assurance Policy
func (r ApiUpdateReplicationAssurancePolicyRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateReplicationAssurancePolicyRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateReplicationAssurancePolicyRequest) Execute() (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	return r.ApiService.UpdateReplicationAssurancePolicyExecute(r)
}

/*
UpdateReplicationAssurancePolicy Update an existing Replication Assurance Policy by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param replicationAssurancePolicyName Name of the Replication Assurance Policy to be updated
 @return ApiUpdateReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyApiService) UpdateReplicationAssurancePolicy(ctx context.Context, replicationAssurancePolicyName string) ApiUpdateReplicationAssurancePolicyRequest {
	return ApiUpdateReplicationAssurancePolicyRequest{
		ApiService: a,
		ctx: ctx,
		replicationAssurancePolicyName: replicationAssurancePolicyName,
	}
}

// Execute executes the request
//  @return ReplicationAssurancePolicyResponse
func (a *ReplicationAssurancePolicyApiService) UpdateReplicationAssurancePolicyExecute(r ApiUpdateReplicationAssurancePolicyRequest) (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReplicationAssurancePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyApiService.UpdateReplicationAssurancePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies/{replication-assurance-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-assurance-policy-name"+"}", url.PathEscape(parameterToString(r.replicationAssurancePolicyName, "")), -1)

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
