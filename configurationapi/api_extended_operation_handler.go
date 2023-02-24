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

// ExtendedOperationHandlerApiService ExtendedOperationHandlerApi service
type ExtendedOperationHandlerApiService service

type ApiAddExtendedOperationHandlerRequest struct {
	ctx                                context.Context
	ApiService                         *ExtendedOperationHandlerApiService
	addExtendedOperationHandlerRequest *AddExtendedOperationHandlerRequest
}

// Create a new Extended Operation Handler in the config
func (r ApiAddExtendedOperationHandlerRequest) AddExtendedOperationHandlerRequest(addExtendedOperationHandlerRequest AddExtendedOperationHandlerRequest) ApiAddExtendedOperationHandlerRequest {
	r.addExtendedOperationHandlerRequest = &addExtendedOperationHandlerRequest
	return r
}

func (r ApiAddExtendedOperationHandlerRequest) Execute() (*AddExtendedOperationHandler200Response, *http.Response, error) {
	return r.ApiService.AddExtendedOperationHandlerExecute(r)
}

/*
AddExtendedOperationHandler Add a new Extended Operation Handler to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddExtendedOperationHandlerRequest
*/
func (a *ExtendedOperationHandlerApiService) AddExtendedOperationHandler(ctx context.Context) ApiAddExtendedOperationHandlerRequest {
	return ApiAddExtendedOperationHandlerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddExtendedOperationHandler200Response
func (a *ExtendedOperationHandlerApiService) AddExtendedOperationHandlerExecute(r ApiAddExtendedOperationHandlerRequest) (*AddExtendedOperationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddExtendedOperationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtendedOperationHandlerApiService.AddExtendedOperationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended-operation-handlers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addExtendedOperationHandlerRequest == nil {
		return localVarReturnValue, nil, reportError("addExtendedOperationHandlerRequest is required and must be specified")
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
	localVarPostBody = r.addExtendedOperationHandlerRequest
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

type ApiDeleteExtendedOperationHandlerRequest struct {
	ctx                          context.Context
	ApiService                   *ExtendedOperationHandlerApiService
	extendedOperationHandlerName string
}

func (r ApiDeleteExtendedOperationHandlerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExtendedOperationHandlerExecute(r)
}

/*
DeleteExtendedOperationHandler Delete a Extended Operation Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param extendedOperationHandlerName Name of the Extended Operation Handler
	@return ApiDeleteExtendedOperationHandlerRequest
*/
func (a *ExtendedOperationHandlerApiService) DeleteExtendedOperationHandler(ctx context.Context, extendedOperationHandlerName string) ApiDeleteExtendedOperationHandlerRequest {
	return ApiDeleteExtendedOperationHandlerRequest{
		ApiService:                   a,
		ctx:                          ctx,
		extendedOperationHandlerName: extendedOperationHandlerName,
	}
}

// Execute executes the request
func (a *ExtendedOperationHandlerApiService) DeleteExtendedOperationHandlerExecute(r ApiDeleteExtendedOperationHandlerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtendedOperationHandlerApiService.DeleteExtendedOperationHandler")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended-operation-handlers/{extended-operation-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"extended-operation-handler-name"+"}", url.PathEscape(parameterValueToString(r.extendedOperationHandlerName, "extendedOperationHandlerName")), -1)

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

type ApiGetExtendedOperationHandlerRequest struct {
	ctx                          context.Context
	ApiService                   *ExtendedOperationHandlerApiService
	extendedOperationHandlerName string
}

func (r ApiGetExtendedOperationHandlerRequest) Execute() (*GetExtendedOperationHandler200Response, *http.Response, error) {
	return r.ApiService.GetExtendedOperationHandlerExecute(r)
}

/*
GetExtendedOperationHandler Returns a single Extended Operation Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param extendedOperationHandlerName Name of the Extended Operation Handler
	@return ApiGetExtendedOperationHandlerRequest
*/
func (a *ExtendedOperationHandlerApiService) GetExtendedOperationHandler(ctx context.Context, extendedOperationHandlerName string) ApiGetExtendedOperationHandlerRequest {
	return ApiGetExtendedOperationHandlerRequest{
		ApiService:                   a,
		ctx:                          ctx,
		extendedOperationHandlerName: extendedOperationHandlerName,
	}
}

// Execute executes the request
//
//	@return GetExtendedOperationHandler200Response
func (a *ExtendedOperationHandlerApiService) GetExtendedOperationHandlerExecute(r ApiGetExtendedOperationHandlerRequest) (*GetExtendedOperationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetExtendedOperationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtendedOperationHandlerApiService.GetExtendedOperationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended-operation-handlers/{extended-operation-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"extended-operation-handler-name"+"}", url.PathEscape(parameterValueToString(r.extendedOperationHandlerName, "extendedOperationHandlerName")), -1)

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

type ApiUpdateExtendedOperationHandlerRequest struct {
	ctx                          context.Context
	ApiService                   *ExtendedOperationHandlerApiService
	extendedOperationHandlerName string
	updateRequest                *UpdateRequest
}

// Update an existing Extended Operation Handler
func (r ApiUpdateExtendedOperationHandlerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateExtendedOperationHandlerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateExtendedOperationHandlerRequest) Execute() (*GetExtendedOperationHandler200Response, *http.Response, error) {
	return r.ApiService.UpdateExtendedOperationHandlerExecute(r)
}

/*
UpdateExtendedOperationHandler Update an existing Extended Operation Handler by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param extendedOperationHandlerName Name of the Extended Operation Handler
	@return ApiUpdateExtendedOperationHandlerRequest
*/
func (a *ExtendedOperationHandlerApiService) UpdateExtendedOperationHandler(ctx context.Context, extendedOperationHandlerName string) ApiUpdateExtendedOperationHandlerRequest {
	return ApiUpdateExtendedOperationHandlerRequest{
		ApiService:                   a,
		ctx:                          ctx,
		extendedOperationHandlerName: extendedOperationHandlerName,
	}
}

// Execute executes the request
//
//	@return GetExtendedOperationHandler200Response
func (a *ExtendedOperationHandlerApiService) UpdateExtendedOperationHandlerExecute(r ApiUpdateExtendedOperationHandlerRequest) (*GetExtendedOperationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetExtendedOperationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtendedOperationHandlerApiService.UpdateExtendedOperationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/extended-operation-handlers/{extended-operation-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"extended-operation-handler-name"+"}", url.PathEscape(parameterValueToString(r.extendedOperationHandlerName, "extendedOperationHandlerName")), -1)

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
