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


// SoftDeletePolicyApiService SoftDeletePolicyApi service
type SoftDeletePolicyApiService service

type ApiAddSoftDeletePolicyRequest struct {
	ctx context.Context
	ApiService *SoftDeletePolicyApiService
	addSoftDeletePolicyRequest *AddSoftDeletePolicyRequest
}

// Create a new Soft Delete Policy in the config
func (r ApiAddSoftDeletePolicyRequest) AddSoftDeletePolicyRequest(addSoftDeletePolicyRequest AddSoftDeletePolicyRequest) ApiAddSoftDeletePolicyRequest {
	r.addSoftDeletePolicyRequest = &addSoftDeletePolicyRequest
	return r
}

func (r ApiAddSoftDeletePolicyRequest) Execute() (*SoftDeletePolicyResponse, *http.Response, error) {
	return r.ApiService.AddSoftDeletePolicyExecute(r)
}

/*
AddSoftDeletePolicy Add a new Soft Delete Policy to the config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddSoftDeletePolicyRequest
*/
func (a *SoftDeletePolicyApiService) AddSoftDeletePolicy(ctx context.Context) ApiAddSoftDeletePolicyRequest {
	return ApiAddSoftDeletePolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SoftDeletePolicyResponse
func (a *SoftDeletePolicyApiService) AddSoftDeletePolicyExecute(r ApiAddSoftDeletePolicyRequest) (*SoftDeletePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftDeletePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftDeletePolicyApiService.AddSoftDeletePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/soft-delete-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addSoftDeletePolicyRequest == nil {
		return localVarReturnValue, nil, reportError("addSoftDeletePolicyRequest is required and must be specified")
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
	localVarPostBody = r.addSoftDeletePolicyRequest
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

type ApiDeleteSoftDeletePolicyRequest struct {
	ctx context.Context
	ApiService *SoftDeletePolicyApiService
	softDeletePolicyName string
}

func (r ApiDeleteSoftDeletePolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSoftDeletePolicyExecute(r)
}

/*
DeleteSoftDeletePolicy Delete a Soft Delete Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softDeletePolicyName Name of the Soft Delete Policy to be deleted
 @return ApiDeleteSoftDeletePolicyRequest
*/
func (a *SoftDeletePolicyApiService) DeleteSoftDeletePolicy(ctx context.Context, softDeletePolicyName string) ApiDeleteSoftDeletePolicyRequest {
	return ApiDeleteSoftDeletePolicyRequest{
		ApiService: a,
		ctx: ctx,
		softDeletePolicyName: softDeletePolicyName,
	}
}

// Execute executes the request
func (a *SoftDeletePolicyApiService) DeleteSoftDeletePolicyExecute(r ApiDeleteSoftDeletePolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftDeletePolicyApiService.DeleteSoftDeletePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/soft-delete-policies/{soft-delete-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"soft-delete-policy-name"+"}", url.PathEscape(parameterToString(r.softDeletePolicyName, "")), -1)

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

type ApiGetSoftDeletePolicyRequest struct {
	ctx context.Context
	ApiService *SoftDeletePolicyApiService
	softDeletePolicyName string
}

func (r ApiGetSoftDeletePolicyRequest) Execute() (*SoftDeletePolicyResponse, *http.Response, error) {
	return r.ApiService.GetSoftDeletePolicyExecute(r)
}

/*
GetSoftDeletePolicy Returns a single Soft Delete Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softDeletePolicyName Name of the Soft Delete Policy to be read
 @return ApiGetSoftDeletePolicyRequest
*/
func (a *SoftDeletePolicyApiService) GetSoftDeletePolicy(ctx context.Context, softDeletePolicyName string) ApiGetSoftDeletePolicyRequest {
	return ApiGetSoftDeletePolicyRequest{
		ApiService: a,
		ctx: ctx,
		softDeletePolicyName: softDeletePolicyName,
	}
}

// Execute executes the request
//  @return SoftDeletePolicyResponse
func (a *SoftDeletePolicyApiService) GetSoftDeletePolicyExecute(r ApiGetSoftDeletePolicyRequest) (*SoftDeletePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftDeletePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftDeletePolicyApiService.GetSoftDeletePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/soft-delete-policies/{soft-delete-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"soft-delete-policy-name"+"}", url.PathEscape(parameterToString(r.softDeletePolicyName, "")), -1)

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

type ApiUpdateSoftDeletePolicyRequest struct {
	ctx context.Context
	ApiService *SoftDeletePolicyApiService
	softDeletePolicyName string
	updateRequest *UpdateRequest
}

// Update an existing Soft Delete Policy
func (r ApiUpdateSoftDeletePolicyRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateSoftDeletePolicyRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateSoftDeletePolicyRequest) Execute() (*SoftDeletePolicyResponse, *http.Response, error) {
	return r.ApiService.UpdateSoftDeletePolicyExecute(r)
}

/*
UpdateSoftDeletePolicy Update an existing Soft Delete Policy by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softDeletePolicyName Name of the Soft Delete Policy to be updated
 @return ApiUpdateSoftDeletePolicyRequest
*/
func (a *SoftDeletePolicyApiService) UpdateSoftDeletePolicy(ctx context.Context, softDeletePolicyName string) ApiUpdateSoftDeletePolicyRequest {
	return ApiUpdateSoftDeletePolicyRequest{
		ApiService: a,
		ctx: ctx,
		softDeletePolicyName: softDeletePolicyName,
	}
}

// Execute executes the request
//  @return SoftDeletePolicyResponse
func (a *SoftDeletePolicyApiService) UpdateSoftDeletePolicyExecute(r ApiUpdateSoftDeletePolicyRequest) (*SoftDeletePolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftDeletePolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftDeletePolicyApiService.UpdateSoftDeletePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/soft-delete-policies/{soft-delete-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"soft-delete-policy-name"+"}", url.PathEscape(parameterToString(r.softDeletePolicyName, "")), -1)

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
