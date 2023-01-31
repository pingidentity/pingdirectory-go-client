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

// LogFileRotationListenerApiService LogFileRotationListenerApi service
type LogFileRotationListenerApiService service

type ApiAddLogFileRotationListenerRequest struct {
	ctx                               context.Context
	ApiService                        *LogFileRotationListenerApiService
	addLogFileRotationListenerRequest *AddLogFileRotationListenerRequest
}

// Create a new Log File Rotation Listener in the config
func (r ApiAddLogFileRotationListenerRequest) AddLogFileRotationListenerRequest(addLogFileRotationListenerRequest AddLogFileRotationListenerRequest) ApiAddLogFileRotationListenerRequest {
	r.addLogFileRotationListenerRequest = &addLogFileRotationListenerRequest
	return r
}

func (r ApiAddLogFileRotationListenerRequest) Execute() (*AddLogFileRotationListener200Response, *http.Response, error) {
	return r.ApiService.AddLogFileRotationListenerExecute(r)
}

/*
AddLogFileRotationListener Add a new Log File Rotation Listener to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddLogFileRotationListenerRequest
*/
func (a *LogFileRotationListenerApiService) AddLogFileRotationListener(ctx context.Context) ApiAddLogFileRotationListenerRequest {
	return ApiAddLogFileRotationListenerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddLogFileRotationListener200Response
func (a *LogFileRotationListenerApiService) AddLogFileRotationListenerExecute(r ApiAddLogFileRotationListenerRequest) (*AddLogFileRotationListener200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLogFileRotationListener200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFileRotationListenerApiService.AddLogFileRotationListener")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-file-rotation-listeners"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addLogFileRotationListenerRequest == nil {
		return localVarReturnValue, nil, reportError("addLogFileRotationListenerRequest is required and must be specified")
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
	localVarPostBody = r.addLogFileRotationListenerRequest
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

type ApiDeleteLogFileRotationListenerRequest struct {
	ctx                         context.Context
	ApiService                  *LogFileRotationListenerApiService
	logFileRotationListenerName string
}

func (r ApiDeleteLogFileRotationListenerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLogFileRotationListenerExecute(r)
}

/*
DeleteLogFileRotationListener Delete a Log File Rotation Listener

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFileRotationListenerName Name of the Log File Rotation Listener to be deleted
	@return ApiDeleteLogFileRotationListenerRequest
*/
func (a *LogFileRotationListenerApiService) DeleteLogFileRotationListener(ctx context.Context, logFileRotationListenerName string) ApiDeleteLogFileRotationListenerRequest {
	return ApiDeleteLogFileRotationListenerRequest{
		ApiService:                  a,
		ctx:                         ctx,
		logFileRotationListenerName: logFileRotationListenerName,
	}
}

// Execute executes the request
func (a *LogFileRotationListenerApiService) DeleteLogFileRotationListenerExecute(r ApiDeleteLogFileRotationListenerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFileRotationListenerApiService.DeleteLogFileRotationListener")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-file-rotation-listeners/{log-file-rotation-listener-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-file-rotation-listener-name"+"}", url.PathEscape(parameterToString(r.logFileRotationListenerName, "")), -1)

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

type ApiGetLogFileRotationListenerRequest struct {
	ctx                         context.Context
	ApiService                  *LogFileRotationListenerApiService
	logFileRotationListenerName string
}

func (r ApiGetLogFileRotationListenerRequest) Execute() (*AddLogFileRotationListener200Response, *http.Response, error) {
	return r.ApiService.GetLogFileRotationListenerExecute(r)
}

/*
GetLogFileRotationListener Returns a single Log File Rotation Listener

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFileRotationListenerName Name of the Log File Rotation Listener to be read
	@return ApiGetLogFileRotationListenerRequest
*/
func (a *LogFileRotationListenerApiService) GetLogFileRotationListener(ctx context.Context, logFileRotationListenerName string) ApiGetLogFileRotationListenerRequest {
	return ApiGetLogFileRotationListenerRequest{
		ApiService:                  a,
		ctx:                         ctx,
		logFileRotationListenerName: logFileRotationListenerName,
	}
}

// Execute executes the request
//
//	@return AddLogFileRotationListener200Response
func (a *LogFileRotationListenerApiService) GetLogFileRotationListenerExecute(r ApiGetLogFileRotationListenerRequest) (*AddLogFileRotationListener200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLogFileRotationListener200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFileRotationListenerApiService.GetLogFileRotationListener")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-file-rotation-listeners/{log-file-rotation-listener-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-file-rotation-listener-name"+"}", url.PathEscape(parameterToString(r.logFileRotationListenerName, "")), -1)

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

type ApiUpdateLogFileRotationListenerRequest struct {
	ctx                         context.Context
	ApiService                  *LogFileRotationListenerApiService
	logFileRotationListenerName string
	updateRequest               *UpdateRequest
}

// Update an existing Log File Rotation Listener
func (r ApiUpdateLogFileRotationListenerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateLogFileRotationListenerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateLogFileRotationListenerRequest) Execute() (*AddLogFileRotationListener200Response, *http.Response, error) {
	return r.ApiService.UpdateLogFileRotationListenerExecute(r)
}

/*
UpdateLogFileRotationListener Update an existing Log File Rotation Listener by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logFileRotationListenerName Name of the Log File Rotation Listener to be updated
	@return ApiUpdateLogFileRotationListenerRequest
*/
func (a *LogFileRotationListenerApiService) UpdateLogFileRotationListener(ctx context.Context, logFileRotationListenerName string) ApiUpdateLogFileRotationListenerRequest {
	return ApiUpdateLogFileRotationListenerRequest{
		ApiService:                  a,
		ctx:                         ctx,
		logFileRotationListenerName: logFileRotationListenerName,
	}
}

// Execute executes the request
//
//	@return AddLogFileRotationListener200Response
func (a *LogFileRotationListenerApiService) UpdateLogFileRotationListenerExecute(r ApiUpdateLogFileRotationListenerRequest) (*AddLogFileRotationListener200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLogFileRotationListener200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogFileRotationListenerApiService.UpdateLogFileRotationListener")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/log-file-rotation-listeners/{log-file-rotation-listener-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"log-file-rotation-listener-name"+"}", url.PathEscape(parameterToString(r.logFileRotationListenerName, "")), -1)

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
