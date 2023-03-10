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

// WebApplicationExtensionApiService WebApplicationExtensionApi service
type WebApplicationExtensionApiService service

type ApiAddWebApplicationExtensionRequest struct {
	ctx                                      context.Context
	ApiService                               *WebApplicationExtensionApiService
	addGenericWebApplicationExtensionRequest *AddGenericWebApplicationExtensionRequest
}

// Create a new Web Application Extension in the config
func (r ApiAddWebApplicationExtensionRequest) AddGenericWebApplicationExtensionRequest(addGenericWebApplicationExtensionRequest AddGenericWebApplicationExtensionRequest) ApiAddWebApplicationExtensionRequest {
	r.addGenericWebApplicationExtensionRequest = &addGenericWebApplicationExtensionRequest
	return r
}

func (r ApiAddWebApplicationExtensionRequest) Execute() (*AddWebApplicationExtension200Response, *http.Response, error) {
	return r.ApiService.AddWebApplicationExtensionExecute(r)
}

/*
AddWebApplicationExtension Add a new Web Application Extension to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddWebApplicationExtensionRequest
*/
func (a *WebApplicationExtensionApiService) AddWebApplicationExtension(ctx context.Context) ApiAddWebApplicationExtensionRequest {
	return ApiAddWebApplicationExtensionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddWebApplicationExtension200Response
func (a *WebApplicationExtensionApiService) AddWebApplicationExtensionExecute(r ApiAddWebApplicationExtensionRequest) (*AddWebApplicationExtension200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddWebApplicationExtension200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebApplicationExtensionApiService.AddWebApplicationExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/web-application-extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addGenericWebApplicationExtensionRequest == nil {
		return localVarReturnValue, nil, reportError("addGenericWebApplicationExtensionRequest is required and must be specified")
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
	localVarPostBody = r.addGenericWebApplicationExtensionRequest
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

type ApiDeleteWebApplicationExtensionRequest struct {
	ctx                         context.Context
	ApiService                  *WebApplicationExtensionApiService
	webApplicationExtensionName string
}

func (r ApiDeleteWebApplicationExtensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWebApplicationExtensionExecute(r)
}

/*
DeleteWebApplicationExtension Delete a Web Application Extension

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webApplicationExtensionName Name of the Web Application Extension
	@return ApiDeleteWebApplicationExtensionRequest
*/
func (a *WebApplicationExtensionApiService) DeleteWebApplicationExtension(ctx context.Context, webApplicationExtensionName string) ApiDeleteWebApplicationExtensionRequest {
	return ApiDeleteWebApplicationExtensionRequest{
		ApiService:                  a,
		ctx:                         ctx,
		webApplicationExtensionName: webApplicationExtensionName,
	}
}

// Execute executes the request
func (a *WebApplicationExtensionApiService) DeleteWebApplicationExtensionExecute(r ApiDeleteWebApplicationExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebApplicationExtensionApiService.DeleteWebApplicationExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/web-application-extensions/{web-application-extension-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"web-application-extension-name"+"}", url.PathEscape(parameterValueToString(r.webApplicationExtensionName, "webApplicationExtensionName")), -1)

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

type ApiGetWebApplicationExtensionRequest struct {
	ctx                         context.Context
	ApiService                  *WebApplicationExtensionApiService
	webApplicationExtensionName string
}

func (r ApiGetWebApplicationExtensionRequest) Execute() (*GetWebApplicationExtension200Response, *http.Response, error) {
	return r.ApiService.GetWebApplicationExtensionExecute(r)
}

/*
GetWebApplicationExtension Returns a single Web Application Extension

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webApplicationExtensionName Name of the Web Application Extension
	@return ApiGetWebApplicationExtensionRequest
*/
func (a *WebApplicationExtensionApiService) GetWebApplicationExtension(ctx context.Context, webApplicationExtensionName string) ApiGetWebApplicationExtensionRequest {
	return ApiGetWebApplicationExtensionRequest{
		ApiService:                  a,
		ctx:                         ctx,
		webApplicationExtensionName: webApplicationExtensionName,
	}
}

// Execute executes the request
//
//	@return GetWebApplicationExtension200Response
func (a *WebApplicationExtensionApiService) GetWebApplicationExtensionExecute(r ApiGetWebApplicationExtensionRequest) (*GetWebApplicationExtension200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetWebApplicationExtension200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebApplicationExtensionApiService.GetWebApplicationExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/web-application-extensions/{web-application-extension-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"web-application-extension-name"+"}", url.PathEscape(parameterValueToString(r.webApplicationExtensionName, "webApplicationExtensionName")), -1)

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

type ApiUpdateWebApplicationExtensionRequest struct {
	ctx                         context.Context
	ApiService                  *WebApplicationExtensionApiService
	webApplicationExtensionName string
	updateRequest               *UpdateRequest
}

// Update an existing Web Application Extension
func (r ApiUpdateWebApplicationExtensionRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateWebApplicationExtensionRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateWebApplicationExtensionRequest) Execute() (*GetWebApplicationExtension200Response, *http.Response, error) {
	return r.ApiService.UpdateWebApplicationExtensionExecute(r)
}

/*
UpdateWebApplicationExtension Update an existing Web Application Extension by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webApplicationExtensionName Name of the Web Application Extension
	@return ApiUpdateWebApplicationExtensionRequest
*/
func (a *WebApplicationExtensionApiService) UpdateWebApplicationExtension(ctx context.Context, webApplicationExtensionName string) ApiUpdateWebApplicationExtensionRequest {
	return ApiUpdateWebApplicationExtensionRequest{
		ApiService:                  a,
		ctx:                         ctx,
		webApplicationExtensionName: webApplicationExtensionName,
	}
}

// Execute executes the request
//
//	@return GetWebApplicationExtension200Response
func (a *WebApplicationExtensionApiService) UpdateWebApplicationExtensionExecute(r ApiUpdateWebApplicationExtensionRequest) (*GetWebApplicationExtension200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetWebApplicationExtension200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebApplicationExtensionApiService.UpdateWebApplicationExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/web-application-extensions/{web-application-extension-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"web-application-extension-name"+"}", url.PathEscape(parameterValueToString(r.webApplicationExtensionName, "webApplicationExtensionName")), -1)

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
