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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// KeyPairApiService KeyPairApi service
type KeyPairApiService service

type ApiAddKeyPairRequest struct {
	ctx               context.Context
	ApiService        *KeyPairApiService
	addKeyPairRequest *AddKeyPairRequest
}

// Create a new Key Pair in the config
func (r ApiAddKeyPairRequest) AddKeyPairRequest(addKeyPairRequest AddKeyPairRequest) ApiAddKeyPairRequest {
	r.addKeyPairRequest = &addKeyPairRequest
	return r
}

func (r ApiAddKeyPairRequest) Execute() (*KeyPairResponse, *http.Response, error) {
	return r.ApiService.AddKeyPairExecute(r)
}

/*
AddKeyPair Add a new Key Pair to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddKeyPairRequest
*/
func (a *KeyPairApiService) AddKeyPair(ctx context.Context) ApiAddKeyPairRequest {
	return ApiAddKeyPairRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KeyPairResponse
func (a *KeyPairApiService) AddKeyPairExecute(r ApiAddKeyPairRequest) (*KeyPairResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeyPairResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyPairApiService.AddKeyPair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-pairs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addKeyPairRequest == nil {
		return localVarReturnValue, nil, reportError("addKeyPairRequest is required and must be specified")
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
	localVarPostBody = r.addKeyPairRequest
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

type ApiDeleteKeyPairRequest struct {
	ctx         context.Context
	ApiService  *KeyPairApiService
	keyPairName string
}

func (r ApiDeleteKeyPairRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteKeyPairExecute(r)
}

/*
DeleteKeyPair Delete a Key Pair

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyPairName Name of the Key Pair
	@return ApiDeleteKeyPairRequest
*/
func (a *KeyPairApiService) DeleteKeyPair(ctx context.Context, keyPairName string) ApiDeleteKeyPairRequest {
	return ApiDeleteKeyPairRequest{
		ApiService:  a,
		ctx:         ctx,
		keyPairName: keyPairName,
	}
}

// Execute executes the request
func (a *KeyPairApiService) DeleteKeyPairExecute(r ApiDeleteKeyPairRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyPairApiService.DeleteKeyPair")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-pairs/{key-pair-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"key-pair-name"+"}", url.PathEscape(parameterToString(r.keyPairName, "")), -1)

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

type ApiGetKeyPairRequest struct {
	ctx         context.Context
	ApiService  *KeyPairApiService
	keyPairName string
}

func (r ApiGetKeyPairRequest) Execute() (*KeyPairResponse, *http.Response, error) {
	return r.ApiService.GetKeyPairExecute(r)
}

/*
GetKeyPair Returns a single Key Pair

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyPairName Name of the Key Pair
	@return ApiGetKeyPairRequest
*/
func (a *KeyPairApiService) GetKeyPair(ctx context.Context, keyPairName string) ApiGetKeyPairRequest {
	return ApiGetKeyPairRequest{
		ApiService:  a,
		ctx:         ctx,
		keyPairName: keyPairName,
	}
}

// Execute executes the request
//
//	@return KeyPairResponse
func (a *KeyPairApiService) GetKeyPairExecute(r ApiGetKeyPairRequest) (*KeyPairResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeyPairResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyPairApiService.GetKeyPair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-pairs/{key-pair-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"key-pair-name"+"}", url.PathEscape(parameterToString(r.keyPairName, "")), -1)

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

type ApiUpdateKeyPairRequest struct {
	ctx           context.Context
	ApiService    *KeyPairApiService
	keyPairName   string
	updateRequest *UpdateRequest
}

// Update an existing Key Pair
func (r ApiUpdateKeyPairRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateKeyPairRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateKeyPairRequest) Execute() (*KeyPairResponse, *http.Response, error) {
	return r.ApiService.UpdateKeyPairExecute(r)
}

/*
UpdateKeyPair Update an existing Key Pair by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyPairName Name of the Key Pair
	@return ApiUpdateKeyPairRequest
*/
func (a *KeyPairApiService) UpdateKeyPair(ctx context.Context, keyPairName string) ApiUpdateKeyPairRequest {
	return ApiUpdateKeyPairRequest{
		ApiService:  a,
		ctx:         ctx,
		keyPairName: keyPairName,
	}
}

// Execute executes the request
//
//	@return KeyPairResponse
func (a *KeyPairApiService) UpdateKeyPairExecute(r ApiUpdateKeyPairRequest) (*KeyPairResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeyPairResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyPairApiService.UpdateKeyPair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key-pairs/{key-pair-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"key-pair-name"+"}", url.PathEscape(parameterToString(r.keyPairName, "")), -1)

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