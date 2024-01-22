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

// LogFieldSyntaxAPIService LogFieldSyntaxAPI service
type LogFieldSyntaxAPIService service

type ApiGetLogFieldSyntaxRequest struct {
	ctx                context.Context
	ApiService         *LogFieldSyntaxAPIService
	logFieldSyntaxName string
}

func (r ApiGetLogFieldSyntaxRequest) Execute() (*GetLogFieldSyntax200Response, *http.Response, error) {
	return r.ApiService.GetLogFieldSyntaxExecute(r)
}

/*
GetLogFieldSyntax Returns a single Log Field Syntax

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFieldSyntaxName Name of the Log Field Syntax
	@return ApiGetLogFieldSyntaxRequest
*/
func (a *LogFieldSyntaxAPIService) GetLogFieldSyntax(ctx context.Context, logFieldSyntaxName string) ApiGetLogFieldSyntaxRequest {
	return ApiGetLogFieldSyntaxRequest{
		ApiService:         a,
		ctx:                ctx,
		logFieldSyntaxName: logFieldSyntaxName,
	}
}

// Execute executes the request
//
//	@return GetLogFieldSyntax200Response
func (a *LogFieldSyntaxAPIService) GetLogFieldSyntaxExecute(r ApiGetLogFieldSyntaxRequest) (*GetLogFieldSyntax200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLogFieldSyntax200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldSyntaxAPIService.GetLogFieldSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-syntaxes/{log-field-syntax-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-field-syntax-name"+"}", url.PathEscape(parameterValueToString(r.logFieldSyntaxName, "logFieldSyntaxName")), -1)

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

type ApiListLogFieldSyntaxesRequest struct {
	ctx        context.Context
	ApiService *LogFieldSyntaxAPIService
	filter     *string
}

// SCIM filter
func (r ApiListLogFieldSyntaxesRequest) Filter(filter string) ApiListLogFieldSyntaxesRequest {
	r.filter = &filter
	return r
}

func (r ApiListLogFieldSyntaxesRequest) Execute() (*LogFieldSyntaxListResponse, *http.Response, error) {
	return r.ApiService.ListLogFieldSyntaxesExecute(r)
}

/*
ListLogFieldSyntaxes Returns a list of all Log Field Syntax objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListLogFieldSyntaxesRequest
*/
func (a *LogFieldSyntaxAPIService) ListLogFieldSyntaxes(ctx context.Context) ApiListLogFieldSyntaxesRequest {
	return ApiListLogFieldSyntaxesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LogFieldSyntaxListResponse
func (a *LogFieldSyntaxAPIService) ListLogFieldSyntaxesExecute(r ApiListLogFieldSyntaxesRequest) (*LogFieldSyntaxListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogFieldSyntaxListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldSyntaxAPIService.ListLogFieldSyntaxes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-syntaxes"

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

type ApiUpdateLogFieldSyntaxRequest struct {
	ctx                context.Context
	ApiService         *LogFieldSyntaxAPIService
	logFieldSyntaxName string
	updateRequest      *UpdateRequest
}

// Update an existing Log Field Syntax
func (r ApiUpdateLogFieldSyntaxRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateLogFieldSyntaxRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateLogFieldSyntaxRequest) Execute() (*GetLogFieldSyntax200Response, *http.Response, error) {
	return r.ApiService.UpdateLogFieldSyntaxExecute(r)
}

/*
UpdateLogFieldSyntax Update an existing Log Field Syntax by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFieldSyntaxName Name of the Log Field Syntax
	@return ApiUpdateLogFieldSyntaxRequest
*/
func (a *LogFieldSyntaxAPIService) UpdateLogFieldSyntax(ctx context.Context, logFieldSyntaxName string) ApiUpdateLogFieldSyntaxRequest {
	return ApiUpdateLogFieldSyntaxRequest{
		ApiService:         a,
		ctx:                ctx,
		logFieldSyntaxName: logFieldSyntaxName,
	}
}

// Execute executes the request
//
//	@return GetLogFieldSyntax200Response
func (a *LogFieldSyntaxAPIService) UpdateLogFieldSyntaxExecute(r ApiUpdateLogFieldSyntaxRequest) (*GetLogFieldSyntax200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLogFieldSyntax200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldSyntaxAPIService.UpdateLogFieldSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-syntaxes/{log-field-syntax-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-field-syntax-name"+"}", url.PathEscape(parameterValueToString(r.logFieldSyntaxName, "logFieldSyntaxName")), -1)

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
