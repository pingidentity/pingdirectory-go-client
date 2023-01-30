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

// MatchingRuleApiService MatchingRuleApi service
type MatchingRuleApiService service

type ApiGetMatchingRuleRequest struct {
	ctx              context.Context
	ApiService       *MatchingRuleApiService
	matchingRuleName string
}

func (r ApiGetMatchingRuleRequest) Execute() (*GetMatchingRule200Response, *http.Response, error) {
	return r.ApiService.GetMatchingRuleExecute(r)
}

/*
GetMatchingRule Returns a single Matching Rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param matchingRuleName Name of the Matching Rule to be read
	@return ApiGetMatchingRuleRequest
*/
func (a *MatchingRuleApiService) GetMatchingRule(ctx context.Context, matchingRuleName string) ApiGetMatchingRuleRequest {
	return ApiGetMatchingRuleRequest{
		ApiService:       a,
		ctx:              ctx,
		matchingRuleName: matchingRuleName,
	}
}

// Execute executes the request
//
//	@return GetMatchingRule200Response
func (a *MatchingRuleApiService) GetMatchingRuleExecute(r ApiGetMatchingRuleRequest) (*GetMatchingRule200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMatchingRule200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MatchingRuleApiService.GetMatchingRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/matching-rules/{matching-rule-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"matching-rule-name"+"}", url.PathEscape(parameterToString(r.matchingRuleName, "")), -1)

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

type ApiUpdateMatchingRuleRequest struct {
	ctx              context.Context
	ApiService       *MatchingRuleApiService
	matchingRuleName string
	updateRequest    *UpdateRequest
}

// Update an existing Matching Rule
func (r ApiUpdateMatchingRuleRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateMatchingRuleRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateMatchingRuleRequest) Execute() (*GetMatchingRule200Response, *http.Response, error) {
	return r.ApiService.UpdateMatchingRuleExecute(r)
}

/*
UpdateMatchingRule Update an existing Matching Rule by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param matchingRuleName Name of the Matching Rule to be updated
	@return ApiUpdateMatchingRuleRequest
*/
func (a *MatchingRuleApiService) UpdateMatchingRule(ctx context.Context, matchingRuleName string) ApiUpdateMatchingRuleRequest {
	return ApiUpdateMatchingRuleRequest{
		ApiService:       a,
		ctx:              ctx,
		matchingRuleName: matchingRuleName,
	}
}

// Execute executes the request
//
//	@return GetMatchingRule200Response
func (a *MatchingRuleApiService) UpdateMatchingRuleExecute(r ApiUpdateMatchingRuleRequest) (*GetMatchingRule200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMatchingRule200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MatchingRuleApiService.UpdateMatchingRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/matching-rules/{matching-rule-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"matching-rule-name"+"}", url.PathEscape(parameterToString(r.matchingRuleName, "")), -1)

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
