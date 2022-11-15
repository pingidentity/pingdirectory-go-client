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


// KeyManagerProviderApiService KeyManagerProviderApi service
type KeyManagerProviderApiService service

type ApiAddKeyManagerProviderRequest struct {
	ctx context.Context
	ApiService *KeyManagerProviderApiService
	addKeyManagerProviderRequest *AddKeyManagerProviderRequest
}

// Create a new Key Manager Provider in the config
func (r ApiAddKeyManagerProviderRequest) AddKeyManagerProviderRequest(addKeyManagerProviderRequest AddKeyManagerProviderRequest) ApiAddKeyManagerProviderRequest {
	r.addKeyManagerProviderRequest = &addKeyManagerProviderRequest
	return r
}

func (r ApiAddKeyManagerProviderRequest) Execute() (*AddKeyManagerProvider200Response, *http.Response, error) {
	return r.ApiService.AddKeyManagerProviderExecute(r)
}

/*
AddKeyManagerProvider Add a new Key Manager Provider to the config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddKeyManagerProviderRequest
*/
func (a *KeyManagerProviderApiService) AddKeyManagerProvider(ctx context.Context) ApiAddKeyManagerProviderRequest {
	return ApiAddKeyManagerProviderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddKeyManagerProvider200Response
func (a *KeyManagerProviderApiService) AddKeyManagerProviderExecute(r ApiAddKeyManagerProviderRequest) (*AddKeyManagerProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddKeyManagerProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagerProviderApiService.AddKeyManagerProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-manager-providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addKeyManagerProviderRequest == nil {
		return localVarReturnValue, nil, reportError("addKeyManagerProviderRequest is required and must be specified")
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
	localVarPostBody = r.addKeyManagerProviderRequest
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

type ApiDeleteKeyManagerProviderRequest struct {
	ctx context.Context
	ApiService *KeyManagerProviderApiService
	keyManagerProviderName string
}

func (r ApiDeleteKeyManagerProviderRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteKeyManagerProviderExecute(r)
}

/*
DeleteKeyManagerProvider Delete a Key Manager Provider

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyManagerProviderName Name of the Key Manager Provider to be deleted
 @return ApiDeleteKeyManagerProviderRequest
*/
func (a *KeyManagerProviderApiService) DeleteKeyManagerProvider(ctx context.Context, keyManagerProviderName string) ApiDeleteKeyManagerProviderRequest {
	return ApiDeleteKeyManagerProviderRequest{
		ApiService: a,
		ctx: ctx,
		keyManagerProviderName: keyManagerProviderName,
	}
}

// Execute executes the request
func (a *KeyManagerProviderApiService) DeleteKeyManagerProviderExecute(r ApiDeleteKeyManagerProviderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagerProviderApiService.DeleteKeyManagerProvider")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-manager-providers/{key-manager-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"key-manager-provider-name"+"}", url.PathEscape(parameterToString(r.keyManagerProviderName, "")), -1)

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

type ApiGetKeyManagerProviderRequest struct {
	ctx context.Context
	ApiService *KeyManagerProviderApiService
	keyManagerProviderName string
}

func (r ApiGetKeyManagerProviderRequest) Execute() (*AddKeyManagerProvider200Response, *http.Response, error) {
	return r.ApiService.GetKeyManagerProviderExecute(r)
}

/*
GetKeyManagerProvider Returns a single Key Manager Provider

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyManagerProviderName Name of the Key Manager Provider to be read
 @return ApiGetKeyManagerProviderRequest
*/
func (a *KeyManagerProviderApiService) GetKeyManagerProvider(ctx context.Context, keyManagerProviderName string) ApiGetKeyManagerProviderRequest {
	return ApiGetKeyManagerProviderRequest{
		ApiService: a,
		ctx: ctx,
		keyManagerProviderName: keyManagerProviderName,
	}
}

// Execute executes the request
//  @return AddKeyManagerProvider200Response
func (a *KeyManagerProviderApiService) GetKeyManagerProviderExecute(r ApiGetKeyManagerProviderRequest) (*AddKeyManagerProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddKeyManagerProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagerProviderApiService.GetKeyManagerProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-manager-providers/{key-manager-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"key-manager-provider-name"+"}", url.PathEscape(parameterToString(r.keyManagerProviderName, "")), -1)

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

type ApiUpdateKeyManagerProviderRequest struct {
	ctx context.Context
	ApiService *KeyManagerProviderApiService
	keyManagerProviderName string
	updateRequest *UpdateRequest
}

// Update an existing Key Manager Provider
func (r ApiUpdateKeyManagerProviderRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateKeyManagerProviderRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateKeyManagerProviderRequest) Execute() (*AddKeyManagerProvider200Response, *http.Response, error) {
	return r.ApiService.UpdateKeyManagerProviderExecute(r)
}

/*
UpdateKeyManagerProvider Update an existing Key Manager Provider by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyManagerProviderName Name of the Key Manager Provider to be updated
 @return ApiUpdateKeyManagerProviderRequest
*/
func (a *KeyManagerProviderApiService) UpdateKeyManagerProvider(ctx context.Context, keyManagerProviderName string) ApiUpdateKeyManagerProviderRequest {
	return ApiUpdateKeyManagerProviderRequest{
		ApiService: a,
		ctx: ctx,
		keyManagerProviderName: keyManagerProviderName,
	}
}

// Execute executes the request
//  @return AddKeyManagerProvider200Response
func (a *KeyManagerProviderApiService) UpdateKeyManagerProviderExecute(r ApiUpdateKeyManagerProviderRequest) (*AddKeyManagerProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddKeyManagerProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyManagerProviderApiService.UpdateKeyManagerProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-manager-providers/{key-manager-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"key-manager-provider-name"+"}", url.PathEscape(parameterToString(r.keyManagerProviderName, "")), -1)

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
