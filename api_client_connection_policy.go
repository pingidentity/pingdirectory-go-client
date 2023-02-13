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

// ClientConnectionPolicyApiService ClientConnectionPolicyApi service
type ClientConnectionPolicyApiService service

type ApiAddClientConnectionPolicyRequest struct {
	ctx                              context.Context
	ApiService                       *ClientConnectionPolicyApiService
	addClientConnectionPolicyRequest *AddClientConnectionPolicyRequest
}

// Create a new Client Connection Policy in the config
func (r ApiAddClientConnectionPolicyRequest) AddClientConnectionPolicyRequest(addClientConnectionPolicyRequest AddClientConnectionPolicyRequest) ApiAddClientConnectionPolicyRequest {
	r.addClientConnectionPolicyRequest = &addClientConnectionPolicyRequest
	return r
}

func (r ApiAddClientConnectionPolicyRequest) Execute() (*ClientConnectionPolicyResponse, *http.Response, error) {
	return r.ApiService.AddClientConnectionPolicyExecute(r)
}

/*
AddClientConnectionPolicy Add a new Client Connection Policy to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddClientConnectionPolicyRequest
*/
func (a *ClientConnectionPolicyApiService) AddClientConnectionPolicy(ctx context.Context) ApiAddClientConnectionPolicyRequest {
	return ApiAddClientConnectionPolicyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClientConnectionPolicyResponse
func (a *ClientConnectionPolicyApiService) AddClientConnectionPolicyExecute(r ApiAddClientConnectionPolicyRequest) (*ClientConnectionPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientConnectionPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientConnectionPolicyApiService.AddClientConnectionPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client-connection-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addClientConnectionPolicyRequest == nil {
		return localVarReturnValue, nil, reportError("addClientConnectionPolicyRequest is required and must be specified")
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
	localVarPostBody = r.addClientConnectionPolicyRequest
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

type ApiDeleteClientConnectionPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *ClientConnectionPolicyApiService
	clientConnectionPolicyName string
}

func (r ApiDeleteClientConnectionPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClientConnectionPolicyExecute(r)
}

/*
DeleteClientConnectionPolicy Delete a Client Connection Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientConnectionPolicyName Name of the Client Connection Policy
	@return ApiDeleteClientConnectionPolicyRequest
*/
func (a *ClientConnectionPolicyApiService) DeleteClientConnectionPolicy(ctx context.Context, clientConnectionPolicyName string) ApiDeleteClientConnectionPolicyRequest {
	return ApiDeleteClientConnectionPolicyRequest{
		ApiService:                 a,
		ctx:                        ctx,
		clientConnectionPolicyName: clientConnectionPolicyName,
	}
}

// Execute executes the request
func (a *ClientConnectionPolicyApiService) DeleteClientConnectionPolicyExecute(r ApiDeleteClientConnectionPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientConnectionPolicyApiService.DeleteClientConnectionPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client-connection-policies/{client-connection-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"client-connection-policy-name"+"}", url.PathEscape(parameterToString(r.clientConnectionPolicyName, "")), -1)

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

type ApiGetClientConnectionPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *ClientConnectionPolicyApiService
	clientConnectionPolicyName string
}

func (r ApiGetClientConnectionPolicyRequest) Execute() (*ClientConnectionPolicyResponse, *http.Response, error) {
	return r.ApiService.GetClientConnectionPolicyExecute(r)
}

/*
GetClientConnectionPolicy Returns a single Client Connection Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientConnectionPolicyName Name of the Client Connection Policy
	@return ApiGetClientConnectionPolicyRequest
*/
func (a *ClientConnectionPolicyApiService) GetClientConnectionPolicy(ctx context.Context, clientConnectionPolicyName string) ApiGetClientConnectionPolicyRequest {
	return ApiGetClientConnectionPolicyRequest{
		ApiService:                 a,
		ctx:                        ctx,
		clientConnectionPolicyName: clientConnectionPolicyName,
	}
}

// Execute executes the request
//
//	@return ClientConnectionPolicyResponse
func (a *ClientConnectionPolicyApiService) GetClientConnectionPolicyExecute(r ApiGetClientConnectionPolicyRequest) (*ClientConnectionPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientConnectionPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientConnectionPolicyApiService.GetClientConnectionPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client-connection-policies/{client-connection-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"client-connection-policy-name"+"}", url.PathEscape(parameterToString(r.clientConnectionPolicyName, "")), -1)

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

type ApiUpdateClientConnectionPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *ClientConnectionPolicyApiService
	clientConnectionPolicyName string
	updateRequest              *UpdateRequest
}

// Update an existing Client Connection Policy
func (r ApiUpdateClientConnectionPolicyRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateClientConnectionPolicyRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateClientConnectionPolicyRequest) Execute() (*ClientConnectionPolicyResponse, *http.Response, error) {
	return r.ApiService.UpdateClientConnectionPolicyExecute(r)
}

/*
UpdateClientConnectionPolicy Update an existing Client Connection Policy by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientConnectionPolicyName Name of the Client Connection Policy
	@return ApiUpdateClientConnectionPolicyRequest
*/
func (a *ClientConnectionPolicyApiService) UpdateClientConnectionPolicy(ctx context.Context, clientConnectionPolicyName string) ApiUpdateClientConnectionPolicyRequest {
	return ApiUpdateClientConnectionPolicyRequest{
		ApiService:                 a,
		ctx:                        ctx,
		clientConnectionPolicyName: clientConnectionPolicyName,
	}
}

// Execute executes the request
//
//	@return ClientConnectionPolicyResponse
func (a *ClientConnectionPolicyApiService) UpdateClientConnectionPolicyExecute(r ApiUpdateClientConnectionPolicyRequest) (*ClientConnectionPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientConnectionPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientConnectionPolicyApiService.UpdateClientConnectionPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client-connection-policies/{client-connection-policy-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"client-connection-policy-name"+"}", url.PathEscape(parameterToString(r.clientConnectionPolicyName, "")), -1)

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
