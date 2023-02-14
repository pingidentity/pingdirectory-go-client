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

// TrustManagerProviderApiService TrustManagerProviderApi service
type TrustManagerProviderApiService service

type ApiAddTrustManagerProviderRequest struct {
	ctx                            context.Context
	ApiService                     *TrustManagerProviderApiService
	addTrustManagerProviderRequest *AddTrustManagerProviderRequest
}

// Create a new Trust Manager Provider in the config
func (r ApiAddTrustManagerProviderRequest) AddTrustManagerProviderRequest(addTrustManagerProviderRequest AddTrustManagerProviderRequest) ApiAddTrustManagerProviderRequest {
	r.addTrustManagerProviderRequest = &addTrustManagerProviderRequest
	return r
}

func (r ApiAddTrustManagerProviderRequest) Execute() (*AddTrustManagerProvider200Response, *http.Response, error) {
	return r.ApiService.AddTrustManagerProviderExecute(r)
}

/*
AddTrustManagerProvider Add a new Trust Manager Provider to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddTrustManagerProviderRequest
*/
func (a *TrustManagerProviderApiService) AddTrustManagerProvider(ctx context.Context) ApiAddTrustManagerProviderRequest {
	return ApiAddTrustManagerProviderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddTrustManagerProvider200Response
func (a *TrustManagerProviderApiService) AddTrustManagerProviderExecute(r ApiAddTrustManagerProviderRequest) (*AddTrustManagerProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddTrustManagerProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustManagerProviderApiService.AddTrustManagerProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trust-manager-providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addTrustManagerProviderRequest == nil {
		return localVarReturnValue, nil, reportError("addTrustManagerProviderRequest is required and must be specified")
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
	localVarPostBody = r.addTrustManagerProviderRequest
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

type ApiDeleteTrustManagerProviderRequest struct {
	ctx                      context.Context
	ApiService               *TrustManagerProviderApiService
	trustManagerProviderName string
}

func (r ApiDeleteTrustManagerProviderRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTrustManagerProviderExecute(r)
}

/*
DeleteTrustManagerProvider Delete a Trust Manager Provider

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trustManagerProviderName Name of the Trust Manager Provider
	@return ApiDeleteTrustManagerProviderRequest
*/
func (a *TrustManagerProviderApiService) DeleteTrustManagerProvider(ctx context.Context, trustManagerProviderName string) ApiDeleteTrustManagerProviderRequest {
	return ApiDeleteTrustManagerProviderRequest{
		ApiService:               a,
		ctx:                      ctx,
		trustManagerProviderName: trustManagerProviderName,
	}
}

// Execute executes the request
func (a *TrustManagerProviderApiService) DeleteTrustManagerProviderExecute(r ApiDeleteTrustManagerProviderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustManagerProviderApiService.DeleteTrustManagerProvider")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trust-manager-providers/{trust-manager-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"trust-manager-provider-name"+"}", url.PathEscape(parameterToString(r.trustManagerProviderName, "")), -1)

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

type ApiGetTrustManagerProviderRequest struct {
	ctx                      context.Context
	ApiService               *TrustManagerProviderApiService
	trustManagerProviderName string
}

func (r ApiGetTrustManagerProviderRequest) Execute() (*AddTrustManagerProvider200Response, *http.Response, error) {
	return r.ApiService.GetTrustManagerProviderExecute(r)
}

/*
GetTrustManagerProvider Returns a single Trust Manager Provider

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trustManagerProviderName Name of the Trust Manager Provider
	@return ApiGetTrustManagerProviderRequest
*/
func (a *TrustManagerProviderApiService) GetTrustManagerProvider(ctx context.Context, trustManagerProviderName string) ApiGetTrustManagerProviderRequest {
	return ApiGetTrustManagerProviderRequest{
		ApiService:               a,
		ctx:                      ctx,
		trustManagerProviderName: trustManagerProviderName,
	}
}

// Execute executes the request
//
//	@return AddTrustManagerProvider200Response
func (a *TrustManagerProviderApiService) GetTrustManagerProviderExecute(r ApiGetTrustManagerProviderRequest) (*AddTrustManagerProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddTrustManagerProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustManagerProviderApiService.GetTrustManagerProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trust-manager-providers/{trust-manager-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"trust-manager-provider-name"+"}", url.PathEscape(parameterToString(r.trustManagerProviderName, "")), -1)

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

type ApiUpdateTrustManagerProviderRequest struct {
	ctx                      context.Context
	ApiService               *TrustManagerProviderApiService
	trustManagerProviderName string
	updateRequest            *UpdateRequest
}

// Update an existing Trust Manager Provider
func (r ApiUpdateTrustManagerProviderRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateTrustManagerProviderRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateTrustManagerProviderRequest) Execute() (*AddTrustManagerProvider200Response, *http.Response, error) {
	return r.ApiService.UpdateTrustManagerProviderExecute(r)
}

/*
UpdateTrustManagerProvider Update an existing Trust Manager Provider by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trustManagerProviderName Name of the Trust Manager Provider
	@return ApiUpdateTrustManagerProviderRequest
*/
func (a *TrustManagerProviderApiService) UpdateTrustManagerProvider(ctx context.Context, trustManagerProviderName string) ApiUpdateTrustManagerProviderRequest {
	return ApiUpdateTrustManagerProviderRequest{
		ApiService:               a,
		ctx:                      ctx,
		trustManagerProviderName: trustManagerProviderName,
	}
}

// Execute executes the request
//
//	@return AddTrustManagerProvider200Response
func (a *TrustManagerProviderApiService) UpdateTrustManagerProviderExecute(r ApiUpdateTrustManagerProviderRequest) (*AddTrustManagerProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddTrustManagerProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustManagerProviderApiService.UpdateTrustManagerProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trust-manager-providers/{trust-manager-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"trust-manager-provider-name"+"}", url.PathEscape(parameterToString(r.trustManagerProviderName, "")), -1)

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
