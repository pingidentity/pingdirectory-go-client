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

// PassThroughAuthenticationHandlerApiService PassThroughAuthenticationHandlerApi service
type PassThroughAuthenticationHandlerApiService service

type ApiAddPassThroughAuthenticationHandlerRequest struct {
	ctx                                        context.Context
	ApiService                                 *PassThroughAuthenticationHandlerApiService
	addPassThroughAuthenticationHandlerRequest *AddPassThroughAuthenticationHandlerRequest
}

// Create a new Pass Through Authentication Handler in the config
func (r ApiAddPassThroughAuthenticationHandlerRequest) AddPassThroughAuthenticationHandlerRequest(addPassThroughAuthenticationHandlerRequest AddPassThroughAuthenticationHandlerRequest) ApiAddPassThroughAuthenticationHandlerRequest {
	r.addPassThroughAuthenticationHandlerRequest = &addPassThroughAuthenticationHandlerRequest
	return r
}

func (r ApiAddPassThroughAuthenticationHandlerRequest) Execute() (*AddPassThroughAuthenticationHandler200Response, *http.Response, error) {
	return r.ApiService.AddPassThroughAuthenticationHandlerExecute(r)
}

/*
AddPassThroughAuthenticationHandler Add a new Pass Through Authentication Handler to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPassThroughAuthenticationHandlerRequest
*/
func (a *PassThroughAuthenticationHandlerApiService) AddPassThroughAuthenticationHandler(ctx context.Context) ApiAddPassThroughAuthenticationHandlerRequest {
	return ApiAddPassThroughAuthenticationHandlerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddPassThroughAuthenticationHandler200Response
func (a *PassThroughAuthenticationHandlerApiService) AddPassThroughAuthenticationHandlerExecute(r ApiAddPassThroughAuthenticationHandlerRequest) (*AddPassThroughAuthenticationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPassThroughAuthenticationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PassThroughAuthenticationHandlerApiService.AddPassThroughAuthenticationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pass-through-authentication-handlers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addPassThroughAuthenticationHandlerRequest == nil {
		return localVarReturnValue, nil, reportError("addPassThroughAuthenticationHandlerRequest is required and must be specified")
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
	localVarPostBody = r.addPassThroughAuthenticationHandlerRequest
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

type ApiDeletePassThroughAuthenticationHandlerRequest struct {
	ctx                                  context.Context
	ApiService                           *PassThroughAuthenticationHandlerApiService
	passThroughAuthenticationHandlerName string
}

func (r ApiDeletePassThroughAuthenticationHandlerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePassThroughAuthenticationHandlerExecute(r)
}

/*
DeletePassThroughAuthenticationHandler Delete a Pass Through Authentication Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passThroughAuthenticationHandlerName Name of the Pass Through Authentication Handler
	@return ApiDeletePassThroughAuthenticationHandlerRequest
*/
func (a *PassThroughAuthenticationHandlerApiService) DeletePassThroughAuthenticationHandler(ctx context.Context, passThroughAuthenticationHandlerName string) ApiDeletePassThroughAuthenticationHandlerRequest {
	return ApiDeletePassThroughAuthenticationHandlerRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		passThroughAuthenticationHandlerName: passThroughAuthenticationHandlerName,
	}
}

// Execute executes the request
func (a *PassThroughAuthenticationHandlerApiService) DeletePassThroughAuthenticationHandlerExecute(r ApiDeletePassThroughAuthenticationHandlerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PassThroughAuthenticationHandlerApiService.DeletePassThroughAuthenticationHandler")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pass-through-authentication-handlers/{pass-through-authentication-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"pass-through-authentication-handler-name"+"}", url.PathEscape(parameterValueToString(r.passThroughAuthenticationHandlerName, "passThroughAuthenticationHandlerName")), -1)

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

type ApiGetPassThroughAuthenticationHandlerRequest struct {
	ctx                                  context.Context
	ApiService                           *PassThroughAuthenticationHandlerApiService
	passThroughAuthenticationHandlerName string
}

func (r ApiGetPassThroughAuthenticationHandlerRequest) Execute() (*AddPassThroughAuthenticationHandler200Response, *http.Response, error) {
	return r.ApiService.GetPassThroughAuthenticationHandlerExecute(r)
}

/*
GetPassThroughAuthenticationHandler Returns a single Pass Through Authentication Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passThroughAuthenticationHandlerName Name of the Pass Through Authentication Handler
	@return ApiGetPassThroughAuthenticationHandlerRequest
*/
func (a *PassThroughAuthenticationHandlerApiService) GetPassThroughAuthenticationHandler(ctx context.Context, passThroughAuthenticationHandlerName string) ApiGetPassThroughAuthenticationHandlerRequest {
	return ApiGetPassThroughAuthenticationHandlerRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		passThroughAuthenticationHandlerName: passThroughAuthenticationHandlerName,
	}
}

// Execute executes the request
//
//	@return AddPassThroughAuthenticationHandler200Response
func (a *PassThroughAuthenticationHandlerApiService) GetPassThroughAuthenticationHandlerExecute(r ApiGetPassThroughAuthenticationHandlerRequest) (*AddPassThroughAuthenticationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPassThroughAuthenticationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PassThroughAuthenticationHandlerApiService.GetPassThroughAuthenticationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pass-through-authentication-handlers/{pass-through-authentication-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"pass-through-authentication-handler-name"+"}", url.PathEscape(parameterValueToString(r.passThroughAuthenticationHandlerName, "passThroughAuthenticationHandlerName")), -1)

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

type ApiListPassThroughAuthenticationHandlersRequest struct {
	ctx        context.Context
	ApiService *PassThroughAuthenticationHandlerApiService
	filter     *string
}

// SCIM filter
func (r ApiListPassThroughAuthenticationHandlersRequest) Filter(filter string) ApiListPassThroughAuthenticationHandlersRequest {
	r.filter = &filter
	return r
}

func (r ApiListPassThroughAuthenticationHandlersRequest) Execute() (*PassThroughAuthenticationHandlerListResponse, *http.Response, error) {
	return r.ApiService.ListPassThroughAuthenticationHandlersExecute(r)
}

/*
ListPassThroughAuthenticationHandlers Returns a list of all Pass Through Authentication Handler objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListPassThroughAuthenticationHandlersRequest
*/
func (a *PassThroughAuthenticationHandlerApiService) ListPassThroughAuthenticationHandlers(ctx context.Context) ApiListPassThroughAuthenticationHandlersRequest {
	return ApiListPassThroughAuthenticationHandlersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PassThroughAuthenticationHandlerListResponse
func (a *PassThroughAuthenticationHandlerApiService) ListPassThroughAuthenticationHandlersExecute(r ApiListPassThroughAuthenticationHandlersRequest) (*PassThroughAuthenticationHandlerListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PassThroughAuthenticationHandlerListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PassThroughAuthenticationHandlerApiService.ListPassThroughAuthenticationHandlers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pass-through-authentication-handlers"

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

type ApiUpdatePassThroughAuthenticationHandlerRequest struct {
	ctx                                  context.Context
	ApiService                           *PassThroughAuthenticationHandlerApiService
	passThroughAuthenticationHandlerName string
	updateRequest                        *UpdateRequest
}

// Update an existing Pass Through Authentication Handler
func (r ApiUpdatePassThroughAuthenticationHandlerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdatePassThroughAuthenticationHandlerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdatePassThroughAuthenticationHandlerRequest) Execute() (*AddPassThroughAuthenticationHandler200Response, *http.Response, error) {
	return r.ApiService.UpdatePassThroughAuthenticationHandlerExecute(r)
}

/*
UpdatePassThroughAuthenticationHandler Update an existing Pass Through Authentication Handler by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param passThroughAuthenticationHandlerName Name of the Pass Through Authentication Handler
	@return ApiUpdatePassThroughAuthenticationHandlerRequest
*/
func (a *PassThroughAuthenticationHandlerApiService) UpdatePassThroughAuthenticationHandler(ctx context.Context, passThroughAuthenticationHandlerName string) ApiUpdatePassThroughAuthenticationHandlerRequest {
	return ApiUpdatePassThroughAuthenticationHandlerRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		passThroughAuthenticationHandlerName: passThroughAuthenticationHandlerName,
	}
}

// Execute executes the request
//
//	@return AddPassThroughAuthenticationHandler200Response
func (a *PassThroughAuthenticationHandlerApiService) UpdatePassThroughAuthenticationHandlerExecute(r ApiUpdatePassThroughAuthenticationHandlerRequest) (*AddPassThroughAuthenticationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddPassThroughAuthenticationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PassThroughAuthenticationHandlerApiService.UpdatePassThroughAuthenticationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pass-through-authentication-handlers/{pass-through-authentication-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"pass-through-authentication-handler-name"+"}", url.PathEscape(parameterValueToString(r.passThroughAuthenticationHandlerName, "passThroughAuthenticationHandlerName")), -1)

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
