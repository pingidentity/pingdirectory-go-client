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


// IdTokenValidatorApiService IdTokenValidatorApi service
type IdTokenValidatorApiService service

type ApiAddIdTokenValidatorRequest struct {
	ctx context.Context
	ApiService *IdTokenValidatorApiService
	addIdTokenValidatorRequest *AddIdTokenValidatorRequest
}

// Create a new ID Token Validator in the config
func (r ApiAddIdTokenValidatorRequest) AddIdTokenValidatorRequest(addIdTokenValidatorRequest AddIdTokenValidatorRequest) ApiAddIdTokenValidatorRequest {
	r.addIdTokenValidatorRequest = &addIdTokenValidatorRequest
	return r
}

func (r ApiAddIdTokenValidatorRequest) Execute() (*AddIdTokenValidator200Response, *http.Response, error) {
	return r.ApiService.AddIdTokenValidatorExecute(r)
}

/*
AddIdTokenValidator Add a new ID Token Validator to the config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddIdTokenValidatorRequest
*/
func (a *IdTokenValidatorApiService) AddIdTokenValidator(ctx context.Context) ApiAddIdTokenValidatorRequest {
	return ApiAddIdTokenValidatorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddIdTokenValidator200Response
func (a *IdTokenValidatorApiService) AddIdTokenValidatorExecute(r ApiAddIdTokenValidatorRequest) (*AddIdTokenValidator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddIdTokenValidator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdTokenValidatorApiService.AddIdTokenValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/id-token-validators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addIdTokenValidatorRequest == nil {
		return localVarReturnValue, nil, reportError("addIdTokenValidatorRequest is required and must be specified")
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
	localVarPostBody = r.addIdTokenValidatorRequest
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

type ApiDeleteIdTokenValidatorRequest struct {
	ctx context.Context
	ApiService *IdTokenValidatorApiService
	idTokenValidatorName string
}

func (r ApiDeleteIdTokenValidatorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIdTokenValidatorExecute(r)
}

/*
DeleteIdTokenValidator Delete a ID Token Validator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idTokenValidatorName Name of the ID Token Validator to be deleted
 @return ApiDeleteIdTokenValidatorRequest
*/
func (a *IdTokenValidatorApiService) DeleteIdTokenValidator(ctx context.Context, idTokenValidatorName string) ApiDeleteIdTokenValidatorRequest {
	return ApiDeleteIdTokenValidatorRequest{
		ApiService: a,
		ctx: ctx,
		idTokenValidatorName: idTokenValidatorName,
	}
}

// Execute executes the request
func (a *IdTokenValidatorApiService) DeleteIdTokenValidatorExecute(r ApiDeleteIdTokenValidatorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdTokenValidatorApiService.DeleteIdTokenValidator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/id-token-validators/{id-token-validator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"id-token-validator-name"+"}", url.PathEscape(parameterToString(r.idTokenValidatorName, "")), -1)

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

type ApiGetIdTokenValidatorRequest struct {
	ctx context.Context
	ApiService *IdTokenValidatorApiService
	idTokenValidatorName string
}

func (r ApiGetIdTokenValidatorRequest) Execute() (*AddIdTokenValidator200Response, *http.Response, error) {
	return r.ApiService.GetIdTokenValidatorExecute(r)
}

/*
GetIdTokenValidator Returns a single ID Token Validator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idTokenValidatorName Name of the ID Token Validator to be read
 @return ApiGetIdTokenValidatorRequest
*/
func (a *IdTokenValidatorApiService) GetIdTokenValidator(ctx context.Context, idTokenValidatorName string) ApiGetIdTokenValidatorRequest {
	return ApiGetIdTokenValidatorRequest{
		ApiService: a,
		ctx: ctx,
		idTokenValidatorName: idTokenValidatorName,
	}
}

// Execute executes the request
//  @return AddIdTokenValidator200Response
func (a *IdTokenValidatorApiService) GetIdTokenValidatorExecute(r ApiGetIdTokenValidatorRequest) (*AddIdTokenValidator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddIdTokenValidator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdTokenValidatorApiService.GetIdTokenValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/id-token-validators/{id-token-validator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"id-token-validator-name"+"}", url.PathEscape(parameterToString(r.idTokenValidatorName, "")), -1)

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

type ApiUpdateIdTokenValidatorRequest struct {
	ctx context.Context
	ApiService *IdTokenValidatorApiService
	idTokenValidatorName string
	updateRequest *UpdateRequest
}

// Update an existing ID Token Validator
func (r ApiUpdateIdTokenValidatorRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateIdTokenValidatorRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateIdTokenValidatorRequest) Execute() (*AddIdTokenValidator200Response, *http.Response, error) {
	return r.ApiService.UpdateIdTokenValidatorExecute(r)
}

/*
UpdateIdTokenValidator Update an existing ID Token Validator by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param idTokenValidatorName Name of the ID Token Validator to be updated
 @return ApiUpdateIdTokenValidatorRequest
*/
func (a *IdTokenValidatorApiService) UpdateIdTokenValidator(ctx context.Context, idTokenValidatorName string) ApiUpdateIdTokenValidatorRequest {
	return ApiUpdateIdTokenValidatorRequest{
		ApiService: a,
		ctx: ctx,
		idTokenValidatorName: idTokenValidatorName,
	}
}

// Execute executes the request
//  @return AddIdTokenValidator200Response
func (a *IdTokenValidatorApiService) UpdateIdTokenValidatorExecute(r ApiUpdateIdTokenValidatorRequest) (*AddIdTokenValidator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddIdTokenValidator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdTokenValidatorApiService.UpdateIdTokenValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/id-token-validators/{id-token-validator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"id-token-validator-name"+"}", url.PathEscape(parameterToString(r.idTokenValidatorName, "")), -1)

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
