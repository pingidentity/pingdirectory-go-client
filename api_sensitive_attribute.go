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

// SensitiveAttributeApiService SensitiveAttributeApi service
type SensitiveAttributeApiService service

type ApiAddSensitiveAttributeRequest struct {
	ctx                          context.Context
	ApiService                   *SensitiveAttributeApiService
	addSensitiveAttributeRequest *AddSensitiveAttributeRequest
}

// Create a new Sensitive Attribute in the config
func (r ApiAddSensitiveAttributeRequest) AddSensitiveAttributeRequest(addSensitiveAttributeRequest AddSensitiveAttributeRequest) ApiAddSensitiveAttributeRequest {
	r.addSensitiveAttributeRequest = &addSensitiveAttributeRequest
	return r
}

func (r ApiAddSensitiveAttributeRequest) Execute() (*SensitiveAttributeResponse, *http.Response, error) {
	return r.ApiService.AddSensitiveAttributeExecute(r)
}

/*
AddSensitiveAttribute Add a new Sensitive Attribute to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddSensitiveAttributeRequest
*/
func (a *SensitiveAttributeApiService) AddSensitiveAttribute(ctx context.Context) ApiAddSensitiveAttributeRequest {
	return ApiAddSensitiveAttributeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SensitiveAttributeResponse
func (a *SensitiveAttributeApiService) AddSensitiveAttributeExecute(r ApiAddSensitiveAttributeRequest) (*SensitiveAttributeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SensitiveAttributeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SensitiveAttributeApiService.AddSensitiveAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sensitive-attributes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addSensitiveAttributeRequest == nil {
		return localVarReturnValue, nil, reportError("addSensitiveAttributeRequest is required and must be specified")
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
	localVarPostBody = r.addSensitiveAttributeRequest
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

type ApiDeleteSensitiveAttributeRequest struct {
	ctx                    context.Context
	ApiService             *SensitiveAttributeApiService
	sensitiveAttributeName string
}

func (r ApiDeleteSensitiveAttributeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSensitiveAttributeExecute(r)
}

/*
DeleteSensitiveAttribute Delete a Sensitive Attribute

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sensitiveAttributeName Name of the Sensitive Attribute to be deleted
	@return ApiDeleteSensitiveAttributeRequest
*/
func (a *SensitiveAttributeApiService) DeleteSensitiveAttribute(ctx context.Context, sensitiveAttributeName string) ApiDeleteSensitiveAttributeRequest {
	return ApiDeleteSensitiveAttributeRequest{
		ApiService:             a,
		ctx:                    ctx,
		sensitiveAttributeName: sensitiveAttributeName,
	}
}

// Execute executes the request
func (a *SensitiveAttributeApiService) DeleteSensitiveAttributeExecute(r ApiDeleteSensitiveAttributeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SensitiveAttributeApiService.DeleteSensitiveAttribute")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sensitive-attributes/{sensitive-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"sensitive-attribute-name"+"}", url.PathEscape(parameterToString(r.sensitiveAttributeName, "")), -1)

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

type ApiGetSensitiveAttributeRequest struct {
	ctx                    context.Context
	ApiService             *SensitiveAttributeApiService
	sensitiveAttributeName string
}

func (r ApiGetSensitiveAttributeRequest) Execute() (*SensitiveAttributeResponse, *http.Response, error) {
	return r.ApiService.GetSensitiveAttributeExecute(r)
}

/*
GetSensitiveAttribute Returns a single Sensitive Attribute

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sensitiveAttributeName Name of the Sensitive Attribute to be read
	@return ApiGetSensitiveAttributeRequest
*/
func (a *SensitiveAttributeApiService) GetSensitiveAttribute(ctx context.Context, sensitiveAttributeName string) ApiGetSensitiveAttributeRequest {
	return ApiGetSensitiveAttributeRequest{
		ApiService:             a,
		ctx:                    ctx,
		sensitiveAttributeName: sensitiveAttributeName,
	}
}

// Execute executes the request
//
//	@return SensitiveAttributeResponse
func (a *SensitiveAttributeApiService) GetSensitiveAttributeExecute(r ApiGetSensitiveAttributeRequest) (*SensitiveAttributeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SensitiveAttributeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SensitiveAttributeApiService.GetSensitiveAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sensitive-attributes/{sensitive-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"sensitive-attribute-name"+"}", url.PathEscape(parameterToString(r.sensitiveAttributeName, "")), -1)

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

type ApiUpdateSensitiveAttributeRequest struct {
	ctx                    context.Context
	ApiService             *SensitiveAttributeApiService
	sensitiveAttributeName string
	updateRequest          *UpdateRequest
}

// Update an existing Sensitive Attribute
func (r ApiUpdateSensitiveAttributeRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateSensitiveAttributeRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateSensitiveAttributeRequest) Execute() (*SensitiveAttributeResponse, *http.Response, error) {
	return r.ApiService.UpdateSensitiveAttributeExecute(r)
}

/*
UpdateSensitiveAttribute Update an existing Sensitive Attribute by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sensitiveAttributeName Name of the Sensitive Attribute to be updated
	@return ApiUpdateSensitiveAttributeRequest
*/
func (a *SensitiveAttributeApiService) UpdateSensitiveAttribute(ctx context.Context, sensitiveAttributeName string) ApiUpdateSensitiveAttributeRequest {
	return ApiUpdateSensitiveAttributeRequest{
		ApiService:             a,
		ctx:                    ctx,
		sensitiveAttributeName: sensitiveAttributeName,
	}
}

// Execute executes the request
//
//	@return SensitiveAttributeResponse
func (a *SensitiveAttributeApiService) UpdateSensitiveAttributeExecute(r ApiUpdateSensitiveAttributeRequest) (*SensitiveAttributeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SensitiveAttributeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SensitiveAttributeApiService.UpdateSensitiveAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sensitive-attributes/{sensitive-attribute-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"sensitive-attribute-name"+"}", url.PathEscape(parameterToString(r.sensitiveAttributeName, "")), -1)

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
