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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RootDnUserApiService RootDnUserApi service
type RootDnUserApiService service

type ApiAddRootDnUserRequest struct {
	ctx                  context.Context
	ApiService           *RootDnUserApiService
	addRootDnUserRequest *AddRootDnUserRequest
}

// Create a new Root DN User in the config
func (r ApiAddRootDnUserRequest) AddRootDnUserRequest(addRootDnUserRequest AddRootDnUserRequest) ApiAddRootDnUserRequest {
	r.addRootDnUserRequest = &addRootDnUserRequest
	return r
}

func (r ApiAddRootDnUserRequest) Execute() (*RootDnUserResponse, *http.Response, error) {
	return r.ApiService.AddRootDnUserExecute(r)
}

/*
AddRootDnUser Add a new Root DN User to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddRootDnUserRequest
*/
func (a *RootDnUserApiService) AddRootDnUser(ctx context.Context) ApiAddRootDnUserRequest {
	return ApiAddRootDnUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RootDnUserResponse
func (a *RootDnUserApiService) AddRootDnUserExecute(r ApiAddRootDnUserRequest) (*RootDnUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RootDnUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootDnUserApiService.AddRootDnUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/root-dn/root-dn-users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addRootDnUserRequest == nil {
		return localVarReturnValue, nil, reportError("addRootDnUserRequest is required and must be specified")
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
	localVarPostBody = r.addRootDnUserRequest
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

type ApiDeleteRootDnUserRequest struct {
	ctx            context.Context
	ApiService     *RootDnUserApiService
	rootDnUserName string
}

func (r ApiDeleteRootDnUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRootDnUserExecute(r)
}

/*
DeleteRootDnUser Delete a Root DN User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param rootDnUserName Name of the Root DN User
	@return ApiDeleteRootDnUserRequest
*/
func (a *RootDnUserApiService) DeleteRootDnUser(ctx context.Context, rootDnUserName string) ApiDeleteRootDnUserRequest {
	return ApiDeleteRootDnUserRequest{
		ApiService:     a,
		ctx:            ctx,
		rootDnUserName: rootDnUserName,
	}
}

// Execute executes the request
func (a *RootDnUserApiService) DeleteRootDnUserExecute(r ApiDeleteRootDnUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootDnUserApiService.DeleteRootDnUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/root-dn/root-dn-users/{root-dn-user-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"root-dn-user-name"+"}", url.PathEscape(parameterToString(r.rootDnUserName, "")), -1)

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

type ApiGetRootDnUserRequest struct {
	ctx            context.Context
	ApiService     *RootDnUserApiService
	rootDnUserName string
}

func (r ApiGetRootDnUserRequest) Execute() (*RootDnUserResponse, *http.Response, error) {
	return r.ApiService.GetRootDnUserExecute(r)
}

/*
GetRootDnUser Returns a single Root DN User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param rootDnUserName Name of the Root DN User
	@return ApiGetRootDnUserRequest
*/
func (a *RootDnUserApiService) GetRootDnUser(ctx context.Context, rootDnUserName string) ApiGetRootDnUserRequest {
	return ApiGetRootDnUserRequest{
		ApiService:     a,
		ctx:            ctx,
		rootDnUserName: rootDnUserName,
	}
}

// Execute executes the request
//
//	@return RootDnUserResponse
func (a *RootDnUserApiService) GetRootDnUserExecute(r ApiGetRootDnUserRequest) (*RootDnUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RootDnUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootDnUserApiService.GetRootDnUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/root-dn/root-dn-users/{root-dn-user-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"root-dn-user-name"+"}", url.PathEscape(parameterToString(r.rootDnUserName, "")), -1)

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

type ApiUpdateRootDnUserRequest struct {
	ctx            context.Context
	ApiService     *RootDnUserApiService
	rootDnUserName string
	updateRequest  *UpdateRequest
}

// Update an existing Root DN User
func (r ApiUpdateRootDnUserRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateRootDnUserRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateRootDnUserRequest) Execute() (*RootDnUserResponse, *http.Response, error) {
	return r.ApiService.UpdateRootDnUserExecute(r)
}

/*
UpdateRootDnUser Update an existing Root DN User by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param rootDnUserName Name of the Root DN User
	@return ApiUpdateRootDnUserRequest
*/
func (a *RootDnUserApiService) UpdateRootDnUser(ctx context.Context, rootDnUserName string) ApiUpdateRootDnUserRequest {
	return ApiUpdateRootDnUserRequest{
		ApiService:     a,
		ctx:            ctx,
		rootDnUserName: rootDnUserName,
	}
}

// Execute executes the request
//
//	@return RootDnUserResponse
func (a *RootDnUserApiService) UpdateRootDnUserExecute(r ApiUpdateRootDnUserRequest) (*RootDnUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RootDnUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootDnUserApiService.UpdateRootDnUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/root-dn/root-dn-users/{root-dn-user-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"root-dn-user-name"+"}", url.PathEscape(parameterToString(r.rootDnUserName, "")), -1)

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