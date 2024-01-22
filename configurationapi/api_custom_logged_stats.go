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

// CustomLoggedStatsAPIService CustomLoggedStatsAPI service
type CustomLoggedStatsAPIService service

type ApiAddCustomLoggedStatsRequest struct {
	ctx                         context.Context
	ApiService                  *CustomLoggedStatsAPIService
	pluginName                  string
	addCustomLoggedStatsRequest *AddCustomLoggedStatsRequest
}

// Create a new Custom Logged Stats in the config
func (r ApiAddCustomLoggedStatsRequest) AddCustomLoggedStatsRequest(addCustomLoggedStatsRequest AddCustomLoggedStatsRequest) ApiAddCustomLoggedStatsRequest {
	r.addCustomLoggedStatsRequest = &addCustomLoggedStatsRequest
	return r
}

func (r ApiAddCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsResponse, *http.Response, error) {
	return r.ApiService.AddCustomLoggedStatsExecute(r)
}

/*
AddCustomLoggedStats Add a new Custom Logged Stats to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pluginName Name of the Plugin
	@return ApiAddCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsAPIService) AddCustomLoggedStats(ctx context.Context, pluginName string) ApiAddCustomLoggedStatsRequest {
	return ApiAddCustomLoggedStatsRequest{
		ApiService: a,
		ctx:        ctx,
		pluginName: pluginName,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsResponse
func (a *CustomLoggedStatsAPIService) AddCustomLoggedStatsExecute(r ApiAddCustomLoggedStatsRequest) (*CustomLoggedStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsAPIService.AddCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin-root/plugins/{plugin-name}/custom-logged-stats"
	localVarPath = strings.Replace(localVarPath, "{"+"plugin-name"+"}", url.PathEscape(parameterValueToString(r.pluginName, "pluginName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addCustomLoggedStatsRequest == nil {
		return localVarReturnValue, nil, reportError("addCustomLoggedStatsRequest is required and must be specified")
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
	localVarPostBody = r.addCustomLoggedStatsRequest
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

type ApiDeleteCustomLoggedStatsRequest struct {
	ctx                   context.Context
	ApiService            *CustomLoggedStatsAPIService
	customLoggedStatsName string
	pluginName            string
}

func (r ApiDeleteCustomLoggedStatsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomLoggedStatsExecute(r)
}

/*
DeleteCustomLoggedStats Delete a Custom Logged Stats

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customLoggedStatsName Name of the Custom Logged Stats
	@param pluginName Name of the Plugin
	@return ApiDeleteCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsAPIService) DeleteCustomLoggedStats(ctx context.Context, customLoggedStatsName string, pluginName string) ApiDeleteCustomLoggedStatsRequest {
	return ApiDeleteCustomLoggedStatsRequest{
		ApiService:            a,
		ctx:                   ctx,
		customLoggedStatsName: customLoggedStatsName,
		pluginName:            pluginName,
	}
}

// Execute executes the request
func (a *CustomLoggedStatsAPIService) DeleteCustomLoggedStatsExecute(r ApiDeleteCustomLoggedStatsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsAPIService.DeleteCustomLoggedStats")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin-root/plugins/{plugin-name}/custom-logged-stats/{custom-logged-stats-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom-logged-stats-name"+"}", url.PathEscape(parameterValueToString(r.customLoggedStatsName, "customLoggedStatsName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"plugin-name"+"}", url.PathEscape(parameterValueToString(r.pluginName, "pluginName")), -1)

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

type ApiGetCustomLoggedStatsRequest struct {
	ctx                   context.Context
	ApiService            *CustomLoggedStatsAPIService
	customLoggedStatsName string
	pluginName            string
}

func (r ApiGetCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsResponse, *http.Response, error) {
	return r.ApiService.GetCustomLoggedStatsExecute(r)
}

/*
GetCustomLoggedStats Returns a single Custom Logged Stats

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customLoggedStatsName Name of the Custom Logged Stats
	@param pluginName Name of the Plugin
	@return ApiGetCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsAPIService) GetCustomLoggedStats(ctx context.Context, customLoggedStatsName string, pluginName string) ApiGetCustomLoggedStatsRequest {
	return ApiGetCustomLoggedStatsRequest{
		ApiService:            a,
		ctx:                   ctx,
		customLoggedStatsName: customLoggedStatsName,
		pluginName:            pluginName,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsResponse
func (a *CustomLoggedStatsAPIService) GetCustomLoggedStatsExecute(r ApiGetCustomLoggedStatsRequest) (*CustomLoggedStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsAPIService.GetCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin-root/plugins/{plugin-name}/custom-logged-stats/{custom-logged-stats-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom-logged-stats-name"+"}", url.PathEscape(parameterValueToString(r.customLoggedStatsName, "customLoggedStatsName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"plugin-name"+"}", url.PathEscape(parameterValueToString(r.pluginName, "pluginName")), -1)

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

type ApiListCustomLoggedStatsRequest struct {
	ctx        context.Context
	ApiService *CustomLoggedStatsAPIService
	pluginName string
	filter     *string
}

// SCIM filter
func (r ApiListCustomLoggedStatsRequest) Filter(filter string) ApiListCustomLoggedStatsRequest {
	r.filter = &filter
	return r
}

func (r ApiListCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsListResponse, *http.Response, error) {
	return r.ApiService.ListCustomLoggedStatsExecute(r)
}

/*
ListCustomLoggedStats Returns a list of all Custom Logged Stats objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pluginName Name of the Plugin
	@return ApiListCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsAPIService) ListCustomLoggedStats(ctx context.Context, pluginName string) ApiListCustomLoggedStatsRequest {
	return ApiListCustomLoggedStatsRequest{
		ApiService: a,
		ctx:        ctx,
		pluginName: pluginName,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsListResponse
func (a *CustomLoggedStatsAPIService) ListCustomLoggedStatsExecute(r ApiListCustomLoggedStatsRequest) (*CustomLoggedStatsListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsAPIService.ListCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin-root/plugins/{plugin-name}/custom-logged-stats"
	localVarPath = strings.Replace(localVarPath, "{"+"plugin-name"+"}", url.PathEscape(parameterValueToString(r.pluginName, "pluginName")), -1)

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

type ApiUpdateCustomLoggedStatsRequest struct {
	ctx                   context.Context
	ApiService            *CustomLoggedStatsAPIService
	customLoggedStatsName string
	pluginName            string
	updateRequest         *UpdateRequest
}

// Update an existing Custom Logged Stats
func (r ApiUpdateCustomLoggedStatsRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateCustomLoggedStatsRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateCustomLoggedStatsRequest) Execute() (*CustomLoggedStatsResponse, *http.Response, error) {
	return r.ApiService.UpdateCustomLoggedStatsExecute(r)
}

/*
UpdateCustomLoggedStats Update an existing Custom Logged Stats by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customLoggedStatsName Name of the Custom Logged Stats
	@param pluginName Name of the Plugin
	@return ApiUpdateCustomLoggedStatsRequest
*/
func (a *CustomLoggedStatsAPIService) UpdateCustomLoggedStats(ctx context.Context, customLoggedStatsName string, pluginName string) ApiUpdateCustomLoggedStatsRequest {
	return ApiUpdateCustomLoggedStatsRequest{
		ApiService:            a,
		ctx:                   ctx,
		customLoggedStatsName: customLoggedStatsName,
		pluginName:            pluginName,
	}
}

// Execute executes the request
//
//	@return CustomLoggedStatsResponse
func (a *CustomLoggedStatsAPIService) UpdateCustomLoggedStatsExecute(r ApiUpdateCustomLoggedStatsRequest) (*CustomLoggedStatsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomLoggedStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomLoggedStatsAPIService.UpdateCustomLoggedStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin-root/plugins/{plugin-name}/custom-logged-stats/{custom-logged-stats-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom-logged-stats-name"+"}", url.PathEscape(parameterValueToString(r.customLoggedStatsName, "customLoggedStatsName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"plugin-name"+"}", url.PathEscape(parameterValueToString(r.pluginName, "pluginName")), -1)

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
