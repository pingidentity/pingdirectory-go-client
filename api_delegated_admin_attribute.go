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

// DelegatedAdminAttributeApiService DelegatedAdminAttributeApi service
type DelegatedAdminAttributeApiService service

type ApiAddDelegatedAdminAttributeRequest struct {
	ctx                               context.Context
	ApiService                        *DelegatedAdminAttributeApiService
	addDelegatedAdminAttributeRequest *AddDelegatedAdminAttributeRequest
}

// Create a new Delegated Admin Attribute in the config
func (r ApiAddDelegatedAdminAttributeRequest) AddDelegatedAdminAttributeRequest(addDelegatedAdminAttributeRequest AddDelegatedAdminAttributeRequest) ApiAddDelegatedAdminAttributeRequest {
	r.addDelegatedAdminAttributeRequest = &addDelegatedAdminAttributeRequest
	return r
}

func (r ApiAddDelegatedAdminAttributeRequest) Execute() (*AddDelegatedAdminAttribute200Response, *http.Response, error) {
	return r.ApiService.AddDelegatedAdminAttributeExecute(r)
}

/*
AddDelegatedAdminAttribute Add a new Delegated Admin Attribute to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddDelegatedAdminAttributeRequest
*/
func (a *DelegatedAdminAttributeApiService) AddDelegatedAdminAttribute(ctx context.Context) ApiAddDelegatedAdminAttributeRequest {
	return ApiAddDelegatedAdminAttributeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddDelegatedAdminAttribute200Response
func (a *DelegatedAdminAttributeApiService) AddDelegatedAdminAttributeExecute(r ApiAddDelegatedAdminAttributeRequest) (*AddDelegatedAdminAttribute200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddDelegatedAdminAttribute200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminAttributeApiService.AddDelegatedAdminAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-attributes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addDelegatedAdminAttributeRequest == nil {
		return localVarReturnValue, nil, reportError("addDelegatedAdminAttributeRequest is required and must be specified")
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
	localVarPostBody = r.addDelegatedAdminAttributeRequest
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

type ApiDeleteDelegatedAdminAttributeRequest struct {
	ctx                         context.Context
	ApiService                  *DelegatedAdminAttributeApiService
	delegatedAdminAttributeName string
}

func (r ApiDeleteDelegatedAdminAttributeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDelegatedAdminAttributeExecute(r)
}

/*
DeleteDelegatedAdminAttribute Delete a Delegated Admin Attribute

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param delegatedAdminAttributeName Name of the Delegated Admin Attribute to be deleted
	@return ApiDeleteDelegatedAdminAttributeRequest
*/
func (a *DelegatedAdminAttributeApiService) DeleteDelegatedAdminAttribute(ctx context.Context, delegatedAdminAttributeName string) ApiDeleteDelegatedAdminAttributeRequest {
	return ApiDeleteDelegatedAdminAttributeRequest{
		ApiService:                  a,
		ctx:                         ctx,
		delegatedAdminAttributeName: delegatedAdminAttributeName,
	}
}

// Execute executes the request
func (a *DelegatedAdminAttributeApiService) DeleteDelegatedAdminAttributeExecute(r ApiDeleteDelegatedAdminAttributeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminAttributeApiService.DeleteDelegatedAdminAttribute")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-attributes/{delegated-admin-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"delegated-admin-attribute-name"+"}", url.PathEscape(parameterToString(r.delegatedAdminAttributeName, "")), -1)

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

type ApiGetDelegatedAdminAttributeRequest struct {
	ctx                         context.Context
	ApiService                  *DelegatedAdminAttributeApiService
	delegatedAdminAttributeName string
}

func (r ApiGetDelegatedAdminAttributeRequest) Execute() (*AddDelegatedAdminAttribute200Response, *http.Response, error) {
	return r.ApiService.GetDelegatedAdminAttributeExecute(r)
}

/*
GetDelegatedAdminAttribute Returns a single Delegated Admin Attribute

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param delegatedAdminAttributeName Name of the Delegated Admin Attribute to be read
	@return ApiGetDelegatedAdminAttributeRequest
*/
func (a *DelegatedAdminAttributeApiService) GetDelegatedAdminAttribute(ctx context.Context, delegatedAdminAttributeName string) ApiGetDelegatedAdminAttributeRequest {
	return ApiGetDelegatedAdminAttributeRequest{
		ApiService:                  a,
		ctx:                         ctx,
		delegatedAdminAttributeName: delegatedAdminAttributeName,
	}
}

// Execute executes the request
//
//	@return AddDelegatedAdminAttribute200Response
func (a *DelegatedAdminAttributeApiService) GetDelegatedAdminAttributeExecute(r ApiGetDelegatedAdminAttributeRequest) (*AddDelegatedAdminAttribute200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddDelegatedAdminAttribute200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminAttributeApiService.GetDelegatedAdminAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-attributes/{delegated-admin-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"delegated-admin-attribute-name"+"}", url.PathEscape(parameterToString(r.delegatedAdminAttributeName, "")), -1)

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

type ApiUpdateDelegatedAdminAttributeRequest struct {
	ctx                         context.Context
	ApiService                  *DelegatedAdminAttributeApiService
	delegatedAdminAttributeName string
	updateRequest               *UpdateRequest
}

// Update an existing Delegated Admin Attribute
func (r ApiUpdateDelegatedAdminAttributeRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateDelegatedAdminAttributeRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateDelegatedAdminAttributeRequest) Execute() (*AddDelegatedAdminAttribute200Response, *http.Response, error) {
	return r.ApiService.UpdateDelegatedAdminAttributeExecute(r)
}

/*
UpdateDelegatedAdminAttribute Update an existing Delegated Admin Attribute by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param delegatedAdminAttributeName Name of the Delegated Admin Attribute to be updated
	@return ApiUpdateDelegatedAdminAttributeRequest
*/
func (a *DelegatedAdminAttributeApiService) UpdateDelegatedAdminAttribute(ctx context.Context, delegatedAdminAttributeName string) ApiUpdateDelegatedAdminAttributeRequest {
	return ApiUpdateDelegatedAdminAttributeRequest{
		ApiService:                  a,
		ctx:                         ctx,
		delegatedAdminAttributeName: delegatedAdminAttributeName,
	}
}

// Execute executes the request
//
//	@return AddDelegatedAdminAttribute200Response
func (a *DelegatedAdminAttributeApiService) UpdateDelegatedAdminAttributeExecute(r ApiUpdateDelegatedAdminAttributeRequest) (*AddDelegatedAdminAttribute200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddDelegatedAdminAttribute200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAdminAttributeApiService.UpdateDelegatedAdminAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delegated-admin-attributes/{delegated-admin-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"delegated-admin-attribute-name"+"}", url.PathEscape(parameterToString(r.delegatedAdminAttributeName, "")), -1)

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
