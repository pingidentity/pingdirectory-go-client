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

// HttpServletExtensionAPIService HttpServletExtensionAPI service
type HttpServletExtensionAPIService service

type ApiAddHttpServletExtensionRequest struct {
	ctx                            context.Context
	ApiService                     *HttpServletExtensionAPIService
	addHttpServletExtensionRequest *AddHttpServletExtensionRequest
}

// Create a new HTTP Servlet Extension in the config
func (r ApiAddHttpServletExtensionRequest) AddHttpServletExtensionRequest(addHttpServletExtensionRequest AddHttpServletExtensionRequest) ApiAddHttpServletExtensionRequest {
	r.addHttpServletExtensionRequest = &addHttpServletExtensionRequest
	return r
}

func (r ApiAddHttpServletExtensionRequest) Execute() (*AddHttpServletExtension200Response, *http.Response, error) {
	return r.ApiService.AddHttpServletExtensionExecute(r)
}

/*
AddHttpServletExtension Add a new HTTP Servlet Extension to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddHttpServletExtensionRequest
*/
func (a *HttpServletExtensionAPIService) AddHttpServletExtension(ctx context.Context) ApiAddHttpServletExtensionRequest {
	return ApiAddHttpServletExtensionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddHttpServletExtension200Response
func (a *HttpServletExtensionAPIService) AddHttpServletExtensionExecute(r ApiAddHttpServletExtensionRequest) (*AddHttpServletExtension200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddHttpServletExtension200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletExtensionAPIService.AddHttpServletExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addHttpServletExtensionRequest == nil {
		return localVarReturnValue, nil, reportError("addHttpServletExtensionRequest is required and must be specified")
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
	localVarPostBody = r.addHttpServletExtensionRequest
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

type ApiDeleteHttpServletExtensionRequest struct {
	ctx                      context.Context
	ApiService               *HttpServletExtensionAPIService
	httpServletExtensionName string
}

func (r ApiDeleteHttpServletExtensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteHttpServletExtensionExecute(r)
}

/*
DeleteHttpServletExtension Delete a HTTP Servlet Extension

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiDeleteHttpServletExtensionRequest
*/
func (a *HttpServletExtensionAPIService) DeleteHttpServletExtension(ctx context.Context, httpServletExtensionName string) ApiDeleteHttpServletExtensionRequest {
	return ApiDeleteHttpServletExtensionRequest{
		ApiService:               a,
		ctx:                      ctx,
		httpServletExtensionName: httpServletExtensionName,
	}
}

// Execute executes the request
func (a *HttpServletExtensionAPIService) DeleteHttpServletExtensionExecute(r ApiDeleteHttpServletExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletExtensionAPIService.DeleteHttpServletExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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

type ApiGetHttpServletExtensionRequest struct {
	ctx                      context.Context
	ApiService               *HttpServletExtensionAPIService
	httpServletExtensionName string
}

func (r ApiGetHttpServletExtensionRequest) Execute() (*GetHttpServletExtension200Response, *http.Response, error) {
	return r.ApiService.GetHttpServletExtensionExecute(r)
}

/*
GetHttpServletExtension Returns a single HTTP Servlet Extension

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiGetHttpServletExtensionRequest
*/
func (a *HttpServletExtensionAPIService) GetHttpServletExtension(ctx context.Context, httpServletExtensionName string) ApiGetHttpServletExtensionRequest {
	return ApiGetHttpServletExtensionRequest{
		ApiService:               a,
		ctx:                      ctx,
		httpServletExtensionName: httpServletExtensionName,
	}
}

// Execute executes the request
//
//	@return GetHttpServletExtension200Response
func (a *HttpServletExtensionAPIService) GetHttpServletExtensionExecute(r ApiGetHttpServletExtensionRequest) (*GetHttpServletExtension200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetHttpServletExtension200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletExtensionAPIService.GetHttpServletExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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

type ApiListHttpServletExtensionsRequest struct {
	ctx        context.Context
	ApiService *HttpServletExtensionAPIService
	filter     *string
}

// SCIM filter
func (r ApiListHttpServletExtensionsRequest) Filter(filter string) ApiListHttpServletExtensionsRequest {
	r.filter = &filter
	return r
}

func (r ApiListHttpServletExtensionsRequest) Execute() (*HttpServletExtensionListResponse, *http.Response, error) {
	return r.ApiService.ListHttpServletExtensionsExecute(r)
}

/*
ListHttpServletExtensions Returns a list of all HTTP Servlet Extension objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListHttpServletExtensionsRequest
*/
func (a *HttpServletExtensionAPIService) ListHttpServletExtensions(ctx context.Context) ApiListHttpServletExtensionsRequest {
	return ApiListHttpServletExtensionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HttpServletExtensionListResponse
func (a *HttpServletExtensionAPIService) ListHttpServletExtensionsExecute(r ApiListHttpServletExtensionsRequest) (*HttpServletExtensionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HttpServletExtensionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletExtensionAPIService.ListHttpServletExtensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions"

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

type ApiUpdateHttpServletExtensionRequest struct {
	ctx                      context.Context
	ApiService               *HttpServletExtensionAPIService
	httpServletExtensionName string
	updateRequest            *UpdateRequest
}

// Update an existing HTTP Servlet Extension
func (r ApiUpdateHttpServletExtensionRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateHttpServletExtensionRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateHttpServletExtensionRequest) Execute() (*GetHttpServletExtension200Response, *http.Response, error) {
	return r.ApiService.UpdateHttpServletExtensionExecute(r)
}

/*
UpdateHttpServletExtension Update an existing HTTP Servlet Extension by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletExtensionName Name of the HTTP Servlet Extension
	@return ApiUpdateHttpServletExtensionRequest
*/
func (a *HttpServletExtensionAPIService) UpdateHttpServletExtension(ctx context.Context, httpServletExtensionName string) ApiUpdateHttpServletExtensionRequest {
	return ApiUpdateHttpServletExtensionRequest{
		ApiService:               a,
		ctx:                      ctx,
		httpServletExtensionName: httpServletExtensionName,
	}
}

// Execute executes the request
//
//	@return GetHttpServletExtension200Response
func (a *HttpServletExtensionAPIService) UpdateHttpServletExtensionExecute(r ApiUpdateHttpServletExtensionRequest) (*GetHttpServletExtension200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetHttpServletExtension200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletExtensionAPIService.UpdateHttpServletExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-extensions/{http-servlet-extension-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-extension-name"+"}", url.PathEscape(parameterValueToString(r.httpServletExtensionName, "httpServletExtensionName")), -1)

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
