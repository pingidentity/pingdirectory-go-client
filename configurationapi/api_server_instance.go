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

// ServerInstanceAPIService ServerInstanceAPI service
type ServerInstanceAPIService service

type ApiGetServerInstanceRequest struct {
	ctx                context.Context
	ApiService         *ServerInstanceAPIService
	serverInstanceName string
}

func (r ApiGetServerInstanceRequest) Execute() (*GetServerInstance200Response, *http.Response, error) {
	return r.ApiService.GetServerInstanceExecute(r)
}

/*
GetServerInstance Returns a single Server Instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serverInstanceName Name of the Server Instance
	@return ApiGetServerInstanceRequest
*/
func (a *ServerInstanceAPIService) GetServerInstance(ctx context.Context, serverInstanceName string) ApiGetServerInstanceRequest {
	return ApiGetServerInstanceRequest{
		ApiService:         a,
		ctx:                ctx,
		serverInstanceName: serverInstanceName,
	}
}

// Execute executes the request
//
//	@return GetServerInstance200Response
func (a *ServerInstanceAPIService) GetServerInstanceExecute(r ApiGetServerInstanceRequest) (*GetServerInstance200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetServerInstance200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerInstanceAPIService.GetServerInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server-instances/{server-instance-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"server-instance-name"+"}", url.PathEscape(parameterValueToString(r.serverInstanceName, "serverInstanceName")), -1)

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

type ApiListServerInstancesRequest struct {
	ctx        context.Context
	ApiService *ServerInstanceAPIService
	filter     *string
}

// SCIM filter
func (r ApiListServerInstancesRequest) Filter(filter string) ApiListServerInstancesRequest {
	r.filter = &filter
	return r
}

func (r ApiListServerInstancesRequest) Execute() (*ServerInstanceListResponse, *http.Response, error) {
	return r.ApiService.ListServerInstancesExecute(r)
}

/*
ListServerInstances Returns a list of all Server Instance objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListServerInstancesRequest
*/
func (a *ServerInstanceAPIService) ListServerInstances(ctx context.Context) ApiListServerInstancesRequest {
	return ApiListServerInstancesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ServerInstanceListResponse
func (a *ServerInstanceAPIService) ListServerInstancesExecute(r ApiListServerInstancesRequest) (*ServerInstanceListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerInstanceListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerInstanceAPIService.ListServerInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server-instances"

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

type ApiUpdateServerInstanceRequest struct {
	ctx                context.Context
	ApiService         *ServerInstanceAPIService
	serverInstanceName string
	updateRequest      *UpdateRequest
}

// Update an existing Server Instance
func (r ApiUpdateServerInstanceRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateServerInstanceRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateServerInstanceRequest) Execute() (*GetServerInstance200Response, *http.Response, error) {
	return r.ApiService.UpdateServerInstanceExecute(r)
}

/*
UpdateServerInstance Update an existing Server Instance by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serverInstanceName Name of the Server Instance
	@return ApiUpdateServerInstanceRequest
*/
func (a *ServerInstanceAPIService) UpdateServerInstance(ctx context.Context, serverInstanceName string) ApiUpdateServerInstanceRequest {
	return ApiUpdateServerInstanceRequest{
		ApiService:         a,
		ctx:                ctx,
		serverInstanceName: serverInstanceName,
	}
}

// Execute executes the request
//
//	@return GetServerInstance200Response
func (a *ServerInstanceAPIService) UpdateServerInstanceExecute(r ApiUpdateServerInstanceRequest) (*GetServerInstance200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetServerInstance200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerInstanceAPIService.UpdateServerInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server-instances/{server-instance-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"server-instance-name"+"}", url.PathEscape(parameterValueToString(r.serverInstanceName, "serverInstanceName")), -1)

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
