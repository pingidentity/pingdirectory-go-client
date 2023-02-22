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

// PasswordGeneratorApiService PasswordGeneratorApi service
type PasswordGeneratorApiService service

type ApiAddPasswordGeneratorRequest struct {
	ctx                         context.Context
	ApiService                  *PasswordGeneratorApiService
	addPasswordGeneratorRequest *AddPasswordGeneratorRequest
}

// Create a new Password Generator in the config
func (r ApiAddPasswordGeneratorRequest) AddPasswordGeneratorRequest(addPasswordGeneratorRequest AddPasswordGeneratorRequest) ApiAddPasswordGeneratorRequest {
	r.addPasswordGeneratorRequest = &addPasswordGeneratorRequest
	return r
}

func (r ApiAddPasswordGeneratorRequest) Execute() (*AddPasswordGenerator200Response, *http.Response, error) {
	return r.ApiService.AddPasswordGeneratorExecute(r)
}

/*
AddPasswordGenerator Add a new Password Generator to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPasswordGeneratorRequest
*/
func (a *PasswordGeneratorApiService) AddPasswordGenerator(ctx context.Context) ApiAddPasswordGeneratorRequest {
	return ApiAddPasswordGeneratorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddPasswordGenerator200Response
func (a *PasswordGeneratorApiService) AddPasswordGeneratorExecute(r ApiAddPasswordGeneratorRequest) (*AddPasswordGenerator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordGenerator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordGeneratorApiService.AddPasswordGenerator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-generators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addPasswordGeneratorRequest == nil {
		return localVarReturnValue, nil, reportError("addPasswordGeneratorRequest is required and must be specified")
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
	localVarPostBody = r.addPasswordGeneratorRequest
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

type ApiDeletePasswordGeneratorRequest struct {
	ctx                   context.Context
	ApiService            *PasswordGeneratorApiService
	passwordGeneratorName string
}

func (r ApiDeletePasswordGeneratorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePasswordGeneratorExecute(r)
}

/*
DeletePasswordGenerator Delete a Password Generator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordGeneratorName Name of the Password Generator
	@return ApiDeletePasswordGeneratorRequest
*/
func (a *PasswordGeneratorApiService) DeletePasswordGenerator(ctx context.Context, passwordGeneratorName string) ApiDeletePasswordGeneratorRequest {
	return ApiDeletePasswordGeneratorRequest{
		ApiService:            a,
		ctx:                   ctx,
		passwordGeneratorName: passwordGeneratorName,
	}
}

// Execute executes the request
func (a *PasswordGeneratorApiService) DeletePasswordGeneratorExecute(r ApiDeletePasswordGeneratorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordGeneratorApiService.DeletePasswordGenerator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-generators/{password-generator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-generator-name"+"}", url.PathEscape(parameterToString(r.passwordGeneratorName, "")), -1)

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

type ApiGetPasswordGeneratorRequest struct {
	ctx                   context.Context
	ApiService            *PasswordGeneratorApiService
	passwordGeneratorName string
}

func (r ApiGetPasswordGeneratorRequest) Execute() (*AddPasswordGenerator200Response, *http.Response, error) {
	return r.ApiService.GetPasswordGeneratorExecute(r)
}

/*
GetPasswordGenerator Returns a single Password Generator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordGeneratorName Name of the Password Generator
	@return ApiGetPasswordGeneratorRequest
*/
func (a *PasswordGeneratorApiService) GetPasswordGenerator(ctx context.Context, passwordGeneratorName string) ApiGetPasswordGeneratorRequest {
	return ApiGetPasswordGeneratorRequest{
		ApiService:            a,
		ctx:                   ctx,
		passwordGeneratorName: passwordGeneratorName,
	}
}

// Execute executes the request
//
//	@return AddPasswordGenerator200Response
func (a *PasswordGeneratorApiService) GetPasswordGeneratorExecute(r ApiGetPasswordGeneratorRequest) (*AddPasswordGenerator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordGenerator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordGeneratorApiService.GetPasswordGenerator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-generators/{password-generator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-generator-name"+"}", url.PathEscape(parameterToString(r.passwordGeneratorName, "")), -1)

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

type ApiUpdatePasswordGeneratorRequest struct {
	ctx                   context.Context
	ApiService            *PasswordGeneratorApiService
	passwordGeneratorName string
	updateRequest         *UpdateRequest
}

// Update an existing Password Generator
func (r ApiUpdatePasswordGeneratorRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdatePasswordGeneratorRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdatePasswordGeneratorRequest) Execute() (*AddPasswordGenerator200Response, *http.Response, error) {
	return r.ApiService.UpdatePasswordGeneratorExecute(r)
}

/*
UpdatePasswordGenerator Update an existing Password Generator by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passwordGeneratorName Name of the Password Generator
	@return ApiUpdatePasswordGeneratorRequest
*/
func (a *PasswordGeneratorApiService) UpdatePasswordGenerator(ctx context.Context, passwordGeneratorName string) ApiUpdatePasswordGeneratorRequest {
	return ApiUpdatePasswordGeneratorRequest{
		ApiService:            a,
		ctx:                   ctx,
		passwordGeneratorName: passwordGeneratorName,
	}
}

// Execute executes the request
//
//	@return AddPasswordGenerator200Response
func (a *PasswordGeneratorApiService) UpdatePasswordGeneratorExecute(r ApiUpdatePasswordGeneratorRequest) (*AddPasswordGenerator200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPasswordGenerator200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordGeneratorApiService.UpdatePasswordGenerator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/password-generators/{password-generator-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"password-generator-name"+"}", url.PathEscape(parameterToString(r.passwordGeneratorName, "")), -1)

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