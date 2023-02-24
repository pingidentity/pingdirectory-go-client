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

// ExternalServerApiService ExternalServerApi service
type ExternalServerApiService service

type ApiAddExternalServerRequest struct {
	ctx                      context.Context
	ApiService               *ExternalServerApiService
	addExternalServerRequest *AddExternalServerRequest
}

// Create a new External Server in the config
func (r ApiAddExternalServerRequest) AddExternalServerRequest(addExternalServerRequest AddExternalServerRequest) ApiAddExternalServerRequest {
	r.addExternalServerRequest = &addExternalServerRequest
	return r
}

func (r ApiAddExternalServerRequest) Execute() (*AddExternalServer200Response, *http.Response, error) {
	return r.ApiService.AddExternalServerExecute(r)
}

/*
AddExternalServer Add a new External Server to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddExternalServerRequest
*/
func (a *ExternalServerApiService) AddExternalServer(ctx context.Context) ApiAddExternalServerRequest {
	return ApiAddExternalServerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddExternalServer200Response
func (a *ExternalServerApiService) AddExternalServerExecute(r ApiAddExternalServerRequest) (*AddExternalServer200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddExternalServer200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalServerApiService.AddExternalServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external-servers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addExternalServerRequest == nil {
		return localVarReturnValue, nil, reportError("addExternalServerRequest is required and must be specified")
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
	localVarPostBody = r.addExternalServerRequest
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

type ApiDeleteExternalServerRequest struct {
	ctx                context.Context
	ApiService         *ExternalServerApiService
	externalServerName string
}

func (r ApiDeleteExternalServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExternalServerExecute(r)
}

/*
DeleteExternalServer Delete a External Server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalServerName Name of the External Server
	@return ApiDeleteExternalServerRequest
*/
func (a *ExternalServerApiService) DeleteExternalServer(ctx context.Context, externalServerName string) ApiDeleteExternalServerRequest {
	return ApiDeleteExternalServerRequest{
		ApiService:         a,
		ctx:                ctx,
		externalServerName: externalServerName,
	}
}

// Execute executes the request
func (a *ExternalServerApiService) DeleteExternalServerExecute(r ApiDeleteExternalServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalServerApiService.DeleteExternalServer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external-servers/{external-server-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"external-server-name"+"}", url.PathEscape(parameterValueToString(r.externalServerName, "externalServerName")), -1)

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

type ApiGetExternalServerRequest struct {
	ctx                context.Context
	ApiService         *ExternalServerApiService
	externalServerName string
}

func (r ApiGetExternalServerRequest) Execute() (*AddExternalServer200Response, *http.Response, error) {
	return r.ApiService.GetExternalServerExecute(r)
}

/*
GetExternalServer Returns a single External Server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalServerName Name of the External Server
	@return ApiGetExternalServerRequest
*/
func (a *ExternalServerApiService) GetExternalServer(ctx context.Context, externalServerName string) ApiGetExternalServerRequest {
	return ApiGetExternalServerRequest{
		ApiService:         a,
		ctx:                ctx,
		externalServerName: externalServerName,
	}
}

// Execute executes the request
//
//	@return AddExternalServer200Response
func (a *ExternalServerApiService) GetExternalServerExecute(r ApiGetExternalServerRequest) (*AddExternalServer200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddExternalServer200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalServerApiService.GetExternalServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external-servers/{external-server-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"external-server-name"+"}", url.PathEscape(parameterValueToString(r.externalServerName, "externalServerName")), -1)

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

type ApiUpdateExternalServerRequest struct {
	ctx                context.Context
	ApiService         *ExternalServerApiService
	externalServerName string
	updateRequest      *UpdateRequest
}

// Update an existing External Server
func (r ApiUpdateExternalServerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateExternalServerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateExternalServerRequest) Execute() (*AddExternalServer200Response, *http.Response, error) {
	return r.ApiService.UpdateExternalServerExecute(r)
}

/*
UpdateExternalServer Update an existing External Server by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalServerName Name of the External Server
	@return ApiUpdateExternalServerRequest
*/
func (a *ExternalServerApiService) UpdateExternalServer(ctx context.Context, externalServerName string) ApiUpdateExternalServerRequest {
	return ApiUpdateExternalServerRequest{
		ApiService:         a,
		ctx:                ctx,
		externalServerName: externalServerName,
	}
}

// Execute executes the request
//
//	@return AddExternalServer200Response
func (a *ExternalServerApiService) UpdateExternalServerExecute(r ApiUpdateExternalServerRequest) (*AddExternalServer200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddExternalServer200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalServerApiService.UpdateExternalServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external-servers/{external-server-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"external-server-name"+"}", url.PathEscape(parameterValueToString(r.externalServerName, "externalServerName")), -1)

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
