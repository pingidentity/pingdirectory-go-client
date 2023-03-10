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

// DelegatedAdminRightsApiService DelegatedAdminRightsApi service
type DelegatedAdminRightsApiService service

type ApiAddDelegatedAdminRightsRequest struct {
	ctx                            context.Context
	ApiService                     *DelegatedAdminRightsApiService
	addDelegatedAdminRightsRequest *AddDelegatedAdminRightsRequest
}

// Create a new Delegated Admin Rights in the config
func (r ApiAddDelegatedAdminRightsRequest) AddDelegatedAdminRightsRequest(addDelegatedAdminRightsRequest AddDelegatedAdminRightsRequest) ApiAddDelegatedAdminRightsRequest {
	r.addDelegatedAdminRightsRequest = &addDelegatedAdminRightsRequest
	return r
}

func (r ApiAddDelegatedAdminRightsRequest) Execute() (*DelegatedAdminRightsResponse, *http.Response, error) {
	return r.ApiService.AddDelegatedAdminRightsExecute(r)
}

/*
AddDelegatedAdminRights Add a new Delegated Admin Rights to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddDelegatedAdminRightsRequest
*/
func (a *DelegatedAdminRightsApiService) AddDelegatedAdminRights(ctx context.Context) ApiAddDelegatedAdminRightsRequest {
	return ApiAddDelegatedAdminRightsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DelegatedAdminRightsResponse
func (a *DelegatedAdminRightsApiService) AddDelegatedAdminRightsExecute(r ApiAddDelegatedAdminRightsRequest) (*DelegatedAdminRightsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DelegatedAdminRightsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminRightsApiService.AddDelegatedAdminRights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-rights"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addDelegatedAdminRightsRequest == nil {
		return localVarReturnValue, nil, reportError("addDelegatedAdminRightsRequest is required and must be specified")
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
	localVarPostBody = r.addDelegatedAdminRightsRequest
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

type ApiDeleteDelegatedAdminRightsRequest struct {
	ctx                      context.Context
	ApiService               *DelegatedAdminRightsApiService
	delegatedAdminRightsName string
}

func (r ApiDeleteDelegatedAdminRightsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDelegatedAdminRightsExecute(r)
}

/*
DeleteDelegatedAdminRights Delete a Delegated Admin Rights

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param delegatedAdminRightsName Name of the Delegated Admin Rights
	@return ApiDeleteDelegatedAdminRightsRequest
*/
func (a *DelegatedAdminRightsApiService) DeleteDelegatedAdminRights(ctx context.Context, delegatedAdminRightsName string) ApiDeleteDelegatedAdminRightsRequest {
	return ApiDeleteDelegatedAdminRightsRequest{
		ApiService:               a,
		ctx:                      ctx,
		delegatedAdminRightsName: delegatedAdminRightsName,
	}
}

// Execute executes the request
func (a *DelegatedAdminRightsApiService) DeleteDelegatedAdminRightsExecute(r ApiDeleteDelegatedAdminRightsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminRightsApiService.DeleteDelegatedAdminRights")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-rights/{delegated-admin-rights-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"delegated-admin-rights-name"+"}", url.PathEscape(parameterValueToString(r.delegatedAdminRightsName, "delegatedAdminRightsName")), -1)

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

type ApiGetDelegatedAdminRightsRequest struct {
	ctx                      context.Context
	ApiService               *DelegatedAdminRightsApiService
	delegatedAdminRightsName string
}

func (r ApiGetDelegatedAdminRightsRequest) Execute() (*DelegatedAdminRightsResponse, *http.Response, error) {
	return r.ApiService.GetDelegatedAdminRightsExecute(r)
}

/*
GetDelegatedAdminRights Returns a single Delegated Admin Rights

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param delegatedAdminRightsName Name of the Delegated Admin Rights
	@return ApiGetDelegatedAdminRightsRequest
*/
func (a *DelegatedAdminRightsApiService) GetDelegatedAdminRights(ctx context.Context, delegatedAdminRightsName string) ApiGetDelegatedAdminRightsRequest {
	return ApiGetDelegatedAdminRightsRequest{
		ApiService:               a,
		ctx:                      ctx,
		delegatedAdminRightsName: delegatedAdminRightsName,
	}
}

// Execute executes the request
//
//	@return DelegatedAdminRightsResponse
func (a *DelegatedAdminRightsApiService) GetDelegatedAdminRightsExecute(r ApiGetDelegatedAdminRightsRequest) (*DelegatedAdminRightsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DelegatedAdminRightsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminRightsApiService.GetDelegatedAdminRights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-rights/{delegated-admin-rights-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"delegated-admin-rights-name"+"}", url.PathEscape(parameterValueToString(r.delegatedAdminRightsName, "delegatedAdminRightsName")), -1)

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

type ApiUpdateDelegatedAdminRightsRequest struct {
	ctx                      context.Context
	ApiService               *DelegatedAdminRightsApiService
	delegatedAdminRightsName string
	updateRequest            *UpdateRequest
}

// Update an existing Delegated Admin Rights
func (r ApiUpdateDelegatedAdminRightsRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateDelegatedAdminRightsRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateDelegatedAdminRightsRequest) Execute() (*DelegatedAdminRightsResponse, *http.Response, error) {
	return r.ApiService.UpdateDelegatedAdminRightsExecute(r)
}

/*
UpdateDelegatedAdminRights Update an existing Delegated Admin Rights by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param delegatedAdminRightsName Name of the Delegated Admin Rights
	@return ApiUpdateDelegatedAdminRightsRequest
*/
func (a *DelegatedAdminRightsApiService) UpdateDelegatedAdminRights(ctx context.Context, delegatedAdminRightsName string) ApiUpdateDelegatedAdminRightsRequest {
	return ApiUpdateDelegatedAdminRightsRequest{
		ApiService:               a,
		ctx:                      ctx,
		delegatedAdminRightsName: delegatedAdminRightsName,
	}
}

// Execute executes the request
//
//	@return DelegatedAdminRightsResponse
func (a *DelegatedAdminRightsApiService) UpdateDelegatedAdminRightsExecute(r ApiUpdateDelegatedAdminRightsRequest) (*DelegatedAdminRightsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DelegatedAdminRightsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminRightsApiService.UpdateDelegatedAdminRights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-rights/{delegated-admin-rights-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"delegated-admin-rights-name"+"}", url.PathEscape(parameterValueToString(r.delegatedAdminRightsName, "delegatedAdminRightsName")), -1)

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
