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

// ChangeSubscriptionApiService ChangeSubscriptionApi service
type ChangeSubscriptionApiService service

type ApiAddChangeSubscriptionRequest struct {
	ctx                          context.Context
	ApiService                   *ChangeSubscriptionApiService
	addChangeSubscriptionRequest *AddChangeSubscriptionRequest
}

// Create a new Change Subscription in the config
func (r ApiAddChangeSubscriptionRequest) AddChangeSubscriptionRequest(addChangeSubscriptionRequest AddChangeSubscriptionRequest) ApiAddChangeSubscriptionRequest {
	r.addChangeSubscriptionRequest = &addChangeSubscriptionRequest
	return r
}

func (r ApiAddChangeSubscriptionRequest) Execute() (*ChangeSubscriptionResponse, *http.Response, error) {
	return r.ApiService.AddChangeSubscriptionExecute(r)
}

/*
AddChangeSubscription Add a new Change Subscription to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddChangeSubscriptionRequest
*/
func (a *ChangeSubscriptionApiService) AddChangeSubscription(ctx context.Context) ApiAddChangeSubscriptionRequest {
	return ApiAddChangeSubscriptionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeSubscriptionResponse
func (a *ChangeSubscriptionApiService) AddChangeSubscriptionExecute(r ApiAddChangeSubscriptionRequest) (*ChangeSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ChangeSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionApiService.AddChangeSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addChangeSubscriptionRequest == nil {
		return localVarReturnValue, nil, reportError("addChangeSubscriptionRequest is required and must be specified")
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
	localVarPostBody = r.addChangeSubscriptionRequest
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

type ApiDeleteChangeSubscriptionRequest struct {
	ctx                    context.Context
	ApiService             *ChangeSubscriptionApiService
	changeSubscriptionName string
}

func (r ApiDeleteChangeSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteChangeSubscriptionExecute(r)
}

/*
DeleteChangeSubscription Delete a Change Subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param changeSubscriptionName Name of the Change Subscription
	@return ApiDeleteChangeSubscriptionRequest
*/
func (a *ChangeSubscriptionApiService) DeleteChangeSubscription(ctx context.Context, changeSubscriptionName string) ApiDeleteChangeSubscriptionRequest {
	return ApiDeleteChangeSubscriptionRequest{
		ApiService:             a,
		ctx:                    ctx,
		changeSubscriptionName: changeSubscriptionName,
	}
}

// Execute executes the request
func (a *ChangeSubscriptionApiService) DeleteChangeSubscriptionExecute(r ApiDeleteChangeSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionApiService.DeleteChangeSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscriptions/{change-subscription-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"change-subscription-name"+"}", url.PathEscape(parameterValueToString(r.changeSubscriptionName, "changeSubscriptionName")), -1)

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

type ApiGetChangeSubscriptionRequest struct {
	ctx                    context.Context
	ApiService             *ChangeSubscriptionApiService
	changeSubscriptionName string
}

func (r ApiGetChangeSubscriptionRequest) Execute() (*ChangeSubscriptionResponse, *http.Response, error) {
	return r.ApiService.GetChangeSubscriptionExecute(r)
}

/*
GetChangeSubscription Returns a single Change Subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param changeSubscriptionName Name of the Change Subscription
	@return ApiGetChangeSubscriptionRequest
*/
func (a *ChangeSubscriptionApiService) GetChangeSubscription(ctx context.Context, changeSubscriptionName string) ApiGetChangeSubscriptionRequest {
	return ApiGetChangeSubscriptionRequest{
		ApiService:             a,
		ctx:                    ctx,
		changeSubscriptionName: changeSubscriptionName,
	}
}

// Execute executes the request
//
//	@return ChangeSubscriptionResponse
func (a *ChangeSubscriptionApiService) GetChangeSubscriptionExecute(r ApiGetChangeSubscriptionRequest) (*ChangeSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ChangeSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionApiService.GetChangeSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscriptions/{change-subscription-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"change-subscription-name"+"}", url.PathEscape(parameterValueToString(r.changeSubscriptionName, "changeSubscriptionName")), -1)

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

type ApiListChangeSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *ChangeSubscriptionApiService
	filter     *string
}

// SCIM filter
func (r ApiListChangeSubscriptionsRequest) Filter(filter string) ApiListChangeSubscriptionsRequest {
	r.filter = &filter
	return r
}

func (r ApiListChangeSubscriptionsRequest) Execute() (*ChangeSubscriptionListResponse, *http.Response, error) {
	return r.ApiService.ListChangeSubscriptionsExecute(r)
}

/*
ListChangeSubscriptions Returns a list of all Change Subscription objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListChangeSubscriptionsRequest
*/
func (a *ChangeSubscriptionApiService) ListChangeSubscriptions(ctx context.Context) ApiListChangeSubscriptionsRequest {
	return ApiListChangeSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeSubscriptionListResponse
func (a *ChangeSubscriptionApiService) ListChangeSubscriptionsExecute(r ApiListChangeSubscriptionsRequest) (*ChangeSubscriptionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ChangeSubscriptionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionApiService.ListChangeSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscriptions"

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

type ApiUpdateChangeSubscriptionRequest struct {
	ctx                    context.Context
	ApiService             *ChangeSubscriptionApiService
	changeSubscriptionName string
	updateRequest          *UpdateRequest
}

// Update an existing Change Subscription
func (r ApiUpdateChangeSubscriptionRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateChangeSubscriptionRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateChangeSubscriptionRequest) Execute() (*ChangeSubscriptionResponse, *http.Response, error) {
	return r.ApiService.UpdateChangeSubscriptionExecute(r)
}

/*
UpdateChangeSubscription Update an existing Change Subscription by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param changeSubscriptionName Name of the Change Subscription
	@return ApiUpdateChangeSubscriptionRequest
*/
func (a *ChangeSubscriptionApiService) UpdateChangeSubscription(ctx context.Context, changeSubscriptionName string) ApiUpdateChangeSubscriptionRequest {
	return ApiUpdateChangeSubscriptionRequest{
		ApiService:             a,
		ctx:                    ctx,
		changeSubscriptionName: changeSubscriptionName,
	}
}

// Execute executes the request
//
//	@return ChangeSubscriptionResponse
func (a *ChangeSubscriptionApiService) UpdateChangeSubscriptionExecute(r ApiUpdateChangeSubscriptionRequest) (*ChangeSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ChangeSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChangeSubscriptionApiService.UpdateChangeSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/change-subscriptions/{change-subscription-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"change-subscription-name"+"}", url.PathEscape(parameterValueToString(r.changeSubscriptionName, "changeSubscriptionName")), -1)

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
