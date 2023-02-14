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

// HttpServletCrossOriginPolicyApiService HttpServletCrossOriginPolicyApi service
type HttpServletCrossOriginPolicyApiService service

type ApiAddHttpServletCrossOriginPolicyRequest struct {
	ctx                                    context.Context
	ApiService                             *HttpServletCrossOriginPolicyApiService
	addHttpServletCrossOriginPolicyRequest *AddHttpServletCrossOriginPolicyRequest
}

// Create a new HTTP Servlet Cross Origin Policy in the config
func (r ApiAddHttpServletCrossOriginPolicyRequest) AddHttpServletCrossOriginPolicyRequest(addHttpServletCrossOriginPolicyRequest AddHttpServletCrossOriginPolicyRequest) ApiAddHttpServletCrossOriginPolicyRequest {
	r.addHttpServletCrossOriginPolicyRequest = &addHttpServletCrossOriginPolicyRequest
	return r
}

func (r ApiAddHttpServletCrossOriginPolicyRequest) Execute() (*HttpServletCrossOriginPolicyResponse, *http.Response, error) {
	return r.ApiService.AddHttpServletCrossOriginPolicyExecute(r)
}

/*
AddHttpServletCrossOriginPolicy Add a new HTTP Servlet Cross Origin Policy to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddHttpServletCrossOriginPolicyRequest
*/
func (a *HttpServletCrossOriginPolicyApiService) AddHttpServletCrossOriginPolicy(ctx context.Context) ApiAddHttpServletCrossOriginPolicyRequest {
	return ApiAddHttpServletCrossOriginPolicyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HttpServletCrossOriginPolicyResponse
func (a *HttpServletCrossOriginPolicyApiService) AddHttpServletCrossOriginPolicyExecute(r ApiAddHttpServletCrossOriginPolicyRequest) (*HttpServletCrossOriginPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HttpServletCrossOriginPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletCrossOriginPolicyApiService.AddHttpServletCrossOriginPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-cross-origin-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addHttpServletCrossOriginPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("addHttpServletCrossOriginPolicyRequest is required and must be specified")
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
	localVarPostBody = r.addHttpServletCrossOriginPolicyRequest
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

type ApiDeleteHttpServletCrossOriginPolicyRequest struct {
	ctx                              context.Context
	ApiService                       *HttpServletCrossOriginPolicyApiService
	httpServletCrossOriginPolicyName string
}

func (r ApiDeleteHttpServletCrossOriginPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteHttpServletCrossOriginPolicyExecute(r)
}

/*
DeleteHttpServletCrossOriginPolicy Delete a HTTP Servlet Cross Origin Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletCrossOriginPolicyName Name of the HTTP Servlet Cross Origin Policy
	@return ApiDeleteHttpServletCrossOriginPolicyRequest
*/
func (a *HttpServletCrossOriginPolicyApiService) DeleteHttpServletCrossOriginPolicy(ctx context.Context, httpServletCrossOriginPolicyName string) ApiDeleteHttpServletCrossOriginPolicyRequest {
	return ApiDeleteHttpServletCrossOriginPolicyRequest{
		ApiService:                       a,
		ctx:                              ctx,
		httpServletCrossOriginPolicyName: httpServletCrossOriginPolicyName,
	}
}

// Execute executes the request
func (a *HttpServletCrossOriginPolicyApiService) DeleteHttpServletCrossOriginPolicyExecute(r ApiDeleteHttpServletCrossOriginPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletCrossOriginPolicyApiService.DeleteHttpServletCrossOriginPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-cross-origin-policies/{http-servlet-cross-origin-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-cross-origin-policy-name"+"}", url.PathEscape(parameterToString(r.httpServletCrossOriginPolicyName, "")), -1)

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

type ApiGetHttpServletCrossOriginPolicyRequest struct {
	ctx                              context.Context
	ApiService                       *HttpServletCrossOriginPolicyApiService
	httpServletCrossOriginPolicyName string
}

func (r ApiGetHttpServletCrossOriginPolicyRequest) Execute() (*HttpServletCrossOriginPolicyResponse, *http.Response, error) {
	return r.ApiService.GetHttpServletCrossOriginPolicyExecute(r)
}

/*
GetHttpServletCrossOriginPolicy Returns a single HTTP Servlet Cross Origin Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletCrossOriginPolicyName Name of the HTTP Servlet Cross Origin Policy
	@return ApiGetHttpServletCrossOriginPolicyRequest
*/
func (a *HttpServletCrossOriginPolicyApiService) GetHttpServletCrossOriginPolicy(ctx context.Context, httpServletCrossOriginPolicyName string) ApiGetHttpServletCrossOriginPolicyRequest {
	return ApiGetHttpServletCrossOriginPolicyRequest{
		ApiService:                       a,
		ctx:                              ctx,
		httpServletCrossOriginPolicyName: httpServletCrossOriginPolicyName,
	}
}

// Execute executes the request
//
//	@return HttpServletCrossOriginPolicyResponse
func (a *HttpServletCrossOriginPolicyApiService) GetHttpServletCrossOriginPolicyExecute(r ApiGetHttpServletCrossOriginPolicyRequest) (*HttpServletCrossOriginPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HttpServletCrossOriginPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletCrossOriginPolicyApiService.GetHttpServletCrossOriginPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-cross-origin-policies/{http-servlet-cross-origin-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-cross-origin-policy-name"+"}", url.PathEscape(parameterToString(r.httpServletCrossOriginPolicyName, "")), -1)

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

type ApiUpdateHttpServletCrossOriginPolicyRequest struct {
	ctx                              context.Context
	ApiService                       *HttpServletCrossOriginPolicyApiService
	httpServletCrossOriginPolicyName string
	updateRequest                    *UpdateRequest
}

// Update an existing HTTP Servlet Cross Origin Policy
func (r ApiUpdateHttpServletCrossOriginPolicyRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateHttpServletCrossOriginPolicyRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateHttpServletCrossOriginPolicyRequest) Execute() (*HttpServletCrossOriginPolicyResponse, *http.Response, error) {
	return r.ApiService.UpdateHttpServletCrossOriginPolicyExecute(r)
}

/*
UpdateHttpServletCrossOriginPolicy Update an existing HTTP Servlet Cross Origin Policy by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param httpServletCrossOriginPolicyName Name of the HTTP Servlet Cross Origin Policy
	@return ApiUpdateHttpServletCrossOriginPolicyRequest
*/
func (a *HttpServletCrossOriginPolicyApiService) UpdateHttpServletCrossOriginPolicy(ctx context.Context, httpServletCrossOriginPolicyName string) ApiUpdateHttpServletCrossOriginPolicyRequest {
	return ApiUpdateHttpServletCrossOriginPolicyRequest{
		ApiService:                       a,
		ctx:                              ctx,
		httpServletCrossOriginPolicyName: httpServletCrossOriginPolicyName,
	}
}

// Execute executes the request
//
//	@return HttpServletCrossOriginPolicyResponse
func (a *HttpServletCrossOriginPolicyApiService) UpdateHttpServletCrossOriginPolicyExecute(r ApiUpdateHttpServletCrossOriginPolicyRequest) (*HttpServletCrossOriginPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HttpServletCrossOriginPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServletCrossOriginPolicyApiService.UpdateHttpServletCrossOriginPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/http-servlet-cross-origin-policies/{http-servlet-cross-origin-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"http-servlet-cross-origin-policy-name"+"}", url.PathEscape(parameterToString(r.httpServletCrossOriginPolicyName, "")), -1)

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
