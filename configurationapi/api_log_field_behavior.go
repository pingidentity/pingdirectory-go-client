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

// LogFieldBehaviorAPIService LogFieldBehaviorAPI service
type LogFieldBehaviorAPIService service

type ApiAddLogFieldBehaviorRequest struct {
	ctx                        context.Context
	ApiService                 *LogFieldBehaviorAPIService
	addLogFieldBehaviorRequest *AddLogFieldBehaviorRequest
}

// Create a new Log Field Behavior in the config
func (r ApiAddLogFieldBehaviorRequest) AddLogFieldBehaviorRequest(addLogFieldBehaviorRequest AddLogFieldBehaviorRequest) ApiAddLogFieldBehaviorRequest {
	r.addLogFieldBehaviorRequest = &addLogFieldBehaviorRequest
	return r
}

func (r ApiAddLogFieldBehaviorRequest) Execute() (*AddLogFieldBehavior200Response, *http.Response, error) {
	return r.ApiService.AddLogFieldBehaviorExecute(r)
}

/*
AddLogFieldBehavior Add a new Log Field Behavior to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddLogFieldBehaviorRequest
*/
func (a *LogFieldBehaviorAPIService) AddLogFieldBehavior(ctx context.Context) ApiAddLogFieldBehaviorRequest {
	return ApiAddLogFieldBehaviorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddLogFieldBehavior200Response
func (a *LogFieldBehaviorAPIService) AddLogFieldBehaviorExecute(r ApiAddLogFieldBehaviorRequest) (*AddLogFieldBehavior200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLogFieldBehavior200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldBehaviorAPIService.AddLogFieldBehavior")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-behaviors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addLogFieldBehaviorRequest == nil {
		return localVarReturnValue, nil, reportError("addLogFieldBehaviorRequest is required and must be specified")
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
	localVarPostBody = r.addLogFieldBehaviorRequest
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

type ApiDeleteLogFieldBehaviorRequest struct {
	ctx                  context.Context
	ApiService           *LogFieldBehaviorAPIService
	logFieldBehaviorName string
}

func (r ApiDeleteLogFieldBehaviorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLogFieldBehaviorExecute(r)
}

/*
DeleteLogFieldBehavior Delete a Log Field Behavior

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFieldBehaviorName Name of the Log Field Behavior
	@return ApiDeleteLogFieldBehaviorRequest
*/
func (a *LogFieldBehaviorAPIService) DeleteLogFieldBehavior(ctx context.Context, logFieldBehaviorName string) ApiDeleteLogFieldBehaviorRequest {
	return ApiDeleteLogFieldBehaviorRequest{
		ApiService:           a,
		ctx:                  ctx,
		logFieldBehaviorName: logFieldBehaviorName,
	}
}

// Execute executes the request
func (a *LogFieldBehaviorAPIService) DeleteLogFieldBehaviorExecute(r ApiDeleteLogFieldBehaviorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldBehaviorAPIService.DeleteLogFieldBehavior")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-behaviors/{log-field-behavior-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-field-behavior-name"+"}", url.PathEscape(parameterValueToString(r.logFieldBehaviorName, "logFieldBehaviorName")), -1)

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

type ApiGetLogFieldBehaviorRequest struct {
	ctx                  context.Context
	ApiService           *LogFieldBehaviorAPIService
	logFieldBehaviorName string
}

func (r ApiGetLogFieldBehaviorRequest) Execute() (*AddLogFieldBehavior200Response, *http.Response, error) {
	return r.ApiService.GetLogFieldBehaviorExecute(r)
}

/*
GetLogFieldBehavior Returns a single Log Field Behavior

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFieldBehaviorName Name of the Log Field Behavior
	@return ApiGetLogFieldBehaviorRequest
*/
func (a *LogFieldBehaviorAPIService) GetLogFieldBehavior(ctx context.Context, logFieldBehaviorName string) ApiGetLogFieldBehaviorRequest {
	return ApiGetLogFieldBehaviorRequest{
		ApiService:           a,
		ctx:                  ctx,
		logFieldBehaviorName: logFieldBehaviorName,
	}
}

// Execute executes the request
//
//	@return AddLogFieldBehavior200Response
func (a *LogFieldBehaviorAPIService) GetLogFieldBehaviorExecute(r ApiGetLogFieldBehaviorRequest) (*AddLogFieldBehavior200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLogFieldBehavior200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldBehaviorAPIService.GetLogFieldBehavior")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-behaviors/{log-field-behavior-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-field-behavior-name"+"}", url.PathEscape(parameterValueToString(r.logFieldBehaviorName, "logFieldBehaviorName")), -1)

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

type ApiListLogFieldBehaviorsRequest struct {
	ctx        context.Context
	ApiService *LogFieldBehaviorAPIService
	filter     *string
}

// SCIM filter
func (r ApiListLogFieldBehaviorsRequest) Filter(filter string) ApiListLogFieldBehaviorsRequest {
	r.filter = &filter
	return r
}

func (r ApiListLogFieldBehaviorsRequest) Execute() (*LogFieldBehaviorListResponse, *http.Response, error) {
	return r.ApiService.ListLogFieldBehaviorsExecute(r)
}

/*
ListLogFieldBehaviors Returns a list of all Log Field Behavior objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListLogFieldBehaviorsRequest
*/
func (a *LogFieldBehaviorAPIService) ListLogFieldBehaviors(ctx context.Context) ApiListLogFieldBehaviorsRequest {
	return ApiListLogFieldBehaviorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LogFieldBehaviorListResponse
func (a *LogFieldBehaviorAPIService) ListLogFieldBehaviorsExecute(r ApiListLogFieldBehaviorsRequest) (*LogFieldBehaviorListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogFieldBehaviorListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldBehaviorAPIService.ListLogFieldBehaviors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-behaviors"

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

type ApiUpdateLogFieldBehaviorRequest struct {
	ctx                  context.Context
	ApiService           *LogFieldBehaviorAPIService
	logFieldBehaviorName string
	updateRequest        *UpdateRequest
}

// Update an existing Log Field Behavior
func (r ApiUpdateLogFieldBehaviorRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateLogFieldBehaviorRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateLogFieldBehaviorRequest) Execute() (*AddLogFieldBehavior200Response, *http.Response, error) {
	return r.ApiService.UpdateLogFieldBehaviorExecute(r)
}

/*
UpdateLogFieldBehavior Update an existing Log Field Behavior by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFieldBehaviorName Name of the Log Field Behavior
	@return ApiUpdateLogFieldBehaviorRequest
*/
func (a *LogFieldBehaviorAPIService) UpdateLogFieldBehavior(ctx context.Context, logFieldBehaviorName string) ApiUpdateLogFieldBehaviorRequest {
	return ApiUpdateLogFieldBehaviorRequest{
		ApiService:           a,
		ctx:                  ctx,
		logFieldBehaviorName: logFieldBehaviorName,
	}
}

// Execute executes the request
//
//	@return AddLogFieldBehavior200Response
func (a *LogFieldBehaviorAPIService) UpdateLogFieldBehaviorExecute(r ApiUpdateLogFieldBehaviorRequest) (*AddLogFieldBehavior200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLogFieldBehavior200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFieldBehaviorAPIService.UpdateLogFieldBehavior")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-field-behaviors/{log-field-behavior-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-field-behavior-name"+"}", url.PathEscape(parameterValueToString(r.logFieldBehaviorName, "logFieldBehaviorName")), -1)

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
