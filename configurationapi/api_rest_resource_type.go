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

// RestResourceTypeApiService RestResourceTypeApi service
type RestResourceTypeApiService service

type ApiAddRestResourceTypeRequest struct {
	ctx                        context.Context
	ApiService                 *RestResourceTypeApiService
	addRestResourceTypeRequest *AddRestResourceTypeRequest
}

// Create a new REST Resource Type in the config
func (r ApiAddRestResourceTypeRequest) AddRestResourceTypeRequest(addRestResourceTypeRequest AddRestResourceTypeRequest) ApiAddRestResourceTypeRequest {
	r.addRestResourceTypeRequest = &addRestResourceTypeRequest
	return r
}

func (r ApiAddRestResourceTypeRequest) Execute() (*AddRestResourceType200Response, *http.Response, error) {
	return r.ApiService.AddRestResourceTypeExecute(r)
}

/*
AddRestResourceType Add a new REST Resource Type to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddRestResourceTypeRequest
*/
func (a *RestResourceTypeApiService) AddRestResourceType(ctx context.Context) ApiAddRestResourceTypeRequest {
	return ApiAddRestResourceTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddRestResourceType200Response
func (a *RestResourceTypeApiService) AddRestResourceTypeExecute(r ApiAddRestResourceTypeRequest) (*AddRestResourceType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddRestResourceType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RestResourceTypeApiService.AddRestResourceType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest-resource-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addRestResourceTypeRequest == nil {
		return localVarReturnValue, nil, reportError("addRestResourceTypeRequest is required and must be specified")
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
	localVarPostBody = r.addRestResourceTypeRequest
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

type ApiDeleteRestResourceTypeRequest struct {
	ctx                  context.Context
	ApiService           *RestResourceTypeApiService
	restResourceTypeName string
}

func (r ApiDeleteRestResourceTypeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRestResourceTypeExecute(r)
}

/*
DeleteRestResourceType Delete a REST Resource Type

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param restResourceTypeName Name of the REST Resource Type
	@return ApiDeleteRestResourceTypeRequest
*/
func (a *RestResourceTypeApiService) DeleteRestResourceType(ctx context.Context, restResourceTypeName string) ApiDeleteRestResourceTypeRequest {
	return ApiDeleteRestResourceTypeRequest{
		ApiService:           a,
		ctx:                  ctx,
		restResourceTypeName: restResourceTypeName,
	}
}

// Execute executes the request
func (a *RestResourceTypeApiService) DeleteRestResourceTypeExecute(r ApiDeleteRestResourceTypeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RestResourceTypeApiService.DeleteRestResourceType")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest-resource-types/{rest-resource-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"rest-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.restResourceTypeName, "restResourceTypeName")), -1)

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

type ApiGetRestResourceTypeRequest struct {
	ctx                  context.Context
	ApiService           *RestResourceTypeApiService
	restResourceTypeName string
}

func (r ApiGetRestResourceTypeRequest) Execute() (*AddRestResourceType200Response, *http.Response, error) {
	return r.ApiService.GetRestResourceTypeExecute(r)
}

/*
GetRestResourceType Returns a single REST Resource Type

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param restResourceTypeName Name of the REST Resource Type
	@return ApiGetRestResourceTypeRequest
*/
func (a *RestResourceTypeApiService) GetRestResourceType(ctx context.Context, restResourceTypeName string) ApiGetRestResourceTypeRequest {
	return ApiGetRestResourceTypeRequest{
		ApiService:           a,
		ctx:                  ctx,
		restResourceTypeName: restResourceTypeName,
	}
}

// Execute executes the request
//
//	@return AddRestResourceType200Response
func (a *RestResourceTypeApiService) GetRestResourceTypeExecute(r ApiGetRestResourceTypeRequest) (*AddRestResourceType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddRestResourceType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RestResourceTypeApiService.GetRestResourceType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest-resource-types/{rest-resource-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"rest-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.restResourceTypeName, "restResourceTypeName")), -1)

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

type ApiUpdateRestResourceTypeRequest struct {
	ctx                  context.Context
	ApiService           *RestResourceTypeApiService
	restResourceTypeName string
	updateRequest        *UpdateRequest
}

// Update an existing REST Resource Type
func (r ApiUpdateRestResourceTypeRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateRestResourceTypeRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateRestResourceTypeRequest) Execute() (*AddRestResourceType200Response, *http.Response, error) {
	return r.ApiService.UpdateRestResourceTypeExecute(r)
}

/*
UpdateRestResourceType Update an existing REST Resource Type by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param restResourceTypeName Name of the REST Resource Type
	@return ApiUpdateRestResourceTypeRequest
*/
func (a *RestResourceTypeApiService) UpdateRestResourceType(ctx context.Context, restResourceTypeName string) ApiUpdateRestResourceTypeRequest {
	return ApiUpdateRestResourceTypeRequest{
		ApiService:           a,
		ctx:                  ctx,
		restResourceTypeName: restResourceTypeName,
	}
}

// Execute executes the request
//
//	@return AddRestResourceType200Response
func (a *RestResourceTypeApiService) UpdateRestResourceTypeExecute(r ApiUpdateRestResourceTypeRequest) (*AddRestResourceType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddRestResourceType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RestResourceTypeApiService.UpdateRestResourceType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest-resource-types/{rest-resource-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"rest-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.restResourceTypeName, "restResourceTypeName")), -1)

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
