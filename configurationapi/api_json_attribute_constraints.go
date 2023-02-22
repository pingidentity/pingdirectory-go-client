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

// JsonAttributeConstraintsApiService JsonAttributeConstraintsApi service
type JsonAttributeConstraintsApiService service

type ApiAddJsonAttributeConstraintsRequest struct {
	ctx                                context.Context
	ApiService                         *JsonAttributeConstraintsApiService
	addJsonAttributeConstraintsRequest *AddJsonAttributeConstraintsRequest
}

// Create a new JSON Attribute Constraints in the config
func (r ApiAddJsonAttributeConstraintsRequest) AddJsonAttributeConstraintsRequest(addJsonAttributeConstraintsRequest AddJsonAttributeConstraintsRequest) ApiAddJsonAttributeConstraintsRequest {
	r.addJsonAttributeConstraintsRequest = &addJsonAttributeConstraintsRequest
	return r
}

func (r ApiAddJsonAttributeConstraintsRequest) Execute() (*JsonAttributeConstraintsResponse, *http.Response, error) {
	return r.ApiService.AddJsonAttributeConstraintsExecute(r)
}

/*
AddJsonAttributeConstraints Add a new JSON Attribute Constraints to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddJsonAttributeConstraintsRequest
*/
func (a *JsonAttributeConstraintsApiService) AddJsonAttributeConstraints(ctx context.Context) ApiAddJsonAttributeConstraintsRequest {
	return ApiAddJsonAttributeConstraintsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return JsonAttributeConstraintsResponse
func (a *JsonAttributeConstraintsApiService) AddJsonAttributeConstraintsExecute(r ApiAddJsonAttributeConstraintsRequest) (*JsonAttributeConstraintsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JsonAttributeConstraintsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JsonAttributeConstraintsApiService.AddJsonAttributeConstraints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/json-attribute-constraints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addJsonAttributeConstraintsRequest == nil {
		return localVarReturnValue, nil, reportError("addJsonAttributeConstraintsRequest is required and must be specified")
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
	localVarPostBody = r.addJsonAttributeConstraintsRequest
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

type ApiDeleteJsonAttributeConstraintsRequest struct {
	ctx                          context.Context
	ApiService                   *JsonAttributeConstraintsApiService
	jsonAttributeConstraintsName string
}

func (r ApiDeleteJsonAttributeConstraintsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJsonAttributeConstraintsExecute(r)
}

/*
DeleteJsonAttributeConstraints Delete a JSON Attribute Constraints

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jsonAttributeConstraintsName Name of the JSON Attribute Constraints
	@return ApiDeleteJsonAttributeConstraintsRequest
*/
func (a *JsonAttributeConstraintsApiService) DeleteJsonAttributeConstraints(ctx context.Context, jsonAttributeConstraintsName string) ApiDeleteJsonAttributeConstraintsRequest {
	return ApiDeleteJsonAttributeConstraintsRequest{
		ApiService:                   a,
		ctx:                          ctx,
		jsonAttributeConstraintsName: jsonAttributeConstraintsName,
	}
}

// Execute executes the request
func (a *JsonAttributeConstraintsApiService) DeleteJsonAttributeConstraintsExecute(r ApiDeleteJsonAttributeConstraintsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JsonAttributeConstraintsApiService.DeleteJsonAttributeConstraints")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/json-attribute-constraints/{json-attribute-constraints-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"json-attribute-constraints-name"+"}", url.PathEscape(parameterToString(r.jsonAttributeConstraintsName, "")), -1)

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

type ApiGetJsonAttributeConstraintsRequest struct {
	ctx                          context.Context
	ApiService                   *JsonAttributeConstraintsApiService
	jsonAttributeConstraintsName string
}

func (r ApiGetJsonAttributeConstraintsRequest) Execute() (*JsonAttributeConstraintsResponse, *http.Response, error) {
	return r.ApiService.GetJsonAttributeConstraintsExecute(r)
}

/*
GetJsonAttributeConstraints Returns a single JSON Attribute Constraints

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jsonAttributeConstraintsName Name of the JSON Attribute Constraints
	@return ApiGetJsonAttributeConstraintsRequest
*/
func (a *JsonAttributeConstraintsApiService) GetJsonAttributeConstraints(ctx context.Context, jsonAttributeConstraintsName string) ApiGetJsonAttributeConstraintsRequest {
	return ApiGetJsonAttributeConstraintsRequest{
		ApiService:                   a,
		ctx:                          ctx,
		jsonAttributeConstraintsName: jsonAttributeConstraintsName,
	}
}

// Execute executes the request
//
//	@return JsonAttributeConstraintsResponse
func (a *JsonAttributeConstraintsApiService) GetJsonAttributeConstraintsExecute(r ApiGetJsonAttributeConstraintsRequest) (*JsonAttributeConstraintsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JsonAttributeConstraintsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JsonAttributeConstraintsApiService.GetJsonAttributeConstraints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/json-attribute-constraints/{json-attribute-constraints-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"json-attribute-constraints-name"+"}", url.PathEscape(parameterToString(r.jsonAttributeConstraintsName, "")), -1)

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

type ApiUpdateJsonAttributeConstraintsRequest struct {
	ctx                          context.Context
	ApiService                   *JsonAttributeConstraintsApiService
	jsonAttributeConstraintsName string
	updateRequest                *UpdateRequest
}

// Update an existing JSON Attribute Constraints
func (r ApiUpdateJsonAttributeConstraintsRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateJsonAttributeConstraintsRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateJsonAttributeConstraintsRequest) Execute() (*JsonAttributeConstraintsResponse, *http.Response, error) {
	return r.ApiService.UpdateJsonAttributeConstraintsExecute(r)
}

/*
UpdateJsonAttributeConstraints Update an existing JSON Attribute Constraints by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jsonAttributeConstraintsName Name of the JSON Attribute Constraints
	@return ApiUpdateJsonAttributeConstraintsRequest
*/
func (a *JsonAttributeConstraintsApiService) UpdateJsonAttributeConstraints(ctx context.Context, jsonAttributeConstraintsName string) ApiUpdateJsonAttributeConstraintsRequest {
	return ApiUpdateJsonAttributeConstraintsRequest{
		ApiService:                   a,
		ctx:                          ctx,
		jsonAttributeConstraintsName: jsonAttributeConstraintsName,
	}
}

// Execute executes the request
//
//	@return JsonAttributeConstraintsResponse
func (a *JsonAttributeConstraintsApiService) UpdateJsonAttributeConstraintsExecute(r ApiUpdateJsonAttributeConstraintsRequest) (*JsonAttributeConstraintsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JsonAttributeConstraintsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JsonAttributeConstraintsApiService.UpdateJsonAttributeConstraints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/json-attribute-constraints/{json-attribute-constraints-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"json-attribute-constraints-name"+"}", url.PathEscape(parameterToString(r.jsonAttributeConstraintsName, "")), -1)

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