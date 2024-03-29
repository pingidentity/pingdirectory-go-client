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

// ReplicationAssurancePolicyAPIService ReplicationAssurancePolicyAPI service
type ReplicationAssurancePolicyAPIService service

type ApiAddReplicationAssurancePolicyRequest struct {
	ctx                                  context.Context
	ApiService                           *ReplicationAssurancePolicyAPIService
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
func (a *ReplicationAssurancePolicyAPIService) AddReplicationAssurancePolicy(ctx context.Context) ApiAddReplicationAssurancePolicyRequest {
	return ApiAddReplicationAssurancePolicyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReplicationAssurancePolicyResponse
func (a *ReplicationAssurancePolicyAPIService) AddReplicationAssurancePolicyExecute(r ApiAddReplicationAssurancePolicyRequest) (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationAssurancePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyAPIService.AddReplicationAssurancePolicy")
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

type ApiDeleteReplicationAssurancePolicyRequest struct {
	ctx                            context.Context
	ApiService                     *ReplicationAssurancePolicyAPIService
	replicationAssurancePolicyName string
}

func (r ApiDeleteReplicationAssurancePolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteReplicationAssurancePolicyExecute(r)
}

/*
DeleteReplicationAssurancePolicy Delete a Replication Assurance Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param replicationAssurancePolicyName Name of the Replication Assurance Policy
	@return ApiDeleteReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyAPIService) DeleteReplicationAssurancePolicy(ctx context.Context, replicationAssurancePolicyName string) ApiDeleteReplicationAssurancePolicyRequest {
	return ApiDeleteReplicationAssurancePolicyRequest{
		ApiService:                     a,
		ctx:                            ctx,
		replicationAssurancePolicyName: replicationAssurancePolicyName,
	}
}

// Execute executes the request
func (a *ReplicationAssurancePolicyAPIService) DeleteReplicationAssurancePolicyExecute(r ApiDeleteReplicationAssurancePolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyAPIService.DeleteReplicationAssurancePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies/{replication-assurance-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-assurance-policy-name"+"}", url.PathEscape(parameterValueToString(r.replicationAssurancePolicyName, "replicationAssurancePolicyName")), -1)

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

type ApiGetReplicationAssurancePolicyRequest struct {
	ctx                            context.Context
	ApiService                     *ReplicationAssurancePolicyAPIService
	replicationAssurancePolicyName string
}

func (r ApiGetReplicationAssurancePolicyRequest) Execute() (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	return r.ApiService.GetReplicationAssurancePolicyExecute(r)
}

/*
GetReplicationAssurancePolicy Returns a single Replication Assurance Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param replicationAssurancePolicyName Name of the Replication Assurance Policy
	@return ApiGetReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyAPIService) GetReplicationAssurancePolicy(ctx context.Context, replicationAssurancePolicyName string) ApiGetReplicationAssurancePolicyRequest {
	return ApiGetReplicationAssurancePolicyRequest{
		ApiService:                     a,
		ctx:                            ctx,
		replicationAssurancePolicyName: replicationAssurancePolicyName,
	}
}

// Execute executes the request
//
//	@return ReplicationAssurancePolicyResponse
func (a *ReplicationAssurancePolicyAPIService) GetReplicationAssurancePolicyExecute(r ApiGetReplicationAssurancePolicyRequest) (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationAssurancePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyAPIService.GetReplicationAssurancePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies/{replication-assurance-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-assurance-policy-name"+"}", url.PathEscape(parameterValueToString(r.replicationAssurancePolicyName, "replicationAssurancePolicyName")), -1)

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

type ApiListReplicationAssurancePoliciesRequest struct {
	ctx        context.Context
	ApiService *ReplicationAssurancePolicyAPIService
	filter     *string
}

// SCIM filter
func (r ApiListReplicationAssurancePoliciesRequest) Filter(filter string) ApiListReplicationAssurancePoliciesRequest {
	r.filter = &filter
	return r
}

func (r ApiListReplicationAssurancePoliciesRequest) Execute() (*ReplicationAssurancePolicyListResponse, *http.Response, error) {
	return r.ApiService.ListReplicationAssurancePoliciesExecute(r)
}

/*
ListReplicationAssurancePolicies Returns a list of all Replication Assurance Policy objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListReplicationAssurancePoliciesRequest
*/
func (a *ReplicationAssurancePolicyAPIService) ListReplicationAssurancePolicies(ctx context.Context) ApiListReplicationAssurancePoliciesRequest {
	return ApiListReplicationAssurancePoliciesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReplicationAssurancePolicyListResponse
func (a *ReplicationAssurancePolicyAPIService) ListReplicationAssurancePoliciesExecute(r ApiListReplicationAssurancePoliciesRequest) (*ReplicationAssurancePolicyListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationAssurancePolicyListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyAPIService.ListReplicationAssurancePolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
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

type ApiUpdateReplicationAssurancePolicyRequest struct {
	ctx                            context.Context
	ApiService                     *ReplicationAssurancePolicyAPIService
	replicationAssurancePolicyName string
	updateRequest                  *UpdateRequest
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
	@param replicationAssurancePolicyName Name of the Replication Assurance Policy
	@return ApiUpdateReplicationAssurancePolicyRequest
*/
func (a *ReplicationAssurancePolicyAPIService) UpdateReplicationAssurancePolicy(ctx context.Context, replicationAssurancePolicyName string) ApiUpdateReplicationAssurancePolicyRequest {
	return ApiUpdateReplicationAssurancePolicyRequest{
		ApiService:                     a,
		ctx:                            ctx,
		replicationAssurancePolicyName: replicationAssurancePolicyName,
	}
}

// Execute executes the request
//
//	@return ReplicationAssurancePolicyResponse
func (a *ReplicationAssurancePolicyAPIService) UpdateReplicationAssurancePolicyExecute(r ApiUpdateReplicationAssurancePolicyRequest) (*ReplicationAssurancePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReplicationAssurancePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplicationAssurancePolicyAPIService.UpdateReplicationAssurancePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replication-assurance-policies/{replication-assurance-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"replication-assurance-policy-name"+"}", url.PathEscape(parameterValueToString(r.replicationAssurancePolicyName, "replicationAssurancePolicyName")), -1)

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
