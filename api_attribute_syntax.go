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


// AttributeSyntaxApiService AttributeSyntaxApi service
type AttributeSyntaxApiService service

type ApiGetAttributeSyntaxRequest struct {
	ctx context.Context
	ApiService *AttributeSyntaxApiService
	attributeSyntaxName string
}

func (r ApiGetAttributeSyntaxRequest) Execute() (*GetAttributeSyntax200Response, *http.Response, error) {
	return r.ApiService.GetAttributeSyntaxExecute(r)
}

/*
GetAttributeSyntax Returns a single Attribute Syntax

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param attributeSyntaxName Name of the Attribute Syntax to be read
 @return ApiGetAttributeSyntaxRequest
*/
func (a *AttributeSyntaxApiService) GetAttributeSyntax(ctx context.Context, attributeSyntaxName string) ApiGetAttributeSyntaxRequest {
	return ApiGetAttributeSyntaxRequest{
		ApiService: a,
		ctx: ctx,
		attributeSyntaxName: attributeSyntaxName,
	}
}

// Execute executes the request
//  @return GetAttributeSyntax200Response
func (a *AttributeSyntaxApiService) GetAttributeSyntaxExecute(r ApiGetAttributeSyntaxRequest) (*GetAttributeSyntax200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAttributeSyntax200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeSyntaxApiService.GetAttributeSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/attribute-syntaxes/{attribute-syntax-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"attribute-syntax-name"+"}", url.PathEscape(parameterToString(r.attributeSyntaxName, "")), -1)

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

type ApiUpdateAttributeSyntaxRequest struct {
	ctx context.Context
	ApiService *AttributeSyntaxApiService
	attributeSyntaxName string
	updateRequest *UpdateRequest
}

// Update an existing Attribute Syntax
func (r ApiUpdateAttributeSyntaxRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateAttributeSyntaxRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateAttributeSyntaxRequest) Execute() (*GetAttributeSyntax200Response, *http.Response, error) {
	return r.ApiService.UpdateAttributeSyntaxExecute(r)
}

/*
UpdateAttributeSyntax Update an existing Attribute Syntax by name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param attributeSyntaxName Name of the Attribute Syntax to be updated
 @return ApiUpdateAttributeSyntaxRequest
*/
func (a *AttributeSyntaxApiService) UpdateAttributeSyntax(ctx context.Context, attributeSyntaxName string) ApiUpdateAttributeSyntaxRequest {
	return ApiUpdateAttributeSyntaxRequest{
		ApiService: a,
		ctx: ctx,
		attributeSyntaxName: attributeSyntaxName,
	}
}

// Execute executes the request
//  @return GetAttributeSyntax200Response
func (a *AttributeSyntaxApiService) UpdateAttributeSyntaxExecute(r ApiUpdateAttributeSyntaxRequest) (*GetAttributeSyntax200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAttributeSyntax200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeSyntaxApiService.UpdateAttributeSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/attribute-syntaxes/{attribute-syntax-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"attribute-syntax-name"+"}", url.PathEscape(parameterToString(r.attributeSyntaxName, "")), -1)

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
