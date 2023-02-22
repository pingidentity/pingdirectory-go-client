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

// ScimSchemaApiService ScimSchemaApi service
type ScimSchemaApiService service

type ApiAddScimSchemaRequest struct {
	ctx                  context.Context
	ApiService           *ScimSchemaApiService
	addScimSchemaRequest *AddScimSchemaRequest
}

// Create a new SCIM Schema in the config
func (r ApiAddScimSchemaRequest) AddScimSchemaRequest(addScimSchemaRequest AddScimSchemaRequest) ApiAddScimSchemaRequest {
	r.addScimSchemaRequest = &addScimSchemaRequest
	return r
}

func (r ApiAddScimSchemaRequest) Execute() (*ScimSchemaResponse, *http.Response, error) {
	return r.ApiService.AddScimSchemaExecute(r)
}

/*
AddScimSchema Add a new SCIM Schema to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddScimSchemaRequest
*/
func (a *ScimSchemaApiService) AddScimSchema(ctx context.Context) ApiAddScimSchemaRequest {
	return ApiAddScimSchemaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ScimSchemaResponse
func (a *ScimSchemaApiService) AddScimSchemaExecute(r ApiAddScimSchemaRequest) (*ScimSchemaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ScimSchemaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScimSchemaApiService.AddScimSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-schemas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addScimSchemaRequest == nil {
		return localVarReturnValue, nil, reportError("addScimSchemaRequest is required and must be specified")
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
	localVarPostBody = r.addScimSchemaRequest
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

type ApiDeleteScimSchemaRequest struct {
	ctx            context.Context
	ApiService     *ScimSchemaApiService
	scimSchemaName string
}

func (r ApiDeleteScimSchemaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteScimSchemaExecute(r)
}

/*
DeleteScimSchema Delete a SCIM Schema

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scimSchemaName Name of the SCIM Schema
	@return ApiDeleteScimSchemaRequest
*/
func (a *ScimSchemaApiService) DeleteScimSchema(ctx context.Context, scimSchemaName string) ApiDeleteScimSchemaRequest {
	return ApiDeleteScimSchemaRequest{
		ApiService:     a,
		ctx:            ctx,
		scimSchemaName: scimSchemaName,
	}
}

// Execute executes the request
func (a *ScimSchemaApiService) DeleteScimSchemaExecute(r ApiDeleteScimSchemaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScimSchemaApiService.DeleteScimSchema")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-schemas/{scim-schema-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"scim-schema-name"+"}", url.PathEscape(parameterToString(r.scimSchemaName, "")), -1)

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

type ApiGetScimSchemaRequest struct {
	ctx            context.Context
	ApiService     *ScimSchemaApiService
	scimSchemaName string
}

func (r ApiGetScimSchemaRequest) Execute() (*ScimSchemaResponse, *http.Response, error) {
	return r.ApiService.GetScimSchemaExecute(r)
}

/*
GetScimSchema Returns a single SCIM Schema

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scimSchemaName Name of the SCIM Schema
	@return ApiGetScimSchemaRequest
*/
func (a *ScimSchemaApiService) GetScimSchema(ctx context.Context, scimSchemaName string) ApiGetScimSchemaRequest {
	return ApiGetScimSchemaRequest{
		ApiService:     a,
		ctx:            ctx,
		scimSchemaName: scimSchemaName,
	}
}

// Execute executes the request
//
//	@return ScimSchemaResponse
func (a *ScimSchemaApiService) GetScimSchemaExecute(r ApiGetScimSchemaRequest) (*ScimSchemaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ScimSchemaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScimSchemaApiService.GetScimSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-schemas/{scim-schema-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"scim-schema-name"+"}", url.PathEscape(parameterToString(r.scimSchemaName, "")), -1)

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

type ApiUpdateScimSchemaRequest struct {
	ctx            context.Context
	ApiService     *ScimSchemaApiService
	scimSchemaName string
	updateRequest  *UpdateRequest
}

// Update an existing SCIM Schema
func (r ApiUpdateScimSchemaRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateScimSchemaRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateScimSchemaRequest) Execute() (*ScimSchemaResponse, *http.Response, error) {
	return r.ApiService.UpdateScimSchemaExecute(r)
}

/*
UpdateScimSchema Update an existing SCIM Schema by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scimSchemaName Name of the SCIM Schema
	@return ApiUpdateScimSchemaRequest
*/
func (a *ScimSchemaApiService) UpdateScimSchema(ctx context.Context, scimSchemaName string) ApiUpdateScimSchemaRequest {
	return ApiUpdateScimSchemaRequest{
		ApiService:     a,
		ctx:            ctx,
		scimSchemaName: scimSchemaName,
	}
}

// Execute executes the request
//
//	@return ScimSchemaResponse
func (a *ScimSchemaApiService) UpdateScimSchemaExecute(r ApiUpdateScimSchemaRequest) (*ScimSchemaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ScimSchemaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScimSchemaApiService.UpdateScimSchema")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-schemas/{scim-schema-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"scim-schema-name"+"}", url.PathEscape(parameterToString(r.scimSchemaName, "")), -1)

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