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


// ConstructedAttributeApiService ConstructedAttributeApi service
type ConstructedAttributeApiService service

type ApiAddConstructedAttributeRequest struct {
	ctx context.Context
	ApiService *ConstructedAttributeApiService
	addConstructedAttributeRequest *AddConstructedAttributeRequest
}

// Create a new Constructed Attribute in the config
func (r ApiAddConstructedAttributeRequest) AddConstructedAttributeRequest(addConstructedAttributeRequest AddConstructedAttributeRequest) ApiAddConstructedAttributeRequest {
	r.addConstructedAttributeRequest = &addConstructedAttributeRequest
	return r
}

func (r ApiAddConstructedAttributeRequest) Execute() (*ConstructedAttributeResponse, *http.Response, error) {
	return r.ApiService.AddConstructedAttributeExecute(r)
}

/*
AddConstructedAttribute Add a new Constructed Attribute to the config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddConstructedAttributeRequest
*/
func (a *ConstructedAttributeApiService) AddConstructedAttribute(ctx context.Context) ApiAddConstructedAttributeRequest {
	return ApiAddConstructedAttributeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConstructedAttributeResponse
func (a *ConstructedAttributeApiService) AddConstructedAttributeExecute(r ApiAddConstructedAttributeRequest) (*ConstructedAttributeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConstructedAttributeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConstructedAttributeApiService.AddConstructedAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/constructed-attributes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addConstructedAttributeRequest == nil {
		return localVarReturnValue, nil, reportError("addConstructedAttributeRequest is required and must be specified")
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
	localVarPostBody = r.addConstructedAttributeRequest
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

type ApiDeleteConstructedAttributeRequest struct {
	ctx context.Context
	ApiService *ConstructedAttributeApiService
	constructedAttributeName string
}

func (r ApiDeleteConstructedAttributeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteConstructedAttributeExecute(r)
}

/*
DeleteConstructedAttribute Delete a Constructed Attribute

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param constructedAttributeName Name of the Constructed Attribute to be deleted
 @return ApiDeleteConstructedAttributeRequest
*/
func (a *ConstructedAttributeApiService) DeleteConstructedAttribute(ctx context.Context, constructedAttributeName string) ApiDeleteConstructedAttributeRequest {
	return ApiDeleteConstructedAttributeRequest{
		ApiService: a,
		ctx: ctx,
		constructedAttributeName: constructedAttributeName,
	}
}

// Execute executes the request
func (a *ConstructedAttributeApiService) DeleteConstructedAttributeExecute(r ApiDeleteConstructedAttributeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConstructedAttributeApiService.DeleteConstructedAttribute")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/constructed-attributes/{constructed-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"constructed-attribute-name"+"}", url.PathEscape(parameterToString(r.constructedAttributeName, "")), -1)

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

type ApiGetConstructedAttributeRequest struct {
	ctx context.Context
	ApiService *ConstructedAttributeApiService
	constructedAttributeName string
}

func (r ApiGetConstructedAttributeRequest) Execute() (*ConstructedAttributeResponse, *http.Response, error) {
	return r.ApiService.GetConstructedAttributeExecute(r)
}

/*
GetConstructedAttribute Returns a single Constructed Attribute

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param constructedAttributeName Name of the Constructed Attribute to be read
 @return ApiGetConstructedAttributeRequest
*/
func (a *ConstructedAttributeApiService) GetConstructedAttribute(ctx context.Context, constructedAttributeName string) ApiGetConstructedAttributeRequest {
	return ApiGetConstructedAttributeRequest{
		ApiService: a,
		ctx: ctx,
		constructedAttributeName: constructedAttributeName,
	}
}

// Execute executes the request
//  @return ConstructedAttributeResponse
func (a *ConstructedAttributeApiService) GetConstructedAttributeExecute(r ApiGetConstructedAttributeRequest) (*ConstructedAttributeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConstructedAttributeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConstructedAttributeApiService.GetConstructedAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/constructed-attributes/{constructed-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"constructed-attribute-name"+"}", url.PathEscape(parameterToString(r.constructedAttributeName, "")), -1)

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

type ApiUpdateConstructedAttributeRequest struct {
	ctx context.Context
	ApiService *ConstructedAttributeApiService
	constructedAttributeName string
	updateRequest *UpdateRequest
}

// Update an existing Constructed Attribute
func (r ApiUpdateConstructedAttributeRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateConstructedAttributeRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateConstructedAttributeRequest) Execute() (*ConstructedAttributeResponse, *http.Response, error) {
	return r.ApiService.UpdateConstructedAttributeExecute(r)
}

/*
UpdateConstructedAttribute Update an existing Constructed Attribute by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param constructedAttributeName Name of the Constructed Attribute to be updated
 @return ApiUpdateConstructedAttributeRequest
*/
func (a *ConstructedAttributeApiService) UpdateConstructedAttribute(ctx context.Context, constructedAttributeName string) ApiUpdateConstructedAttributeRequest {
	return ApiUpdateConstructedAttributeRequest{
		ApiService: a,
		ctx: ctx,
		constructedAttributeName: constructedAttributeName,
	}
}

// Execute executes the request
//  @return ConstructedAttributeResponse
func (a *ConstructedAttributeApiService) UpdateConstructedAttributeExecute(r ApiUpdateConstructedAttributeRequest) (*ConstructedAttributeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConstructedAttributeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConstructedAttributeApiService.UpdateConstructedAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/constructed-attributes/{constructed-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"constructed-attribute-name"+"}", url.PathEscape(parameterToString(r.constructedAttributeName, "")), -1)

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
