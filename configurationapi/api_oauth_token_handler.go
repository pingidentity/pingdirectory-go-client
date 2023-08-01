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

// OauthTokenHandlerApiService OauthTokenHandlerApi service
type OauthTokenHandlerApiService service

type ApiAddOauthTokenHandlerRequest struct {
	ctx                         context.Context
	ApiService                  *OauthTokenHandlerApiService
	addOauthTokenHandlerRequest *AddOauthTokenHandlerRequest
}

// Create a new OAuth Token Handler in the config
func (r ApiAddOauthTokenHandlerRequest) AddOauthTokenHandlerRequest(addOauthTokenHandlerRequest AddOauthTokenHandlerRequest) ApiAddOauthTokenHandlerRequest {
	r.addOauthTokenHandlerRequest = &addOauthTokenHandlerRequest
	return r
}

func (r ApiAddOauthTokenHandlerRequest) Execute() (*AddOauthTokenHandler200Response, *http.Response, error) {
	return r.ApiService.AddOauthTokenHandlerExecute(r)
}

/*
AddOauthTokenHandler Add a new OAuth Token Handler to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddOauthTokenHandlerRequest
*/
func (a *OauthTokenHandlerApiService) AddOauthTokenHandler(ctx context.Context) ApiAddOauthTokenHandlerRequest {
	return ApiAddOauthTokenHandlerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddOauthTokenHandler200Response
func (a *OauthTokenHandlerApiService) AddOauthTokenHandlerExecute(r ApiAddOauthTokenHandlerRequest) (*AddOauthTokenHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddOauthTokenHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenHandlerApiService.AddOauthTokenHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth-token-handlers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addOauthTokenHandlerRequest == nil {
		return localVarReturnValue, nil, reportError("addOauthTokenHandlerRequest is required and must be specified")
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
	localVarPostBody = r.addOauthTokenHandlerRequest
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

type ApiDeleteOauthTokenHandlerRequest struct {
	ctx                   context.Context
	ApiService            *OauthTokenHandlerApiService
	oauthTokenHandlerName string
}

func (r ApiDeleteOauthTokenHandlerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOauthTokenHandlerExecute(r)
}

/*
DeleteOauthTokenHandler Delete a OAuth Token Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param oauthTokenHandlerName Name of the OAuth Token Handler
	@return ApiDeleteOauthTokenHandlerRequest
*/
func (a *OauthTokenHandlerApiService) DeleteOauthTokenHandler(ctx context.Context, oauthTokenHandlerName string) ApiDeleteOauthTokenHandlerRequest {
	return ApiDeleteOauthTokenHandlerRequest{
		ApiService:            a,
		ctx:                   ctx,
		oauthTokenHandlerName: oauthTokenHandlerName,
	}
}

// Execute executes the request
func (a *OauthTokenHandlerApiService) DeleteOauthTokenHandlerExecute(r ApiDeleteOauthTokenHandlerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenHandlerApiService.DeleteOauthTokenHandler")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth-token-handlers/{oauth-token-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"oauth-token-handler-name"+"}", url.PathEscape(parameterValueToString(r.oauthTokenHandlerName, "oauthTokenHandlerName")), -1)

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

type ApiGetOauthTokenHandlerRequest struct {
	ctx                   context.Context
	ApiService            *OauthTokenHandlerApiService
	oauthTokenHandlerName string
}

func (r ApiGetOauthTokenHandlerRequest) Execute() (*AddOauthTokenHandler200Response, *http.Response, error) {
	return r.ApiService.GetOauthTokenHandlerExecute(r)
}

/*
GetOauthTokenHandler Returns a single OAuth Token Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param oauthTokenHandlerName Name of the OAuth Token Handler
	@return ApiGetOauthTokenHandlerRequest
*/
func (a *OauthTokenHandlerApiService) GetOauthTokenHandler(ctx context.Context, oauthTokenHandlerName string) ApiGetOauthTokenHandlerRequest {
	return ApiGetOauthTokenHandlerRequest{
		ApiService:            a,
		ctx:                   ctx,
		oauthTokenHandlerName: oauthTokenHandlerName,
	}
}

// Execute executes the request
//
//	@return AddOauthTokenHandler200Response
func (a *OauthTokenHandlerApiService) GetOauthTokenHandlerExecute(r ApiGetOauthTokenHandlerRequest) (*AddOauthTokenHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddOauthTokenHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenHandlerApiService.GetOauthTokenHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth-token-handlers/{oauth-token-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"oauth-token-handler-name"+"}", url.PathEscape(parameterValueToString(r.oauthTokenHandlerName, "oauthTokenHandlerName")), -1)

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

type ApiListOauthTokenHandlersRequest struct {
	ctx        context.Context
	ApiService *OauthTokenHandlerApiService
	filter     *string
}

// SCIM filter
func (r ApiListOauthTokenHandlersRequest) Filter(filter string) ApiListOauthTokenHandlersRequest {
	r.filter = &filter
	return r
}

func (r ApiListOauthTokenHandlersRequest) Execute() (*OauthTokenHandlerListResponse, *http.Response, error) {
	return r.ApiService.ListOauthTokenHandlersExecute(r)
}

/*
ListOauthTokenHandlers Returns a list of all OAuth Token Handler objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListOauthTokenHandlersRequest
*/
func (a *OauthTokenHandlerApiService) ListOauthTokenHandlers(ctx context.Context) ApiListOauthTokenHandlersRequest {
	return ApiListOauthTokenHandlersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OauthTokenHandlerListResponse
func (a *OauthTokenHandlerApiService) ListOauthTokenHandlersExecute(r ApiListOauthTokenHandlersRequest) (*OauthTokenHandlerListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OauthTokenHandlerListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenHandlerApiService.ListOauthTokenHandlers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth-token-handlers"

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

type ApiUpdateOauthTokenHandlerRequest struct {
	ctx                   context.Context
	ApiService            *OauthTokenHandlerApiService
	oauthTokenHandlerName string
	updateRequest         *UpdateRequest
}

// Update an existing OAuth Token Handler
func (r ApiUpdateOauthTokenHandlerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateOauthTokenHandlerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateOauthTokenHandlerRequest) Execute() (*AddOauthTokenHandler200Response, *http.Response, error) {
	return r.ApiService.UpdateOauthTokenHandlerExecute(r)
}

/*
UpdateOauthTokenHandler Update an existing OAuth Token Handler by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param oauthTokenHandlerName Name of the OAuth Token Handler
	@return ApiUpdateOauthTokenHandlerRequest
*/
func (a *OauthTokenHandlerApiService) UpdateOauthTokenHandler(ctx context.Context, oauthTokenHandlerName string) ApiUpdateOauthTokenHandlerRequest {
	return ApiUpdateOauthTokenHandlerRequest{
		ApiService:            a,
		ctx:                   ctx,
		oauthTokenHandlerName: oauthTokenHandlerName,
	}
}

// Execute executes the request
//
//	@return AddOauthTokenHandler200Response
func (a *OauthTokenHandlerApiService) UpdateOauthTokenHandlerExecute(r ApiUpdateOauthTokenHandlerRequest) (*AddOauthTokenHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddOauthTokenHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenHandlerApiService.UpdateOauthTokenHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth-token-handlers/{oauth-token-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"oauth-token-handler-name"+"}", url.PathEscape(parameterValueToString(r.oauthTokenHandlerName, "oauthTokenHandlerName")), -1)

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
