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

// CipherStreamProviderAPIService CipherStreamProviderAPI service
type CipherStreamProviderAPIService service

type ApiAddCipherStreamProviderRequest struct {
	ctx                            context.Context
	ApiService                     *CipherStreamProviderAPIService
	addCipherStreamProviderRequest *AddCipherStreamProviderRequest
}

// Create a new Cipher Stream Provider in the config
func (r ApiAddCipherStreamProviderRequest) AddCipherStreamProviderRequest(addCipherStreamProviderRequest AddCipherStreamProviderRequest) ApiAddCipherStreamProviderRequest {
	r.addCipherStreamProviderRequest = &addCipherStreamProviderRequest
	return r
}

func (r ApiAddCipherStreamProviderRequest) Execute() (*AddCipherStreamProvider200Response, *http.Response, error) {
	return r.ApiService.AddCipherStreamProviderExecute(r)
}

/*
AddCipherStreamProvider Add a new Cipher Stream Provider to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddCipherStreamProviderRequest
*/
func (a *CipherStreamProviderAPIService) AddCipherStreamProvider(ctx context.Context) ApiAddCipherStreamProviderRequest {
	return ApiAddCipherStreamProviderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddCipherStreamProvider200Response
func (a *CipherStreamProviderAPIService) AddCipherStreamProviderExecute(r ApiAddCipherStreamProviderRequest) (*AddCipherStreamProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCipherStreamProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherStreamProviderAPIService.AddCipherStreamProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cipher-stream-providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addCipherStreamProviderRequest == nil {
		return localVarReturnValue, nil, reportError("addCipherStreamProviderRequest is required and must be specified")
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
	localVarPostBody = r.addCipherStreamProviderRequest
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

type ApiDeleteCipherStreamProviderRequest struct {
	ctx                      context.Context
	ApiService               *CipherStreamProviderAPIService
	cipherStreamProviderName string
}

func (r ApiDeleteCipherStreamProviderRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCipherStreamProviderExecute(r)
}

/*
DeleteCipherStreamProvider Delete a Cipher Stream Provider

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cipherStreamProviderName Name of the Cipher Stream Provider
	@return ApiDeleteCipherStreamProviderRequest
*/
func (a *CipherStreamProviderAPIService) DeleteCipherStreamProvider(ctx context.Context, cipherStreamProviderName string) ApiDeleteCipherStreamProviderRequest {
	return ApiDeleteCipherStreamProviderRequest{
		ApiService:               a,
		ctx:                      ctx,
		cipherStreamProviderName: cipherStreamProviderName,
	}
}

// Execute executes the request
func (a *CipherStreamProviderAPIService) DeleteCipherStreamProviderExecute(r ApiDeleteCipherStreamProviderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherStreamProviderAPIService.DeleteCipherStreamProvider")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cipher-stream-providers/{cipher-stream-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cipher-stream-provider-name"+"}", url.PathEscape(parameterValueToString(r.cipherStreamProviderName, "cipherStreamProviderName")), -1)

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

type ApiGetCipherStreamProviderRequest struct {
	ctx                      context.Context
	ApiService               *CipherStreamProviderAPIService
	cipherStreamProviderName string
}

func (r ApiGetCipherStreamProviderRequest) Execute() (*AddCipherStreamProvider200Response, *http.Response, error) {
	return r.ApiService.GetCipherStreamProviderExecute(r)
}

/*
GetCipherStreamProvider Returns a single Cipher Stream Provider

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cipherStreamProviderName Name of the Cipher Stream Provider
	@return ApiGetCipherStreamProviderRequest
*/
func (a *CipherStreamProviderAPIService) GetCipherStreamProvider(ctx context.Context, cipherStreamProviderName string) ApiGetCipherStreamProviderRequest {
	return ApiGetCipherStreamProviderRequest{
		ApiService:               a,
		ctx:                      ctx,
		cipherStreamProviderName: cipherStreamProviderName,
	}
}

// Execute executes the request
//
//	@return AddCipherStreamProvider200Response
func (a *CipherStreamProviderAPIService) GetCipherStreamProviderExecute(r ApiGetCipherStreamProviderRequest) (*AddCipherStreamProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCipherStreamProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherStreamProviderAPIService.GetCipherStreamProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cipher-stream-providers/{cipher-stream-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cipher-stream-provider-name"+"}", url.PathEscape(parameterValueToString(r.cipherStreamProviderName, "cipherStreamProviderName")), -1)

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

type ApiListCipherStreamProvidersRequest struct {
	ctx        context.Context
	ApiService *CipherStreamProviderAPIService
	filter     *string
}

// SCIM filter
func (r ApiListCipherStreamProvidersRequest) Filter(filter string) ApiListCipherStreamProvidersRequest {
	r.filter = &filter
	return r
}

func (r ApiListCipherStreamProvidersRequest) Execute() (*CipherStreamProviderListResponse, *http.Response, error) {
	return r.ApiService.ListCipherStreamProvidersExecute(r)
}

/*
ListCipherStreamProviders Returns a list of all Cipher Stream Provider objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCipherStreamProvidersRequest
*/
func (a *CipherStreamProviderAPIService) ListCipherStreamProviders(ctx context.Context) ApiListCipherStreamProvidersRequest {
	return ApiListCipherStreamProvidersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CipherStreamProviderListResponse
func (a *CipherStreamProviderAPIService) ListCipherStreamProvidersExecute(r ApiListCipherStreamProvidersRequest) (*CipherStreamProviderListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CipherStreamProviderListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherStreamProviderAPIService.ListCipherStreamProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cipher-stream-providers"

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

type ApiUpdateCipherStreamProviderRequest struct {
	ctx                      context.Context
	ApiService               *CipherStreamProviderAPIService
	cipherStreamProviderName string
	updateRequest            *UpdateRequest
}

// Update an existing Cipher Stream Provider
func (r ApiUpdateCipherStreamProviderRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateCipherStreamProviderRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateCipherStreamProviderRequest) Execute() (*AddCipherStreamProvider200Response, *http.Response, error) {
	return r.ApiService.UpdateCipherStreamProviderExecute(r)
}

/*
UpdateCipherStreamProvider Update an existing Cipher Stream Provider by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cipherStreamProviderName Name of the Cipher Stream Provider
	@return ApiUpdateCipherStreamProviderRequest
*/
func (a *CipherStreamProviderAPIService) UpdateCipherStreamProvider(ctx context.Context, cipherStreamProviderName string) ApiUpdateCipherStreamProviderRequest {
	return ApiUpdateCipherStreamProviderRequest{
		ApiService:               a,
		ctx:                      ctx,
		cipherStreamProviderName: cipherStreamProviderName,
	}
}

// Execute executes the request
//
//	@return AddCipherStreamProvider200Response
func (a *CipherStreamProviderAPIService) UpdateCipherStreamProviderExecute(r ApiUpdateCipherStreamProviderRequest) (*AddCipherStreamProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCipherStreamProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CipherStreamProviderAPIService.UpdateCipherStreamProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cipher-stream-providers/{cipher-stream-provider-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cipher-stream-provider-name"+"}", url.PathEscape(parameterValueToString(r.cipherStreamProviderName, "cipherStreamProviderName")), -1)

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
