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

// VaultAuthenticationMethodApiService VaultAuthenticationMethodApi service
type VaultAuthenticationMethodApiService service

type ApiAddVaultAuthenticationMethodRequest struct {
	ctx                                 context.Context
	ApiService                          *VaultAuthenticationMethodApiService
	addVaultAuthenticationMethodRequest *AddVaultAuthenticationMethodRequest
}

// Create a new Vault Authentication Method in the config
func (r ApiAddVaultAuthenticationMethodRequest) AddVaultAuthenticationMethodRequest(addVaultAuthenticationMethodRequest AddVaultAuthenticationMethodRequest) ApiAddVaultAuthenticationMethodRequest {
	r.addVaultAuthenticationMethodRequest = &addVaultAuthenticationMethodRequest
	return r
}

func (r ApiAddVaultAuthenticationMethodRequest) Execute() (*AddVaultAuthenticationMethod200Response, *http.Response, error) {
	return r.ApiService.AddVaultAuthenticationMethodExecute(r)
}

/*
AddVaultAuthenticationMethod Add a new Vault Authentication Method to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddVaultAuthenticationMethodRequest
*/
func (a *VaultAuthenticationMethodApiService) AddVaultAuthenticationMethod(ctx context.Context) ApiAddVaultAuthenticationMethodRequest {
	return ApiAddVaultAuthenticationMethodRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddVaultAuthenticationMethod200Response
func (a *VaultAuthenticationMethodApiService) AddVaultAuthenticationMethodExecute(r ApiAddVaultAuthenticationMethodRequest) (*AddVaultAuthenticationMethod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddVaultAuthenticationMethod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VaultAuthenticationMethodApiService.AddVaultAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vault-authentication-methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addVaultAuthenticationMethodRequest == nil {
		return localVarReturnValue, nil, reportError("addVaultAuthenticationMethodRequest is required and must be specified")
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
	localVarPostBody = r.addVaultAuthenticationMethodRequest
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

type ApiDeleteVaultAuthenticationMethodRequest struct {
	ctx                           context.Context
	ApiService                    *VaultAuthenticationMethodApiService
	vaultAuthenticationMethodName string
}

func (r ApiDeleteVaultAuthenticationMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVaultAuthenticationMethodExecute(r)
}

/*
DeleteVaultAuthenticationMethod Delete a Vault Authentication Method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vaultAuthenticationMethodName Name of the Vault Authentication Method
	@return ApiDeleteVaultAuthenticationMethodRequest
*/
func (a *VaultAuthenticationMethodApiService) DeleteVaultAuthenticationMethod(ctx context.Context, vaultAuthenticationMethodName string) ApiDeleteVaultAuthenticationMethodRequest {
	return ApiDeleteVaultAuthenticationMethodRequest{
		ApiService:                    a,
		ctx:                           ctx,
		vaultAuthenticationMethodName: vaultAuthenticationMethodName,
	}
}

// Execute executes the request
func (a *VaultAuthenticationMethodApiService) DeleteVaultAuthenticationMethodExecute(r ApiDeleteVaultAuthenticationMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VaultAuthenticationMethodApiService.DeleteVaultAuthenticationMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vault-authentication-methods/{vault-authentication-method-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"vault-authentication-method-name"+"}", url.PathEscape(parameterValueToString(r.vaultAuthenticationMethodName, "vaultAuthenticationMethodName")), -1)

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

type ApiGetVaultAuthenticationMethodRequest struct {
	ctx                           context.Context
	ApiService                    *VaultAuthenticationMethodApiService
	vaultAuthenticationMethodName string
}

func (r ApiGetVaultAuthenticationMethodRequest) Execute() (*AddVaultAuthenticationMethod200Response, *http.Response, error) {
	return r.ApiService.GetVaultAuthenticationMethodExecute(r)
}

/*
GetVaultAuthenticationMethod Returns a single Vault Authentication Method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vaultAuthenticationMethodName Name of the Vault Authentication Method
	@return ApiGetVaultAuthenticationMethodRequest
*/
func (a *VaultAuthenticationMethodApiService) GetVaultAuthenticationMethod(ctx context.Context, vaultAuthenticationMethodName string) ApiGetVaultAuthenticationMethodRequest {
	return ApiGetVaultAuthenticationMethodRequest{
		ApiService:                    a,
		ctx:                           ctx,
		vaultAuthenticationMethodName: vaultAuthenticationMethodName,
	}
}

// Execute executes the request
//
//	@return AddVaultAuthenticationMethod200Response
func (a *VaultAuthenticationMethodApiService) GetVaultAuthenticationMethodExecute(r ApiGetVaultAuthenticationMethodRequest) (*AddVaultAuthenticationMethod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddVaultAuthenticationMethod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VaultAuthenticationMethodApiService.GetVaultAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vault-authentication-methods/{vault-authentication-method-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"vault-authentication-method-name"+"}", url.PathEscape(parameterValueToString(r.vaultAuthenticationMethodName, "vaultAuthenticationMethodName")), -1)

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

type ApiUpdateVaultAuthenticationMethodRequest struct {
	ctx                           context.Context
	ApiService                    *VaultAuthenticationMethodApiService
	vaultAuthenticationMethodName string
	updateRequest                 *UpdateRequest
}

// Update an existing Vault Authentication Method
func (r ApiUpdateVaultAuthenticationMethodRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateVaultAuthenticationMethodRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateVaultAuthenticationMethodRequest) Execute() (*AddVaultAuthenticationMethod200Response, *http.Response, error) {
	return r.ApiService.UpdateVaultAuthenticationMethodExecute(r)
}

/*
UpdateVaultAuthenticationMethod Update an existing Vault Authentication Method by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param vaultAuthenticationMethodName Name of the Vault Authentication Method
	@return ApiUpdateVaultAuthenticationMethodRequest
*/
func (a *VaultAuthenticationMethodApiService) UpdateVaultAuthenticationMethod(ctx context.Context, vaultAuthenticationMethodName string) ApiUpdateVaultAuthenticationMethodRequest {
	return ApiUpdateVaultAuthenticationMethodRequest{
		ApiService:                    a,
		ctx:                           ctx,
		vaultAuthenticationMethodName: vaultAuthenticationMethodName,
	}
}

// Execute executes the request
//
//	@return AddVaultAuthenticationMethod200Response
func (a *VaultAuthenticationMethodApiService) UpdateVaultAuthenticationMethodExecute(r ApiUpdateVaultAuthenticationMethodRequest) (*AddVaultAuthenticationMethod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddVaultAuthenticationMethod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VaultAuthenticationMethodApiService.UpdateVaultAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vault-authentication-methods/{vault-authentication-method-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"vault-authentication-method-name"+"}", url.PathEscape(parameterValueToString(r.vaultAuthenticationMethodName, "vaultAuthenticationMethodName")), -1)

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
