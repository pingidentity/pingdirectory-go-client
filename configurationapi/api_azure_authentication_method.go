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

// AzureAuthenticationMethodApiService AzureAuthenticationMethodApi service
type AzureAuthenticationMethodApiService service

type ApiAddAzureAuthenticationMethodRequest struct {
	ctx                                 context.Context
	ApiService                          *AzureAuthenticationMethodApiService
	addAzureAuthenticationMethodRequest *AddAzureAuthenticationMethodRequest
}

// Create a new Azure Authentication Method in the config
func (r ApiAddAzureAuthenticationMethodRequest) AddAzureAuthenticationMethodRequest(addAzureAuthenticationMethodRequest AddAzureAuthenticationMethodRequest) ApiAddAzureAuthenticationMethodRequest {
	r.addAzureAuthenticationMethodRequest = &addAzureAuthenticationMethodRequest
	return r
}

func (r ApiAddAzureAuthenticationMethodRequest) Execute() (*AddAzureAuthenticationMethod200Response, *http.Response, error) {
	return r.ApiService.AddAzureAuthenticationMethodExecute(r)
}

/*
AddAzureAuthenticationMethod Add a new Azure Authentication Method to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddAzureAuthenticationMethodRequest
*/
func (a *AzureAuthenticationMethodApiService) AddAzureAuthenticationMethod(ctx context.Context) ApiAddAzureAuthenticationMethodRequest {
	return ApiAddAzureAuthenticationMethodRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddAzureAuthenticationMethod200Response
func (a *AzureAuthenticationMethodApiService) AddAzureAuthenticationMethodExecute(r ApiAddAzureAuthenticationMethodRequest) (*AddAzureAuthenticationMethod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddAzureAuthenticationMethod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AzureAuthenticationMethodApiService.AddAzureAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/azure-authentication-methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addAzureAuthenticationMethodRequest == nil {
		return localVarReturnValue, nil, reportError("addAzureAuthenticationMethodRequest is required and must be specified")
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
	localVarPostBody = r.addAzureAuthenticationMethodRequest
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

type ApiDeleteAzureAuthenticationMethodRequest struct {
	ctx                           context.Context
	ApiService                    *AzureAuthenticationMethodApiService
	azureAuthenticationMethodName string
}

func (r ApiDeleteAzureAuthenticationMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAzureAuthenticationMethodExecute(r)
}

/*
DeleteAzureAuthenticationMethod Delete a Azure Authentication Method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param azureAuthenticationMethodName Name of the Azure Authentication Method
	@return ApiDeleteAzureAuthenticationMethodRequest
*/
func (a *AzureAuthenticationMethodApiService) DeleteAzureAuthenticationMethod(ctx context.Context, azureAuthenticationMethodName string) ApiDeleteAzureAuthenticationMethodRequest {
	return ApiDeleteAzureAuthenticationMethodRequest{
		ApiService:                    a,
		ctx:                           ctx,
		azureAuthenticationMethodName: azureAuthenticationMethodName,
	}
}

// Execute executes the request
func (a *AzureAuthenticationMethodApiService) DeleteAzureAuthenticationMethodExecute(r ApiDeleteAzureAuthenticationMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AzureAuthenticationMethodApiService.DeleteAzureAuthenticationMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/azure-authentication-methods/{azure-authentication-method-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"azure-authentication-method-name"+"}", url.PathEscape(parameterValueToString(r.azureAuthenticationMethodName, "azureAuthenticationMethodName")), -1)

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

type ApiGetAzureAuthenticationMethodRequest struct {
	ctx                           context.Context
	ApiService                    *AzureAuthenticationMethodApiService
	azureAuthenticationMethodName string
}

func (r ApiGetAzureAuthenticationMethodRequest) Execute() (*AddAzureAuthenticationMethod200Response, *http.Response, error) {
	return r.ApiService.GetAzureAuthenticationMethodExecute(r)
}

/*
GetAzureAuthenticationMethod Returns a single Azure Authentication Method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param azureAuthenticationMethodName Name of the Azure Authentication Method
	@return ApiGetAzureAuthenticationMethodRequest
*/
func (a *AzureAuthenticationMethodApiService) GetAzureAuthenticationMethod(ctx context.Context, azureAuthenticationMethodName string) ApiGetAzureAuthenticationMethodRequest {
	return ApiGetAzureAuthenticationMethodRequest{
		ApiService:                    a,
		ctx:                           ctx,
		azureAuthenticationMethodName: azureAuthenticationMethodName,
	}
}

// Execute executes the request
//
//	@return AddAzureAuthenticationMethod200Response
func (a *AzureAuthenticationMethodApiService) GetAzureAuthenticationMethodExecute(r ApiGetAzureAuthenticationMethodRequest) (*AddAzureAuthenticationMethod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddAzureAuthenticationMethod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AzureAuthenticationMethodApiService.GetAzureAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/azure-authentication-methods/{azure-authentication-method-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"azure-authentication-method-name"+"}", url.PathEscape(parameterValueToString(r.azureAuthenticationMethodName, "azureAuthenticationMethodName")), -1)

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

type ApiUpdateAzureAuthenticationMethodRequest struct {
	ctx                           context.Context
	ApiService                    *AzureAuthenticationMethodApiService
	azureAuthenticationMethodName string
	updateRequest                 *UpdateRequest
}

// Update an existing Azure Authentication Method
func (r ApiUpdateAzureAuthenticationMethodRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateAzureAuthenticationMethodRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateAzureAuthenticationMethodRequest) Execute() (*AddAzureAuthenticationMethod200Response, *http.Response, error) {
	return r.ApiService.UpdateAzureAuthenticationMethodExecute(r)
}

/*
UpdateAzureAuthenticationMethod Update an existing Azure Authentication Method by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param azureAuthenticationMethodName Name of the Azure Authentication Method
	@return ApiUpdateAzureAuthenticationMethodRequest
*/
func (a *AzureAuthenticationMethodApiService) UpdateAzureAuthenticationMethod(ctx context.Context, azureAuthenticationMethodName string) ApiUpdateAzureAuthenticationMethodRequest {
	return ApiUpdateAzureAuthenticationMethodRequest{
		ApiService:                    a,
		ctx:                           ctx,
		azureAuthenticationMethodName: azureAuthenticationMethodName,
	}
}

// Execute executes the request
//
//	@return AddAzureAuthenticationMethod200Response
func (a *AzureAuthenticationMethodApiService) UpdateAzureAuthenticationMethodExecute(r ApiUpdateAzureAuthenticationMethodRequest) (*AddAzureAuthenticationMethod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddAzureAuthenticationMethod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AzureAuthenticationMethodApiService.UpdateAzureAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/azure-authentication-methods/{azure-authentication-method-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"azure-authentication-method-name"+"}", url.PathEscape(parameterValueToString(r.azureAuthenticationMethodName, "azureAuthenticationMethodName")), -1)

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
