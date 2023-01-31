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

// ChangeSubscriptionHandlerApiService ChangeSubscriptionHandlerApi service
type ChangeSubscriptionHandlerApiService service

type ApiAddChangeSubscriptionHandlerRequest struct {
	ctx                                 context.Context
	ApiService                          *ChangeSubscriptionHandlerApiService
	addChangeSubscriptionHandlerRequest *AddChangeSubscriptionHandlerRequest
}

// Create a new Change Subscription Handler in the config
func (r ApiAddChangeSubscriptionHandlerRequest) AddChangeSubscriptionHandlerRequest(addChangeSubscriptionHandlerRequest AddChangeSubscriptionHandlerRequest) ApiAddChangeSubscriptionHandlerRequest {
	r.addChangeSubscriptionHandlerRequest = &addChangeSubscriptionHandlerRequest
	return r
}

func (r ApiAddChangeSubscriptionHandlerRequest) Execute() (*AddChangeSubscriptionHandler200Response, *http.Response, error) {
	return r.ApiService.AddChangeSubscriptionHandlerExecute(r)
}

/*
AddChangeSubscriptionHandler Add a new Change Subscription Handler to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddChangeSubscriptionHandlerRequest
*/
func (a *ChangeSubscriptionHandlerApiService) AddChangeSubscriptionHandler(ctx context.Context) ApiAddChangeSubscriptionHandlerRequest {
	return ApiAddChangeSubscriptionHandlerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddChangeSubscriptionHandler200Response
func (a *ChangeSubscriptionHandlerApiService) AddChangeSubscriptionHandlerExecute(r ApiAddChangeSubscriptionHandlerRequest) (*AddChangeSubscriptionHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddChangeSubscriptionHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionHandlerApiService.AddChangeSubscriptionHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscription-handlers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addChangeSubscriptionHandlerRequest == nil {
		return localVarReturnValue, nil, reportError("addChangeSubscriptionHandlerRequest is required and must be specified")
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
	localVarPostBody = r.addChangeSubscriptionHandlerRequest
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

type ApiDeleteChangeSubscriptionHandlerRequest struct {
	ctx                           context.Context
	ApiService                    *ChangeSubscriptionHandlerApiService
	changeSubscriptionHandlerName string
}

func (r ApiDeleteChangeSubscriptionHandlerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteChangeSubscriptionHandlerExecute(r)
}

/*
DeleteChangeSubscriptionHandler Delete a Change Subscription Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param changeSubscriptionHandlerName Name of the Change Subscription Handler to be deleted
	@return ApiDeleteChangeSubscriptionHandlerRequest
*/
func (a *ChangeSubscriptionHandlerApiService) DeleteChangeSubscriptionHandler(ctx context.Context, changeSubscriptionHandlerName string) ApiDeleteChangeSubscriptionHandlerRequest {
	return ApiDeleteChangeSubscriptionHandlerRequest{
		ApiService:                    a,
		ctx:                           ctx,
		changeSubscriptionHandlerName: changeSubscriptionHandlerName,
	}
}

// Execute executes the request
func (a *ChangeSubscriptionHandlerApiService) DeleteChangeSubscriptionHandlerExecute(r ApiDeleteChangeSubscriptionHandlerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionHandlerApiService.DeleteChangeSubscriptionHandler")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscription-handlers/{change-subscription-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"change-subscription-handler-name"+"}", url.PathEscape(parameterToString(r.changeSubscriptionHandlerName, "")), -1)

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

type ApiGetChangeSubscriptionHandlerRequest struct {
	ctx                           context.Context
	ApiService                    *ChangeSubscriptionHandlerApiService
	changeSubscriptionHandlerName string
}

func (r ApiGetChangeSubscriptionHandlerRequest) Execute() (*AddChangeSubscriptionHandler200Response, *http.Response, error) {
	return r.ApiService.GetChangeSubscriptionHandlerExecute(r)
}

/*
GetChangeSubscriptionHandler Returns a single Change Subscription Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param changeSubscriptionHandlerName Name of the Change Subscription Handler to be read
	@return ApiGetChangeSubscriptionHandlerRequest
*/
func (a *ChangeSubscriptionHandlerApiService) GetChangeSubscriptionHandler(ctx context.Context, changeSubscriptionHandlerName string) ApiGetChangeSubscriptionHandlerRequest {
	return ApiGetChangeSubscriptionHandlerRequest{
		ApiService:                    a,
		ctx:                           ctx,
		changeSubscriptionHandlerName: changeSubscriptionHandlerName,
	}
}

// Execute executes the request
//
//	@return AddChangeSubscriptionHandler200Response
func (a *ChangeSubscriptionHandlerApiService) GetChangeSubscriptionHandlerExecute(r ApiGetChangeSubscriptionHandlerRequest) (*AddChangeSubscriptionHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddChangeSubscriptionHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionHandlerApiService.GetChangeSubscriptionHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscription-handlers/{change-subscription-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"change-subscription-handler-name"+"}", url.PathEscape(parameterToString(r.changeSubscriptionHandlerName, "")), -1)

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

type ApiUpdateChangeSubscriptionHandlerRequest struct {
	ctx                           context.Context
	ApiService                    *ChangeSubscriptionHandlerApiService
	changeSubscriptionHandlerName string
	updateRequest                 *UpdateRequest
}

// Update an existing Change Subscription Handler
func (r ApiUpdateChangeSubscriptionHandlerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateChangeSubscriptionHandlerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateChangeSubscriptionHandlerRequest) Execute() (*AddChangeSubscriptionHandler200Response, *http.Response, error) {
	return r.ApiService.UpdateChangeSubscriptionHandlerExecute(r)
}

/*
UpdateChangeSubscriptionHandler Update an existing Change Subscription Handler by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param changeSubscriptionHandlerName Name of the Change Subscription Handler to be updated
	@return ApiUpdateChangeSubscriptionHandlerRequest
*/
func (a *ChangeSubscriptionHandlerApiService) UpdateChangeSubscriptionHandler(ctx context.Context, changeSubscriptionHandlerName string) ApiUpdateChangeSubscriptionHandlerRequest {
	return ApiUpdateChangeSubscriptionHandlerRequest{
		ApiService:                    a,
		ctx:                           ctx,
		changeSubscriptionHandlerName: changeSubscriptionHandlerName,
	}
}

// Execute executes the request
//
//	@return AddChangeSubscriptionHandler200Response
func (a *ChangeSubscriptionHandlerApiService) UpdateChangeSubscriptionHandlerExecute(r ApiUpdateChangeSubscriptionHandlerRequest) (*AddChangeSubscriptionHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddChangeSubscriptionHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionHandlerApiService.UpdateChangeSubscriptionHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscription-handlers/{change-subscription-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"change-subscription-handler-name"+"}", url.PathEscape(parameterToString(r.changeSubscriptionHandlerName, "")), -1)

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
