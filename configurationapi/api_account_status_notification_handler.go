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

// AccountStatusNotificationHandlerApiService AccountStatusNotificationHandlerApi service
type AccountStatusNotificationHandlerApiService service

type ApiAddAccountStatusNotificationHandlerRequest struct {
	ctx                                        context.Context
	ApiService                                 *AccountStatusNotificationHandlerApiService
	addAccountStatusNotificationHandlerRequest *AddAccountStatusNotificationHandlerRequest
}

// Create a new Account Status Notification Handler in the config
func (r ApiAddAccountStatusNotificationHandlerRequest) AddAccountStatusNotificationHandlerRequest(addAccountStatusNotificationHandlerRequest AddAccountStatusNotificationHandlerRequest) ApiAddAccountStatusNotificationHandlerRequest {
	r.addAccountStatusNotificationHandlerRequest = &addAccountStatusNotificationHandlerRequest
	return r
}

func (r ApiAddAccountStatusNotificationHandlerRequest) Execute() (*AddAccountStatusNotificationHandler200Response, *http.Response, error) {
	return r.ApiService.AddAccountStatusNotificationHandlerExecute(r)
}

/*
AddAccountStatusNotificationHandler Add a new Account Status Notification Handler to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddAccountStatusNotificationHandlerRequest
*/
func (a *AccountStatusNotificationHandlerApiService) AddAccountStatusNotificationHandler(ctx context.Context) ApiAddAccountStatusNotificationHandlerRequest {
	return ApiAddAccountStatusNotificationHandlerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddAccountStatusNotificationHandler200Response
func (a *AccountStatusNotificationHandlerApiService) AddAccountStatusNotificationHandlerExecute(r ApiAddAccountStatusNotificationHandlerRequest) (*AddAccountStatusNotificationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddAccountStatusNotificationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountStatusNotificationHandlerApiService.AddAccountStatusNotificationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-status-notification-handlers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addAccountStatusNotificationHandlerRequest == nil {
		return localVarReturnValue, nil, reportError("addAccountStatusNotificationHandlerRequest is required and must be specified")
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
	localVarPostBody = r.addAccountStatusNotificationHandlerRequest
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

type ApiDeleteAccountStatusNotificationHandlerRequest struct {
	ctx                                  context.Context
	ApiService                           *AccountStatusNotificationHandlerApiService
	accountStatusNotificationHandlerName string
}

func (r ApiDeleteAccountStatusNotificationHandlerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAccountStatusNotificationHandlerExecute(r)
}

/*
DeleteAccountStatusNotificationHandler Delete a Account Status Notification Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountStatusNotificationHandlerName Name of the Account Status Notification Handler
	@return ApiDeleteAccountStatusNotificationHandlerRequest
*/
func (a *AccountStatusNotificationHandlerApiService) DeleteAccountStatusNotificationHandler(ctx context.Context, accountStatusNotificationHandlerName string) ApiDeleteAccountStatusNotificationHandlerRequest {
	return ApiDeleteAccountStatusNotificationHandlerRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		accountStatusNotificationHandlerName: accountStatusNotificationHandlerName,
	}
}

// Execute executes the request
func (a *AccountStatusNotificationHandlerApiService) DeleteAccountStatusNotificationHandlerExecute(r ApiDeleteAccountStatusNotificationHandlerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountStatusNotificationHandlerApiService.DeleteAccountStatusNotificationHandler")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-status-notification-handlers/{account-status-notification-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"account-status-notification-handler-name"+"}", url.PathEscape(parameterValueToString(r.accountStatusNotificationHandlerName, "accountStatusNotificationHandlerName")), -1)

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

type ApiGetAccountStatusNotificationHandlerRequest struct {
	ctx                                  context.Context
	ApiService                           *AccountStatusNotificationHandlerApiService
	accountStatusNotificationHandlerName string
}

func (r ApiGetAccountStatusNotificationHandlerRequest) Execute() (*AddAccountStatusNotificationHandler200Response, *http.Response, error) {
	return r.ApiService.GetAccountStatusNotificationHandlerExecute(r)
}

/*
GetAccountStatusNotificationHandler Returns a single Account Status Notification Handler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountStatusNotificationHandlerName Name of the Account Status Notification Handler
	@return ApiGetAccountStatusNotificationHandlerRequest
*/
func (a *AccountStatusNotificationHandlerApiService) GetAccountStatusNotificationHandler(ctx context.Context, accountStatusNotificationHandlerName string) ApiGetAccountStatusNotificationHandlerRequest {
	return ApiGetAccountStatusNotificationHandlerRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		accountStatusNotificationHandlerName: accountStatusNotificationHandlerName,
	}
}

// Execute executes the request
//
//	@return AddAccountStatusNotificationHandler200Response
func (a *AccountStatusNotificationHandlerApiService) GetAccountStatusNotificationHandlerExecute(r ApiGetAccountStatusNotificationHandlerRequest) (*AddAccountStatusNotificationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddAccountStatusNotificationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountStatusNotificationHandlerApiService.GetAccountStatusNotificationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-status-notification-handlers/{account-status-notification-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"account-status-notification-handler-name"+"}", url.PathEscape(parameterValueToString(r.accountStatusNotificationHandlerName, "accountStatusNotificationHandlerName")), -1)

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

type ApiUpdateAccountStatusNotificationHandlerRequest struct {
	ctx                                  context.Context
	ApiService                           *AccountStatusNotificationHandlerApiService
	accountStatusNotificationHandlerName string
	updateRequest                        *UpdateRequest
}

// Update an existing Account Status Notification Handler
func (r ApiUpdateAccountStatusNotificationHandlerRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateAccountStatusNotificationHandlerRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateAccountStatusNotificationHandlerRequest) Execute() (*AddAccountStatusNotificationHandler200Response, *http.Response, error) {
	return r.ApiService.UpdateAccountStatusNotificationHandlerExecute(r)
}

/*
UpdateAccountStatusNotificationHandler Update an existing Account Status Notification Handler by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountStatusNotificationHandlerName Name of the Account Status Notification Handler
	@return ApiUpdateAccountStatusNotificationHandlerRequest
*/
func (a *AccountStatusNotificationHandlerApiService) UpdateAccountStatusNotificationHandler(ctx context.Context, accountStatusNotificationHandlerName string) ApiUpdateAccountStatusNotificationHandlerRequest {
	return ApiUpdateAccountStatusNotificationHandlerRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		accountStatusNotificationHandlerName: accountStatusNotificationHandlerName,
	}
}

// Execute executes the request
//
//	@return AddAccountStatusNotificationHandler200Response
func (a *AccountStatusNotificationHandlerApiService) UpdateAccountStatusNotificationHandlerExecute(r ApiUpdateAccountStatusNotificationHandlerRequest) (*AddAccountStatusNotificationHandler200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddAccountStatusNotificationHandler200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountStatusNotificationHandlerApiService.UpdateAccountStatusNotificationHandler")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-status-notification-handlers/{account-status-notification-handler-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"account-status-notification-handler-name"+"}", url.PathEscape(parameterValueToString(r.accountStatusNotificationHandlerName, "accountStatusNotificationHandlerName")), -1)

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
