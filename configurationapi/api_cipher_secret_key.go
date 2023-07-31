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

// CipherSecretKeyApiService CipherSecretKeyApi service
type CipherSecretKeyApiService service

type ApiGetCipherSecretKeyRequest struct {
	ctx                 context.Context
	ApiService          *CipherSecretKeyApiService
	cipherSecretKeyName string
	serverInstanceName  string
}

func (r ApiGetCipherSecretKeyRequest) Execute() (*CipherSecretKeyResponse, *http.Response, error) {
	return r.ApiService.GetCipherSecretKeyExecute(r)
}

/*
GetCipherSecretKey Returns a single Cipher Secret Key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cipherSecretKeyName Name of the Cipher Secret Key
	@param serverInstanceName Name of the Server Instance
	@return ApiGetCipherSecretKeyRequest
*/
func (a *CipherSecretKeyApiService) GetCipherSecretKey(ctx context.Context, cipherSecretKeyName string, serverInstanceName string) ApiGetCipherSecretKeyRequest {
	return ApiGetCipherSecretKeyRequest{
		ApiService:          a,
		ctx:                 ctx,
		cipherSecretKeyName: cipherSecretKeyName,
		serverInstanceName:  serverInstanceName,
	}
}

// Execute executes the request
//
//	@return CipherSecretKeyResponse
func (a *CipherSecretKeyApiService) GetCipherSecretKeyExecute(r ApiGetCipherSecretKeyRequest) (*CipherSecretKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CipherSecretKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherSecretKeyApiService.GetCipherSecretKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server-instances/{server-instance-name}/cipher-secret-keys/{cipher-secret-key-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cipher-secret-key-name"+"}", url.PathEscape(parameterValueToString(r.cipherSecretKeyName, "cipherSecretKeyName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"server-instance-name"+"}", url.PathEscape(parameterValueToString(r.serverInstanceName, "serverInstanceName")), -1)

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

type ApiListCipherSecretKeysRequest struct {
	ctx                context.Context
	ApiService         *CipherSecretKeyApiService
	serverInstanceName string
	filter             *string
}

// SCIM filter
func (r ApiListCipherSecretKeysRequest) Filter(filter string) ApiListCipherSecretKeysRequest {
	r.filter = &filter
	return r
}

func (r ApiListCipherSecretKeysRequest) Execute() (*CipherSecretKeyListResponse, *http.Response, error) {
	return r.ApiService.ListCipherSecretKeysExecute(r)
}

/*
ListCipherSecretKeys Returns a list of all Cipher Secret Key objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serverInstanceName Name of the Server Instance
	@return ApiListCipherSecretKeysRequest
*/
func (a *CipherSecretKeyApiService) ListCipherSecretKeys(ctx context.Context, serverInstanceName string) ApiListCipherSecretKeysRequest {
	return ApiListCipherSecretKeysRequest{
		ApiService:         a,
		ctx:                ctx,
		serverInstanceName: serverInstanceName,
	}
}

// Execute executes the request
//
//	@return CipherSecretKeyListResponse
func (a *CipherSecretKeyApiService) ListCipherSecretKeysExecute(r ApiListCipherSecretKeysRequest) (*CipherSecretKeyListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CipherSecretKeyListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherSecretKeyApiService.ListCipherSecretKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server-instances/{server-instance-name}/cipher-secret-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"server-instance-name"+"}", url.PathEscape(parameterValueToString(r.serverInstanceName, "serverInstanceName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
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

type ApiUpdateCipherSecretKeyRequest struct {
	ctx                 context.Context
	ApiService          *CipherSecretKeyApiService
	cipherSecretKeyName string
	serverInstanceName  string
	updateRequest       *UpdateRequest
}

// Update an existing Cipher Secret Key
func (r ApiUpdateCipherSecretKeyRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateCipherSecretKeyRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateCipherSecretKeyRequest) Execute() (*CipherSecretKeyResponse, *http.Response, error) {
	return r.ApiService.UpdateCipherSecretKeyExecute(r)
}

/*
UpdateCipherSecretKey Update an existing Cipher Secret Key by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cipherSecretKeyName Name of the Cipher Secret Key
	@param serverInstanceName Name of the Server Instance
	@return ApiUpdateCipherSecretKeyRequest
*/
func (a *CipherSecretKeyApiService) UpdateCipherSecretKey(ctx context.Context, cipherSecretKeyName string, serverInstanceName string) ApiUpdateCipherSecretKeyRequest {
	return ApiUpdateCipherSecretKeyRequest{
		ApiService:          a,
		ctx:                 ctx,
		cipherSecretKeyName: cipherSecretKeyName,
		serverInstanceName:  serverInstanceName,
	}
}

// Execute executes the request
//
//	@return CipherSecretKeyResponse
func (a *CipherSecretKeyApiService) UpdateCipherSecretKeyExecute(r ApiUpdateCipherSecretKeyRequest) (*CipherSecretKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CipherSecretKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherSecretKeyApiService.UpdateCipherSecretKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server-instances/{server-instance-name}/cipher-secret-keys/{cipher-secret-key-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cipher-secret-key-name"+"}", url.PathEscape(parameterValueToString(r.cipherSecretKeyName, "cipherSecretKeyName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"server-instance-name"+"}", url.PathEscape(parameterValueToString(r.serverInstanceName, "serverInstanceName")), -1)

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
